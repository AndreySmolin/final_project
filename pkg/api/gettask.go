package api

import (
	"final_project/pkg/db"
	"net/http"
)

// getTaskHandler функция получает задачу по id
func getTaskHandler(w http.ResponseWriter, r *http.Request) {
	var errorMessage ErrorMessage
	id := r.URL.Query().Get("id")
	task, err := db.GetTask(id)
	if err != nil {
		errorMessage.Error = "error retrieving task:" + err.Error()
		writeJson(w, errorMessage, http.StatusBadRequest)
		return
	}
	writeJson(w, task, http.StatusOK)
}
