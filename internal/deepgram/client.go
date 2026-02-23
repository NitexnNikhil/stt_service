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

// NewClient creates a new Deepgram client
func NewClient(apiKey, model string) *Client {
	return &Client{
		APIKey: apiKey,
		Model:  model,
	}
}

// Transcribe sends audio to Deepgram and returns the raw JSON response
func (c *Client) Transcribe(filePath string) (string, error) {
	// Read the audio file
	audioData, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to read audio file: %w", err)
	}

	// Build URL with parameters for high accuracy
	url := fmt.Sprintf(
		"https://api.deepgram.com/v1/listen?"+
			"model=%s&"+
			"punctuate=true&"+
			"smart_format=true&"+
			"diarize=false&"+
			"numerals=true&"+
			"profanity_filter=false&"+
			"redact=false&"+
			"utterances=false",
		c.Model,
	)

	// Create HTTP request
	req, err := http.NewRequest("POST", url, bytes.NewReader(audioData))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	req.Header.Set("Authorization", "Token "+c.APIKey)
	req.Header.Set("Content-Type", "audio/wav")

	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %w", err)
	}

	// Check for API errors
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Deepgram API error (status %d): %s", resp.StatusCode, string(body))
	}

	// Return raw JSON response
	return string(body), nil
}
