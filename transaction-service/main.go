package main

import (
	"context"
	"log"

	pb "github.com/your-repo/notification-service/proto"
	"google.golang.org/grpc"
)

type server struct {}

func (s *server) SendNotification(ctx context.Context, req *pb.NotificationRequest) (*pb.NotificationResponse, error) {
	// Implement notification sending logic here
	return &pb.NotificationResponse{Status: "sent"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50053")
	if err!= nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterNotificationServiceServer(s, &server{})
	log.Println("Notification service started at :50053")
	if err := s.Serve(lis); err!= nil {
		log.Fatalf("failed to serve: %v", err)
	}
}