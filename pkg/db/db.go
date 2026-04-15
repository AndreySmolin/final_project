package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "modernc.org/sqlite"
)

var db *sql.DB

const chema string = `CREATE TABLE scheduler (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	date CHAR(8) NOT NULL DEFAULT "",
	title VARCHAR(256) NOT NULL DEFAULT "",
	comment TEXT NOT NULL DEFAULT "",
	repeat VARCHAR(128) NOT NULL DEFAULT ""
);
CREATE INDEX IDX_DATE ON scheduler(date);`

func Init(dbNAME string) error {
	_, err := os.Stat(dbNAME)
	var install bool
	if err != nil {
		install = true
	}
	if install {
		_, err := os.Create(dbNAME)
		if err != nil {
			return err
		}
		db, err = sql.Open("sqlite", dbNAME)
		if err != nil {
			return err
		}
		defer db.Close()
		_, err = CreateTable(db)
		if err != nil {
			return err
		}
		return nil
	}
	db, err = sql.Open("sqlite", dbNAME)
	if err != nil {
		return err
	}
	defer db.Close()
	return nil
}

func CreateTable(db *sql.DB) (sql.Result, error) {
	result, err := db.Exec(chema)
	if err != nil {
		return nil, fmt.Errorf("error creating the table: %w", err)
	}
	return result, nil
}
