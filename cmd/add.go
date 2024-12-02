package cmd

import (
	"fmt"
	"gitctx/internal/ssh"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [account] [email]",
	Short: "Add a new account with an SSH key",
	Long:  `Adds a new Git account with a corresponding SSH key and sets up the necessary configurations.`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		account := args[0]
		email := args[1]
		err := ssh.AddSSHKey(account, email)
		if err != nil {
			fmt.Printf("Error adding account %s: %v\n", account, err)
		} else {
			fmt.Printf("Account %s added with SSH key.\n", account)
		}
	},
}
