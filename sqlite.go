package gomc

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

/*
CREATE TABLE user (
    id integer PRIMARY KEY AUTOINCREMENT,
    name varchar(64) NULL
);

INSERT INTO user(name) VALUES('Kiichiro NAKA');
 */

func SqliteMain(_ []string) {
	db, err := sql.Open("sqlite3", "/tmp/foo.db")
	checkErr(err)
	defer func() {
		if db.Close() != nil {
			log.Println("Failed to close DB")
		}
	}()
	rows, err := db.Query("SELECT * FROM user")
	checkErr(err)
	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		checkErr(err)
		fmt.Println(id)
		fmt.Println(name)
	}
}
