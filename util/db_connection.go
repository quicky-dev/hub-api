package util

import (
	"context"
	"log"
	"os"
)

func obtainDatabaseClient() *mongo.Client {
	clientOptions := new(options.ClientOptions)

	if os.Getenv("PRODUCTION_MODE") != "" {
		credentials := new(options.Credential)
		credentials.Username = os.Getenv("AUTH_DB_USER")
		credentials.Password = os.Getenv("AUTH_DB_PASSWORD")
		clientOptions.Auth = credentials
	}

	clientOptions.ApplyURI(os.Getenv("MONGODB_URI"))

	err := clientOptions.Validate()

	if err != nil {
		log.Fatal(err)
	}

	client, err := mongo.Connect(context.TODO(), clientOptions)
	log.Println("Successfully connected to mongo.")

	return client
}

var MongoClient = obtainDatabaseClient()
