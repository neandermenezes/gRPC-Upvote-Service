package database

import (
	"context"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

type MongoDatabase struct {
	Ctx      context.Context
	Database *mongo.Database
}

func NewMongoDatabase() *MongoDatabase {
	mongoDb := &MongoDatabase{}

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		log.Fatal("No mongo uri set in env variables")
	}

	client, err := mongo.Connect(mongoDb.Ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(mongoDb.Ctx, nil)
	if err != nil {
		panic(err)
	}

	mongoDb.Database = client.Database("upvote-db")
	return mongoDb
}
