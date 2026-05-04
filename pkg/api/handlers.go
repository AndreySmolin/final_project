package api

import (
	"io"
	"net/http"
	"time"
)

func NextDayHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	stringDateNow := r.FormValue("now")
	if stringDateNow == "" {
		stringDateNow = time.Now().Format("20060102")
	}
	now, err := time.Parse("20060102", stringDateNow)
	if err != nil {
		http.Error(w, "form  parsing Error", http.StatusInternalServerError)
		return
	}
	nextd, err := NextDate(now, r.FormValue("date"), r.FormValue("repeat"))
	if err != nil {
		http.Error(w, "error of the NextDate function:"+err.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(w, nextd)
}

func taskHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodPost:
		addTaskHandler(w, r)
	case http.MethodGet:
		getTaskHandler(w, r)
	case http.MethodPut:
		putTaskHandler(w, r)
	case http.MethodDelete:
		deleteTaskHandler(w, r)
	}
}
func tasksHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	getTasksHandler(w, r)
}
func doneHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	postDoneHandler(w, r)
}
