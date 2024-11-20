package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Yandex-Practicum/final-project-encoding-go/tasks"
	"github.com/go-chi/chi/v5"
)

func GetAllTasksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(tasks.GetTasks()); err != nil {
		http.Error(w, "Ошибка при обработке данных", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func AddTaskHandler(w http.ResponseWriter, r *http.Request) {
	var newTask tasks.Task
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewDecoder(r.Body).Decode(&newTask); err != nil {
		http.Error(w, "Некорректный формат запроса", http.StatusBadRequest)
		return
	}

	if newTask.ID == "" || newTask.Description == "" {
		http.Error(w, "Необходимо указать ID и описание задачи", http.StatusBadRequest)
		return
	}

	tasks.AddTask(newTask)
	w.WriteHeader(http.StatusCreated)
}

func GetTaskByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	w.Header().Set("Content-Type", "application/json")

	task, exists := tasks.GetTaskByID(id)
	if !exists {
		http.Error(w, "Задача не найдена", http.StatusNotFound)
		return
	}

	if err := json.NewEncoder(w).Encode(task); err != nil {
		http.Error(w, "Ошибка при обработке данных", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	w.Header().Set("Content-Type", "application/json")

	if _, exists := tasks.GetTaskByID(id); !exists {
		http.Error(w, "Задача с таким ID не найдена.", http.StatusNotFound)
		return
	}

	tasks.DeleteTask(id)

	w.WriteHeader(http.StatusOK)
}
