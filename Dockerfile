# Stage 1: Build
FROM golang:1.25-alpine AS build

# Dependências do sistema
RUN apk add --no-cache git

# Diretório de trabalho
WORKDIR /app

# Copia go.mod e go.sum e instala dependências
COPY go.mod go.sum ./
RUN go mod download

# Copia o restante do código
COPY . .

# Compila o binário
RUN go build -o server ./src/main.go

# Stage 2: Runtime
FROM alpine:latest

# Dependências mínimas
RUN apk add --no-cache ca-certificates

WORKDIR /app

# Copia binário do build
COPY --from=build /app/server .

# Variável de ambiente para Firebase (será setada pelo Cloud Run)
# O valor virá do GitHub Actions (secrets)
ENV FIREBASE_CREDENTIALS=""

# Comando para rodar
CMD ["./server"]
