package database

import (
	"context"
	"log"
	"os"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

var (
	FirebaseApp *firebase.App
	Ctx         context.Context
	Client      *db.Client
)

func InitFirebase() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error: .env not found, proceeding with environment variables")
	}

	credJSON := os.Getenv("FIREBASE_CREDENTIALS")
	dbURL := os.Getenv("FIREBASE_DB_URL")

	if credJSON == "" || dbURL == "" {
		log.Fatal("Error: FIREBASE_CREDENTIALS or FIREBASE_DB_URL not set in environment variables")
	}

	opt := option.WithCredentialsJSON([]byte(credJSON))
	Ctx = context.Background()

	var err error
	FirebaseApp, err = firebase.NewApp(Ctx, &firebase.Config{
		DatabaseURL: dbURL,
	}, opt)
	if err != nil {
		log.Fatalf("Error initializing Firebase: %v", err)
	}

	Client, err = FirebaseApp.Database(Ctx)
	if err != nil {
		log.Fatalf("Error initializing Realtime Database client: %v", err)
	}

	log.Println("Firebase initialized successfully!")
}
