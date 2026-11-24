package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type healthResp struct {
	Status  string `json:"status"`
	Service string `json:"service"`
	Version string `json:"version"`
}

type infoResp struct {
	Service   string `json:"service"`
	Version   string `json:"version"`
	Hostname  string `json:"hostname"`
	Timestamp string `json:"timestamp"`
}

func envOr(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}

func newMux(version string) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(w, http.StatusOK, healthResp{
			Status:  "ok",
			Service: "sunmao-test-go",
			Version: version,
		})
	})

	infoHandler := func(w http.ResponseWriter, _ *http.Request) {
		host, _ := os.Hostname()
		writeJSON(w, http.StatusOK, infoResp{
			Service:   "sunmao-test-go",
			Version:   version,
			Hostname:  host,
			Timestamp: time.Now().UTC().Format(time.RFC3339),
		})
	}
	mux.HandleFunc("/", infoHandler)
	mux.HandleFunc("/api/info", infoHandler)
	return mux
}

func main() {
	port := envOr("PORT", "8080")
	version := envOr("APP_VERSION", "0.1.0")
	srv := &http.Server{
		Addr:              ":" + port,
		Handler:           newMux(version),
		ReadHeaderTimeout: 5 * time.Second,
	}

	go func() {
		fmt.Printf("sunmao-test-go v%s listening on :%s\n", version, port)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			fmt.Fprintf(os.Stderr, "server error: %v\n", err)
			os.Exit(1)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_ = srv.Shutdown(ctx)
}
