package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

type DB struct {
	client *mongo.Client
}

func NewDB(connectionString string) (*DB, error) {
	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to MongoDB!")

	MongoClient = client
	return &DB{client: client}, nil
}

func (db *DB) Disconnect() {
	if err := db.client.Disconnect(context.TODO()); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Disconnected from MongoDB.")
}

func (db *DB) GetClient() *mongo.Client {
	return db.client
}

func (db *DB) GetDatabase() *mongo.Database {
	return db.client.Database("me-alerte")
}
