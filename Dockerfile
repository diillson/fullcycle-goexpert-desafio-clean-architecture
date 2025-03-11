FROM golang:1.24-alpine

# Instala dependências necessárias
RUN apk add --no-cache postgresql-client netcat-openbsd protobuf-dev gcc musl-dev

# Instala as ferramentas necessárias
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
RUN go install github.com/99designs/gqlgen@latest

WORKDIR /app

# Copia os arquivos necessários
COPY . .
COPY entrypoint.sh /entrypoint.sh

# Dá permissão de execução ao entrypoint
RUN chmod +x /entrypoint.sh

# Garante as dependencias do projeto
RUN go mod download

# Gera os arquivos do gRPC e GraphQL
RUN protoc --go_out=. --go-grpc_out=. internal/grpc/proto/order.proto
RUN go run github.com/99designs/gqlgen generate

# Compila a aplicação
RUN go build -o main ./cmd/server/main.go

# Expõe as portas necessárias
EXPOSE 8080 50051 8081

# Define o entrypoint
ENTRYPOINT ["/entrypoint.sh"]