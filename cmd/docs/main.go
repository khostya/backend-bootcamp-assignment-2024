package main

import (
	"backend-bootcamp-assignment-2024/internal/config"
	"fmt"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"log"
	"net/http"
)

func main() {
	cfg := config.MustNewConfig()

	swaggerHandler := httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
	)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.HTTP.SwaggerPort), swaggerHandler); err != nil {
		log.Fatalln(err)
	}
}
