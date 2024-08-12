FROM golang:1.22.5 as modules
COPY go.mod go.sum /modules/
WORKDIR /modules
RUN go mod download

FROM golang:1.22.5
COPY --from=modules /go/pkg /go/pkg
COPY .. /app
WORKDIR /app

RUN go env -w CGO_ENABLED=0
RUN go env -w GOOS=linux
RUN go env -w GOARCH=amd64

ENV CONFIG_PATH=/config/config.yml

COPY /config /config
CMD ["go", "test", "./tests/postgres/...", "-tags=integration"]