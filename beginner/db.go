package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type MySQL struct {
	DB *sql.DB
}

// Setup database, save db instance in state, migrate models
func (msql MySQL) Setup() (MySQL, error) {
	// wire up db
	db, err := sql.Open("mysql",
		"root:@tcp(127.0.0.1:3306)/test_db")
	if err != nil {
		return msql, err
	}

	// store instance of db in state
	msql.DB = db

	// create database if not exists
	err = msql.CreateDB("test_db")
	if err != nil {
		return msql, err
	}
	log.Println("successfully created database")

	// check conection
	err = msql.Pong()
	if err != nil {
		return msql, err
	}
	log.Println("successfully connected to database")

	// set models
	models := map[string]interface{}{
		"posts":    &Post{},
		"comments": &Comment{},
	}

	// migrate models naively
	err = msql.MigrateModels(models)
	if err != nil {
		return msql, err
	}
	log.Println("successfully migrated models")

	return msql, nil
}

// Check database connection
func (msql *MySQL) Pong() error {
	err := msql.DB.Ping()
	if err != nil {
		return err
	}
	return nil
}

// Close database connection
func (msql *MySQL) Close() error {
	err := msql.DB.Close()
	if err != nil {
		return err
	}
	return nil
}

// Creates database if not exists
func (msql *MySQL) CreateDB(dbName string) error {
	// create db if not exists
	stmt := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s;", dbName)
	_, err := msql.DB.Exec(stmt)
	if err != nil {
		return err
	}
	return nil
}

// Naive implementation of migration
func (msql *MySQL) MigrateModel(tableName string, i interface{}) error {
	// marshal model
	bs, err := json.Marshal(&i)
	if err != nil {
		return err
	}

	// unmarshal model to map
	var model map[string]interface{}
	err = json.Unmarshal(bs, &model)
	if err != nil {
		return err
	}

	stmt := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s(", tableName)

	for k, v := range model {
		switch v.(type) {
		case float64:
			stmt += fmt.Sprintf("%s INT NOT NULL,", msql.sqlifyString(k))

		case string:
			stmt += fmt.Sprintf("%s VARCHAR(1000) NOT NULL,", msql.sqlifyString(k))
		}
	}

	stmt += `
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP);
	`

	_, err = msql.DB.Exec(stmt)
	if err != nil {
		return err
	}

	return nil
}

func (msql *MySQL) MigrateModels(models map[string]interface{}) error {
	for k, v := range models {
		err := msql.MigrateModel(k, v)
		if err != nil {
			return err
		}
	}
	return nil
}

// Naive implementation of insertion
func (msql *MySQL) Insert(tableName string, i interface{}) error {
	// marshal model
	bs, err := json.Marshal(&i)
	if err != nil {
		return err
	}

	// unmarshal model to map
	var model map[string]interface{}
	err = json.Unmarshal(bs, &model)
	if err != nil {
		return err
	}

	stmt := fmt.Sprintf("INSERT INTO %s(", tableName)

	keys := []string{}
	vals := []string{}

	for k, v := range model {
		keys = append(keys, fmt.Sprintf("%s,", msql.sqlifyString(k)))
		switch v.(type) {
		case float64:
			vals = append(vals, fmt.Sprintf("%.0f,", v))
		case string:
			vals = append(vals, fmt.Sprintf("'%s',", v))
		}
	}

	stmt += fmt.Sprintf("%s%s", strings.TrimSuffix(strings.Join(keys, ""), ","), ") VALUES (")
	stmt += fmt.Sprintf("%s%s", strings.TrimSuffix(strings.Join(vals, ""), ","), ");")

	_, err = msql.DB.Exec(stmt)
	if err != nil {
		return err
	}

	return nil
}

// SQLify string, turns 'userId' into 'user_id' etc
func (msql *MySQL) sqlifyString(s string) string {
	r := ""
	for _, c := range s {
		if c >= 65 && c <= 90 {
			r += fmt.Sprintf("_%s", string(c+32))
			continue
		}
		r += string(c)
	}
	return r
}

// http://go-database-sql.org/modifying.html
