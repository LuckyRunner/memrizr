package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/LuckyRunner/memrizr/account/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("Starting server...")

	router := gin.Default()

	handler.NewHandler(&handler.Config{
		R: router,
	})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to initialize server: %v", err)
		}
	}()

	log.Printf("Listening on port %s\n", srv.Addr)

	// Wait for kill signal of channel
	quit := make(chan os.Signal)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// This blocks until a signal is passed into the quit channel
	<-quit

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Failed to shutdown server: %v", err)
	}

	log.Println("Server shut down")
}
