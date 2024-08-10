package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Set up a simple HTTP server
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(2 * time.Second) // Simulate a long-running request
		fmt.Fprintf(w, "Hello, world!")
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	// Channel to listen for interrupt signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	// Goroutine to handle shutdown signal
	go func() {
		<-quit
		log.Println("Shutting down server...")

		// Create a context with a timeout to allow existing requests to complete
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		// Shut down the server
		if err := server.Shutdown(ctx); err != nil {
			log.Fatalf("Server forced to shutdown: %v", err)
		}

		log.Println("Server exiting")
	}()

	// Start the server
	log.Println("Starting server on :8080")
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("Server error: %v", err)
	}

	log.Println("Server stopped")
}
