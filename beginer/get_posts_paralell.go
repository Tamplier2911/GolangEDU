package main

import (
	"fmt"
	"log"
	"sync"
)

func getPostsInParallel(n int) {
	// setup client
	cl := HTTPClient{}.Setup()

	var wg sync.WaitGroup

	// get posts in paralell
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, i int) {
			defer wg.Done()

			post := Post{}
			err := cl.Get(fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%d", i), &post)
			if err != nil {
				log.Fatal("failed to get post:", err)
			}

		}(&wg, i)
	}

	wg.Wait()
}
