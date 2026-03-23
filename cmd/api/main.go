package main

import (
	"bff-dashboard-api/internal"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	cfg := loadConfig()
	slog.Info("config loaded", "addr", cfg.addr, "baseURL", cfg.baseURL, "timeout", cfg.timeout)

	addr := cfg.addr
	if addr[0] != ':' {
		addr = ":" + addr
	}

	dummyClient := internal.NewDummyJSONClient(cfg.baseURL, cfg.timeout)
	dashboardService := internal.NewDashboardService(dummyClient)
	dashboardHandler := internal.NewDashboardHandler(dashboardService)

	mux := http.NewServeMux()
	mux.Handle("GET /dashboard/", dashboardHandler)

	slog.Info("bff-dashboard-api listening", "addr", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		slog.Error("server failed", "err", err)
		os.Exit(1)
	}
}
