package firebase

import (
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

var App *firebase.App

func InitFirebase() {
	// Provide the path to the downloaded service account key file
	opt := option.WithCredentialsFile("../config/svc.json")

	// Initialize the Firebase app
	var err error
	App, err = firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing Firebase app: %v\n", err)
	}
	fmt.Print("Firebase connected successfully")
}
