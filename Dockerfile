# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copiar arquivos de dependência
COPY go.mod go.sum ./

# Baixar dependências
RUN go mod download

# Copiar código fonte
COPY . .

# Compilar a aplicação
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/api

# Final stage
FROM alpine:latest

WORKDIR /app

# Copiar o binário compilado
COPY --from=builder /app/main .

# Expor porta
EXPOSE 8080

# Executar a aplicação
CMD ["./main"] 