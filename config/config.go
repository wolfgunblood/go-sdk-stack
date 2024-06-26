package config

import (
	"os"
)

type Config struct {
    APIKey    string
    APISecret string
    BaseURL   string
}

func LoadConfig() Config {
    return Config{
        APIKey:    os.Getenv("STACK_API_KEY"),
        APISecret: os.Getenv("STACK_API_SECRET"),
        BaseURL:   os.Getenv("STACK_BASE_URL"),
    }
}
