package discordclient

import (
	"bytes"
	"context"
	"dapp-moderator/utils/logger"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"io"
	"net/http"

	"golang.org/x/net/context/ctxhttp"
)

type Client struct {
}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) SendMessage(ctx context.Context, webhookURL string, message Message) error {
	logger.AtLog.Logger.Info(fmt.Sprintf("SendMessage - %s", message.Content), zap.Any("message", message))

	var buf bytes.Buffer
	// err := json.NewEncoder(&buf).Encode(Test{X: 1})
	err := json.NewEncoder(&buf).Encode(message)
	if err != nil {

		logger.AtLog.Logger.Error(fmt.Sprintf("SendMessage - %s", message.Content), zap.Any("message", message), zap.Error(err))
		return err
	}

	resp, err := ctxhttp.Post(ctx, http.DefaultClient, webhookURL, "application/json", &buf)
	if err != nil {
		logger.AtLog.Logger.Error(fmt.Sprintf("SendMessage - %s", message.Content), zap.Any("message", message), zap.Error(err))
		return err
	}

	if resp.StatusCode != 200 && resp.StatusCode != 204 {
		defer resp.Body.Close()

		responseBody, err := io.ReadAll(resp.Body)
		if err != nil {
			logger.AtLog.Logger.Error(fmt.Sprintf("SendMessage - %s", message.Content), zap.Any("message", message), zap.Error(err))
			return err
		}

		err = fmt.Errorf(string(responseBody))
		logger.AtLog.Logger.Error(fmt.Sprintf("SendMessage - %s", message.Content), zap.Any("message", message), zap.Error(err))
		return err
	}

	return nil
}
