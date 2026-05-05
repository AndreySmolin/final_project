package api

import (
	"final_project/pkg/db"
	"net/http"
	"time"
)

// postDoneHandler функция отмечает задачу выполненной и если есть правило то дублирует с новой датой
func postDoneHandler(w http.ResponseWriter, r *http.Request) {
	var errorMessage ErrorMessage
	var message struct{}
	id := r.URL.Query().Get("id")
	task, err := db.GetTask(id)
	if err != nil {
		errorMessage.Error = "error retrieving task:" + err.Error()
		writeJson(w, errorMessage, http.StatusBadRequest)
		return
	}
	date, err := NextDate(time.Now(), task.Date, task.Repeat)
	if date == "DELETE" {
		err = db.DeleteTask(id)
		if err != nil {
			errorMessage.Error = "error delete :" + err.Error()
			writeJson(w, errorMessage, http.StatusBadRequest)
			return
		}
		writeJson(w, message, http.StatusOK)
		return
	}
	err = db.UpdateDate(date, id)
	if err != nil {
		errorMessage.Error = err.Error()
		writeJson(w, errorMessage, http.StatusBadRequest)
		return
	}
	writeJson(w, message, http.StatusOK)
}
