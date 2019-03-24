package main

import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

type Task struct {
    ID        string   `json:"id,omitempty"`
    Description string   `json:"description,omitempty"`
    Duedate  string   `json:"duedate,omitempty"`
    Class   string `json:"class,omitempty"`
}

var tasks []Task

func GetTasksEndpoint(w http.ResponseWriter, req *http.Request) {
    json.NewEncoder(w).Encode(tasks)
}

func CreateTaskEndpoint(w http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    var task Task
    _ = json.NewDecoder(req.Body).Decode(&task)
    task.ID = params["id"]
    tasks = append(tasks, task)
    json.NewEncoder(w).Encode(tasks)
}

func DeleteTaskEndpoint(w http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    for index, item := range tasks {
        if item.ID == params["id"] {
            tasks = append(tasks[:index], tasks[index+1:]...)
            break
        }
    }
    json.NewEncoder(w).Encode(tasks)
}

func main() {
    router := mux.NewRouter()
    tasks = append(tasks, Task{ID: "1", Description: "Homework #3", Duedate: "03/26/19", Class: "Math"})
    tasks = append(tasks, Task{ID: "2", Description: "Homework #4", Duedate: "04/05/19", Class: "Physics"})
    router.HandleFunc("/tasks", GetTasksEndpoint).Methods("GET")
    router.HandleFunc("/tasks/{id}", CreateTaskEndpoint).Methods("POST")
    router.HandleFunc("/tasks/{id}", DeleteTaskEndpoint).Methods("DELETE")
    log.Fatal(http.ListenAndServe(":12345", router))
}