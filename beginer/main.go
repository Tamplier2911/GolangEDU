package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
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
	getPostsWritePostsDB(1)

}
