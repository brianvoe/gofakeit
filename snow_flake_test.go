package gofakeit

import (
	"fmt"
	"testing"
)

func TestSnowFlake(t *testing.T) {
	worker, err := NewSFWorker(0)
	if err != nil {
		fmt.Println(err)
	}
	for i := 0; i < 5; i++ {
		id := worker.NextId()
		fmt.Println(id)
	}
	// Output:
	// 183945744284123136
	// 183945744284123137
	// 183945744284123138
	// 183945744284123139
	// 183945744284123140
}

func BenchmarkID(b *testing.B) {
	worker, _ := NewSFWorker(0)
	for i := 0; i < b.N; i++ {
		worker.NextId()
	}

	// cpu: 12th Gen Intel(R) Core(TM) i5-12400
	// BenchmarkID
	// BenchmarkID-12           4911304               243.9 ns/op
}
