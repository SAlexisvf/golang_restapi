package main

import (
	"fmt"
	"go_restapi/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Starting the application...")
	router := mux.NewRouter()
	router.HandleFunc("/task", handlers.AddTaskEndpoint).Methods("POST", "OPTIONS")
	router.HandleFunc("/tasks", handlers.GetAllTasksEndpoint).Methods("GET", "OPTIONS")
	router.HandleFunc("/deleteTask", handlers.DeleteTaskEndpoint).Methods("POST", "OPTIONS")
	log.Fatal(http.ListenAndServe(":3001", router))
}
