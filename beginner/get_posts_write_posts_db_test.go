package main

import "testing"

// go test -v -run BenchmarkGetPostsWritePostsDB -bench=.
func BenchmarkGetPostsWritePostsDB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getPostsWritePostsDB(1)
	}
}

// 4.914s
// 4.044s
