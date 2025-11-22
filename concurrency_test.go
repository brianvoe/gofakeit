package gofakeit

import (
	"sync"
	"testing"
)

// TestConcurrentFuncLookupOperations tests concurrent operations on func lookups
func TestConcurrentFuncLookupOperations(t *testing.T) {
	// Clean up test functions
	t.Cleanup(func() {
		RemoveFuncLookup("test-func")
		RemoveFuncLookup("simple1")
		RemoveFuncLookup("simple2")
		RemoveFuncLookup("removable-func")
		RemoveFuncLookup("concurrent-func")
	})

	t.Run("Concurrent Add and Get", func(t *testing.T) {
		var wg sync.WaitGroup
		iterations := 100

		for i := 0; i < iterations; i++ {
			wg.Add(2)

			go func(id int) {
				defer wg.Done()
				AddFuncLookup("test-func", Info{
					Category:    "test",
					Description: "Test function",
					Output:      "string",
					Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
						return "test", nil
					},
				})
			}(i)

			go func() {
				defer wg.Done()
				_ = GetFuncLookup("test-func")
			}()
		}

		wg.Wait()
	})

	t.Run("Concurrent GetRandomSimple", func(t *testing.T) {
		AddFuncLookup("simple1", Info{
			Category: "test",
			Output:   "string",
			Params:   nil,
			Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
				return "test1", nil
			},
		})
		AddFuncLookup("simple2", Info{
			Category: "test",
			Output:   "string",
			Params:   nil,
			Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
				return "test2", nil
			},
		})

		var wg sync.WaitGroup
		for i := 0; i < 100; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				faker := New(0)
				_, _ = GetRandomSimpleFunc(faker)
			}()
		}

		wg.Wait()
	})

	t.Run("Concurrent Add and Remove", func(t *testing.T) {
		var wg sync.WaitGroup
		for i := 0; i < 100; i++ {
			wg.Add(2)

			go func(id int) {
				defer wg.Done()
				AddFuncLookup("removable-func", Info{
					Category: "test",
					Output:   "string",
					Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
						return "test", nil
					},
				})
			}(i)

			go func() {
				defer wg.Done()
				RemoveFuncLookup("removable-func")
			}()
		}

		wg.Wait()
	})

	t.Run("All operations simultaneously", func(t *testing.T) {
		var wg sync.WaitGroup
		iterations := 50

		for i := 0; i < iterations; i++ {
			wg.Add(4)

			go func(id int) {
				defer wg.Done()
				AddFuncLookup("concurrent-func", Info{
					Category: "test",
					Output:   "string",
					Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
						return "test", nil
					},
				})
			}(i)

			go func() {
				defer wg.Done()
				_ = GetFuncLookup("concurrent-func")
			}()

			go func() {
				defer wg.Done()
				faker := New(0)
				_, _ = GetRandomSimpleFunc(faker)
			}()

			go func() {
				defer wg.Done()
				RemoveFuncLookup("concurrent-func")
			}()
		}

		wg.Wait()
	})
}

// TestConcurrentNilMapInitialization tests concurrent initialization from nil map
func TestConcurrentNilMapInitialization(t *testing.T) {
	t.Cleanup(func() {
		RemoveFuncLookup("init-test")
	})

	original := FuncLookups
	defer func() { FuncLookups = original }()

	FuncLookups = nil

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			AddFuncLookup("init-test", Info{
				Category: "test",
				Output:   "string",
				Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
					return "test", nil
				},
			})
		}(i)
	}

	wg.Wait()

	if FuncLookups == nil {
		t.Error("FuncLookups should not be nil after concurrent AddFuncLookup calls")
	}
}

// TestConcurrentStructGeneration tests Faker.Struct() concurrent usage
func TestConcurrentStructGeneration(t *testing.T) {
	t.Cleanup(func() {
		RemoveFuncLookup("customfield")
	})

	AddFuncLookup("customfield", Info{
		Category:    "custom",
		Description: "Custom field",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return "custom-value", nil
		},
	})

	type TestStruct struct {
		Name   string `fake:"{name}"`
		Email  string `fake:"{email}"`
		Age    int    `fake:"{number:1,100}"`
		Custom string `fake:"{customfield}"`
	}

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			var s TestStruct
			faker := New(0)
			_ = faker.Struct(&s)
		}()
	}

	wg.Wait()
}

// BenchmarkConcurrentAccess benchmarks concurrent access to func lookups
func BenchmarkConcurrentAccess(b *testing.B) {
	b.Cleanup(func() {
		RemoveFuncLookup("bench-test")
	})

	AddFuncLookup("bench-test", Info{
		Category: "test",
		Output:   "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return "test", nil
		},
	})

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = GetFuncLookup("bench-test")
		}
	})
}
