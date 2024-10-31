package routes

import (
	"github.com/gorilla/mux"
	"github.com/jitendravee/golang/internals/controllers"
)

var RegisterTodoRoutes = func(router *mux.Router) {
	router.HandleFunc("/todo", controllers.CreateTodoHandler).Methods("POST")
	router.HandleFunc("/todo", controllers.GetAllTodoHandler).Methods("GET")
	router.HandleFunc("/todo/{id}", controllers.UpdateTodoByIdHandler).Methods("PUT")
}
