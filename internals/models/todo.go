// package models

// import (
// 	"context"
// 	"log"
// 	"time"

// 	"go.mongodb.org/mongo-driver/bson/primitive"
// 	"go.mongodb.org/mongo-driver/mongo"
// )

// type Todo struct {
// 	ID      primitive.ObjectID `bson:"_id,omitempty" json:"id"`
// 	Title   string             `json:"title"`
// 	Content string             `json:"content"`
// }

// func (t *Todo) CreateTodo(collection *mongo.Collection) error {
// 	t.ID = primitive.NewObjectID() // Generate a new ID for the todo item

// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	_, err := collection.InsertOne(ctx, t)
// 	if err != nil {
// 		log.Printf("Failed to insert todo: %v", err)
// 		return err
// 	}

// 	log.Println("Todo created successfully!")
// 	return nil
// }

package models

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Todo struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title   string             `json:"title"`
	Content string             `json:"content"`
}

// CreateTodo inserts the Todo into the MongoDB collection and returns an error if any.
func (t *Todo) CreateTodo(collection *mongo.Collection) error {
	t.ID = primitive.NewObjectID()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, t)
	if err != nil {
		log.Printf("Failed to insert todo: %v", err)
		return err
	}

	log.Println("Todo created successfully!")
	return nil
}

func GetAllTodos(collection *mongo.Collection) ([]Todo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var todos []Todo
	if err := cursor.All(ctx, &todos); err != nil {
		return nil, err
	}

	return todos, nil
}
