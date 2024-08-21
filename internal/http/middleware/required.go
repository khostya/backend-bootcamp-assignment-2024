package middleware

import (
	"context"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/khostya/backend-bootcamp-assignment-2024/pkg/auth"
	"strings"
)

func AuthenticateRequest(ctx context.Context, input *openapi3filter.AuthenticationInput, manager auth.TokenManager) error {
	request := input.RequestValidationInput.Request
	tokenHeader := strings.Split(request.Header.Get("Authorization"), " ")

	_, err := manager.ExtractUserId(tokenHeader)
	return err
}
