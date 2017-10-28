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

func TestMorphologicalComputationCW(t *testing.T) {
	data := discrete.Create2DInt(100, 3)
	pw2w1 := discrete.Create2D(10, 10)
	pw2a1 := discrete.Create2D(10, 10)

	for i := 0; i < 100; i++ {
		data[i] = make([]int, 3, 3)
		data[i][0] = rand.Intn(10)
		data[i][1] = rand.Intn(10)
		data[i][2] = rand.Intn(10)
	}
	pw2w1a1 := discrete.Emperical3D(data)
	w2Dim := len(pw2w1a1)
	w1Dim := len(pw2w1a1[0])
	a1Dim := len(pw2w1a1[0][0])

	for w2 := 0; w2 < w2Dim; w2++ {
		for w1 := 0; w1 < w1Dim; w1++ {
			for a1 := 0; a1 < a1Dim; a1++ {
				pw2w1[w2][w1] += pw2w1a1[w2][w1][a1]
				pw2a1[w2][a1] += pw2w1a1[w2][w1][a1]
			}
		}
	}

	r := discrete.MorphologicalComputationCW(pw2w1, pw2a1)
	s := state.MorphologicalComputationCW(data)

	q := 0.0
	for _, v := range s {
		q += v
	}
	q /= float64(len(s))

	if math.Abs(r-q) > 0.00001 {
		t.Errorf("MorphologicalComputationCW should be %f but is %f", r, q)
	}
}

func TestMorphologicalComputationWA(t *testing.T) {
	data := discrete.Create2DInt(100, 3)
	for i := 0; i < 100; i++ {
		data[i] = make([]int, 3, 3)
		data[i][0] = rand.Intn(10)
		data[i][1] = rand.Intn(10)
		data[i][2] = rand.Intn(10)
	}
	pw2w1a1 := discrete.Emperical3D(data)

	r := discrete.MorphologicalComputationWA(pw2w1a1)
	s := state.MorphologicalComputationWA(data)

	q := 0.0
	for _, v := range s {
		q += v
	}
	q /= float64(len(s))

	if math.Abs(r-q) > 0.00001 {
		t.Errorf("MorphologicalComputationWA should be %f but is %f", r, q)
	}
}

func TestMorphologicalComputationWS(t *testing.T) {
	data := discrete.Create2DInt(100, 3)
	for i := 0; i < 100; i++ {
		data[i] = make([]int, 3, 3)
		data[i][0] = rand.Intn(10)
		data[i][1] = rand.Intn(10)
		data[i][2] = rand.Intn(10)
	}
	pw2w1s1 := discrete.Emperical3D(data)

	r := discrete.MorphologicalComputationWS(pw2w1s1)
	s := state.MorphologicalComputationWS(data)

	q := 0.0
	for _, v := range s {
		q += v
	}
	q /= float64(len(s))

	if math.Abs(r-q) > 0.00001 {
		t.Errorf("MorphologicalComputationWS should be %f but is %f", r, q)
	}
}

func TestMorphologicalComputationMI(t *testing.T) {
	data := discrete.Create2DInt(100, 4)
	pw2w1 := discrete.Create2D(10, 10)
	pa1s1 := discrete.Create2D(10, 10)
	for i := 0; i < 100; i++ {
		data[i] = make([]int, 4, 4)
		data[i][0] = rand.Intn(10)
		data[i][1] = rand.Intn(10)
		data[i][2] = rand.Intn(10)
		data[i][3] = rand.Intn(10)
	}
	pw2w1s1a1 := discrete.Emperical4D(data)
	w2Dim := len(pw2w1s1a1)
	w1Dim := len(pw2w1s1a1[0])
	s1Dim := len(pw2w1s1a1[0][0])
	a1Dim := len(pw2w1s1a1[0][0][0])

	for w2 := 0; w2 < w2Dim; w2++ {
		for w1 := 0; w1 < w1Dim; w1++ {
			for a1 := 0; a1 < a1Dim; a1++ {
				for s1 := 0; s1 < s1Dim; s1++ {
					pw2w1[w2][w1] += pw2w1s1a1[w2][w1][s1][a1]
					pa1s1[a1][s1] += pw2w1s1a1[w2][w1][s1][a1]
				}
			}
		}
	}

	r := discrete.MorphologicalComputationMI(pw2w1, pa1s1)
	s := state.MorphologicalComputationMI(data)

	q := 0.0
	for _, v := range s {
		q += v
	}
	q /= float64(len(s))

	if math.Abs(r-q) > 0.00001 {
		t.Errorf("MorphologicalComputationMI should be %f but is %f", r, q)
	}
}
