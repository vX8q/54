package main

import (
	"fmt"
	"net/http"

	"github.com/Yandex-Practicum/final-project-encoding-go/handlers"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	r.Get("/tasks", handlers.GetAllTasksHandler)
	r.Post("/tasks", handlers.AddTaskHandler)
	r.Get("/tasks/{id}", handlers.GetTaskByIDHandler)
	r.Delete("/tasks/{id}", handlers.DeleteTaskHandler)

	fmt.Println("Сервер запущен на порту 8080...")
	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Printf("Ошибка при запуске сервера: %s", err.Error())
		return
	}
}
