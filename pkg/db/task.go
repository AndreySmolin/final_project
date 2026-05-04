package db

import (
	"database/sql"
	"errors"
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

func GetTask(id string) (*Task, error) {
	task := Task{}
	query := `SELECT date, title, comment, repeat FROM scheduler WHERE id = :id`
	err := db.QueryRow(query, sql.Named("id", id)).Scan(&task.Date, &task.Title, &task.Comment, &task.Repeat)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("The user was not found")
		} else {
			return nil, fmt.Errorf("error reading the table:%w", err)
		}
	}
	task.ID = id
	return &task, nil
}

func UpdateTask(task *Task) error {
	query := `UPDATE scheduler SET date=$1, title=$2, comment=$3, repeat=$4 WHERE id =$5`
	res, err := db.Exec(query, task.Date, task.Title, task.Comment, task.Repeat, task.ID)
	if err != nil {
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf(`incorrect id for updating task`)
	}
	return nil
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
