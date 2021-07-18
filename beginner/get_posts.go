package main

import (
	"log"
)

// Get all posts from json placeholder
func getAllPosts() {
	// setup client
	cl := HTTPClient{}.Setup()

	// get posts
	posts := []Post{}
	err := cl.Get("https://jsonplaceholder.typicode.com/posts", &posts)
	if err != nil {
		log.Fatal("failed to get posts:", err)
	}
}
