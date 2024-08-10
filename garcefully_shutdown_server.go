package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"meggase": "hello"})
	})

	// creating context with cancel to automatically cancel the request processing once server is shut down
/*	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()*/

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	// channel for interrupt signal
	quite := make(chan os.Signal, 1)
	signal.Notify(quite, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			os.Exit(1)
		}
	}()

	<- quite
	log.Info("shutdown the server ...")

	// create a conetxt with timeout to allow ongoing task is completed
	ctx, cancel := context.WithTimeout(context.Background(), 5 *time.Second)

	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("server forced to shutdown:", err)
	}

	log.Info("server closed")

}