package chat

import (
	pb "chat_app/api/proto/gen"
	"chat_app/internal/models"
	"chat_app/pkg/db"
	"chat_app/utils"
	"context"
	"log"
)

type Server struct { // Change this line
	pb.UnimplementedChatServiceServer
}

func (s *Server) SendMessage(ctx context.Context, in *pb.SendMessageRequest) (*pb.SendMessageResponse, error) { // And this line
	log.Printf("Received message from %s: %s", in.GetUser(), in.GetMessage())
	return &pb.SendMessageResponse{Success: true}, nil
}

func (s *Server) RegisterUser(ctx context.Context, in *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error) {
	// Hash the password (using bcrypt or a similar library)
	hashedPassword, err := utils.HashPassword(in.GetPassword())
	if err != nil {
		return nil, err
	}

	// Create a new user in MongoDB
	user := models.User{
		Username: in.GetUsername(),
		Email:    in.GetEmail(),
		Password: hashedPassword,
	}
	_, err = db.CreateUser(&user) // Implement this function to interact with MongoDB
	if err != nil {
		return nil, err
	}

	return &pb.RegisterUserResponse{Success: true}, nil
}

func (s *Server) AuthenticateUser(ctx context.Context, in *pb.AuthenticateUserRequest) (*pb.AuthenticateUserResponse, error) {
	user, err := db.FindUserByUsername(in.GetUsername())
	if err != nil {
		// Handle error (user not found or other error)
		return nil, err
	}

	// Assuming you have a function to check passwords that returns a boolean
	if utils.CheckPasswordHash(in.GetPassword(), user.Password) {
		// Mark user as online
		db.MarkUserOnline(user.Email) // Assuming user.ID is of type primitive.ObjectID

		return &pb.AuthenticateUserResponse{Success: true}, nil
	} else {
		return &pb.AuthenticateUserResponse{Success: false}, nil
	}
}

func (s *Server) GetUserDetails(userEmail string) (bool, error) {
	return db.IsUserOnline(userEmail), nil
}
