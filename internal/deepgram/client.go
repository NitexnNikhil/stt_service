package deepgram

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Client struct {
	APIKey string
	Model  string
}

func NewClient(apiKey, model string) *Client {
	return &Client{
		APIKey: apiKey,
		Model:  model,
	}
}

func (c *Client) Transcribe(filePath string) (string, error) {

	audioBytes, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf(
		"https://api.deepgram.com/v1/listen?model=%s&punctuate=true&smart_format=true",
		c.Model,
	)

	req, err := http.NewRequest("POST", url, bytes.NewReader(audioBytes))
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Token "+c.APIKey)
	req.Header.Set("Content-Type", "audio/wav")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	return string(body), nil
}
