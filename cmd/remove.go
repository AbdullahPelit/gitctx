package cmd

import (
	"fmt"
	"gitctx/internal/ssh"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove [account]",
	Short: "Remove an account and its SSH key",
	Long:  `Removes the specified account and its associated SSH key from the configuration.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		account := args[0]
		err := ssh.RemoveSSHKey(account)
		if err != nil {
			fmt.Printf("Error removing account %s: %v\n", account, err)
		} else {
			fmt.Printf("Account %s removed.\n", account)
		}
	},
}
