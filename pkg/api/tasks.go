package api

import (
	"final_project/pkg/db"
	"net/http"
)

// TasksResp структура для вывода задач списком
type TasksResp struct {
	Tasks []*db.Task `json:"tasks"`
}

// getTasksHandler функция получает задачу и заносит в список TasksResp с последующим отправление клиенту
func getTasksHandler(w http.ResponseWriter, r *http.Request) {
	var errorMessage ErrorMessage
	tasks, err := db.Tasks(50)
	if err != nil {
		errorMessage.Error = "error fetching task list:" + err.Error()
		writeJson(w, errorMessage)
		return
	}
	writeJson(w, TasksResp{
		Tasks: tasks,
	})
}
