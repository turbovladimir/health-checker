package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const BaseUrl = "https://api.telegram.org/bot%s"

type Message struct {
	ChatID int64  `json:"chat_id"`
	Text   string `json:"text"`
}

type Sender struct {
	token  string
	client *http.Client
}

type MyRoundTripper struct{}

func (t MyRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {

	// Do work before the request is sent
	fmt.Println(req)

	resp, err := http.DefaultTransport.RoundTrip(req)
	fmt.Println(resp)

	if err != nil {
		return resp, err
	}
	// Do work after the response is received

	return resp, err
}

func NewSender(token string) *Sender {
	rt := MyRoundTripper{}
	client := http.Client{Transport: rt}

	return &Sender{
		token:  token,
		client: &client,
	}
}

func (s *Sender) SendMessage(message *Message) error {
	payload, err := json.Marshal(message)
	if err != nil {
		return err
	}

	response, err := s.client.Post(fmt.Sprintf(BaseUrl, s.token)+"/sendMessage", "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	defer func(body io.ReadCloser) {
		if err := body.Close(); err != nil {
			log.Println("failed to close response body")
		}
	}(response.Body)
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send successful request. Status was %q", response.Status)
	}
	return nil
}
