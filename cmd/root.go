package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	version     = "1.0.0" // Versiyon numarası
	versionFlag bool
)

var rootCmd = &cobra.Command{
	Use:   "gitctx",
	Short: "Git context manager",
	Long: `Git context manager for handling multiple SSH keys and Git accounts.
Complete documentation is available at https://github.com/AbdullahPelit/gitctx`,
	Run: func(cmd *cobra.Command, args []string) {
		if versionFlag {
			fmt.Printf("gitctx version %s\n", version)
			os.Exit(0)
		}
		// Eğer hiçbir komut verilmemişse help göster
		cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	// Version flag'i ekle
	rootCmd.Flags().BoolVarP(&versionFlag, "version", "v", false, "Print version information")

	// Alt komutları ekle
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(removeCmd)
	rootCmd.AddCommand(switchCmd)
	rootCmd.AddCommand(listCmd)
}
