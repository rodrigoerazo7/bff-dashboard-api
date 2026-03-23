package main

import (
	"os"
	"time"
)

type config struct {
	addr    string
	baseURL string
	timeout time.Duration
}

func loadConfig() config {
	timeout, err := time.ParseDuration(envOrDefault("DASHBOARD_TIMEOUT", "2s"))
	if err != nil {
		timeout = 2 * time.Second
	}

	return config{
		addr:    envOrDefault("PORT", ":8080"),
		baseURL: envOrDefault("DUMMYJSON_BASE_URL", "https://dummyjson.com"),
		timeout: timeout,
	}
}

func envOrDefault(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
