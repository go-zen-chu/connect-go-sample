package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-zen-chu/connect-go-sample/httpserver"
)

const (
	gracefulPeriod = 30 * time.Second
)

func main() {
	s := httpserver.NewHttpServer(":8080")

	// start signal handling
	runningCtx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	go func() {
		if err := s.Run(); err != nil {
			if !errors.Is(http.ErrServerClosed, err) {
				log.Fatalf("Failed to run server: %v", err)
			}
			log.Println("closing server...")
		}
	}()
	<-runningCtx.Done()

	log.Printf("Start shutdown server with graceful shutdown period %v", gracefulPeriod)
	gracefulCtx, cancel := context.WithTimeout(context.Background(), gracefulPeriod)
	defer cancel()
	if err := s.Shutdown(gracefulCtx); err != nil {
		log.Fatalf("Failed to shutdown server: %v", err)
	}
	log.Println("Server shutdown gracefully")
}
