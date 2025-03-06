# Order Service - REST, GraphQL e gRPC

Este projeto implementa um serviço de gerenciamento de pedidos (orders) utilizando três diferentes interfaces de API: REST, GraphQL e gRPC.
Portas Utilizadas

    REST API: 8080
    GraphQL: 8081
    gRPC: 50051

Pré-requisitos

    Docker
    Docker Compose
    Go 1.19 ou superior (para desenvolvimento local)
    make (opcional, para comandos de conveniência)

Tecnologias Utilizadas

    Go
    PostgreSQL
    gRPC
    GraphQL (gqlgen)
    Docker
    Migrate (para migrações do banco de dados)

Como Executar

Clone o repositório:

    git clone https://github.com/diillson/fullcycle-goexpert-desafio-clean-architecture.git
    cd fullcycle-goexpert-desafio-clean-architecture

Inicie os serviços usando Docker Compose:

    docker compose up --build

Endpoints e Exemplos de Uso
## REST API

Listar Orders:

    curl http://localhost:8080/order

Criar Order:

    curl -X POST http://localhost:8080/order \
    -H "Content-Type: application/json" \
    -d '{"customer_id": 1, "amount": 150.50, "status": "pending"}'

## GraphQL

Listar Orders:

    query {
        listOrders {
            id
            customerId
            amount
            status
            createdAt
        }
    }

Criar Order:

    mutation createOrder($input: CreateOrderInput!) {
        createOrder(input: $input) {
        id
        customerId
        amount
        status
        createdAt
        }
    }

### Variables:
    {
    "input": {
    "customerId": 1,
    "amount": 150.50,
    "status": "pending"
        }
    }

## gRPC

Listar Orders:

    grpcurl -plaintext localhost:50051 order.OrderService/ListOrders

Criar Order:

    grpcurl -plaintext -d '{"customer_id": 1, "amount": 150.50, "status": "pending"}' \
    localhost:50051 order.OrderService/CreateOrder

## Estrutura do Projeto
```bash
.
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── database/
│   │   └── database.go
│   ├── graphql/
│   │   ├── schema.graphql
│   │   └── resolver.go
│   ├── grpc/
│   │   ├── proto/
│   │   │   └── order.proto
│   │   └── server.go
│   └── http/
│       └── handlers.go
├── migrations/
│   └── 000001_create_orders_table.up.sql
├── Dockerfile
├── docker-compose.yml
├── entrypoint.sh
├── api.http
└── README.md
```

# Desenvolvimento
Gerando Código

Gerar código GraphQL:

    go run github.com/99designs/gqlgen generate

Gerar código gRPC:

    protoc --go_out=. --go-grpc_out=. internal/grpc/proto/order.proto

### Migrações

As migrações são executadas automaticamente quando o container inicia, mas você pode executá-las manualmente:

migrate -path migrations -database "postgresql://root:root@localhost:5432/orders?sslmode=disable" up

### Testes

Use o arquivo api.http incluído no projeto para testar todos os endpoints. Se estiver usando VS Code, instale a extensão "REST Client" para executar as requisições diretamente do editor.
Troubleshooting
Logs dos Containers

# Ver logs de todos os serviços
docker compose logs -f

# Ver logs de um serviço específico
docker compose logs -f app

Reiniciar Serviços

# Reiniciar todos os serviços
docker compose restart

# Reiniciar um serviço específico
docker compose restart app

Limpar Tudo

# Parar e remover containers, redes e volumes
docker compose down -v
