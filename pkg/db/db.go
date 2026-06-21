package db

import (
	"fmt"
    "database/sql"
	"os"

    _ "modernc.org/sqlite"
)

var db *sql.DB

var schema = (`
		CREATE TABLE scheduler (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			date CHAR(8) NOT NULL DEFAULT "",
			comment TEXT,
			title VARCHAR,
			repeat VARCHAR
		)
	`)

func Init(dbFile string) error {

	fmt.Println("Проверяю наличие файла")
	_, err := os.Stat(dbFile)

	install := os.IsNotExist(err)

	fmt.Println("Открываю файл")

	db, err = sql.Open("sqlite", dbFile)
	if err != nil {
		panic(err)
	}

	if install {
			fmt.Println("Создаю базу")
			_, err = db.Exec(schema)
			if err != nil {
				return err
			}
	}
	return nil
}