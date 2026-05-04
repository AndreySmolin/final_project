package api

import (
	"final_project/pkg/db"
	"net/http"
)

func deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	var errorMessage ErrorMessage
	var message struct{}
	id := r.URL.Query().Get("id")
	err := db.DeleteTask(id)
	if err != nil {
		errorMessage.Error = "error delete :" + err.Error()
		writeJson(w, errorMessage)
		return
	}
	writeJson(w, message)
}
