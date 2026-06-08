package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"whatsapp-cli/internal/whatsapp"
)

var sendCmd = &cobra.Command{
	Use:   "send [phone] [message]",
	Short: "Send a WhatsApp message",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		client, err := whatsapp.NewClient()
		if err != nil {
			fmt.Fprintf(os.Stderr, "error creating client: %v\n", err)
			os.Exit(1)
		}
		defer client.Disconnect()

		resp, err := client.SendTextMessage(context.Background(), args[0], args[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "error sending message: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Message sent! ID: %s\n", resp)
	},
}
