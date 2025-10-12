package main

import (
	"context"
	"log"
	"os"

	firebase "firebase.google.com/go"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

func InitFirebase() *firebase.App {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
	}

	credPath := os.Getenv("FIREBASE_CREDENTIALS")
	if credPath == "" {
		log.Fatal("Error: Firebase credential path not found in .env")
	}

	opt := option.WithCredentialsFile(credPath)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("Error initializing Firebase: %v", err)
	}
	return app
}
