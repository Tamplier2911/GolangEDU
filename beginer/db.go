package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type DataBase struct {
	DB *sql.DB
}

func (d DataBase) Setup() DataBase {
	db, err := sql.Open("mysql",
		"root:@tcp(127.0.0.1:3306)/test_db")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	return d
}

func (d *DataBase) Close() {
	d.DB.Close()
}
