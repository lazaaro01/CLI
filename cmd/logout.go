package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"whatsapp-cli/internal/whatsapp"
)

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Logout and remove session",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := whatsapp.NewClient()
		if err != nil {
			fmt.Fprintf(os.Stderr, "error creating client: %v\n", err)
			os.Exit(1)
		}
		defer client.Disconnect()

		if err := client.Logout(context.Background()); err != nil {
			fmt.Fprintf(os.Stderr, "error logging out: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Logged out successfully")
	},
}
