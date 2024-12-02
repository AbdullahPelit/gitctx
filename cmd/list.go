package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all configured accounts",
	Long:  `List all accounts that have been added and configured with SSH keys.`,
	Run: func(cmd *cobra.Command, args []string) {
		listAccounts()
	},
}

func listAccounts() {
	configFile := filepath.Join(os.Getenv("HOME"), ".gitctx_config")
	currentAccountFile := filepath.Join(os.Getenv("HOME"), ".gitctx_current")

	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		fmt.Println("No accounts configured yet. Use 'gitctx add' to add an account.")
		return
	}

	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		fmt.Printf("Error reading config file: %v\n", err)
		return
	}

	currentAccount := ""
	if data, err := ioutil.ReadFile(currentAccountFile); err == nil {
		currentAccount = strings.TrimSpace(string(data))
	}

	accounts := strings.Split(string(data), "\n")
	if len(accounts) == 0 || (len(accounts) == 1 && accounts[0] == "") {
		fmt.Println("No accounts configured yet. Use 'gitctx add' to add an account.")
		return
	}

	fmt.Println("\nConfigured Git Accounts:")
	fmt.Println("------------------------")
	for _, account := range accounts {
		if account != "" {
			accountParts := strings.Split(account, ":")
			accountName := accountParts[0]
			if accountName == currentAccount {
				fmt.Printf("\033[1;32m✓ %s\033[0m\n", accountName)
				fmt.Printf("  Provider: %s\n", accountParts[1])
				fmt.Printf("  SSH Key: %s\n", accountParts[2])
				fmt.Println("  Status: Active")
			} else {
				fmt.Printf("• %s\n", accountName)
				fmt.Printf("  Provider: %s\n", accountParts[1])
				fmt.Printf("  SSH Key: %s\n", accountParts[2])
				fmt.Println("  Status: Inactive")
			}
			fmt.Println()
		}
	}
}
