# FROM golang:1.20-alpine3.17 AS builder

# RUN apk add --no-cache git

# RUN apk add --update --no-cache curl

# # COPY . /go/src
# WORKDIR /app

# COPY go.mod go.sum ./
# RUN go mod download

# COPY . .

# RUN CGO_ENABLED=0 go build -o server .

# FROM alpine:latest

# # Instalar pacotes necessários
# RUN apk --no-cache add ca-certificates

# COPY --from=builder /app/server .

# # Expor a porta que o aplicativo irá usar
# EXPOSE 8081

# # Comando para rodar o servidor
# CMD ["./server"]



FROM golang:1.23.2-alpine AS builder

RUN apk add --no-cache git

RUN apk add --update --no-cache curl

WORKDIR /app

RUN go install github.com/pressly/goose/v3/cmd/goose@latest

COPY go.mod go.sum ./
RUN go mod download

# Expor a porta que o aplicativo irá usar
EXPOSE 8082

# Comando para rodar o servidor
CMD ["go", "run", "cmd/server/main.go"]
