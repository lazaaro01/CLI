package whatsapp

import (
	"context"
	"fmt"

	"go.mau.fi/whatsmeow/proto/waE2E"
	"google.golang.org/protobuf/proto"
)

func (c *Client) SendTextMessage(ctx context.Context, phone, text string) (string, error) {
	if c.waCli.Store.ID == nil {
		return "", fmt.Errorf("not logged in. Run 'login' first")
	}

	if !c.waCli.IsConnected() {
		if err := c.waCli.Connect(); err != nil {
			return "", fmt.Errorf("failed to connect: %w", err)
		}
	}

	existences, err := c.waCli.IsOnWhatsApp(ctx, []string{"+" + phone})
	if err != nil {
		return "", fmt.Errorf("failed to check number: %w", err)
	}
	if len(existences) == 0 || !existences[0].IsIn {
		return "", fmt.Errorf("phone number %s is not registered on WhatsApp", phone)
	}

	jid := existences[0].JID
	msg := &waE2E.Message{
		Conversation: proto.String(text),
	}

	resp, err := c.waCli.SendMessage(ctx, jid, msg)
	if err != nil {
		return "", fmt.Errorf("failed to send message: %w", err)
	}

	return resp.ID, nil
}
