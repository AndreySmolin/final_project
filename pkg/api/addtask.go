package api

import (
	"encoding/json"
	"final_project/pkg/db"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

// ErrorMessage структура для отправки клиенту ошибки
type ErrorMessage struct {
	Error string `json:"error"`
}

// IDMessage структура для отправки клиенту id задачи
type IDMessage struct {
	Id string `json:"id"`
}

// addTaskHandler функция для добавления задачи
func addTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task db.Task
	var errorMessage ErrorMessage
	var idMessage IDMessage
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "error reading json"+err.Error(), http.StatusBadRequest)
		return
	}
	if err = json.Unmarshal(body, &task); err != nil {
		errorMessage.Error = "error deserializing JSON:" + err.Error()
		writeJson(w, errorMessage, http.StatusBadRequest)
		return
	}
	if task.Title == "" {
		errorMessage.Error = "Task title is missing"
		writeJson(w, errorMessage, http.StatusBadRequest)
		return
	}
	if err = checkDate(&task); err != nil {
		errorMessage.Error = err.Error()
		writeJson(w, errorMessage, http.StatusBadRequest)
		return
	}
	id, err := db.AddTask(&task)
	if err != nil {
		http.Error(w, "error Addtask"+err.Error(), http.StatusBadRequest)
		return
	}
	idMessage.Id = strconv.FormatInt(id, 10)
	writeJson(w, idMessage, http.StatusOK)
}

// writeJson функция для сериализации и отправки JSON клиенту
func writeJson(w http.ResponseWriter, data any, code int) {
	resp, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(code)
	_, err = w.Write(resp)
	if err != nil {
		log.Println("write failed:", err)
	}
}

// checkDate проверяет на корректность полученное значение task.Date
func checkDate(task *db.Task) error {
	now := time.Now()
	if task.Date == "" {
		task.Date = now.Format(FormatDate)
	}
	t, err := time.Parse(FormatDate, task.Date)
	if err != nil {
		return fmt.Errorf("error format:%w", err)
	}
	var next string
	if len(task.Repeat) > 0 {
		next, err = NextDate(now, task.Date, task.Repeat)
		if err != nil {
			return fmt.Errorf("the repetition rule is formatted incorrectly:%w", err)
		}
	}
	if afterNow(now, t) {
		if len(task.Repeat) == 0 {
			task.Date = now.Format(FormatDate)
		} else {
			task.Date = next
		}
	}
	return nil
}
