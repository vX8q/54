package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Yandex-Practicum/final-project-encoding-go/tasks"
	"github.com/go-chi/chi/v5"
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	tasks, err := tasks.GetTasks()
	if err != nil {
		http.Error(w, "Ошибка при получении задач: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(tasks); err != nil {
		http.Error(w, "Ошибка при кодировании данных", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func AddTask(w http.ResponseWriter, r *http.Request) {
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

	if err := tasks.AddTask(newTask); err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func GetTask(w http.ResponseWriter, r *http.Request) {
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

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	w.Header().Set("Content-Type", "application/json")

	if err := tasks.DeleteTask(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}
