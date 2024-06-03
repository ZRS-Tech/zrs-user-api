// helper/helper.go
package helper

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SayHello() {
	fmt.Println("Hello from the helper package!")
}

func ConnectDB() *mongo.Collection {
	config := GetConfiguration()
	// Set client options
	clientOptions := options.Client().ApplyURI(config.ConnectionString)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	collection := client.Database("go_rest_api").Collection("books")

	return collection
}

type ErrorResponse struct {
	StatusCode   int    `json:"status"`
	ErrorMessage string `json:"message"`
}

func GetError(err error, w http.ResponseWriter) {

	log.Fatal(err.Error())
	var response = ErrorResponse{
		ErrorMessage: err.Error(),
		StatusCode:   http.StatusInternalServerError,
	}

	message, _ := json.Marshal(response)

	w.WriteHeader(response.StatusCode)
	w.Write(message)
}

// Configuration model
type Configuration struct {
	Port             string
	ConnectionString string
}

func GetConfiguration() Configuration {
	err := godotenv.Load("./.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Get the PORT and CONNECTION_STRING environment variables
	port := os.Getenv("PORT")
	connectionString := os.Getenv("CONNECTION_STRING") + "/?retryWrites=true&w=majority&appName=dev"

	if port == "" {
		log.Fatal("PORT must be set in the .env file")
	}

	if connectionString == "" {
		log.Fatal("CONNECTION_STRING must be set in the .env file")
	}

	// For demonstration, just print the connection string
	fmt.Println("Connection String:", connectionString)

	configuration := Configuration{
		port,
		connectionString,
	}

	return configuration
}
