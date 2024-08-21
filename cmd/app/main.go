package main

import (
	"context"
	"errors"
	"github.com/khostya/backend-bootcamp-assignment-2024/internal/app"
	"github.com/khostya/backend-bootcamp-assignment-2024/internal/config"
	"log"
	"net/http"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalln(err)
	}

	if err := app.Run(ctx, cfg); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalln(err)
	}
}
