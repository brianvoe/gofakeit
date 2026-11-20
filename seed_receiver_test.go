package gofakeit

import (
	"math/rand/v2"
	"sync"
	"testing"

	"github.com/brianvoe/gofakeit/v7/source"
)

func preserveGlobalFaker(t *testing.T) {
	t.Helper()
	original := GlobalFaker
	t.Cleanup(func() { GlobalFaker = original })
}

func TestSeedUsesReceiverSource(t *testing.T) {
	for name, factory := range map[string]func() *Faker{
		"PCG":  func() *Faker { return NewFaker(rand.NewPCG(1, 2), true) },
		"JSF":  func() *Faker { return NewFaker(source.NewJSF(1), true) },
		"SFC":  func() *Faker { return NewFaker(source.NewSFC(1), true) },
		"Dumb": func() *Faker { return NewFaker(source.NewDumb(1), true) },
	} {
		t.Run(name, func(t *testing.T) {
			preserveGlobalFaker(t)
			GlobalFaker = New(91)
			expectedGlobal := New(91)
			local := factory()
			expected := factory()
			for i := 0; i < 17; i++ {
				local.Uint64()
			}
			// Seed the independent oracle directly, not through Faker.Seed.
			switch src := expected.Rand.(type) {
			case *rand.PCG:
				src.Seed(77, 77)
			case *source.JSF:
				src.Seed(77)
			case *source.SFC:
				src.Seed(77)
			case *source.Dumb:
				src.Seed(77)
			}
			if err := local.Seed(77); err != nil {
				t.Fatal(err)
			}
			for i := 0; i < 10; i++ {
				if got, want := local.Uint64(), expected.Uint64(); got != want {
					t.Fatalf("local output %d=%d, want %d", i, got, want)
				}
			}
			for i := 0; i < 10; i++ {
				if got, want := GlobalFaker.Uint64(), expectedGlobal.Uint64(); got != want {
					t.Fatalf("local Seed changed GlobalFaker output %d", i)
				}
			}
		})
	}
}

func TestSeedReceiverIndependentOfGlobal(t *testing.T) {
	for name, global := range map[string]*Faker{"nil faker": nil, "nil source": {}, "unseedable source": NewFaker(source.NewCrypto(), true)} {
		t.Run(name, func(t *testing.T) {
			preserveGlobalFaker(t)
			GlobalFaker = global
			local := New(1)
			if err := local.Seed(77); err != nil {
				t.Fatalf("local Seed depends on global: %v", err)
			}
			if got, want := local.Uint64(), New(77).Uint64(); got != want {
				t.Fatalf("local output=%d, want %d", got, want)
			}
		})
	}
}

func TestSeedRejectsReceiverWithoutSource(t *testing.T) {
	preserveGlobalFaker(t)
	GlobalFaker = New(11)
	local := &Faker{Locked: true}
	if err := local.Seed(77); err == nil {
		t.Fatal("Seed accepted receiver with nil source")
	}
	if got, want := GlobalFaker.Uint64(), New(11).Uint64(); got != want {
		t.Fatal("invalid local Seed changed global generator")
	}
}

func TestSeedRejectsUnseedableReceiver(t *testing.T) {
	preserveGlobalFaker(t)
	GlobalFaker = New(11)
	if err := NewFaker(source.NewCrypto(), true).Seed(77); err == nil {
		t.Fatal("Seed accepted unseedable receiver")
	}
	if got, want := GlobalFaker.Uint64(), New(11).Uint64(); got != want {
		t.Fatal("unseedable local Seed changed global generator")
	}
}

func TestSeedPackageFunctionStillSeedsGlobal(t *testing.T) {
	preserveGlobalFaker(t)
	GlobalFaker = New(11)
	if err := Seed(77); err != nil {
		t.Fatal(err)
	}
	if got, want := GlobalFaker.Uint64(), New(77).Uint64(); got != want {
		t.Fatalf("global output=%d, want %d", got, want)
	}
}

func TestSeedLocalSourcesConcurrently(t *testing.T) {
	preserveGlobalFaker(t)
	GlobalFaker = New(11)
	var wg sync.WaitGroup
	for worker := 0; worker < 4; worker++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			local := New(91)
			for i := 0; i < 50; i++ {
				if err := local.Seed(77); err != nil {
					t.Error(err)
					return
				}
				local.Uint64()
			}
		}()
	}
	wg.Wait()
}

func TestSeedLocalDoesNotReseedGlobal(t *testing.T) {
	preserveGlobalFaker(t)
	GlobalFaker = New(91)
	want := New(91)
	if err := New(1).Seed(77); err != nil {
		t.Fatal(err)
	}
	for i := 0; i < 10; i++ {
		if got := GlobalFaker.Uint64(); got != want.Uint64() {
			t.Fatalf("local Seed changed global output %d", i)
		}
	}
}

func TestSeedReceiverTypedArguments(t *testing.T) {
	t.Run("PCG explicit pair", func(t *testing.T) {
		preserveGlobalFaker(t)
		GlobalFaker = New(91)
		local := New(1)
		want := NewFaker(rand.NewPCG(77, 88), true)
		if err := local.Seed(uint32(77), int64(88)); err != nil {
			t.Fatal(err)
		}
		for i := 0; i < 10; i++ {
			if local.Uint64() != want.Uint64() {
				t.Fatal("explicit PCG seed pair ignored")
			}
		}
	})
	t.Run("ChaCha8 array", func(t *testing.T) {
		preserveGlobalFaker(t)
		GlobalFaker = New(91)
		seed := [32]byte{1, 2, 3, 4, 5, 6, 7, 8}
		local := NewFaker(rand.NewChaCha8([32]byte{}), true)
		want := NewFaker(rand.NewChaCha8(seed), true)
		if err := local.Seed(seed); err != nil {
			t.Fatal(err)
		}
		for i := 0; i < 10; i++ {
			if local.Uint64() != want.Uint64() {
				t.Fatal("receiver's ChaCha8 seed signature ignored")
			}
		}
	})
}

type recordingSeedSource struct{ seeds []uint64 }

func (s *recordingSeedSource) Uint64() uint64   { return 1 }
func (s *recordingSeedSource) Seed(seed uint64) { s.seeds = append(s.seeds, seed) }

func TestSeedRandomDefaultsUseReceiver(t *testing.T) {
	preserveGlobalFaker(t)
	GlobalFaker = New(91)
	want := New(91)
	src := &recordingSeedSource{}
	local := NewFaker(src, true)
	if err := local.Seed(); err != nil {
		t.Fatal(err)
	}
	if err := local.Seed(0); err != nil {
		t.Fatal(err)
	}
	if len(src.seeds) != 2 {
		t.Fatalf("receiver seeded %d times, want 2", len(src.seeds))
	}
	for i := 0; i < 10; i++ {
		if GlobalFaker.Uint64() != want.Uint64() {
			t.Fatal("random default Seed changed global")
		}
	}
}

func TestSeedSharedReceiverConcurrently(t *testing.T) {
	preserveGlobalFaker(t)
	GlobalFaker = New(11)
	local := New(91)
	var wg sync.WaitGroup
	for worker := 0; worker < 4; worker++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 50; i++ {
				if err := local.Seed(77); err != nil {
					t.Error(err)
					return
				}
				local.Uint64()
			}
		}()
	}
	wg.Wait()
}
