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
    fmt.Println("Available accounts:")
    for _, account := range accounts {
        if account != "" {
            accountParts := strings.Split(account, ":")
            accountName := accountParts[0]
            if accountName == currentAccount {
                fmt.Printf("\033[1;32m- %s (SSH key: %s) [current]\033[0m\n", accountName, accountParts[2])
            } else {
                fmt.Printf("- %s (SSH key: %s)\n", accountName, accountParts[2])
            }
        }
    }
}
