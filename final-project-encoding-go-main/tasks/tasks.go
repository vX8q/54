package tasks

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

func GetTasks() map[string]Task {
	return tasks
}

func AddTask(newTask Task) {
	tasks[newTask.ID] = newTask
}

func GetTaskByID(id string) (Task, bool) {
	task, exists := tasks[id]
	return task, exists
}

func DeleteTask(id string) {
	delete(tasks, id)
}
