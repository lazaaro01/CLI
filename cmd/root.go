package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "whatsapp-cli",
	Short: "WhatsApp CLI",
	Long:  "A CLI tool for interacting with WhatsApp Web",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(loginCmd)
	rootCmd.AddCommand(sendCmd)
	rootCmd.AddCommand(logoutCmd)
}
