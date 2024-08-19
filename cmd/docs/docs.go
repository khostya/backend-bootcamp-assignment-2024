package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/khostya/backend-bootcamp-assignment-2024/internal/config"
	"github.com/khostya/backend-bootcamp-assignment-2024/internal/http/openapi"
	"github.com/swaggo/swag"
	"strings"
)

func init() {
	openapi, err := openapi.GetOpenapiV3(context.Background())
	if err != nil {
		panic(err)
	}

	api := config.MustNewConfig()
	openapi.AddServer(&openapi3.Server{URL: fmt.Sprintf("http://localhost:%s", api.HTTP.Port)})

	swaggerJSON, err := openapi.MarshalJSON()
	if err != nil {
		panic(err)
	}

	template := string(swaggerJSON)

	m := make(map[string]any)
	err = json.NewDecoder(strings.NewReader(template)).Decode(&m)
	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	err = json.NewEncoder(&buf).Encode(m)
	if err != nil {
		panic(err)
	}

	swagger := &swag.Spec{
		InfoInstanceName: "swagger",
		SwaggerTemplate:  buf.String(),
	}

	swag.Register(swagger.InstanceName(), swagger)
}
