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
var DB *mongo.Client = ConnectDB()

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
	return client
}

//getting database collections
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("golangAPI").Collection(collectionName)
	return collection
}
