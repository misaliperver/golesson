package db

import (
	"context"
	"log"
	"sync"

	"github.com/misaliperver/golesson/lesson2/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/* Used to create a singleton object of MongoDB client.
Initialized and exposed through  GetMongoClient().*/
var clientInstance *mongo.Client

// Used during creation of singleton client object in GetMongoClient().
var clientInstanceError error

// Used to execute client creation procedure only once.
var mongoOnce sync.Once

var (
	URI  string
	NAME string
)

func initClient() {
	env, err := config.Get()

	URI = env.Db.Uri
	NAME = env.Db.Name

	// Set client options
	clientOptions := options.Client().ApplyURI(URI)
	if err != nil {
		panic(err)
	}

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		clientInstanceError = err
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		clientInstanceError = err
	}

	clientInstance = client
}

// GetMongoClient - Return mongodb connection to work with
func GetMongoClient() *mongo.Client {
	// Perform connection creation operation only once.
	// mongoOnce.Do(initClient)

	return clientInstance
}

//getting database collections
func GetCollection(collectionName string) *mongo.Collection {
	log.Println("Getting collection: " + collectionName)
	collection := clientInstance.Database(NAME).Collection(collectionName)
	return collection
}

func init() {
	log.Println("[db/db.go] init()")
	// GetMongoClient()
	initClient()
}
