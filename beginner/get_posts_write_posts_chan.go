package main

import (
	"fmt"
	"log"
	"sync"
)

// looks ugly as hell, but works

// Get all posts from json placeholder with user id of <n>
// get all comments for each post in parallel, save everything in database in in parallel
// utilize channels
func getPostsWritePostsChan(userId int) {
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
	log.Println("getting all posts for user with id:", userId)
	posts := []Post{}
	err = c.Get(fmt.Sprintf("https://jsonplaceholder.typicode.com/posts?userId=%d", userId), &posts)
	if err != nil {
		log.Fatal("failed to get posts:", err)
	}
	log.Println("successfully fetched posts")

	// create wait group
	var wg sync.WaitGroup

	// create channels
	c_posts := make(chan Post)
	c_comment := make(chan Comment)

	for _, p := range posts {
		wg.Add(2)
		go postWriter(c_posts, &wg, p)
		go postReader(c_posts, c_comment, &wg, &db, &c)
	}

	// chan that takes whatever and writes to db
	wg.Wait()
}

func postWriter(c chan<- Post, wg *sync.WaitGroup, post Post) {
	defer wg.Done()

	c <- post
	log.Println("Written post with id of:", post.ID)

}

func commentWriter(c chan<- Comment, wg *sync.WaitGroup, comment Comment) {
	defer wg.Done()

	c <- comment
	log.Println("Written comment with id of:", comment.ID)
}

func postReader(cp <-chan Post, cc chan Comment, wg *sync.WaitGroup, db *MySQL, client *HTTPClient) {
	defer wg.Done()
	post := <-cp
	log.Println("Read post from channel with id of:", post.ID)

	// save post to database
	wg.Add(1)
	go dbInsert("posts", post, wg, db)

	// get post comments
	wg.Add(1)
	go func(post Post, wg *sync.WaitGroup, client *HTTPClient) {
		defer wg.Done()

		var comments []Comment
		err := client.Get(fmt.Sprintf("https://jsonplaceholder.typicode.com/comments?postId=%d", post.ID), &comments)
		if err != nil {
			log.Fatal("failed to get comments:", err)
		}
		log.Println("successfully fetched comments")

		for _, c := range comments {
			wg.Add(2)
			go commentWriter(cc, wg, c)
			go commentReader(cc, wg, db)
		}

	}(post, wg, client)
}

func commentReader(cc <-chan Comment, wg *sync.WaitGroup, db *MySQL) {
	defer wg.Done()
	comment := <-cc
	log.Println("Read comment from channel with id of:", comment.ID)

	// save comments to database
	wg.Add(1)
	go dbInsert("comments", comment, wg, db)
}

func dbInsert(tableName string, model interface{}, wg *sync.WaitGroup, db *MySQL) {
	defer wg.Done()
	// write data to database
	err := db.Insert(tableName, model)
	if err != nil {
		log.Fatal("failed to insert data to database:", err)
	}
	log.Println("successfully insert data to database")
}
