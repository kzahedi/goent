package state_test

import (
	"math"
	"math/rand"
	"testing"

	"github.com/kzahedi/goent/continuous"
	"github.com/kzahedi/goent/continuous/state"
	"github.com/kzahedi/goent/discrete"
)

func TestMorphologicalComputationW(t *testing.T) {
	n := 1000
	data := discrete.Create2D(n, 9)
	w2Indices := []int{0, 1, 2}
	w1Indices := []int{3, 4, 5}
	a1Indices := []int{6, 7, 8}
	k := 30

	for i := 0; i < n; i++ {
		for j := 0; j < 9; j++ {
			data[i][j] = rand.Float64() * 100.0
		}
	}

	r := continuous.MorphologicalComputationW(data, w2Indices, w1Indices, a1Indices, k, false)
	s := state.MorphologicalComputationW(data, w2Indices, w1Indices, a1Indices, k, false)

	q := 0.0
	for _, v := range s {
		q += v
	}
	q /= float64(len(s))

	if math.Abs(r-q) > 0.000001 {
		t.Errorf("Continuous MC_W should be %f but it is %f", r, q)
	}
}

func TestMorphologicalComputationA(t *testing.T) {
	n := 1000
	data := discrete.Create2D(n, 9)
	w2Indices := []int{0, 1, 2}
	w1Indices := []int{3, 4, 5}
	a1Indices := []int{6, 7, 8}
	k := 30

	for i := 0; i < n; i++ {
		for j := 0; j < 9; j++ {
			data[i][j] = rand.Float64() * 100.0
		}
	}

	r := continuous.MorphologicalComputationA(data, w2Indices, w1Indices, a1Indices, k, false)
	s := state.MorphologicalComputationA(data, w2Indices, w1Indices, a1Indices, k, false)

	q := 0.0
	for _, v := range s {
		q += v
	}
	q /= float64(len(s))

	if math.Abs(r-q) > 0.000001 {
		t.Errorf("Continuous MC_A should be %f but it is %f", r, q)
	}

}

// func TestMorphologicalComputationCW1(t *testing.T) {
// n := 1000
// data := discrete.Create2D(n, 9)
// w2Indices := []int{0, 1, 2}
// w1Indices := []int{3, 4, 5}
// a1Indices := []int{6, 7, 8}
// k := 30

// for i := 0; i < n; i++ {
// for j := 0; j < 9; j++ {
// data[i][j] = rand.Float64() * 100.0
// }
// }

// r := continuous.MorphologicalComputationCW1(data, w2Indices, w1Indices, a1Indices, k, false)
// s := state.MorphologicalComputationCW1(data, w2Indices, w1Indices, a1Indices, k, false)

// q := 0.0
// for _, v := range s {
// q += v
// }
// q /= float64(len(s))

// if math.Abs(r-q) > 0.000001 {
// t.Errorf("Continuous MC_CW1 should be %f but it is %f", r, q)
// }

// }

// func TestMorphologicalComputationCW2(t *testing.T) {
// n := 1000
// data := discrete.Create2D(n, 9)
// w2Indices := []int{0, 1, 2}
// w1Indices := []int{3, 4, 5}
// a1Indices := []int{6, 7, 8}
// k := 30

// for i := 0; i < n; i++ {
// for j := 0; j < 9; j++ {
// data[i][j] = rand.Float64() * 100.0
// }
// }

// r := continuous.MorphologicalComputationCW2(data, w2Indices, w1Indices, a1Indices, k, false)
// s := state.MorphologicalComputationCW2(data, w2Indices, w1Indices, a1Indices, k, false)

// q := 0.0
// for _, v := range s {
// q += v
// }
// q /= float64(len(s))

// if math.Abs(r-q) > 0.000001 {
// t.Errorf("Continuous MC_CW2 should be %f but it is %f", r, q)
// }

// }

// func TestMorphologicalComputationWA1(t *testing.T) {
// n := 1000
// data := discrete.Create2D(n, 9)
// w2Indices := []int{0, 1, 2}
// w1Indices := []int{3, 4, 5}
// a1Indices := []int{6, 7, 8}
// k := 30

// for i := 0; i < n; i++ {
// for j := 0; j < 9; j++ {
// data[i][j] = rand.Float64() * 100.0
// }
// }

// r := continuous.MorphologicalComputationWA1(data, w2Indices, w1Indices, a1Indices, k, false)
// s := state.MorphologicalComputationWA1(data, w2Indices, w1Indices, a1Indices, k, false)

// q := 0.0
// for _, v := range s {
// q += v
// }
// q /= float64(len(s))

// if math.Abs(r-q) > 0.000001 {
// t.Errorf("Continuous MC_WA1 should be %f but it is %f", r, q)
// }

// }

// func TestMorphologicalComputationWA2(t *testing.T) {
// n := 1000
// data := discrete.Create2D(n, 9)
// w2Indices := []int{0, 1, 2}
// w1Indices := []int{3, 4, 5}
// a1Indices := []int{6, 7, 8}
// k := 30

// for i := 0; i < n; i++ {
// for j := 0; j < 9; j++ {
// data[i][j] = rand.Float64() * 100.0
// }
// }

// r := continuous.MorphologicalComputationWA2(data, w2Indices, w1Indices, a1Indices, k, false)
// s := state.MorphologicalComputationWA2(data, w2Indices, w1Indices, a1Indices, k, false)

// q := 0.0
// for _, v := range s {
// q += v
// }
// q /= float64(len(s))

// if math.Abs(r-q) > 0.000001 {
// t.Errorf("Continuous MC_WA2 should be %f but it is %f", r, q)
// }

// }

// func TestMorphologicalComputationWS1(t *testing.T) {
// n := 1000
// data := discrete.Create2D(n, 9)
// w2Indices := []int{0, 1, 2}
// w1Indices := []int{3, 4, 5}
// s1Indices := []int{6, 7, 8}
// k := 30

// for i := 0; i < n; i++ {
// for j := 0; j < 9; j++ {
// data[i][j] = rand.Float64() * 100.0
// }
// }

// r := continuous.MorphologicalComputationWS1(data, w2Indices, w1Indices, s1Indices, k, false)
// s := state.MorphologicalComputationWS1(data, w2Indices, w1Indices, s1Indices, k, false)

// q := 0.0
// for _, v := range s {
// q += v
// }
// q /= float64(len(s))

// if math.Abs(r-q) > 0.000001 {
// t.Errorf("Continuous MC_WS1 should be %f but it is %f", r, q)
// }

// }

// func TestMorphologicalComputationWS2(t *testing.T) {
// n := 1000
// data := discrete.Create2D(n, 9)
// w2Indices := []int{0, 1, 2}
// w1Indices := []int{3, 4, 5}
// s1Indices := []int{6, 7, 8}
// k := 30

// for i := 0; i < n; i++ {
// for j := 0; j < 9; j++ {
// data[i][j] = rand.Float64() * 100.0
// }
// }

// r := continuous.MorphologicalComputationWS2(data, w2Indices, w1Indices, s1Indices, k, false)
// s := state.MorphologicalComputationWS2(data, w2Indices, w1Indices, s1Indices, k, false)

// q := 0.0
// for _, v := range s {
// q += v
// }
// q /= float64(len(s))

// if math.Abs(r-q) > 0.000001 {
// t.Errorf("Continuous MC_WS2 should be %f but it is %f", r, q)
// }

// }

// func TestMorphologicalComputationMI1(t *testing.T) {
// n := 1000
// data := discrete.Create2D(n, 12)
// w2Indices := []int{0, 1, 2}
// w1Indices := []int{3, 4, 5}
// a1Indices := []int{6, 7, 8}
// s1Indices := []int{9, 10, 11}
// k := 30

// for i := 0; i < n; i++ {
// for j := 0; j < 12; j++ {
// data[i][j] = rand.Float64() * 100.0
// }
// }

// r := continuous.MorphologicalComputationMI1(data, w2Indices, w1Indices, s1Indices, a1Indices, k, false)
// s := state.MorphologicalComputationMI1(data, w2Indices, w1Indices, s1Indices, a1Indices, k, false)

// q := 0.0
// for _, v := range s {
// q += v
// }
// q /= float64(len(s))

// if math.Abs(r-q) > 0.000001 {
// t.Errorf("Continuous MC_MI1 should be %f but it is %f", r, q)
// }

// }

// func TestMorphologicalComputationMI2(t *testing.T) {
// n := 1000
// data := discrete.Create2D(n, 12)
// w2Indices := []int{0, 1, 2}
// w1Indices := []int{3, 4, 5}
// a1Indices := []int{6, 7, 8}
// s1Indices := []int{9, 10, 11}
// k := 30

// for i := 0; i < n; i++ {
// for j := 0; j < 12; j++ {
// data[i][j] = rand.Float64() * 100.0
// }
// }

// r := continuous.MorphologicalComputationMI2(data, w2Indices, w1Indices, s1Indices, a1Indices, k, false)
// s := state.MorphologicalComputationMI2(data, w2Indices, w1Indices, s1Indices, a1Indices, k, false)

// q := 0.0
// for _, v := range s {
// q += v
// }
// q /= float64(len(s))

// if math.Abs(r-q) > 0.000001 {
// t.Errorf("Continuous MC_MI2 should be %f but it is %f", r, q)
// }

// }
