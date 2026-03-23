package main

import (
	"bff-dashboard-api/internal"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"
)

func main() {
	addr := envOrDefault("PORT", ":8080")
	if addr[0] != ':' {
		addr = ":" + addr
	}

	httpClient := &http.Client{Timeout: 2 * time.Second}
	dummyClient := internal.NewDummyJSONClient(httpClient, "https://dummyjson.com")
	dashboardService := internal.NewDashboardService(dummyClient)
	dashboardHandler := internal.NewDashboardHandler(dashboardService)

	mux := http.NewServeMux()
	mux.Handle("GET /dashboard/", dashboardHandler)

	slog.Info("bff-dashboard-api listening", "addr", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}

func envOrDefault(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
