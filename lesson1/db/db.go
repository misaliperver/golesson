package db

import (
	"context"
	"log"
	"time"

	"github.com/misaliperver/golesson/lesson1/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.TODO()

//Client instance
var DB mongo.Client

func ConnectDB() *mongo.Client {
	env, errEnv := config.Get()

	if errEnv != nil {
		panic(errEnv)
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(env.Db.Uri))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	//ping the database
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Print("Connected to MongoDB")

	DB = *client

	return &DB
}

//getting database collections
func GetCollection(collectionName string) *mongo.Collection {
	println("Getting collection: " + collectionName)

	collection := DB.Database("golesson").Collection(collectionName)
	return collection
}
