# ===========================
# Stage 1: Build
# ===========================
FROM golang:1.25-alpine AS builder

RUN apk add --no-cache git ca-certificates

WORKDIR /app

# Copiar go.mod/go.sum e baixar dependências
COPY go.mod go.sum ./
RUN go mod download

# Copiar todo o código fonte
COPY ./src ./src

# Compilar binário
WORKDIR /app/src
RUN CGO_ENABLED=0 GOOS=linux go build -a -o ../server main.go

# ===========================
# Stage 2: Minimal image
# ===========================
FROM alpine:latest

RUN apk add --no-cache ca-certificates

WORKDIR /app

# Copiar binário
COPY --from=builder /app/server .

# Porta usada pelo Cloud Run
ENV PORT=8080

CMD ["./server"]
