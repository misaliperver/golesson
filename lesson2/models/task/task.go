package task

import (
	"log"
	"time"

	"github.com/misaliperver/golesson/lesson2/db"
	"github.com/misaliperver/golesson/lesson2/models/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Task - struct to map with mongodb documents
type Task struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at"`
	Title       string             `json:"title" bson:"title"`
	Code        string             `json:"code" bson:"code"`
	Description string             `json:"description" bson:"description"`
	Completed   bool               `json:"completed" bson:"completed"`
}

var collection *mongo.Collection
var taskRepository *repository.MongoRepository

func init() {
	log.Println("[models/task.go] init()")
	collection = db.GetCollection("task")
	taskRepository = repository.NewMongoRepository("task", collection)
}

func Create(task Task) error {
	err := taskRepository.CreateOne(task)

	if err != nil {
		return err
	}

	return nil
}

func FindOne(_id string) (*Task, error) {
	var result Task

	bsonResult, err := taskRepository.FindOne(_id)

	bsonBytes, _ := bson.Marshal(bsonResult)

	bson.Unmarshal(bsonBytes, &result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func FindAll() ([]Task, error) {
	var result []Task

	bsonResult, err := taskRepository.FindAll()

	for _, s := range bsonResult {
		var iterator Task
		bsonBytes, _ := bson.Marshal(s)
		bson.Unmarshal(bsonBytes, &result)
		result = append(result, iterator)
	}

	if err != nil {
		return nil, err
	}

	return result, nil
}
