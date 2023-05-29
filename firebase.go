package common

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
)

func GetApp() *firebase.App {
	var config = &firebase.Config{ProjectID: Constants.FIREBASE_PROJECT_ID}

	app, err := firebase.NewApp(context.Background(), config)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	return app
}

var (
	Firebase           = GetApp()
	Firestore, _       = Firebase.Firestore(context.Background())
	FirebaseStorage, _ = Firebase.Storage(context.Background())
	FirebaseAuth, _    = Firebase.Auth(context.Background())
)
