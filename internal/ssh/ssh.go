package ssh

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const (
	defaultSSHKeyType = "ed25519"
	defaultKeyBits    = 4096
)

// SwitchSSHKey switches the active SSH key
func SwitchSSHKey(account string) error {
	if account == "" {
		return fmt.Errorf("account name cannot be empty")
	}

	configFile := filepath.Join(os.Getenv("HOME"), ".gitctx_config")

	// Mevcut konfigürasyonu yedekle
	backupFile := filepath.Join(os.Getenv("HOME"), ".gitctx_config.backup")
	if err := copyFile(configFile, backupFile); err != nil {
		fmt.Printf("Warning: Failed to create backup: %v\n", err)
	}

	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		return fmt.Errorf("Error reading config file: %v", err)
	}

	accounts := strings.Split(string(data), "\n")
	var sshKeyPath string
	for _, acc := range accounts {
		if strings.HasPrefix(acc, account+":") {
			accountParts := strings.Split(acc, ":")
			if len(accountParts) > 2 {
				sshKeyPath = accountParts[2]
			}
		}
	}

	if sshKeyPath == "" {
		return fmt.Errorf("SSH key for account %s does not exist", account)
	}

	if _, err := os.Stat(sshKeyPath); os.IsNotExist(err) {
		return fmt.Errorf("SSH key does not exist at %s", sshKeyPath)
	}

	gitConfigCmd := exec.Command("git", "config", "--global", "core.sshCommand", fmt.Sprintf("ssh -i %s -o IdentitiesOnly=yes", sshKeyPath))
	return gitConfigCmd.Run()
}

// AddSSHKey adds a new SSH key for the specified account
func AddSSHKey(account, email string) error {
	if account == "" || email == "" {
		return fmt.Errorf("account and email cannot be empty")
	}

	sshDir := filepath.Join(os.Getenv("HOME"), ".ssh")
	if err := os.MkdirAll(sshDir, 0700); err != nil {
		return fmt.Errorf("failed to create .ssh directory: %v", err)
	}

	sshKeyPath := filepath.Join(sshDir, fmt.Sprintf("id_rsa_%s", account))

	// Check if SSH key already exists
	if _, err := os.Stat(sshKeyPath); err == nil {
		return fmt.Errorf("SSH key for account %s already exists", account)
	}

	// Generate the SSH key
	sshKeygenCmd := exec.Command("ssh-keygen", "-t", defaultSSHKeyType, "-C", email, "-f", sshKeyPath)
	if err := sshKeygenCmd.Run(); err != nil {
		return fmt.Errorf("failed to generate SSH key: %v", err)
	}

	fmt.Println("SSH key generated:", sshKeyPath)

	// Ensure the SSH config is set to use the new key for github.com
	configEntry := fmt.Sprintf(`
Host github.com
    HostName github.com
    User git
    IdentityFile %s
    IdentitiesOnly yes
`, sshKeyPath)

	sshConfigPath := filepath.Join(sshDir, "config")
	file, err := os.OpenFile(sshConfigPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return fmt.Errorf("failed to open SSH config: %v", err)
	}
	defer file.Close()

	if _, err := file.WriteString(configEntry); err != nil {
		return fmt.Errorf("failed to update SSH config: %v", err)
	}

	// Update gitctx config file
	configFile := filepath.Join(os.Getenv("HOME"), ".gitctx_config")
	configEntry = fmt.Sprintf("%s:%s:%s\n", account, email, sshKeyPath)

	configFileHandle, err := os.OpenFile(configFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return fmt.Errorf("failed to open gitctx config: %v", err)
	}
	defer configFileHandle.Close()

	if _, err := configFileHandle.WriteString(configEntry); err != nil {
		return fmt.Errorf("failed to update gitctx config: %v", err)
	}

	fmt.Printf("SSH config and gitctx config updated for account %s\n", account)
	return nil
}

// RemoveSSHKey removes the SSH key for the specified account
func RemoveSSHKey(account string) error {
	if account == "" {
		return fmt.Errorf("account name cannot be empty")
	}

	sshDir := filepath.Join(os.Getenv("HOME"), ".ssh")
	sshKeyPath := filepath.Join(sshDir, fmt.Sprintf("id_rsa_%s", account))

	if _, err := os.Stat(sshKeyPath); os.IsNotExist(err) {
		return fmt.Errorf("SSH key for account %s does not exist", account)
	}

	// Remove SSH key files
	if err := os.Remove(sshKeyPath); err != nil {
		return fmt.Errorf("failed to remove private key: %v", err)
	}
	if err := os.Remove(sshKeyPath + ".pub"); err != nil {
		return fmt.Errorf("failed to remove public key: %v", err)
	}

	// Remove the SSH config entry
	sshConfigPath := filepath.Join(sshDir, "config")
	if err := removeSSHConfigEntry(sshConfigPath, sshKeyPath); err != nil {
		fmt.Printf("Warning: Failed to remove SSH config entry: %v\n", err)
	}

	// Remove from gitctx config
	if err := removeGitctxConfigEntry(account); err != nil {
		fmt.Printf("Warning: Failed to remove gitctx config entry: %v\n", err)
	}

	// Reset git SSH command if this was the active key
	gitConfigCmd := exec.Command("git", "config", "--global", "--unset", "core.sshCommand")
	if err := gitConfigCmd.Run(); err != nil {
		fmt.Printf("Warning: Failed to unset git SSH command: %v\n", err)
	}

	fmt.Printf("Successfully removed SSH key and configurations for account %s\n", account)
	return nil
}

// Helper functions
func copyFile(src, dst string) error {
	input, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(dst, input, 0644)
}

func removeSSHConfigEntry(configPath, keyPath string) error {
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return err
	}

	lines := strings.Split(string(data), "\n")
	var newLines []string
	skip := false

	for _, line := range lines {
		if strings.Contains(line, keyPath) {
			skip = true
			continue
		}
		if skip && line == "" {
			skip = false
			continue
		}
		if !skip {
			newLines = append(newLines, line)
		}
	}

	return ioutil.WriteFile(configPath, []byte(strings.Join(newLines, "\n")), 0600)
}

func removeGitctxConfigEntry(account string) error {
	configFile := filepath.Join(os.Getenv("HOME"), ".gitctx_config")
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		return err
	}

	lines := strings.Split(string(data), "\n")
	var newLines []string

	for _, line := range lines {
		if !strings.HasPrefix(line, account+":") && line != "" {
			newLines = append(newLines, line)
		}
	}

	return ioutil.WriteFile(configFile, []byte(strings.Join(newLines, "\n")), 0600)
}
