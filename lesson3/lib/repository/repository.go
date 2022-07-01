package repository

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	CreateOne(data interface{}) error
	DeleteOne(id string) error
	UpdateOne(_id string, data interface{}) error
	FindOne(_id string) (*bson.M, error)
	FindAll(_id string) ([]bson.D, error)
}

type MongoRepository struct {
	name       string
	collection *mongo.Collection
}

var dbRequestTimeout time.Duration = 10 * time.Second

func NewMongoRepository(name string, collection *mongo.Collection) *MongoRepository {
	return &MongoRepository{name, collection}
}

func (r *MongoRepository) CreateOne(data interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbRequestTimeout)

	defer cancel()

	_, err := r.collection.InsertOne(ctx, data)

	if err != nil {
		log.Fatal("[", r.name, "CreateRepo]", err)
		return err
	}

	return nil
}

func (r *MongoRepository) DeleteOneById(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbRequestTimeout)

	defer cancel()

	idPrimitive, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		log.Fatal("[", r.name, "CreateRepo]", err)
		return err
	} else {
		_, err := r.collection.DeleteOne(ctx, bson.M{"_id": idPrimitive})

		if err != nil {
			log.Fatal("[", r.name, "CreateRepo]", err)
			return err
		}

		return err
	}
}

func (r *MongoRepository) DeleteOne(filter *bson.D) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbRequestTimeout)

	defer cancel()

	_, err := r.collection.DeleteOne(ctx, filter)

	if err != nil {
		log.Fatal("[", r.name, "CreateRepo]", err)
		return err
	}

	return err
}

func (r *MongoRepository) UpdateOneById(_id string, data interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbRequestTimeout)

	defer cancel()

	idPrimitive, err := primitive.ObjectIDFromHex(_id)

	if err != nil {
		log.Fatal("[", r.name, "UpdateRepo]", err)
		return err
	} else {
		filter := bson.M{"_id": idPrimitive}

		_, err := r.collection.UpdateOne(
			ctx,
			filter,
			data,
		)

		if err != nil {
			log.Fatal("[", r.name, "UpdateRepo]", err)
			return err
		}

		return nil
	}
}

func (r *MongoRepository) UpdateOne(filter *bson.D, data interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbRequestTimeout)

	defer cancel()

	_, err := r.collection.UpdateOne(
		ctx,
		filter,
		data,
	)

	if err != nil {
		log.Fatal("[", r.name, "UpdateRepo]", err)
		return err
	}

	return nil
}

func (r *MongoRepository) FindOneById(_id string) (*bson.M, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbRequestTimeout)

	defer cancel()

	idPrimitive, err := primitive.ObjectIDFromHex(_id)

	if err != nil {
		log.Fatal("[", r.name, "DeleteRepo]", err)
		return nil, err
	} else {
		var result bson.M
		var filter = bson.M{"_id": idPrimitive}

		err := r.collection.FindOne(ctx, filter).Decode(&result)

		if err != nil {
			return nil, err
		}

		return &result, nil
	}
}

func (r *MongoRepository) FindOne(filter *bson.D) (*bson.M, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbRequestTimeout)

	defer cancel()

	var result bson.M

	err := r.collection.FindOne(ctx, filter).Decode(&result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *MongoRepository) FindAll() ([]bson.D, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbRequestTimeout)

	defer cancel()

	result := []bson.D{}
	var filter = bson.D{}

	cur, err := r.collection.Find(ctx, filter)

	if err != nil {
		return nil, err
	}

	//Map result to slice
	for cur.Next(context.TODO()) {
		t := bson.D{}
		err := cur.Decode(&t)

		if err != nil {
			return result, err
		}

		result = append(result, t)
	}

	// once exhausted, close the cursor
	cur.Close(context.TODO())

	if len(result) == 0 {
		return result, mongo.ErrNoDocuments
	}

	return result, nil
}

func (r *MongoRepository) FindMany(filter *bson.D) ([]bson.D, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbRequestTimeout)

	defer cancel()

	result := []bson.D{}

	if filter == nil {
		filter = &bson.D{}
	}

	cur, err := r.collection.Find(ctx, filter)

	if err != nil {
		return nil, err
	}

	//Map result to slice
	for cur.Next(context.TODO()) {
		t := bson.D{}
		err := cur.Decode(&t)

		if err != nil {
			return result, err
		}

		result = append(result, t)
	}

	// once exhausted, close the cursor
	cur.Close(context.TODO())

	if len(result) == 0 {
		return result, mongo.ErrNoDocuments
	}

	return result, nil
}
