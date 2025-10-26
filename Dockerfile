# Etapa 1: Build do Go
FROM golang:1.21-alpine AS build

# Diretório de trabalho
WORKDIR /app

# Copia os módulos e faz o download das dependências
COPY go.mod go.sum ./
RUN go mod download

# Copia o restante do código
COPY . .

# Build da aplicação
RUN go build -o app ./src/main.go

# Etapa 2: Runtime mínimo
FROM alpine:3.18

# Diretório de trabalho
WORKDIR /app

# Copia binário do build
COPY --from=build /app/app .

# Copia service account do Firebase (se necessário)
COPY ./src/infra/database/serviceAccountKey.json ./serviceAccountKey.json

# Variáveis de ambiente (opcional: Cloud Run pode setar direto)
ENV GOOGLE_APPLICATION_CREDENTIALS=/app/serviceAccountKey.json

# Porta que Cloud Run vai expor
EXPOSE 8080

# Comando de execução
CMD ["./app"]
