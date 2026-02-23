package service

import (
	"encoding/json"
	"os"
	"path/filepath"

	"stt-service/internal/deepgram"
)

type Transcriber struct {
	Client *deepgram.Client
}

// Constructor
func NewTranscriber(client *deepgram.Client) *Transcriber {
	return &Transcriber{
		Client: client,
	}
}

// Calls Deepgram
func (t *Transcriber) Process(filePath string) (string, error) {
	return t.Client.Transcribe(filePath)
}

type DeepgramResponse struct {
	Results struct {
		Channels []struct {
			Alternatives []struct {
				Transcript string `json:"transcript"`
			} `json:"alternatives"`
		} `json:"channels"`
	} `json:"results"`
}

func SaveTranscript(rawJSON []byte, fileName string) error {

	var dgResp DeepgramResponse

	err := json.Unmarshal(rawJSON, &dgResp)
	if err != nil {
		return err
	}

	transcript := dgResp.Results.Channels[0].Alternatives[0].Transcript

	dir := "storage/transcripts"
	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return err
	}

	filePath := filepath.Join(dir, fileName+".txt")

	return os.WriteFile(filePath, []byte(transcript), 0644)
}
