package config

import (
	"database/sql"
	"log"
)

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/learn_go")
	if err != nil {
		log.Fatal(err)
	}
	
	return db
}
