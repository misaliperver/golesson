package question

import (
	"log"
	"time"

	"github.com/misaliperver/golesson/lesson3/db"
	"github.com/misaliperver/golesson/lesson3/lib/logger"
	"github.com/misaliperver/golesson/lesson3/lib/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Answer struct {
	Title string `json:"username" bson:"username"`
	Media string `json:"media" bson:"media"`
}

type Flags struct {
	IsMultiReply bool   `json:"isMultiReply" bson:"isMultiReply"`
	RightAnswer  string `json:"rightAnswer" bson:"rightAnswer"`
}

// Question - struct to map with mongodb documents
type Question struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
	Title     string             `json:"title" bson:"title"`
	Media     string             `json:"media" bson:"media"`
	Tags      []string           `json:"tags" bson:"tags"`
	Answers   []Answer           `json:"answers" bson:"answers"`
	Flags     []Flags            `json:"flags" bson:"flags"`
}

var (
	l                  = logger.NewLogger()
	collection         *mongo.Collection
	questionRepository *repository.MongoRepository
)

func init() {
	log.Println("[models/question.go] init()")
	collection = db.GetCollection("question")
	questionRepository = repository.NewMongoRepository("question", collection)
}

func Create(question Question) error {
	err := questionRepository.CreateOne(question)

	if err != nil {
		return err
	}

	return nil
}

func FindOneById(_id string) (*Question, error) {
	var result Question

	bsonResult, err := questionRepository.FindOneById(_id)

	bsonBytes, _ := bson.Marshal(bsonResult)

	bson.Unmarshal(bsonBytes, &result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func FindAll() ([]Question, error) {
	var result []Question

	bsonResult, err := questionRepository.FindAll()

	for _, s := range bsonResult {
		var iterator Question
		bsonBytes, _ := bson.Marshal(s)
		bson.Unmarshal(bsonBytes, &result)
		result = append(result, iterator)
	}

	if err != nil {
		return nil, err
	}

	return result, nil
}
