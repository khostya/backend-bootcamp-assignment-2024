package main

import (
	"backend-bootcamp-assignment-2024/internal/app"
	"backend-bootcamp-assignment-2024/internal/config"
	"context"
	"log"
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

	if err := app.Run(ctx, cfg); err != nil {
		log.Fatalln(err)
	}
}
