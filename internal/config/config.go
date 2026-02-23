package config

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Config struct {
	DeepgramAPIKey string
	DeepgramModel  string
}

func Load() (*Config, error) {
	loadEnvFile(".env")

	config := &Config{
		DeepgramAPIKey: os.Getenv("DEEPGRAM_API_KEY"),
		DeepgramModel:  getEnvOrDefault("DEEPGRAM_MODEL", "nova-2"),
	}

	if err := config.Validate(); err != nil {
		return nil, err
	}

	return config, nil
}

func (c *Config) Validate() error {
	if c.DeepgramAPIKey == "" {
		return fmt.Errorf("DEEPGRAM_API_KEY is required")
	}
	if c.DeepgramModel == "" {
		return fmt.Errorf("DEEPGRAM_MODEL is required")
	}
	return nil
}

func (c *Config) Print() {
	log.Println("=== Configuration ===")
	log.Printf("DEEPGRAM_API_KEY: %s", maskKey(c.DeepgramAPIKey))
	log.Printf("DEEPGRAM_MODEL: %s", c.DeepgramModel)
	log.Println("====================")
}

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

		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		value = strings.Trim(value, `"'`)

		os.Setenv(key, value)
	}
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func maskKey(key string) string {
	if key == "" {
		return "[NOT SET]"
	}
	if len(key) <= 8 {
		return "****"
	}
	return key[:4] + "****" + key[len(key)-4:]
}
