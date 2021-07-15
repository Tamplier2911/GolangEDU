package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sync"
)

// TODO:
// implement basic layer of abstraction for I/O utilities
// implement teardown - removing not needed files and directories recursively
// wire up sql database
// finish last task

func getPostsWritePosts(n int) {
	// setup client
	cl := HTTPClient{}.Setup()

	var wg sync.WaitGroup

	// create folder structure
	p := "storage/posts"
	err := os.MkdirAll(p, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	// get posts in paralell
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, i int) {
			defer wg.Done()

			bs, err := cl.GetRaw(fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%d", i))
			if err != nil {
				log.Fatal("failed to get post", err)
			}
			log.Println(string(bs[:]))

			// using iotil writer
			err = ioutil.WriteFile(fmt.Sprintf("%s/%d.txt", p, i), bs, os.ModePerm)
			if err != nil {
				log.Fatal("failed to create and write a file", err)
			}

			/*

				// using os
				f, err := os.Create(fmt.Sprintf("%s/%d.txt", p, i))
				if err != nil {
					log.Fatal("failed to create file", err)
				}
				defer f.Close()

				// write
				_, err = f.Write(bs)
				if err != nil {
					log.Fatal("failed to write a file", err)
				}

				// or write as string
				_, err = f.WriteString(string(bs[:]))
				if err != nil {
					log.Fatal("failed to write a file as string", err)
				}
				f.Sync()

			*/

			/*

				// using bufio
				f, err := os.Create(fmt.Sprintf("%s/%d.txt", p, i))
				if err != nil {
					log.Fatal("failed to create file", err)
				}
				defer f.Close()

				w := bufio.NewWriter(f)
				_, err = w.WriteString(string(bs[:]))
				if err != nil {
					log.Fatal("failed to create file", err)
				}
				w.Flush()

			*/

		}(&wg, i)
	}

	wg.Wait()
}

// https://stackoverflow.com/questions/1821811/how-to-read-write-from-to-a-file-using-go
