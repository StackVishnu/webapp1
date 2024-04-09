package dbconn

import (
	"database/sql"
	"fmt"
)

func Dbconnect() *sql.DB {
	db, err := sql.Open("postgres", "postgres://ndbwvocg:zMdN0KsyfHI9SOg40ZmnQfRUhgMm3gbz@isilo.db.elephantsql.com/ndbwvocg")
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to the web app database")
	return db
}
