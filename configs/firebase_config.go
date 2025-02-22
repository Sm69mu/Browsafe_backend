package configs

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

var (
	FirestoreClient *firestore.Client
	AuthClient      *auth.Client
)

func InitFirebase() {
	ctx := context.Background()
	opt := option.WithCredentialsFile("firebase.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("error init firebase: %v", err)
	}
	//init firestore db
	FirestoreClient, err = app.Firestore(ctx)
	if err != nil {
		log.Fatalf("error init firestoreClient: %v", err)
	}

	//init auth
	AuthClient, err = app.Auth(ctx)
	if err != nil {
		log.Fatalf("error init auth: %v", err)
	}

	log.Println("Firebase init success")
}
