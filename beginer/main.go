package main

import (
	"database/sql"
	"fmt"
	"log"
)

func main() {
	fmt.Println("Let's go!")

	// get all posts
	// getAllPosts()

	// get n posts in parallel
	// getPostsInParallel(5)

	// get n posts and write in file
	// getPostsWritePostsIO(5)

	// get all posts of user with id 7
	// get all comments for each post and write them in db in paralel
	// getPostsWritePostsDB()

	// db := DataBase{}.Setup()
	// defer db.Close()

	// http://go-database-sql.org/modifying.html

	// setup driver
	db, err := sql.Open("mysql",
		"root:@tcp(127.0.0.1:3306)/test_db")
	if err != nil {
		log.Fatal(err)
	}

	// ping for connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// close db
	defer db.Close()
}
