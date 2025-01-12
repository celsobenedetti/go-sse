package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/celsobenedetti/go-sse/internal/server"
)

func main() {
	ctx := context.Background()

	if err := run(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "%\n", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	s := server.NewServer()

	listenAndServe := func() {
		fmt.Println("Listening on :3000")
		if err := s.ListenAndServe(); err != nil {
			fmt.Fprintf(os.Stderr, "error listening and serving: %s\n", err)
		}
	}

	gracefulShutdown := func() {
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := s.Shutdown(shutdownCtx); err != nil {
			fmt.Fprintf(os.Stderr, "error shutting down http server %d\n", err)
		}
	}

	go listenAndServe()
	<-ctx.Done()
	gracefulShutdown()

	return nil
}
