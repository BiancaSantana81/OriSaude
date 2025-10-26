# ===========================
# Stage 1: Build
# ===========================
FROM golang:1.25-alpine AS builder

RUN apk add --no-cache git ca-certificates

WORKDIR /app

# Baixar dependências
COPY go.mod go.sum ./
RUN go mod download

# Copiar código
COPY . .

# Compilar binário estático
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server .

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
