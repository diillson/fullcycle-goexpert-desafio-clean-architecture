package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"fullcycle-goexpert-desafio-clean-architecture/internal/database"
	"fullcycle-goexpert-desafio-clean-architecture/internal/graphql"
	grpcServer "fullcycle-goexpert-desafio-clean-architecture/internal/grpc"
	httpHandler "fullcycle-goexpert-desafio-clean-architecture/internal/http"
	"fullcycle-goexpert-desafio-clean-architecture/proto"

	"github.com/99designs/gqlgen/graphql/handler"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// Configuração do contexto para shutdown graceful
	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Canal para sinais do sistema
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	// Inicializa conexão com o banco de dados
	db, err := database.NewDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Configuração do servidor HTTP
	orderHandler := httpHandler.NewOrderHandler(db)
	mux := http.NewServeMux()
	mux.HandleFunc("/order", orderHandler.HandleOrder)

	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	// Configuração do servidor gRPC
	grpcSrv := grpc.NewServer()
	orderService := grpcServer.NewOrderService(db)
	proto.RegisterOrderServiceServer(grpcSrv, orderService)
	reflection.Register(grpcSrv)

	// Configuração do servidor GraphQL
	resolver := graphql.NewResolver(db)
	graphqlSrv := handler.NewDefaultServer(graphql.NewExecutableSchema(graphql.Config{Resolvers: resolver}))

	graphqlHandler := http.NewServeMux()
	graphqlHandler.Handle("/query", graphqlSrv)

	graphqlServer := &http.Server{
		Addr:    ":8081",
		Handler: graphqlHandler,
	}

	// Inicia os servidores em goroutines separadas
	go func() {
		log.Printf("Starting HTTP server on :8080")
		if err := httpServer.ListenAndServe(); err != http.ErrServerClosed {
			log.Printf("HTTP server error: %v", err)
		}
	}()

	go func() {
		lis, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatalf("Failed to listen: %v", err)
		}
		log.Printf("Starting gRPC server on :50051")
		if err := grpcSrv.Serve(lis); err != nil {
			log.Printf("gRPC server error: %v", err)
		}
	}()

	go func() {
		log.Printf("Starting GraphQL server on :8081")
		if err := graphqlServer.ListenAndServe(); err != http.ErrServerClosed {
			log.Printf("GraphQL server error: %v", err)
		}
	}()

	// Aguarda sinal de término
	<-signalChan
	log.Println("Shutting down servers...")

	// Shutdown graceful
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownCancel()

	// Fecha os servidores
	if err := httpServer.Shutdown(shutdownCtx); err != nil {
		log.Printf("HTTP server shutdown error: %v", err)
	}

	if err := graphqlServer.Shutdown(shutdownCtx); err != nil {
		log.Printf("GraphQL server shutdown error: %v", err)
	}

	grpcSrv.GracefulStop()

	log.Println("Servers shutdown completed")
}
