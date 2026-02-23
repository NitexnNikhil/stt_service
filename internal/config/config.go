package config

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type Config struct {
	DeepgramAPIKey string
	DeepgramModel  string
}

func Load() *Config {
	loadEnvFile(".env")

	return &Config{
		DeepgramAPIKey: os.Getenv("DEEPGRAM_API_KEY"),
		DeepgramModel:  os.Getenv("DEEPGRAM_MODEL"),
	}
}

// loadEnvFile reads a .env file and sets environment variables
func loadEnvFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Println("No .env file found")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Skip empty lines and comments
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		// Split on first '=' only
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		// Remove surrounding quotes if present
		value = strings.Trim(value, `"'`)

		os.Setenv(key, value)
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Error reading .env file: %v", err)
	}
}

// ```

// This implementation:
// - Uses only Go standard library (`bufio`, `os`, `strings`)
// - Handles comments (lines starting with `#`)
// - Strips quotes from values
// - Skips malformed lines gracefully
// - Is completely dependency-free

// Your `.env` file would look like:
// ```
// DEEPGRAM_API_KEY=your_key_here
// DEEPGRAM_MODEL=nova-2
