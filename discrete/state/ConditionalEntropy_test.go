package state_test

import (
	"math"
	"testing"

	"github.com/kzahedi/goent/discrete"
	"github.com/kzahedi/goent/discrete/state"
)

func TestConditionalEntropyBase2(t *testing.T) {
	t.Log("Testing Conditional Entropy")
	p1 := [][]float64{
		{1.0 / 16.0, 1.0 / 16.0, 1.0 / 16.0, 1.0 / 16.0},
		{1.0 / 16.0, 1.0 / 16.0, 1.0 / 16.0, 1.0 / 16.0},
		{1.0 / 16.0, 1.0 / 16.0, 1.0 / 16.0, 1.0 / 16.0},
		{1.0 / 16.0, 1.0 / 16.0, 1.0 / 16.0, 1.0 / 16.0}}

	d1 := [][]int{
		{0, 0}, {0, 1}, {0, 2}, {0, 3},
		{1, 0}, {1, 1}, {1, 2}, {1, 3},
		{2, 0}, {2, 1}, {2, 2}, {2, 3},
		{3, 0}, {3, 1}, {3, 2}, {3, 3}}

	avg := discrete.ConditionalEntropyBase2(p1)
	s := state.ConditionalEntropyBase2(d1)
	r := 0.0
	for _, v := range s {
		r += v
	}

	r /= float64(len(s))

	diff := avg - r

	if diff > 0.00001 {
		t.Errorf("Discrete (%f) vs. state (%f) must be equal the difference is %f", avg, r, diff)
	}
}

func TestConditionalEntropyBase2Zero(t *testing.T) {
	t.Log("Testing Conditional Entropy")

	p1 := [][]float64{
		{1.0 / 4.0, 0.0, 0.0, 0.0},
		{0.0, 1.0 / 4.0, 0.0, 0.0},
		{0.0, 0.0, 1.0 / 4.0, 0.0},
		{0.0, 0.0, 0.0, 1.0 / 4.0}}

	d1 := [][]int{
		{0, 0}, {0, 0}, {0, 0}, {0, 0},
		{1, 1}, {1, 1}, {1, 1}, {1, 1},
		{2, 2}, {2, 2}, {2, 2}, {2, 2},
		{3, 3}, {3, 3}, {3, 3}, {3, 3}}

	avg := discrete.ConditionalEntropyBase2(p1)
	s := state.ConditionalEntropyBase2(d1)
	r := 0.0
	for _, v := range s {
		r += v
	}
	r /= float64(len(s))

	diff := avg - r

	if math.Abs(avg-r) > 0.00001 {
		t.Errorf("Discrete (%f) vs. state (%f) must be equal the difference is %f", avg, r, diff)
	}
}

func TestConditionalEntropyBaseE(t *testing.T) {
	t.Log("Testing Conditional Entropy")
	p1 := [][]float64{
		{1.0 / 16.0, 1.0 / 16.0, 1.0 / 16.0, 1.0 / 16.0},
		{1.0 / 16.0, 1.0 / 16.0, 1.0 / 16.0, 1.0 / 16.0},
		{1.0 / 16.0, 1.0 / 16.0, 1.0 / 16.0, 1.0 / 16.0},
		{1.0 / 16.0, 1.0 / 16.0, 1.0 / 16.0, 1.0 / 16.0}}

	d1 := [][]int{
		{0, 0}, {0, 1}, {0, 2}, {0, 3},
		{1, 0}, {1, 1}, {1, 2}, {1, 3},
		{2, 0}, {2, 1}, {2, 2}, {2, 3},
		{3, 0}, {3, 1}, {3, 2}, {3, 3}}

	avg := discrete.ConditionalEntropyBaseE(p1)
	s := state.ConditionalEntropyBaseE(d1)
	r := 0.0
	for _, v := range s {
		r += v
	}
	r /= float64(len(s))

	diff := avg - r

	if math.Abs(avg-r) > 0.00001 {
		t.Errorf("Discrete (%f) vs. state (%f) must be equal the difference is %f", avg, r, diff)
	}
}

func TestConditionalEntropyBaseEZero(t *testing.T) {
	t.Log("Testing Conditional Entropy")

	p1 := [][]float64{
		{1.0 / 4.0, 0.0, 0.0, 0.0},
		{0.0, 1.0 / 4.0, 0.0, 0.0},
		{0.0, 0.0, 1.0 / 4.0, 0.0},
		{0.0, 0.0, 0.0, 1.0 / 4.0}}

	d1 := [][]int{
		{0, 0}, {0, 0}, {0, 0}, {0, 0},
		{1, 1}, {1, 1}, {1, 1}, {1, 1},
		{2, 2}, {2, 2}, {2, 2}, {2, 2},
		{3, 3}, {3, 3}, {3, 3}, {3, 3}}

	avg := discrete.ConditionalEntropyBaseE(p1)
	s := state.ConditionalEntropyBaseE(d1)
	r := 0.0
	for _, v := range s {
		r += v
	}
	r /= float64(len(s))

	diff := avg - r

	if math.Abs(avg-r) > 0.00001 {
		t.Errorf("Discrete (%f) vs. state (%f) must be equal the difference is %f", avg, r, diff)
	}
}
