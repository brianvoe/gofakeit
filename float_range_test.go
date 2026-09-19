package gofakeit

import (
	"math"
	"testing"
)

func TestFloat32RangeLargeFiniteBounds(t *testing.T) {
	for _, bounds := range [][2]float32{
		{-math.MaxFloat32, math.MaxFloat32},
		{-math.MaxFloat32, math.MaxFloat32 / 2},
		{-math.MaxFloat32 / 2, math.MaxFloat32},
		{math.MaxFloat32, -math.MaxFloat32},
	} {
		f := New(42)
		low, high := math.Min(float64(bounds[0]), float64(bounds[1])), math.Max(float64(bounds[0]), float64(bounds[1]))
		for i := 0; i < 100; i++ {
			value := float64(f.Float32Range(bounds[0], bounds[1]))
			if math.IsNaN(value) || math.IsInf(value, 0) || value < low || value > high {
				t.Fatalf("Float32Range(%v, %v) = %v, want a finite value within the bounds", bounds[0], bounds[1], value)
			}
		}
	}
}

func TestFloat64RangeLargeFiniteBounds(t *testing.T) {
	for _, bounds := range [][2]float64{
		{-math.MaxFloat64, math.MaxFloat64},
		{-math.MaxFloat64, math.MaxFloat64 / 2},
		{-math.MaxFloat64 / 2, math.MaxFloat64},
		{math.MaxFloat64, -math.MaxFloat64},
	} {
		f := New(42)
		low, high := math.Min(bounds[0], bounds[1]), math.Max(bounds[0], bounds[1])
		for i := 0; i < 100; i++ {
			value := f.Float64Range(bounds[0], bounds[1])
			if math.IsNaN(value) || math.IsInf(value, 0) || value < low || value > high {
				t.Fatalf("Float64Range(%v, %v) = %v, want a finite value within the bounds", bounds[0], bounds[1], value)
			}
		}
	}
}

type zeroFloatRangeSource struct{}

func (zeroFloatRangeSource) Uint64() uint64 { return 0 }

func TestFloatRangeLargeFiniteBoundsZeroDraw(t *testing.T) {
	f := NewFaker(zeroFloatRangeSource{}, false)
	if got := f.Float32Range(-math.MaxFloat32, math.MaxFloat32); got != -math.MaxFloat32 {
		t.Errorf("Float32Range with zero draw = %v, want %v", got, -math.MaxFloat32)
	}
	if got := f.Float64Range(-math.MaxFloat64, math.MaxFloat64); got != -math.MaxFloat64 {
		t.Errorf("Float64Range with zero draw = %v, want %v", got, -math.MaxFloat64)
	}
}
