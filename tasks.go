package tasks

import (
	"fmt"
)

type Task struct {
	ID          string   `json:"id"`          // ID задачи
	Description string   `json:"description"` // Заголовок
	Note        string   `json:"note"`        // Описание задачи
	Application []string `json:"application"` // Приложения, которыми будете пользоваться
}

var tasks = map[string]Task{
	"1": {
		ID:          "1",
		Description: "Сделать финальное задание темы REST API",
		Note:        "Если сегодня сделаю, то завтра будет свободный день. Ура!",
		Application: []string{
			"VS Code",
			"Terminal",
			"git",
		},
	},
	"2": {
		ID:          "2",
		Description: "Протестировать финальное задание с помощью Postman",
		Note:        "Лучше это делать в процессе разработки, каждый раз, когда запускаешь сервер и проверяешь хендлер",
		Application: []string{
			"VS Code",
			"Terminal",
			"git",
			"Postman",
		},
	},
}

func GetTasks() (map[string]Task, error) {
	if tasks == nil {
		return nil, fmt.Errorf("список задач не определён")
	}
	return tasks, nil
}

func AddTask(newTask Task) error {
	if _, exists := tasks[newTask.ID]; exists {
		return fmt.Errorf("задача с ID '%s' уже существует", newTask.ID)
	}
	tasks[newTask.ID] = newTask
	return nil
}

func GetTaskByID(id string) (Task, bool) {
	task, exists := tasks[id]
	return task, exists
}

func DeleteTask(id string) error {
	if _, exists := tasks[id]; !exists {
		return fmt.Errorf("задача с ID %s не найдена", id)
	}
	delete(tasks, id)
	return nil
}
