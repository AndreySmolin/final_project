package api

import (
	"final_project/pkg/db"
	"net/http"
)

// deleteTaskHandler функция удаляет задачу по id
func deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	var errorMessage ErrorMessage
	var message struct{}
	id := r.URL.Query().Get("id")
	err := db.DeleteTask(id)
	if err != nil {
		errorMessage.Error = "error delete :" + err.Error()
		writeJson(w, errorMessage, http.StatusBadRequest)
		return
	}
	writeJson(w, message, http.StatusOK)
}
