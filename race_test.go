package gofakeit

import (
	"sync"
	"testing"
)

// TestConcurrentFuncLookupOperations tests that all FuncLookup operations
// are safe for concurrent use. This test will fail if there are any
// unprotected map accesses.
func TestConcurrentFuncLookupOperations(t *testing.T) {
	// Test concurrent AddFuncLookup + GetFuncLookup
	t.Run("AddAndGet", func(t *testing.T) {
		var wg sync.WaitGroup
		iterations := 100
		
		for i := 0; i < iterations; i++ {
			wg.Add(2)
			
			// Writer goroutine
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
			
			// Reader goroutine
			go func() {
				defer wg.Done()
				_ = GetFuncLookup("test-func")
			}()
		}
		
		wg.Wait()
	})
	
	// Test concurrent GetRandomSimpleFunc
	t.Run("GetRandomSimple", func(t *testing.T) {
		// Add some simple functions first
		AddFuncLookup("simple1", Info{
			Category:    "test",
			Description: "Simple test 1",
			Output:      "string",
			Params:      nil, // No params = simple
			Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
				return "test1", nil
			},
		})
		AddFuncLookup("simple2", Info{
			Category:    "test",
			Description: "Simple test 2",
			Output:      "string",
			Params:      nil,
			Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
				return "test2", nil
			},
		})
		
		var wg sync.WaitGroup
		iterations := 100
		
		for i := 0; i < iterations; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				faker := New(0)
				_, _ = GetRandomSimpleFunc(faker)
			}()
		}
		
		wg.Wait()
	})
	
	// Test concurrent AddFuncLookup + RemoveFuncLookup
	t.Run("AddAndRemove", func(t *testing.T) {
		var wg sync.WaitGroup
		iterations := 100
		
		for i := 0; i < iterations; i++ {
			wg.Add(2)
			
			// Writer goroutine
			go func(id int) {
				defer wg.Done()
				AddFuncLookup("removable-func", Info{
					Category:    "test",
					Description: "Removable function",
					Output:      "string",
					Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
						return "test", nil
					},
				})
			}(i)
			
			// Remover goroutine
			go func() {
				defer wg.Done()
				RemoveFuncLookup("removable-func")
			}()
		}
		
		wg.Wait()
	})
	
	// Test all operations concurrently
	t.Run("AllOperations", func(t *testing.T) {
		var wg sync.WaitGroup
		iterations := 50
		
		for i := 0; i < iterations; i++ {
			wg.Add(4)
			
			// Add
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
			
			// Get
			go func() {
				defer wg.Done()
				_ = GetFuncLookup("concurrent-func")
			}()
			
			// GetRandom
			go func() {
				defer wg.Done()
				faker := New(0)
				_, _ = GetRandomSimpleFunc(faker)
			}()
			
			// Remove
			go func() {
				defer wg.Done()
				RemoveFuncLookup("concurrent-func")
			}()
		}
		
		wg.Wait()
	})
}

// TestNilMapInitialization tests that concurrent AddFuncLookup calls
// safely initialize the FuncLookups map
func TestNilMapInitialization(t *testing.T) {
	// Save original map
	original := FuncLookups
	defer func() { FuncLookups = original }()
	
	// Set to nil to test initialization
	FuncLookups = nil
	
	var wg sync.WaitGroup
	iterations := 100
	
	for i := 0; i < iterations; i++ {
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
	
	// Verify map was initialized
	if FuncLookups == nil {
		t.Error("FuncLookups should not be nil after concurrent AddFuncLookup calls")
	}
}

// TestConcurrentStructGeneration tests that using Faker.Struct()
// concurrently doesn't cause races (it uses GetFuncLookup internally)
func TestConcurrentStructGeneration(t *testing.T) {
	// Register a custom function
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
	iterations := 100
	
	for i := 0; i < iterations; i++ {
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
