package main

import (
	"fmt"
)

func main() {
	fmt.Println("Let's go!")

	// get all posts
	getAllPosts()

	// get all posts in parallel
	getAllPostsInParallel()
}
