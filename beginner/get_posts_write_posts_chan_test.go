package main

import "testing"

// go test -v -run BenchmarkGetPostsWritePostsChan -bench=.
func BenchmarkGetPostsWritePostsChan(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getPostsWritePostsChan(1)
	}
}

// 4.632s
// 3.801s
