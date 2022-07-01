package user

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

// User - struct to map with mongodb documents
type User struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
	Username  string             `json:"username" bson:"username"`
	Password  string             `json:"password" bson:"password"`
	Email     string             `json:"email" bson:"email"`
}

var (
	l              = logger.NewLogger("user", "models/user.go")
	collection     *mongo.Collection
	userRepository *repository.MongoRepository
)

func init() {
	log.Println("[models/user.go] init()")
	collection = db.GetCollection("user")
	userRepository = repository.NewMongoRepository("user", collection)
}

func Create(user User) error {
	err := userRepository.CreateOne(user)

	if err != nil {
		return err
	}

	return nil
}

func FindOneById(_id string) (*User, error) {
	var result User

	bsonResult, err := userRepository.FindOneById(_id)

	bsonBytes, _ := bson.Marshal(bsonResult)

	bson.Unmarshal(bsonBytes, &result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func FindAll() ([]User, error) {
	var result []User

	bsonResult, err := userRepository.FindAll()

	for _, s := range bsonResult {
		var iterator User
		bsonBytes, _ := bson.Marshal(s)
		bson.Unmarshal(bsonBytes, &result)
		result = append(result, iterator)
	}

	if err != nil {
		return nil, err
	}

	return result, nil
}
