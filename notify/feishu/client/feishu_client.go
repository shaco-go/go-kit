package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/shaco-go/go-kit/notify/feishu/message"
	"io"
	"net/http"
	"strings"
)

type Client struct {
	webhook string
}

func NewFeiShuClient(webhook string) *Client {
	return &Client{
		webhook: webhook,
	}
}

func (client *Client) send(msg any) error {
	messageContent, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	payload := strings.NewReader(string(messageContent))
	request, err := http.NewRequest(http.MethodPost, client.webhook, payload)
	if err != nil {
		return err
	}

	httpClient := http.Client{}
	request.Header.Add("Content-Type", "application/json")
	response, err := httpClient.Do(request)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}

	if response.StatusCode != http.StatusOK {
		errMessageResponse := &message.ErrMessageResponse{}
		err := json.Unmarshal(body, errMessageResponse)
		if err != nil {
			return err
		}
		return errors.New(errMessageResponse.Msg)
	}

	return nil
}

func (client *Client) SendTextMessage(text string) error {
	msg := message.NewTextMessage(text)
	return client.send(msg)
}

func (client *Client) SendPostMessage(title string, content [][]message.PostMessageContentPostZhCnContent) error {
	msg := message.NewPostMessage(title, content)
	return client.send(msg)
}

func (client *Client) SendImageMessage(imageKey string) error {
	msg := message.NewImageMessage(imageKey)
	return client.send(msg)
}

func (client *Client) SendShareChatMessage(shareChatId string) error {
	msg := message.NewShareChatMessage(shareChatId)
	return client.send(msg)
}

func (client *Client) SendInteractiveMessage() error {
	return nil
}
