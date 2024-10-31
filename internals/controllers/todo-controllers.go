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

	"github.com/gorilla/mux"
	"github.com/jitendravee/golang/internals/config"
	"github.com/jitendravee/golang/internals/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var todoCollection *mongo.Collection // Define todoCollection as a package-level variable

// Initialize the MongoDB connection and assign the todoCollection
func init() {
	todoCollection = config.Connect()
}

// CreateTodoHandler handles the POST request to create a new todo
func CreateTodoHandler(w http.ResponseWriter, r *http.Request) {
	todo := &models.Todo{}
	err := json.NewDecoder(r.Body).Decode(todo)
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
func GetAllTodoHandler(w http.ResponseWriter, r *http.Request) {
	todos, err := models.GetAllTodos(todoCollection)

	if err != nil {
		http.Error(w, "Failed to fetch todos", http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(todos)

}
func UpdateTodoByIdHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the ID from the request URL (assuming the ID is passed as a URL parameter)
	idParam := mux.Vars(r)["id"] // If using Gorilla Mux
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	// Decode the updated Todo fields from the request body
	todo := &models.Todo{}
	err = json.NewDecoder(r.Body).Decode(todo)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Call the UpdateTodoById method
	updatedTodo, err := todo.UpdateTodoById(todoCollection, id)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			http.Error(w, "Todo not found", http.StatusNotFound)
		} else {
			http.Error(w, "Failed to update todo", http.StatusInternalServerError)
		}
		return
	}

	// If the update was successful, set the response header and encode the updated document as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedTodo)
}
