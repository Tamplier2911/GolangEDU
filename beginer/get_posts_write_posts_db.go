package main

import (
	"fmt"
	"log"
	"sync"
)

func getPostsWritePostsDB() {

	// setup client
	c := HTTPClient{}.Setup()

	// setup db

	// get all posts
	log.Println("getting all posts for user with id of 7")
	posts := []Post{}
	err := c.Get("https://jsonplaceholder.typicode.com/posts?userId=7", &posts)
	if err != nil {
		log.Fatal("failed to get posts:", err)
	}
	log.Println("successfully fetched posts")
	log.Println("getting comments for each post")

	var wg sync.WaitGroup

	// get comments for each post in paralell
	for _, p := range posts {
		wg.Add(1)

		go func(p *Post, wg *sync.WaitGroup) {
			defer wg.Done()

			var comments []Comment
			err = c.Get(fmt.Sprintf("https://jsonplaceholder.typicode.com/comments?postId=%d", p.ID), &comments)
			if err != nil {
				log.Fatal("failed to get comments:", err)
			}
			log.Println("successfully fetched comments")

			// save each post and each comment in database in paralell
			for _, c := range comments {
				wg.Add(1)

				go func(c *Comment, p *Post, wg *sync.WaitGroup) {
					defer wg.Done()
					// save posts to database
					log.Println("post:", *p)

					// save comment to database
					log.Println("comment:", *c)

				}(&c, p, wg)
			}

		}(&p, &wg)

	}

	wg.Wait()
}
