ARG GO_VERSION="1.23-alpine"

# Stage 1: build the app
FROM golang:${GO_VERSION} AS builder

WORKDIR /app

COPY ./cmd/api/main.go .
COPY go.mod .
COPY go.sum .
COPY internal internal
COPY pkg pkg

RUN go build -o ./api-bin .

# Stage 2: Run the app
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/api-bin .

ENTRYPOINT [ "/app/api-bin" ]
