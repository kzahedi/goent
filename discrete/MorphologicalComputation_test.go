package discrete_test

import (
	"math"
	"math/rand"
	"testing"

	"github.com/kzahedi/goent/discrete"
)

func TestMorphologicalComputationW(t *testing.T) {
	pw2w1a1 := discrete.Create3D(10, 10, 10)
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			for k := 0; k < 10; k++ {
				pw2w1a1[i][j][k] = rand.Float64()
			}
		}
	}
	r := discrete.MorphologicalComputationW(pw2w1a1)
	s := discrete.ConditionalMutualInformationBase2(pw2w1a1)

	if math.Abs(r-s) > 0.00001 {
		t.Errorf("MorphologicalComputationW should be %f but is %f", s, r)
	}
}

func TestMorphologicalComputationA(t *testing.T) {

	pw2a1w1 := discrete.Create3D(10, 10, 10)
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			for k := 0; k < 10; k++ {
				pw2a1w1[i][j][k] = rand.Float64()
			}
		}
	}
	r := discrete.MorphologicalComputationA(pw2a1w1)
	s := discrete.ConditionalMutualInformationBase2(pw2a1w1)

	if math.Abs(r-s) > 0.00001 {
		t.Errorf("MorphologicalComputationA should be %f but is %f", s, r)
	}

}

func TestMorphologicalComputationCW(t *testing.T) {
	pw2w1 := discrete.Create2D(10, 10)
	pw2a1 := discrete.Create2D(10, 10)
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			pw2w1[i][j] = rand.Float64()
			pw2a1[i][j] = rand.Float64()
		}
	}
	r := discrete.MorphologicalComputationCW(pw2w1, pw2a1)
	s := discrete.MutualInformationBase2(pw2w1) - discrete.MutualInformationBase2(pw2a1)

	if math.Abs(r-s) > 0.00001 {
		t.Errorf("MorphologicalComputationCW should be %f but is %f", s, r)
	}
}

func TestMorphologicalComputationWA(t *testing.T) {

	pw2w1a1 := discrete.Create3D(10, 10, 10)
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			for k := 0; k < 10; k++ {
				pw2w1a1[i][j][k] = rand.Float64()
			}
		}
	}

	w2Dim := len(pw2w1a1)
	w1Dim := len(pw2w1a1[0])
	a1Dim := len(pw2w1a1[0][0])
	pw2a1 := discrete.Create2D(w2Dim, a1Dim)
	for w2 := 0; w2 < w2Dim; w2++ {
		for w1 := 0; w1 < w1Dim; w1++ {
			for a1 := 0; a1 < a1Dim; a1++ {
				pw2a1[w2][a1] += pw2w1a1[w2][w1][a1]
			}
		}
	}
	r := discrete.MorphologicalComputationWA(pw2w1a1)
	s := discrete.ConditionalMutualInformationBase2(pw2w1a1) - discrete.MutualInformationBase2(pw2a1)

	if math.Abs(r-s) > 0.00001 {
		t.Errorf("MorphologicalComputationWA should be %f but is %f", s, r)
	}

}

func TestMorphologicalComputationWS(t *testing.T) {
	pw2w1s1 := discrete.Create3D(10, 10, 10)
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			for k := 0; k < 10; k++ {
				pw2w1s1[i][j][k] = rand.Float64()
			}
		}
	}

	w2Dim := len(pw2w1s1)
	w1Dim := len(pw2w1s1[0])
	s1Dim := len(pw2w1s1[0][0])
	pw2s1 := discrete.Create2D(w2Dim, s1Dim)
	for w2 := 0; w2 < w2Dim; w2++ {
		for w1 := 0; w1 < w1Dim; w1++ {
			for s1 := 0; s1 < s1Dim; s1++ {
				pw2s1[w2][s1] += pw2w1s1[w2][w1][s1]
			}
		}
	}
	r := discrete.MorphologicalComputationWS(pw2w1s1)
	s := discrete.ConditionalMutualInformationBase2(pw2w1s1) - discrete.MutualInformationBase2(pw2s1)

	if math.Abs(r-s) > 0.00001 {
		t.Errorf("MorphologicalComputationWS should be %f but is %f", s, r)
	}

}

func TestMorphologicalComputationMI(t *testing.T) {
	pw2w1 := discrete.Create2D(10, 10)
	pa1s1 := discrete.Create2D(10, 10)
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			pw2w1[i][j] = rand.Float64()
			pa1s1[i][j] = rand.Float64()
		}
	}
	r := discrete.MorphologicalComputationMI(pw2w1, pa1s1)
	s := discrete.MutualInformationBase2(pw2w1) - discrete.MutualInformationBase2(pa1s1)

	if math.Abs(r-s) > 0.00001 {
		t.Errorf("MorphologicalComputationMI should be %f but is %f", s, r)
	}
}

func TestMorphologicalComputationP(t *testing.T) {
	pw2w1a1 := discrete.Create3D(10, 10, 10)
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			for k := 0; k < 10; k++ {
				pw2w1a1[i][j][k] = rand.Float64()
			}
		}
	}

	r := discrete.MorphologicalComputationP(pw2w1a1, 100, false)
	s := discrete.MorphologicalComputationW(pw2w1a1) - discrete.MorphologicalComputationSY(pw2w1a1, 100, false)

	if math.Abs(r-s) > 0.00001 {
		t.Errorf("MorphologicalComputationP should be %f but is %f", s, r)
	}
}
