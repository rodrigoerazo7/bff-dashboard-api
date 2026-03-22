package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	addr := envOrDefault("PORT", ":8080")
	if addr[0] != ':' {
		addr = ":" + addr
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /dashboard/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	log.Printf("bff-dashboard-api listening on %s", addr)
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
