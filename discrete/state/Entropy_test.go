package state_test

import (
	"math"
	"testing"

	"github.com/kzahedi/goent/discrete"
	"github.com/kzahedi/goent/discrete/state"
)

func TestEntropyBase2(t *testing.T) {
	t.Log("Testing Entropy")
	p1 := []float64{0.5, 0.5, 0.5, 0.5}
	d1 := []int{0, 1, 2, 3}

	s := state.EntropyBase2(d1)
	q := discrete.EntropyBase2(p1)

	r := 0.0
	for _, v := range s {
		r += v
	}
	r /= 4.0

	if math.Abs(r-q) > 0.00001 {
		t.Errorf("Entropy per state should have the same averaged value, but %f != %f", q, r)
	}

	p2 := []float64{1.0, 0.0, 0.0, 0.0}
	d2 := []int{0, 0, 0, 0}

	s = state.EntropyBase2(d2)
	q = discrete.EntropyBase2(p2)

	r = 0.0
	for _, v := range s {
		r += v
	}
	r /= 4.0

	if math.Abs(r-q) > 0.00001 {
		t.Errorf("Entropy per state should have the same averaged value, but %f != %f", q, r)
	}
}

func TestEntropyBaseE(t *testing.T) {
	t.Log("Testing Entropy")
	p1 := []float64{0.5, 0.5, 0.5, 0.5}
	d1 := []int{0, 1, 2, 3}

	s := state.EntropyBaseE(d1)
	q := discrete.EntropyBaseE(p1)

	r := 0.0
	for _, v := range s {
		r += v
	}
	r /= 4.0

	if math.Abs(r-q) > 0.00001 {
		t.Errorf("Entropy per state should have the same averaged value, but %f != %f", q, r)
	}

	p2 := []float64{1.0, 0.0, 0.0, 0.0}
	d2 := []int{0, 0, 0, 0}

	s = state.EntropyBaseE(d2)
	q = discrete.EntropyBaseE(p2)

	r = 0.0
	for _, v := range s {
		r += v
	}
	r /= 4.0

	if math.Abs(r-q) > 0.00001 {
		t.Errorf("Entropy per state should have the same averaged value, but %f != %f", q, r)
	}
}
