package db

import (
	"chat_app/internal/models"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoInstance contains the Mongo client and database names
type MongoInstance struct {
	Client *mongo.Client
	DB     string
}

var MI MongoInstance

func CreateUser(user *models.User) (*models.User, error) {
	collection := MI.Client.Database(MI.DB).Collection("users")
	user.ID = primitive.NewObjectID() // Assign a new ID
	_, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func FindUserByUsername(username string) (*models.User, error) {
	var user models.User
	collection := MI.Client.Database(MI.DB).Collection("users")
	err := collection.FindOne(context.Background(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// ConnectMongo initializes a client connection to MongoDB.
func ConnectMongo(uri, dbName string) {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf("Failed to create new MongoDB client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to MongoDB
	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Ping MongoDB
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	log.Println("Connected to MongoDB!")

	MI = MongoInstance{
		Client: client,
		DB:     dbName,
	}
}

func FindUserByEmail(email string) (*models.User, error) {
	var user models.User
	collection := MI.Client.Database(MI.DB).Collection("users")

	// Create a context with a timeout that suits your application
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err := collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			// Handle the case where the user is not found
			log.Printf("User with email %s not found", email)
			return nil, nil // Or return an error as per your error handling strategy
		}
		// Handle other potential errors
		log.Printf("Error finding user by email: %v", err)
		return nil, err
	}

	return &user, nil
}
