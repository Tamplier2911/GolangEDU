package main

import (
	"fmt"
	"log"
	"sync"
)

// TODO:
// consider refactoring this function to leverage use of channels
// or make another one that leverages channel to accomplish some goal

func getPostsWritePostsDB(userId int) {
	// setup client
	c := HTTPClient{}.Setup()

	// setup db
	db, err := MySQL{}.Setup()
	if err != nil {
		log.Fatal("failed to setup database:", err)
	}
	// close db
	defer db.Close()

	// get all posts
	log.Println("getting all posts for user with id of 7")
	posts := []Post{}
	err = c.Get(fmt.Sprintf("https://jsonplaceholder.typicode.com/posts?userId=%d", userId), &posts)
	if err != nil {
		log.Fatal("failed to get posts:", err)
	}
	log.Println("successfully fetched posts")

	var wg sync.WaitGroup

	// get comments for each post in paralell
	log.Println("getting comments for each post")
	for _, p := range posts {
		wg.Add(1)

		go func(p Post, wg *sync.WaitGroup) {
			defer wg.Done()

			// save post to database
			err = db.Insert("posts", p)
			if err != nil {
				log.Fatal("failed to create post:", err)
			}
			log.Println("successfully insert data to database")

			// fetch comments for each post in paralell
			var comments []Comment
			err = c.Get(fmt.Sprintf("https://jsonplaceholder.typicode.com/comments?postId=%d", p.ID), &comments)
			if err != nil {
				log.Fatal("failed to get comments:", err)
			}
			log.Println("successfully fetched comments")

			// save each comment in database in paralell
			for _, c := range comments {
				wg.Add(1)

				go func(c Comment, wg *sync.WaitGroup) {
					defer wg.Done()
					// save comment to database
					err = db.Insert("comments", c)
					if err != nil {
						log.Fatal("failed to create post:", err)
					}
					log.Println("successfully insert data to database")

				}(c, wg)
			}

		}(p, &wg)

	}

	wg.Wait()
}
