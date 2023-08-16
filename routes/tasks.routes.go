package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lace04/go-gorm-restapi/db"
	"github.com/lace04/go-gorm-restapi/models"
)

func GetTasksRequest(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Task
	db.DB.Find(&tasks)
	json.NewEncoder(w).Encode(&tasks)
}

func GetTaskRequest(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	params := mux.Vars(r)
	db.DB.First(&task, params["id"])
	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Task not found"))
		return
	} else {
		json.NewEncoder(w).Encode(&task)
	}
}

func CreateTaskRequest(w http.ResponseWriter, r *http.Request) {
	//solo debe crear si el usuario existe en su propiedad user_id
	var task models.Task
	json.NewDecoder(r.Body).Decode(&task)
	newTask := db.DB.Create(&task)
	err := newTask.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	} else {
		json.NewEncoder(w).Encode(&task)
	}
}

func UpdateTaskRequest(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	params := mux.Vars(r)
	db.DB.First(&task, params["id"])
	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Task not found"))
		return
	} else {
		json.NewDecoder(r.Body).Decode(&task)
		db.DB.Save(&task)
		json.NewEncoder(w).Encode(&task)
	}
}

func DeleteTaskRequest(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	params := mux.Vars(r)
	db.DB.First(&task, params["id"])
	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Task not found"))
		return
	} else {
		db.DB.Unscoped().Delete(&task)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Task deleted"))
	}
}
