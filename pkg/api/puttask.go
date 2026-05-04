package api

import (
	"encoding/json"
	"final_project/pkg/db"
	"io"
	"net/http"
)

func putTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task db.Task
	var errorMessage ErrorMessage
	var message struct{}
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, "error reading json"+err.Error(), http.StatusBadRequest)
		return
	}
	if err = json.Unmarshal(body, &task); err != nil {
		errorMessage.Error = "error deserializing JSON:" + err.Error()
		writeJson(w, errorMessage)
		return
	}
	if task.Title == "" {
		errorMessage.Error = "Task title is missing"
		writeJson(w, errorMessage)
		return
	}
	if err = checkDate(&task); err != nil {
		errorMessage.Error = err.Error()
		writeJson(w, errorMessage)
		return
	}
	err = db.UpdateTask(&task)
	if err != nil {
		errorMessage.Error = err.Error()
		writeJson(w, errorMessage)
		return
	}
	writeJson(w, message)
}
