package models

import (
	"context"
	"log"
	"time"

	"github.com/misaliperver/golesson/lesson1/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Todo struct {
	Id     primitive.ObjectID `json:"id,omitempty"`
	Title  string             `json:"title,omitempty" validate:"required"`
	Status bool               `json:"status,omitempty" validate:"required"`
}

var todoCollection *mongo.Collection = db.GetCollection("todos")

func Create(title string, status bool) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	newTodo := Todo{
		Title:  title,
		Status: status,
	}

	result, err := todoCollection.InsertOne(ctx, newTodo)

	if err != nil {
		log.Fatal("[createTodo]", err)
		return nil, err
	}

	return result, nil
}

func Delete(_id string) (*mongo.DeleteResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	idPrimitive, errIdPrimitive := primitive.ObjectIDFromHex(_id)

	if errIdPrimitive != nil {
		log.Fatal("[deleteTodo]", errIdPrimitive)
		return nil, errIdPrimitive
	} else {
		result, errDeleteOne := todoCollection.DeleteOne(ctx, bson.M{"_id": idPrimitive})

		if errDeleteOne != nil {
			log.Fatal("[deleteTodo]", errDeleteOne)
			return nil, errDeleteOne
		}

		return result, nil
	}
}

func Find(_id string) (*bson.M, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	idPrimitive, errIdPrimitive := primitive.ObjectIDFromHex(_id)

	if errIdPrimitive != nil {
		log.Fatal("[findTodo]", errIdPrimitive)
		return nil, errIdPrimitive
	} else {
		var result bson.M

		errFindOne := todoCollection.FindOne(ctx, bson.M{"_id": idPrimitive}).Decode(&result)

		if errFindOne != nil {
			log.Fatal("[findTodo]", errFindOne)
			return nil, errFindOne
		}

		return &result, nil
	}
}

func Update(_id string, data *Todo) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	idPrimitive, errIdPrimitive := primitive.ObjectIDFromHex(_id)

	if errIdPrimitive != nil {
		log.Fatal("[findTodo]", errIdPrimitive)
		return nil, errIdPrimitive
	} else {
		filter := bson.M{"_id": idPrimitive}
		updateData := bson.M{"title": data.Title, "status": data.Status}

		result, errUpdateOne := todoCollection.UpdateOne(
			ctx,
			filter,
			updateData,
		)

		if errUpdateOne != nil {
			log.Fatal("[updateOneTodo]", errUpdateOne)
			return nil, errUpdateOne
		}

		return result, nil
	}
}
