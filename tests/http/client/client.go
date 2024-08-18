//go:build integration

package client

import (
	"fmt"
	"github.com/khostya/backend-bootcamp-assignment-2024/internal/http/api"
	"os"
)

const endpointENV = "TEST_HTTP_ENDPOINT"

func New() *api.Client {
	server := os.Getenv(endpointENV)
	if server == "" {
		panic(fmt.Sprintf("undeined env %s", endpointENV))
	}

	client, err := api.NewClient(server)
	if err != nil {
		panic(fmt.Sprintf("undeined env %s", endpointENV))
	}

	return client
}
