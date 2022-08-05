package routes

import (
	"encoding/json"
	"net/http"

	"github.com/DuvanR11/api/db"
	"github.com/DuvanR11/api/models"
	"github.com/gorilla/mux"
)

func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Task
	db.DB.Find(&tasks)
	json.NewEncoder(w).Encode(tasks)
	w.WriteHeader(http.StatusOK)
}

func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	params := mux.Vars(r)
	db.DB.First(&task, params["id"])
	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("No task found"))
		return
	}
	json.NewEncoder(w).Encode(task)
	w.WriteHeader(http.StatusOK)
}

func PostTaskHandler(w http.ResponseWriter, r *http.Request) {
	var tasks models.Task
	json.NewDecoder(r.Body).Decode(&tasks)
	create := db.DB.Create(&tasks)
	err := create.Error
	if err != nil {
		w.WriteHeader(http.StatusConflict)
		w.Write([]byte("Task error"))
	}
	json.NewEncoder(w).Encode(&tasks)
}

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	params := mux.Vars(r)
	db.DB.First(&task, params["id"])
	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("No task found"))
		return
	}
	db.DB.Delete(&task)
	w.WriteHeader(http.StatusOK)
}

func UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {

}
