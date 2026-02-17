package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"example.com/crudapp/internal/api"
	"example.com/crudapp/internal/store"
	memorystore "example.com/crudapp/internal/store/memory"
	sqlitestore "example.com/crudapp/internal/store/sqlite"
)

func main() {
    port := envOrDefault("PORT", "8080")
    storage := envOrDefault("STORAGE", "memory")

    // Initialize the store based on the STORAGE environment variable
    st, err := buildStore(storage)
    if err != nil {
        log.Fatalf("store init failed: %v", err)
    }

    // Create the HTTP server
    srv := &http.Server{
        Addr:              ":" + port,
        Handler:           api.NewServer(st).Handler(),
        ReadHeaderTimeout: 5 * time.Second,
    }

    shutdownCh := make(chan os.Signal, 1)
    signal.Notify(shutdownCh, syscall.SIGINT, syscall.SIGTERM)

    // Start the server in a separate goroutine
    go func() {
        <-shutdownCh
        ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
        defer cancel()
        if err := srv.Shutdown(ctx); err != nil {
            log.Printf("shutdown error: %v", err)
        }
    }()

    // Start the server
    log.Printf("listening on :%s (storage=%s)", port, storage)
    if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
        log.Fatalf("server error: %v", err)
    }
}

func envOrDefault(key, fallback string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return fallback
}

func buildStore(storage string) (store.Store, error) {
    switch storage {
    case "memory":
        return memorystore.New(), nil
    case "sqlite":
        return sqlitestore.New("file:crudapp.db?_busy_timeout=5000"), nil
    default:
        return nil, store.ErrUnsupportedStore
    }
}
