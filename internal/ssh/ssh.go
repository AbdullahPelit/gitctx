package ssh

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// SwitchSSHKey switches the active SSH key
func SwitchSSHKey(account string) error {
	configFile := filepath.Join(os.Getenv("HOME"), ".gitctx_config")
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
				sshKeyPath = accountParts[2] // SSH anahtarının yolu burada
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
	sshKeyPath := filepath.Join(os.Getenv("HOME"), ".ssh", fmt.Sprintf("id_rsa_%s", account))

	// Check if SSH key already exists
	if _, err := os.Stat(sshKeyPath); err == nil {
		return fmt.Errorf("SSH key for account %s already exists", account)
	}

	// Generate the SSH key
	sshKeygenCmd := exec.Command("ssh-keygen", "-t", "ed25519", "-C", email, "-f", sshKeyPath)
	if err := sshKeygenCmd.Run(); err != nil {
		return err
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

	sshConfigPath := filepath.Join(os.Getenv("HOME"), ".ssh", "config")
	file, err := os.OpenFile(sshConfigPath, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err := file.WriteString(configEntry); err != nil {
		return err
	}

	fmt.Println("SSH config updated with new entry:", configEntry)

	return nil
}

// RemoveSSHKey removes the SSH key for the specified account
func RemoveSSHKey(account string) error {
	sshDir := filepath.Join(os.Getenv("HOME"), ".ssh")
	sshKeyPath := filepath.Join(sshDir, fmt.Sprintf("id_rsa_%s", account))

	if _, err := os.Stat(sshKeyPath); os.IsNotExist(err) {
		return fmt.Errorf("SSH key for account %s does not exist", account)
	}

	os.Remove(sshKeyPath)
	os.Remove(sshKeyPath + ".pub")

	// Remove the SSH config entry for github.com
	sshConfigCmd := exec.Command("sed", "-i.bak", "/Host github.com/,/^$/d", filepath.Join(sshDir, "config"))
	if err := sshConfigCmd.Run(); err != nil {
		return err
	}

	fmt.Println("SSH config entry removed.")

	gitConfigCmd := exec.Command("git", "config", "--global", "--unset", "core.sshCommand")
	return gitConfigCmd.Run()
}
