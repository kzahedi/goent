package state_test

import (
	"math"
	"math/rand"
	"testing"

	"github.com/kzahedi/goent/discrete"
	"github.com/kzahedi/goent/discrete/state"
)

func TestMorphologicalComputationW(t *testing.T) {
	data := make([][]int, 100, 100)
	for i := 0; i < 100; i++ {
		data[i] = make([]int, 3, 3)
		for j := 0; j < 3; j++ {
			data[i][j] = rand.Intn(10)
		}
	}
	pw2w1a1 := discrete.Emperical3D(data)

	r := discrete.MorphologicalComputationW(pw2w1a1)
	s := state.MorphologicalComputationW(data)

	q := 0.0
	for _, v := range s {
		q += v
	}
	q /= float64(len(s))

	if math.Abs(r-q) > 0.00001 {
		t.Errorf("MorphologicalComputationW should be %f but is %f", r, q)
	}
}

func TestMorphologicalComputationA(t *testing.T) {
	data := make([][]int, 100, 100)
	for i := 0; i < 100; i++ {
		data[i] = make([]int, 3, 3)
		for j := 0; j < 3; j++ {
			data[i][j] = rand.Intn(10)
		}
	}
	pw2w1a1 := discrete.Emperical3D(data)

	r := discrete.MorphologicalComputationA(pw2w1a1)
	s := state.MorphologicalComputationA(data)

	q := 0.0
	for _, v := range s {
		q += v
	}
	q /= float64(len(s))

	if math.Abs(r-q) > 0.00001 {
		t.Errorf("MorphologicalComputationA should be %f but is %f", r, q)
	}

}

// func TestMorphologicalComputationCW(t *testing.T) {
// dw2w1 := make([][]int, 100, 100)
// dw2a1 := make([][]int, 100, 100)
// data := make([][]int, 100, 100)

// for i := 0; i < 100; i++ {
// a := rand.Intn(10)
// b := rand.Intn(10)
// c := rand.Intn(10)
// dw2w1 := make([]int, 2, 2)
// dw2a1 := make([]int, 2, 2)
// data := make([]int, 4, 4)
// dw2w1[0] = a
// dw2w1[1] = b
// dw2a1[0] = a
// dw2a1[1] = c
// data[0] = a
// data[1] = b
// data[2] = c
// }

// pw2w1 := discrete.Create2D(10, 10)
// pw2a1 := discrete.Create2D(10, 10)

// r := discrete.MorphologicalComputationCW(pw2w1, pw2a1)
// s := state.MorphologicalComputationCW(pw2w1, pw2a1)

// q := 0.0
// for _, v := range s {
// q += v
// }
// q /= float64(len(s))

// if math.Abs(r-q) > 0.00001 {
// t.Errorf("MorphologicalComputationCW should be %f but is %f", r, q)
// }
// }

// func TestMorphologicalComputationWA(t *testing.T) {

// pw2w1a1 := discrete.Create3D(10, 10, 10)
// for i := 0; i < 10; i++ {
// for j := 0; j < 10; j++ {
// for k := 0; k < 10; k++ {
// pw2w1a1[i][j][k] = rand.Float64()
// }
// }
// }

// w2Dim := len(pw2w1a1)
// w1Dim := len(pw2w1a1[0])
// a1Dim := len(pw2w1a1[0][0])
// pw2a1 := discrete.Create2D(w2Dim, a1Dim)
// for w2 := 0; w2 < w2Dim; w2++ {
// for w1 := 0; w1 < w1Dim; w1++ {
// for a1 := 0; a1 < a1Dim; a1++ {
// pw2a1[w2][a1] += pw2w1a1[w2][w1][a1]
// }
// }
// }
// r := discrete.MorphologicalComputationWA(pw2w1a1)
// s := state.MorphologicalComputationWA(pw2w1a1)

// q := 0.0
// for _, v := range s {
// q += v
// }
// q /= float64(len(s))

// if math.Abs(r-q) > 0.00001 {
// t.Errorf("MorphologicalComputationWA should be %f but is %f", r, q)
// }
// }

// func TestMorphologicalComputationWS(t *testing.T) {
// pw2w1s1 := discrete.Create3D(10, 10, 10)
// for i := 0; i < 10; i++ {
// for j := 0; j < 10; j++ {
// for k := 0; k < 10; k++ {
// pw2w1s1[i][j][k] = rand.Float64()
// }
// }
// }

// w2Dim := len(pw2w1s1)
// w1Dim := len(pw2w1s1[0])
// s1Dim := len(pw2w1s1[0][0])
// pw2s1 := discrete.Create2D(w2Dim, s1Dim)
// for w2 := 0; w2 < w2Dim; w2++ {
// for w1 := 0; w1 < w1Dim; w1++ {
// for s1 := 0; s1 < s1Dim; s1++ {
// pw2s1[w2][s1] += pw2w1s1[w2][w1][s1]
// }
// }
// }
// r := discrete.MorphologicalComputationWS(pw2w1s1)
// s := state.MorphologicalComputationWS(pw2w1s1)

// q := 0.0
// for _, v := range s {
// q += v
// }
// q /= float64(len(s))

// if math.Abs(r-q) > 0.00001 {
// t.Errorf("MorphologicalComputationWS should be %f but is %f", r, q)
// }

// }

// func TestMorphologicalComputationMI(t *testing.T) {
// pw2w1 := discrete.Create2D(10, 10)
// pa1s1 := discrete.Create2D(10, 10)
// for i := 0; i < 10; i++ {
// for j := 0; j < 10; j++ {
// pw2w1[i][j] = rand.Float64()
// pa1s1[i][j] = rand.Float64()
// }
// }
// r := discrete.MorphologicalComputationMI(pw2w1, pa1s1)
// s := state.MorphologicalComputationMI(pw2w1, pa1s1)

// q := 0.0
// for _, v := range s {
// q += v
// }
// q /= float64(len(s))

// if math.Abs(r-q) > 0.00001 {
// t.Errorf("MorphologicalComputationMI should be %f but is %f", r, q)
// }
// }

// func TestMorphologicalComputationP(t *testing.T) {
// pw2w1a1 := discrete.Create3D(10, 10, 10)
// for i := 0; i < 10; i++ {
// for j := 0; j < 10; j++ {
// for k := 0; k < 10; k++ {
// pw2w1a1[i][j][k] = rand.Float64()
// }
// }
// }

// r := discrete.MorphologicalComputationP(pw2w1a1, 100, false)
// s := state.MorphologicalComputationP(pw2w1a1, 100, false)

// q := 0.0
// for _, v := range s {
// q += v
// }
// q /= float64(len(s))

// if math.Abs(r-q) > 0.00001 {
// t.Errorf("MorphologicalComputationP should be %f but is %f", r, q)
// }
// }
