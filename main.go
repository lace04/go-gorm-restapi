package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lace04/go-gorm-restapi/db"
	"github.com/lace04/go-gorm-restapi/models"
	"github.com/lace04/go-gorm-restapi/routes"
)

func main() {
	db.DBConnection()
	db.DB.AutoMigrate(models.Task{})
	db.DB.AutoMigrate(models.User{})
	r := mux.NewRouter()

	r.HandleFunc("/", routes.HomeHandler)

	// Users
	r.HandleFunc("/users", routes.GetUsersRequest).Methods("GET")
	r.HandleFunc("/users/{id}", routes.GetUserRequest).Methods("GET")
	r.HandleFunc("/users", routes.CreateUserRequest).Methods("POST")
	r.HandleFunc("/users/{id}", routes.UpdateUserRequest).Methods("PATCH")
	r.HandleFunc("/users/{id}", routes.DeleteUserRequest).Methods("DELETE")

	// Tasks
	r.HandleFunc("/tasks", routes.GetTasksRequest).Methods("GET")
	r.HandleFunc("/tasks/{id}", routes.GetTaskRequest).Methods("GET")
	r.HandleFunc("/tasks", routes.CreateTaskRequest).Methods("POST")
	r.HandleFunc("/tasks/{id}", routes.UpdateTaskRequest).Methods("PATCH")
	r.HandleFunc("/tasks/{id}", routes.DeleteTaskRequest).Methods("DELETE")

	http.ListenAndServe(":3000", r)
}
