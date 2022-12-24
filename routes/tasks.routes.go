package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nawesan12/go-api/db"
	"github.com/nawesan12/go-api/models"
)

func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Task

	db.DB.Find(&tasks)
	json.NewEncoder(w).Encode(&tasks)
}

func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	json.NewDecoder(r.Body).Decode(&task)

	createdTask := db.DB.Create(&task)
	err := createdTask.Error

	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}

	json.NewEncoder(w).Encode(&task)
}

func GetUniqueTaskHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var task models.Task

	db.DB.First(&task, params["id"])

	if task.ID == 0 {
		json.NewEncoder(w).Encode("Task not found")
		return
	}

	json.NewEncoder(w).Encode(&task)
}

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var task models.Task

	db.DB.First(&task, params["id"])

	if task.ID == 0 {
		json.NewEncoder(w).Encode("Task not found")
		return
	}

	db.DB.Delete(&task)
	json.NewEncoder(w).Encode("Task deleted successfully")
}
