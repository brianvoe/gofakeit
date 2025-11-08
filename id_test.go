// Benchmark history:
// 2025-11-07 baseline: BenchmarkID-10  3,980,024   296.5 ns/op   24 B/op   1 allocs/op
// 2025-11-07 optimized: BenchmarkID-10 12,351,492    90.7 ns/op   24 B/op   1 allocs/op
// 2025-11-07 base32 optimized: BenchmarkIDPCG-10 14,558,895    73.3 ns/op   24 B/op   1 allocs/op
// 2025-11-07 base32 crypto  : BenchmarkIDCrypto-10  7,499,304   168.5 ns/op   24 B/op   1 allocs/op

package gofakeit

import (
	"bytes"
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func TestID(t *testing.T) {
	idValue := ID()

	if len(idValue) != idLength {
		t.Fatalf("id length = %d, want %d", len(idValue), idLength)
	}

	// Generate a 100,000 and check that it only contains the allowed characters
	for i := 0; i < 100000; i++ {
		idValue := ID()
		for j := 0; j < len(idValue); j++ {
			if bytes.IndexByte(idAlphabet, idValue[j]) == -1 {
				t.Fatalf("id contains invalid character %q at position %d", idValue[j], j)
			}
		}
	}
}

func TestIDUniquenessConcurrent(t *testing.T) {
	const goroutines = 60
	const idsPerRoutine = 10000

	seen := make(map[string]struct{}, goroutines*idsPerRoutine)
	var seenMu sync.Mutex
	var duplicate atomic.Bool
	errCh := make(chan string, 1)

	var wg sync.WaitGroup
	worker := func() {
		defer wg.Done()
		for i := 0; i < idsPerRoutine; i++ {
			if duplicate.Load() {
				return
			}

			idValue := ID()

			seenMu.Lock()
			if _, exists := seen[idValue]; exists {
				seenMu.Unlock()
				if duplicate.CompareAndSwap(false, true) {
					errCh <- idValue
				}
				return
			}
			seen[idValue] = struct{}{}
			seenMu.Unlock()
		}
	}

	wg.Add(goroutines)
	for i := 0; i < goroutines; i++ {
		go worker()
	}

	go func() {
		wg.Wait()
		close(errCh)
	}()

	if dup, ok := <-errCh; ok {
		t.Fatalf("duplicate id generated: %s", dup)
	}
}

func ExampleID() {
	Seed(11)
	fmt.Println(ID())

	// Output: nfrfzjsb87qckh6bpga2
}

func ExampleFaker_ID() {
	f := New(11)
	fmt.Println(f.ID())

	// Output: nfrfzjsb87qckh6bpga2
}

func BenchmarkID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = ID()
	}
}

func TestUUID(t *testing.T) {
	idValue := UUID()

	if len(idValue) != 36 {
		t.Fatalf("uuid length = %d, want 36", len(idValue))
	}

	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_ = UUID()
		}()
	}
	wg.Wait()
}

func ExampleUUID() {
	Seed(11)
	fmt.Println(UUID())

	// Output: b4ddf623-4ea6-48e5-9292-541f028d1fdb
}

func ExampleFaker_UUID() {
	f := New(11)
	fmt.Println(f.UUID())

	// Output: b4ddf623-4ea6-48e5-9292-541f028d1fdb
}

func BenchmarkUUID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = UUID()
	}
}
