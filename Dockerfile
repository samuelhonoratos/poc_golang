FROM golang:1.23.2-alpine AS builder

RUN apk add --no-cache git

RUN apk add --update --no-cache curl

WORKDIR /app

RUN go install -tags 'mysql' github.com/pressly/goose/v3/cmd/goose@latest

COPY go.mod go.sum ./
RUN go mod download

EXPOSE 8082

CMD ["go", "run", "cmd/server/main.go"]
