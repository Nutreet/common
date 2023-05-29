package common

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
	auth "firebase.google.com/go/v4/auth"
	storage "firebase.google.com/go/v4/storage"
)

type FirebaseConstant struct {
	Firebase  *firebase.App
	Firestore *firestore.Client
	Storage   *storage.Client
	Auth      *auth.Client
}

var instance = FirebaseConstant{}

func GetApp() *firebase.App {
	var config = &firebase.Config{ProjectID: Constants.FIREBASE_PROJECT_ID}

	app, err := firebase.NewApp(context.Background(), config)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	return app
}

func init() {
	log.Println("Initializing Firebase")

	app := GetApp()

	var err error

	instance.Firebase = app
	instance.Firestore, err = app.Firestore(context.Background())

	if err != nil {
		panic(err)
	}

	instance.Storage, err = app.Storage(context.Background())

	if err != nil {
		panic(err)
	}

	instance.Auth, err = app.Auth(context.Background())

	if err != nil {
		panic(err)
	}

}
