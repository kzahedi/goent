package discrete_test

import (
	"math"
	"math/rand"
	"testing"

	"github.com/kzahedi/goent/discrete"
	stat "gonum.org/v1/gonum/stat"
)

func normalise1D(p []float64) []float64 {
	sum := 0.0
	for i := 0; i < len(p); i++ {
		sum += p[i]
	}

	for i := 0; i < len(p); i++ {
		p[i] /= sum
	}
	return p
}

func normalise2D(p [][]float64) [][]float64 {
	sum := 0.0
	for i := 0; i < len(p); i++ {
		for j := 0; j < len(p[i]); j++ {
			sum += p[i][j]
		}
	}

	for i := 0; i < len(p); i++ {
		for j := 0; j < len(p[i]); j++ {
			p[i][j] /= sum
		}
	}
	return p
}

func normalise3D(p [][][]float64) [][][]float64 {
	sum := 0.0
	for i := 0; i < len(p); i++ {
		for j := 0; j < len(p[i]); j++ {
			for k := 0; k < len(p[i][j]); k++ {
				sum += p[i][j][k]
			}
		}
	}

	for i := 0; i < len(p); i++ {
		for j := 0; j < len(p[i]); j++ {
			for k := 0; k < len(p[i][j]); k++ {
				p[i][j][k] /= sum
			}
		}
	}
	return p
}

func TestMorphologicalComputationW(t *testing.T) {
	pw2w1a1 := discrete.Create3D(10, 10, 10)
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			for k := 0; k < 10; k++ {
				pw2w1a1[i][j][k] = rand.Float64()
			}
		}
	}
	pw2w1a1 = normalise3D(pw2w1a1)

	r := discrete.MorphologicalComputationW(pw2w1a1)
	s := discrete.ConditionalMutualInformationBase2(pw2w1a1)

	if math.Abs(r-s) > 0.00001 {
		t.Errorf("MorphologicalComputationW should be %f but is %f", s, r)
	}
	if r < 0.0 {
		t.Errorf("MorphologicalComputationW should non-negativ but is %f", r)
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
	pw2a1w1 = normalise3D(pw2a1w1)

	r := discrete.MorphologicalComputationA(pw2a1w1)
	s := discrete.ConditionalMutualInformationBase2(pw2a1w1)

	if math.Abs(r-s) > 0.00001 {
		t.Errorf("MorphologicalComputationA should be %f but is %f", s, r)
	}
	if r < 0.0 {
		t.Errorf("MorphologicalComputationA should non-negativ but is %f", r)
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
	pw2w1 = normalise2D(pw2w1)
	pw2a1 = normalise2D(pw2a1)

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
	pw2w1a1 = normalise3D(pw2w1a1)

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
	pw2a1 = normalise2D(pw2a1)

	r := discrete.MorphologicalComputationWA(pw2w1a1)
	s := discrete.ConditionalMutualInformationBase2(pw2w1a1) - discrete.MutualInformationBase2(pw2a1)

	if math.Abs(r-s) > 0.00001 {
		t.Errorf("MorphologicalComputationWA should be %f but is %f", s, r)
	}
	if r < 0.0 {
		t.Errorf("MorphologicalComputationWA should non-negativ but is %f", r)
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
	pw2w1s1 = normalise3D(pw2w1s1)

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
	pw2s1 = normalise2D(pw2s1)

	r := discrete.MorphologicalComputationWS(pw2w1s1)
	s := discrete.ConditionalMutualInformationBase2(pw2w1s1) - discrete.MutualInformationBase2(pw2s1)

	if math.Abs(r-s) > 0.00001 {
		t.Errorf("MorphologicalComputationWS should be %f but is %f", s, r)
	}
	if r < 0.0 {
		t.Errorf("MorphologicalComputationWS should non-negativ but is %f", r)
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
	pw2w1 = normalise2D(pw2w1)
	pa1s1 = normalise2D(pa1s1)

	r := discrete.MorphologicalComputationMI(pw2w1, pa1s1)
	s := discrete.MutualInformationBase2(pw2w1) - discrete.MutualInformationBase2(pa1s1)

	if math.Abs(r-s) > 0.00001 {
		t.Errorf("MorphologicalComputationMI should be %f but is %f", s, r)
	}
}

func TestMorphologicalComputationSyNid(t *testing.T) {
	iterations := 100
	pw2w1a1 := discrete.Create3D(10, 10, 10)
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			for k := 0; k < 10; k++ {
				pw2w1a1[i][j][k] = rand.Float64()
			}
		}
	}
	pw2w1a1 = normalise3D(pw2w1a1)

	split := discrete.IterativeScaling{}

	split.NrOfVariables = 3
	w2Dim := len(pw2w1a1)
	w1Dim := len(pw2w1a1[0])
	a1Dim := len(pw2w1a1[0][0])
	split.NrOfStates = []int{w2Dim, w1Dim, a1Dim}

	split.CreateAlphabet()

	split.PTarget = make([]float64, w2Dim*w1Dim*a1Dim, w2Dim*w1Dim*a1Dim)
	for i, a := range split.Alphabet {
		split.PTarget[i] = pw2w1a1[a[0]][a[1]][a[2]]
	}

	split.Features = make(map[string][]int)
	split.Features["W,W'"] = []int{0, 2}
	split.Features["A,W'"] = []int{1, 2}

	split.Init()
	for i := 0; i < iterations; i++ {
		split.Iterate()
	}

	r := stat.KullbackLeibler(split.PTarget, split.PEstimate) / math.Log(2)
	s := discrete.MorphologicalComputationSyNid(pw2w1a1, iterations)

	if math.Abs(r-s) > 0.001 {
		t.Errorf("MorphologicalComputationSyNid should be %f but is %f", r, s)
	}
	if r < 0.0 {
		t.Errorf("MorphologicalComputationSyNid should non-negativ but is %f", r)
	}
}

func TestMorphologicalComputationSY(t *testing.T) {
	iterations := 100
	pw2w1a1 := discrete.Create3D(10, 10, 10)
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			for k := 0; k < 10; k++ {
				pw2w1a1[i][j][k] = rand.Float64()
			}
		}
	}
	pw2w1a1 = normalise3D(pw2w1a1)

	split := discrete.IterativeScaling{}

	split.NrOfVariables = 3
	w2Dim := len(pw2w1a1)
	w1Dim := len(pw2w1a1[0])
	a1Dim := len(pw2w1a1[0][0])
	split.NrOfStates = []int{w2Dim, w1Dim, a1Dim}

	split.CreateAlphabet()

	split.PTarget = make([]float64, w2Dim*w1Dim*a1Dim, w2Dim*w1Dim*a1Dim)
	for i, a := range split.Alphabet {
		split.PTarget[i] = pw2w1a1[a[0]][a[1]][a[2]]
	}

	split.Features = make(map[string][]int)
	split.Features["W,W'"] = []int{0, 2}
	split.Features["A,W'"] = []int{1, 2}
	split.Features["W,A"] = []int{0, 1}

	split.Init()
	for i := 0; i < iterations; i++ {
		split.Iterate()
	}

	r := stat.KullbackLeibler(split.PTarget, split.PEstimate) / math.Log(2)
	s := discrete.MorphologicalComputationSY(pw2w1a1, iterations, false)

	if math.Abs(r-s) > 0.001 {
		t.Errorf("MorphologicalComputationSY should be %f but is %f", r, s)
	}
	if r < 0.0 {
		t.Errorf("MorphologicalComputationSY should non-negativ but is %f", r)
	}

}

func TestMorphologicalComputationWp(t *testing.T) {
	iterations := 100
	pw2w1a1 := discrete.Create3D(10, 10, 10)
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			for k := 0; k < 10; k++ {
				pw2w1a1[i][j][k] = rand.Float64()
			}
		}
	}
	pw2w1a1 = normalise3D(pw2w1a1)

	r := discrete.MorphologicalComputationW(pw2w1a1) - discrete.MorphologicalComputationSY(pw2w1a1, iterations, false)
	s := discrete.MorphologicalComputationWp(pw2w1a1, iterations, false)

	if math.Abs(r-s) > 0.001 {
		t.Errorf("MorphologicalComputationP should be %f but is %f", r, s)
	}
}
