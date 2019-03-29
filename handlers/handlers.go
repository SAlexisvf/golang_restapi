package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"go_restapi/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetAllTasksEndpoint gets all the documents(tasks) on the DB
func GetAllTasksEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Access-Control-Allow-Origin", "*")
	response.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	response.Header().Set("content-type", "application/json")
	var tasks []models.Task
	collection := client.Database("tasksdb").Collection("tasks")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		fmt.Println("Get endpoint error")
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var task models.Task
		cursor.Decode(&task)
		tasks = append(tasks, task)
	}
	if err := cursor.Err(); err != nil {
		fmt.Println("Get endpoint error (cursor)")
	}
	json.NewEncoder(response).Encode(tasks)
}

// AddTaskEndpoint create a document(task) on the DB
func AddTaskEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Access-Control-Allow-Origin", "*")
	response.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	response.Header().Set("content-type", "application/json")
	var task models.Task
	_ = json.NewDecoder(request.Body).Decode(&task)
	if task != (models.Task{}) {
		collection := client.Database("tasksdb").Collection("tasks")
		if err != nil {
			fmt.Println("Create endpoint error")
		}
		result, _ := collection.InsertOne(ctx, task)
		json.NewEncoder(response).Encode(result)
	}
}

// DeleteTaskEndpoint delete a document(task) on the DB
func DeleteTaskEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Access-Control-Allow-Origin", "*")
	response.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	response.Header().Set("content-type", "application/json")
	var task models.Task
	_ = json.NewDecoder(request.Body).Decode(&task)
	if task != (models.Task{}) {
		collection := client.Database("tasksdb").Collection("tasks")
		if err != nil {
			fmt.Println("Delete endpoint error")
		}
		result, _ := collection.DeleteOne(ctx, task)
		json.NewEncoder(response).Encode(result)
	}
}
