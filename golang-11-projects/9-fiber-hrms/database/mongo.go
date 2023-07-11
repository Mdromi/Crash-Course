package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

var mg MongoInstance

const (
	dbName   = "fiber-hrms"
	mongoURI = "mongodb://localhost:27017/" + dbName
)

func GetMongoInstance() *MongoInstance {
	return &mg
}

func Connect() error {
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return err
	}

	db := client.Database(dbName)
	mg = MongoInstance{
		Client: client,
		Db:     db,
	}
	return nil
}

func init() {
	if err := Connect(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")
}
