package grpc

import (
	"context"
	"database/sql"
	"fullcycle-goexpert-desafio-clean-architecture/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type OrderService struct {
	proto.UnimplementedOrderServiceServer
	db *sql.DB
}

func NewOrderService(db *sql.DB) *OrderService {
	return &OrderService{db: db}
}

func (s *OrderService) ListOrders(ctx context.Context, req *proto.ListOrdersRequest) (*proto.ListOrdersResponse, error) {
	rows, err := s.db.QueryContext(ctx, "SELECT id, customer_id, amount, status, created_at FROM orders")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []*proto.Order
	for rows.Next() {
		order := &proto.Order{}
		if err := rows.Scan(&order.Id, &order.CustomerId, &order.Amount, &order.Status, &order.CreatedAt); err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	return &proto.ListOrdersResponse{Orders: orders}, nil
}

func (s *OrderService) CreateOrder(ctx context.Context, req *proto.CreateOrderRequest) (*proto.Order, error) {
	query := `INSERT INTO orders (customer_id, amount, status) VALUES ($1, $2, $3) RETURNING id, created_at`
	var id int64
	var createdAt string
	err := s.db.QueryRow(query, req.CustomerId, req.Amount, req.Status).Scan(&id, &createdAt)
	if err != nil {
		return nil, err
	}

	return &proto.Order{
		Id:         id,
		CustomerId: req.CustomerId,
		Amount:     req.Amount,
		Status:     req.Status,
		CreatedAt:  createdAt,
	}, nil
}

func RegisterGRPCServer(db *sql.DB) *grpc.Server {
	server := grpc.NewServer()
	orderService := NewOrderService(db)
	proto.RegisterOrderServiceServer(server, orderService)
	reflection.Register(server)
	return server
}
