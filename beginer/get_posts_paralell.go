package main

import (
	"fmt"
	"log"
	"sync"
)

func getAllPostsInParallel() {
	// setup client
	cl := HTTPClient{}.Setup()

	var wg sync.WaitGroup

	// get posts in paralell
	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, i int) {
			defer wg.Done()

			post := Post{}
			err := cl.Get(fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%d", i), &post)
			if err != nil {
				log.Panicln("failed to get post")
			}
			log.Printf("Post: \n %+v \n", post)

		}(&wg, i)
	}

	wg.Wait()
}
