package whatsapp

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/mdp/qrterminal/v3"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/store/sqlstore"
	_ "modernc.org/sqlite"
)

type Client struct {
	waCli *whatsmeow.Client
}

func NewClient() (*Client, error) {
	sessionDir := filepath.Join(".", "session")
	if err := os.MkdirAll(sessionDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create session directory: %w", err)
	}

	dbPath := filepath.Join(sessionDir, "whatsapp.db")

	container, err := sqlstore.New(context.Background(), "sqlite", fmt.Sprintf("file:%s?_pragma=foreign_keys(1)", dbPath), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create store container: %w", err)
	}

	deviceStore, err := container.GetFirstDevice(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get device store: %w", err)
	}

	cli := whatsmeow.NewClient(deviceStore, nil)
	return &Client{waCli: cli}, nil
}

func (c *Client) Login(ctx context.Context) error {
	if c.waCli.Store.ID != nil {
		fmt.Println("Restoring existing session...")
		return c.waCli.Connect()
	}

	qrChan, err := c.waCli.GetQRChannel(ctx)
	if err != nil {
		return fmt.Errorf("failed to get QR channel: %w", err)
	}

	if err := c.waCli.Connect(); err != nil {
		return fmt.Errorf("failed to connect: %w", err)
	}

	for evt := range qrChan {
		switch evt.Event {
		case whatsmeow.QRChannelEventCode:
			qrterminal.GenerateHalfBlock(evt.Code, qrterminal.L, os.Stdout)
			fmt.Println("\nScan the QR code with your phone")
		case whatsmeow.QRChannelSuccess.Event:
			fmt.Println("Login successful!")
		default:
			if evt.Event == whatsmeow.QRChannelEventError {
				fmt.Fprintf(os.Stderr, "Error: %v\n", evt.Error)
			}
		}
	}

	return nil
}

func (c *Client) Logout(ctx context.Context) error {
	if c.waCli.Store.ID == nil {
		return fmt.Errorf("not logged in")
	}

	if err := c.waCli.Connect(); err != nil {
		return fmt.Errorf("failed to connect: %w", err)
	}

	return c.waCli.Logout(ctx)
}

func (c *Client) Disconnect() {
	c.waCli.Disconnect()
}
