package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lace04/go-gorm-restapi/db"
	"github.com/lace04/go-gorm-restapi/models"
)

func GetUsersRequest(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	db.DB.Find(&users)
	json.NewEncoder(w).Encode(&users)
}

func GetUserRequest(w http.ResponseWriter, r *http.Request) {
	var user models.User
	params := mux.Vars(r)
	db.DB.First(&user, params["id"])
	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	} else {
		db.DB.Model(&user).Association("Tasks").Find(&user.Tasks)
		json.NewEncoder(w).Encode(&user)
	}
}

func CreateUserRequest(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	newUser := db.DB.Create(&user)
	err := newUser.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	} else {
		json.NewEncoder(w).Encode(&user)
	}
}

func UpdateUserRequest(w http.ResponseWriter, r *http.Request) {
	var user models.User
	params := mux.Vars(r)
	db.DB.First(&user, params["id"])
	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	} else {
		json.NewDecoder(r.Body).Decode(&user)
		db.DB.Save(&user)
		json.NewEncoder(w).Encode(&user)
	}
}

func DeleteUserRequest(w http.ResponseWriter, r *http.Request) {
	var user models.User
	params := mux.Vars(r)
	db.DB.First(&user, params["id"])
	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	} else {
		db.DB.Unscoped().Delete(&user)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("User deleted"))
	}
}
