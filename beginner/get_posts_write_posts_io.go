package main

import (
	"fmt"
	"log"
	"sync"
)

func getPostsWritePostsIO(postsNum int) {
	// setup client
	cl := HTTPClient{}.Setup()

	// setup io
	io := InputOutput{}.Setup()

	var wg sync.WaitGroup

	// create folder structure
	p := "storage/posts"
	err := io.MkDir(p)
	if err != nil {
		log.Fatal("failed to create directory:", err)
	}

	// get posts in parallel
	for i := 1; i <= postsNum; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, i int) {
			defer wg.Done()

			bs, err := cl.GetRaw(fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%d", i))
			if err != nil {
				log.Fatal("failed to get post:", err)
			}

			wp := fmt.Sprintf("%s/%d.txt", p, i)

			// write files with ioutil
			err = io.WriteIOUtil(wp, bs)
			if err != nil {
				log.Fatal("failed to write file with ioutil:", err)
			}

			// write files with bufio
			err = io.WriteBUFIO(wp, bs)
			if err != nil {
				log.Fatal("failed to write file with bufio:", err)
			}

			// write files with os
			err = io.WriteOS(wp, bs)
			if err != nil {
				log.Fatal("failed to write file with os package:", err)
			}

		}(&wg, i)
	}

	wg.Wait()
}
