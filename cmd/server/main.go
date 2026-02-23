package main

import (
	"fmt"
	"log"

	"stt-service/internal/config"
	"stt-service/internal/deepgram"
	"stt-service/internal/service"
)

func main() {

	cfg := config.Load()

	client := deepgram.NewClient(
		cfg.DeepgramAPIKey,
		cfg.DeepgramModel,
	)
	log.Printf("Using Deepgram API Key: %s", cfg.DeepgramAPIKey)
	log.Printf("Using Deepgram Model: %s", cfg.DeepgramModel)

	transcriber := service.NewTranscriber(client)
	log.Println("Transcriber initialized")

	resultJSON, err := transcriber.Process("audio/playground-iBgb-kfwf.wav")
	if err != nil {
		log.Fatal(err)
	}

	err = service.SaveTranscript([]byte(resultJSON), "playground-iBgb-kfwf")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Transcript saved in storage/transcripts/playground-iBgb-kfwf.txt")
}
