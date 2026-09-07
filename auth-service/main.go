package main

import (
	"context"
	"log"

	pb "github.com/your-repo/transaction-service/proto"
	"google.golang.org/grpc"
)

type server struct {}

func (s *server) ProcessTransaction(ctx context.Context, req *pb.TransactionRequest) (*pb.TransactionResponse, error) {
	// Implement transaction processing logic here
	return &pb.TransactionResponse{Status: "success"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err!= nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTransactionServiceServer(s, &server{})
	log.Println("Transaction service started at :50052")
	if err := s.Serve(lis); err!= nil {
		log.Fatalf("failed to serve: %v", err)
	}
}