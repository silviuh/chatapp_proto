package chat

import (
	pb "chat_app/api/proto/gen"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct { // Change this line
	pb.UnimplementedChatServiceServer
}

func (s *Server) SendMessage(ctx context.Context, in *pb.SendMessageRequest) (*pb.SendMessageResponse, error) { // And this line
	log.Printf("Received message from %s: %s", in.GetUser(), in.GetMessage())
	return &pb.SendMessageResponse{Success: true}, nil
}

func StartGRPCServer() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterChatServiceServer(s, &Server{}) // And this line
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
