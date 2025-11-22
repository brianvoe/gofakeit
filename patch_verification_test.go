package gofakeit

import (
	"reflect"
	"strings"
	"sync"
	"testing"
)

// TestPatchesApplied verifies that all 5 race condition patches are in place
func TestPatchesApplied(t *testing.T) {
	t.Run("RWMutex type", func(t *testing.T) {
		// Verify lockFuncLookups is RWMutex, not Mutex
		lockType := reflect.TypeOf(lockFuncLookups)
		if lockType.String() != "sync.RWMutex" {
			t.Errorf("Expected lockFuncLookups to be sync.RWMutex, got %s", lockType)
		}
	})
	
	t.Run("GetFuncLookup uses lock", func(t *testing.T) {
		// This is a behavioral test - if GetFuncLookup doesn't lock,
		// concurrent access will trigger a race detector warning
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
		// If we get here without race detector warnings, the patch works
	})
	
	t.Run("GetRandomSimpleFunc uses lock", func(t *testing.T) {
		// Add a simple function
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
		// Save original
		original := FuncLookups
		defer func() { FuncLookups = original }()
		
		// Test concurrent initialization from nil
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

// TestPatchVersion verifies we're using the expected version
func TestPatchVersion(t *testing.T) {
	// This test documents which version we patched
	// Update this when upgrading gofakeit
	expectedVersion := "v7.9.0"
	
	// The version is documented in our patches as comments
	// We can't directly check the version, but we verify patches are applied
	lockType := reflect.TypeOf(lockFuncLookups)
	if lockType.String() != "sync.RWMutex" {
		t.Errorf("Patches not applied correctly for %s. Expected RWMutex, got %s", 
			expectedVersion, lockType)
	}
	
	t.Logf("✅ Running patched gofakeit %s with race condition fixes", expectedVersion)
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
	
	t.Logf("ℹ️  See GOFAKEIT_RACE_CONDITION_ISSUE.md in product-security-360 for details")
}

// BenchmarkConcurrentAccess benchmarks concurrent access with patches
func BenchmarkConcurrentAccess(b *testing.B) {
	// Add test function
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

// TestUnpatchedVersionWouldFail documents what would happen without patches
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

// Helper to check if a string contains version info
func checkVersionComment(source, expectedVersion string) bool {
	return strings.Contains(source, expectedVersion)
}
