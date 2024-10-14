package cmd

import (
    "fmt"
    "os"
    "path/filepath"
    "gitctx/internal/ssh"

    "github.com/spf13/cobra"
)

var switchCmd = &cobra.Command{
    Use:   "switch [account]",
    Short: "Switch to a different account",
    Long:  `Switch to a different Git account by changing the SSH key used for git operations.`,
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        account := args[0]
        err := ssh.SwitchSSHKey(account)
        if err != nil {
            fmt.Printf("Error switching to account %s: %v\n", account, err)
            return
        }

        // Geçerli hesabı sakla
        currentAccountFile := filepath.Join(os.Getenv("HOME"), ".gitctx_current")
        if err := os.WriteFile(currentAccountFile, []byte(account), 0644); err != nil {
            fmt.Printf("Error saving current account: %v\n", err)
        } else {
            fmt.Printf("Switched to %s account.\n", account)
        }
    },
}
