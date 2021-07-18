package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Let's go!")

	// get all posts
	getAllPosts()

	// get <n> posts in parallel
	getPostsInParallel(5)

	// get <n> posts and write in file - utilizing io
	getPostsWritePostsIO(5)

	// get all posts of user with id n
	// get all comments for each post and write them in db in parallel
	getPostsWritePostsDB(7)

	// get all posts of user with id n
	// get all comments for each post and write them in db in parallel - utilizing channels
	getPostsWritePostsChan(7)
}
