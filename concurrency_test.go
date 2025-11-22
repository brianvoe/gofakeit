package gofakeit

import (
	"reflect"
	"sync"
	"testing"
)

// TestPatchesApplied verifies that all 5 race condition patches are in place
func TestPatchesApplied(t *testing.T) {
	// Clean up any test functions we add
	t.Cleanup(func() {
		RemoveFuncLookup("race-test")
		RemoveFuncLookup("simple-race-test")
		RemoveFuncLookup("another-simple")
		RemoveFuncLookup("init-race-test")
		RemoveFuncLookup("remove-race-test")
	})
	
	t.Run("RWMutex type", func(t *testing.T) {
		lockType := reflect.TypeOf(lockFuncLookups)
		if lockType.String() != "sync.RWMutex" {
			t.Errorf("Expected lockFuncLookups to be sync.RWMutex, got %s", lockType)
		}
	})

	t.Run("GetFuncLookup uses lock", func(t *testing.T) {
		// Behavioral test - if GetFuncLookup doesn't lock, race detector catches it
		var wg sync.WaitGroup
		for i := 0; i < 50; i++ {
			wg.Add(2)
			go func() {
				defer wg.Done()
				AddFuncLookup("race-test", Info{
					Category: "test",
					Output:   "string",
					Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
						return "test", nil
					},
				})
			}()
			go func() {
				defer wg.Done()
				_ = GetFuncLookup("race-test")
			}()
		}
		wg.Wait()
	})

	t.Run("GetRandomSimpleFunc uses lock", func(t *testing.T) {
		AddFuncLookup("simple-race-test", Info{
			Category: "test",
			Output:   "string",
			Params:   nil,
			Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
				return "test", nil
			},
		})

		var wg sync.WaitGroup
		for i := 0; i < 50; i++ {
			wg.Add(2)
			go func() {
				defer wg.Done()
				AddFuncLookup("another-simple", Info{
					Category: "test",
					Output:   "string",
					Params:   nil,
					Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
						return "test", nil
					},
				})
			}()
			go func() {
				defer wg.Done()
				faker := New(0)
				_, _ = GetRandomSimpleFunc(faker)
			}()
		}
		wg.Wait()
	})

	t.Run("AddFuncLookup locks before nil check", func(t *testing.T) {
		original := FuncLookups
		defer func() { FuncLookups = original }()

		FuncLookups = nil
		var wg sync.WaitGroup
		for i := 0; i < 100; i++ {
			wg.Add(1)
			go func(id int) {
				defer wg.Done()
				AddFuncLookup("init-race-test", Info{
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
			t.Error("FuncLookups should be initialized")
		}
	})

	t.Run("RemoveFuncLookup locks before existence check", func(t *testing.T) {
		var wg sync.WaitGroup
		for i := 0; i < 50; i++ {
			wg.Add(2)
			go func() {
				defer wg.Done()
				AddFuncLookup("remove-race-test", Info{
					Category: "test",
					Output:   "string",
					Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
						return "test", nil
					},
				})
			}()
			go func() {
				defer wg.Done()
				RemoveFuncLookup("remove-race-test")
			}()
		}
		wg.Wait()
	})
}

// TestConcurrentOperations tests realistic concurrent usage patterns
func TestConcurrentOperations(t *testing.T) {
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

// TestNilMapInitialization tests concurrent initialization from nil
func TestNilMapInitialization(t *testing.T) {
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

// TestStructGeneration tests Faker.Struct() concurrent usage
func TestStructGeneration(t *testing.T) {
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

// TestPatchVersion documents which version we patched
func TestPatchVersion(t *testing.T) {
	expectedVersion := "v7.9.0"

	lockType := reflect.TypeOf(lockFuncLookups)
	if lockType.String() != "sync.RWMutex" {
		t.Errorf("Patches not applied correctly for %s. Expected RWMutex, got %s",
			expectedVersion, lockType)
	}

	t.Logf("✅ Running patched gofakeit %s with TR race condition fixes", expectedVersion)
}

// TestDocumentedPatches ensures developers know about the patches
func TestDocumentedPatches(t *testing.T) {
	patches := []string{
		"Changed sync.Mutex to sync.RWMutex",
		"Added RLock to GetFuncLookup",
		"Added RLock to GetRandomSimpleFunc",
		"Moved nil check inside lock in AddFuncLookup",
		"Moved existence check inside lock in RemoveFuncLookup",
	}

	t.Logf("📝 This gofakeit copy includes the following patches:")
	for i, patch := range patches {
		t.Logf("   %d. %s", i+1, patch)
	}

	t.Logf("ℹ️  See TR_PATCHES.md for detailed documentation")
}

// BenchmarkConcurrentAccess benchmarks concurrent access with patches
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

// TestUnpatchedVersionWouldFail documents what would fail without patches
func TestUnpatchedVersionWouldFail(t *testing.T) {
	t.Skip("This test documents the race condition - run with -race on unpatched version to see failure")

	// This is what would fail without patches:
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(2)

		// Writer (AddFuncLookup)
		go func() {
			defer wg.Done()
			AddFuncLookup("race", Info{
				Category: "test",
				Output:   "string",
				Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
					return "test", nil
				},
			})
		}()

		// Reader (GetFuncLookup) - would race without patch
		go func() {
			defer wg.Done()
			_ = GetFuncLookup("race")
		}()
	}
	wg.Wait()
}
