package cmd

import (
    "github.com/spf13/cobra"
    "os"
)

var rootCmd = &cobra.Command{
    Use:   "gitctx",
    Short: "Git context manager",
    Long:  `Git context manager for handling multiple SSH keys and Git accounts.`,
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        os.Exit(1)
    }
}

func init() {
    rootCmd.AddCommand(addCmd)
    rootCmd.AddCommand(removeCmd)
    rootCmd.AddCommand(switchCmd)
    rootCmd.AddCommand(listCmd)  // List komutunu burada ekliyoruz
}
