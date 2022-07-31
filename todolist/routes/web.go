package routes

import (
	"github.com/gorilla/mux"
	"github.com/mezz-ir/golang-mini-project/todolist/controllers"
)

func Init() *mux.Router {
	route := mux.NewRouter()

	route.HandleFunc("/", controllers.Show)
	route.HandleFunc("/add", controllers.Add).Methods("POST")
	route.HandleFunc("/delete/{id}", controllers.Delete)
	route.HandleFunc("/complete/{id}", controllers.Complete)

	return route
}
