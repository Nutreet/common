package common

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type ConstantsType struct {
	FIREBASE_PROJECT_ID            string
	TEST                           bool
	FIREBASE_AUTH_EMULATOR_HOST    string
	FIREBASE_STORAGE_EMULATOR_HOST string
	FIRESTORE_EMULATOR_HOST        string
	API_GATEWAY_PORT               string
	USER_GRPC_PORT                 string
}

var Constants = ConstantsType{}

func init() {
	log.Println("Loading .env file...")
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("error loading .env file: %v\n", err)
	}

	var (
		FIREBASE_PROJECT_ID            = os.Getenv("FIREBASE_PROJECT_ID")
		FIREBASE_AUTH_EMULATOR_HOST    = os.Getenv("FIREBASE_AUTH_EMULATOR_HOST")
		FIREBASE_STORAGE_EMULATOR_HOST = os.Getenv("FIREBASE_STORAGE_EMULATOR_HOST")
		FIRESTORE_EMULATOR_HOST        = os.Getenv("FIRESTORE_EMULATOR_HOST")
		TEST                           = os.Getenv("GO_ENV") == "test"
		API_GATEWAY_PORT               = os.Getenv("API_GATEWAY_PORT")
		USER_GRPC_PORT                 = os.Getenv("USER_GRPC_PORT")
	)

	Constants = ConstantsType{
		FIREBASE_PROJECT_ID,
		TEST,
		FIREBASE_AUTH_EMULATOR_HOST,
		FIREBASE_STORAGE_EMULATOR_HOST,
		FIRESTORE_EMULATOR_HOST,
		API_GATEWAY_PORT,
		USER_GRPC_PORT,
	}
}
