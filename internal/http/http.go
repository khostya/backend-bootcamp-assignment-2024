package http

import (
	"context"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/go-chi/cors"
	"github.com/khostya/backend-bootcamp-assignment-2024/internal/config"
	"github.com/khostya/backend-bootcamp-assignment-2024/internal/http/api"
	"github.com/khostya/backend-bootcamp-assignment-2024/internal/http/middleware"
	"github.com/khostya/backend-bootcamp-assignment-2024/internal/http/openapi"
	"github.com/khostya/backend-bootcamp-assignment-2024/pkg/auth"
	"github.com/khostya/backend-bootcamp-assignment-2024/pkg/httpserver"
	"github.com/oapi-codegen/nethttp-middleware"
	"log"
)

func MustRun(ctx context.Context, cfg config.HTTP, cache Cache, useCases UseCases, tokenManager auth.TokenManager) <-chan error {
	httpserver, err := newHttpServer(ctx, cfg, cache, useCases, tokenManager)
	if err != nil {
		res := make(chan error, 2)
		res <- err
		return res
	}

	httpserver.Start()
	return httpserver.Notify()
}

func newHttpServer(ctx context.Context, cfg config.HTTP, cache Cache, useCases UseCases, tokenManager auth.TokenManager) (*httpserver.Server, error) {
	openapi, err := openapi.GetOpenapiV3(ctx)
	if err != nil {
		return nil, err
	}

	options := &nethttpmiddleware.Options{
		SilenceServersWarning: true,
		Options: openapi3filter.Options{
			ExcludeRequestBody: true,
			AuthenticationFunc: func(ctx context.Context, input *openapi3filter.AuthenticationInput) error {
				return middleware.AuthenticateRequest(ctx, input, tokenManager)
			},
		},
	}

	server, err := newServer(useCases, cache)
	if err != nil {
		return nil, err
	}

	handler := api.HandlerWithOptions(server, api.StdHTTPServerOptions{
		Middlewares: []api.MiddlewareFunc{
			cors.AllowAll().Handler,
			nethttpmiddleware.OapiRequestValidatorWithOptions(openapi, options),
			middleware.AuthData(tokenManager),
		},
	})

	httpserver := httpserver.New(handler,
		httpserver.Port(cfg.Port),
		httpserver.IdleTimeout(cfg.IdleTimeout),
		httpserver.MaxHeaderBytes(cfg.MaxHeaderBytes),
		httpserver.WriteTimeout(cfg.WriteTimeout),
	)

	go func() {
		<-ctx.Done()

		if err := httpserver.Shutdown(); err != nil {
			log.Fatalf("HTTP handler Shutdown: %s", err)
		}
	}()

	return httpserver, nil
}
