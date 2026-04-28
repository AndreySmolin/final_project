package db

import (
	"fmt"
)

type Task struct {
	ID      string `json:"id"`
	Date    string `json:"date"`
	Title   string `json:"title"`
	Comment string `json:"comment,omitempty"`
	Repeat  string `json:"repeat"`
}

func AddTask(task *Task) (int64, error) {
	var id int64
	query := `INSERT INTO scheduler (date,title,comment,repeat) VALUES ($1, $2, $3, $4)`
	res, err := db.Exec(query, task.Date, task.Title, task.Comment, task.Repeat)
	if err == nil {
		id, err = res.LastInsertId()
	}
	return id, err
}
func Tasks(limit int) ([]*Task, error) {
	rows, err := db.Query(`SELECT * FROM scheduler`)
	if err != nil {
		return nil, fmt.Errorf("error in SELECT statement:%w", err)
	}
	defer rows.Close()
	sliceTask := make([]*Task, 0, limit)
	for rows.Next() {
		task := Task{}
		err = rows.Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat)
		if err != nil {
			return nil, fmt.Errorf("column copy error:%w", err)
		}
		sliceTask = append(sliceTask, &task)
	}
	return sliceTask, nil
}
