package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/celsobenedetti/go-sse/internal/server"
)

func main() {
	fmt.Println("main.go")
	//
	// ctx := context.Background()
	//
	// if err := run(ctx); err != nil {
	// 	fmt.Fprintf(os.Stderr, "%\n", err)
	// 	os.Exit(1)
	// }
}

func run(ctx context.Context) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	s := server.NewServer()

	go listenAndServe(s)
	<-ctx.Done()
	gracefulShutdown(s)

	return nil
}

func listenAndServe(s *http.Server) {
	fmt.Println("Listening on :3000")
	if err := s.ListenAndServe(); err != nil {
		fmt.Fprintf(os.Stderr, "error listening and serving: %s\n", err)
	}
}

func gracefulShutdown(s *http.Server) {
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := s.Shutdown(shutdownCtx); err != nil {
		fmt.Fprintf(os.Stderr, "error shutting down http server %d\n", err)
	}
}
