package cmd

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
	"whatsapp-cli/internal/whatsapp"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Connect to WhatsApp via QR Code",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := whatsapp.NewClient()
		if err != nil {
			fmt.Fprintf(os.Stderr, "error creating client: %v\n", err)
			os.Exit(1)
		}

		if err := client.Login(context.Background()); err != nil {
			fmt.Fprintf(os.Stderr, "error logging in: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("Connected! Press Ctrl+C to disconnect.")
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
		<-ch
	},
}
