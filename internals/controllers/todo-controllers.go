// package controllers

// import (
// 	"encoding/json"
// 	"net/http"

// 	"github.com/jitendravee/golang/internals/config"
// 	"github.com/jitendravee/golang/internals/models"
// )

// func GetTodo(w http.ResponseWriter, r *http.Request) {

// }
// func init() {
// 	// Initialize the MongoDB connection and get the todo collection
// 	todoCollection = config.Connect() // Assuming Connect() returns the "test" collection
// }
// func CreateTodoHandler(w http.ResponseWriter, r *http.Request) {
// 	var todo models.Todo
// 	err := json.NewDecoder(r.Body).Decode(&todo)
// 	if err != nil {
// 		http.Error(w, "Invalid request payload", http.StatusBadRequest)
// 		return
// 	}

// 	err = todo.CreateTodo(todoCollection)
// 	if err != nil {
// 		http.Error(w, "Failed to create todo", http.StatusInternalServerError)
// 		return
// 	}

//		w.WriteHeader(http.StatusCreated)
//		json.NewEncoder(w).Encode(todo)
//	}
package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/jitendravee/golang/internals/config"
	"github.com/jitendravee/golang/internals/models"
	"go.mongodb.org/mongo-driver/mongo"
)

var todoCollection *mongo.Collection // Define todoCollection as a package-level variable

// Initialize the MongoDB connection and assign the todoCollection
func init() {
	todoCollection = config.Connect()
}

// CreateTodoHandler handles the POST request to create a new todo
func CreateTodoHandler(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = todo.CreateTodo(todoCollection)
	if err != nil {
		http.Error(w, "Failed to create todo", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)
}
