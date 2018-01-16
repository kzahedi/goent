package discrete

import (
	"math"

	stat "gonum.org/v1/gonum/stat"
	pb "gopkg.in/cheggaaa/pb.v1"
)

// MorphologicalComputationW quantifies morphological computation as the information that is contained in
// W about W' that is not contained in A. For more details, please read
// K. Zahedi and N. Ay. Quantifying morphological computation. Entropy, 15(5):1887–1915, 2013.
// http://www.mdpi.com/1099-4300/15/5/1887 (open access)
// and
// K. Ghazi-Zahedi, D. F. Haeufle, G. F. Montufar, S. Schmitt, and N. Ay. Evaluating
// morphological computation in muscle and dc-motor driven models of hopping movements.
// Frontiers in Robotics and AI, 3(42), 2016.
// http://journal.frontiersin.org/article/10.3389/frobt.2016.00042/full (open access)
func MorphologicalComputationW(pw2w1a1 [][][]float64) float64 {
	return ConditionalMutualInformationBase2(pw2w1a1)
}

// MorphologicalComputationA quantifies morphological computation as the information that is contained in
// A about W' that is not contained in W. For more details, please read
// K. Zahedi and N. Ay. Quantifying morphological computation. Entropy, 15(5):1887–1915, 2013.
// http://www.mdpi.com/1099-4300/15/5/1887 (open access)
func MorphologicalComputationA(pw2a1w1 [][][]float64) float64 {
	return ConditionalMutualInformationBase2(pw2a1w1)
}

// MorphologicalComputationCW quantifies morphological computation as the causal information flow from
// W to W' that does pass through A
// MorphologicalComputationCW = CIF(W -> W') - CIF(A -> W') = I(W';W) - I(W'|A)
func MorphologicalComputationCW(pw2w1, pw2a1 [][]float64) float64 {
	return MutualInformationBase2(pw2w1) - MutualInformationBase2(pw2a1)
}

// MorphologicalComputationWA = I(W;{W,A}) - I(W';A)
func MorphologicalComputationWA(pw2w1a1 [][][]float64) float64 {
	w2Dim := len(pw2w1a1)
	w1Dim := len(pw2w1a1[0])
	a1Dim := len(pw2w1a1[0][0])
	pw2a1 := Create2D(w2Dim, a1Dim)
	for w2 := 0; w2 < w2Dim; w2++ {
		for w1 := 0; w1 < w1Dim; w1++ {
			for a1 := 0; a1 < a1Dim; a1++ {
				pw2a1[w2][a1] += pw2w1a1[w2][w1][a1]
			}
		}
	}
	return ConditionalMutualInformationBase2(pw2w1a1) - MutualInformationBase2(pw2a1)
}

// MorphologicalComputationWS = I(W;{W,S}) - I(W';S)
func MorphologicalComputationWS(pw2w1s1 [][][]float64) float64 {
	w2Dim := len(pw2w1s1)
	w1Dim := len(pw2w1s1[0])
	s1Dim := len(pw2w1s1[0][0])
	pw2s1 := Create2D(w2Dim, s1Dim)
	for w2 := 0; w2 < w2Dim; w2++ {
		for w1 := 0; w1 < w1Dim; w1++ {
			for s1 := 0; s1 < s1Dim; s1++ {
				pw2s1[w2][s1] += pw2w1s1[w2][w1][s1]
			}
		}
	}
	return ConditionalMutualInformationBase2(pw2w1s1) - MutualInformationBase2(pw2s1)
}

// MorphologicalComputationMI quantifies morphological computation as the information that is contained in
// W about W' that is not contained in A. For more details, please read
// K. Ghazi-Zahedi, D. F. Haeufle, G. F. Montufar, S. Schmitt, and N. Ay. Evaluating
// morphological computation in muscle and dc-motor driven models of hopping movements.
// Frontiers in Robotics and AI, 3(42), 2016.
// http://journal.frontiersin.org/article/10.3389/frobt.2016.00042/full (open access)
func MorphologicalComputationMI(pw2w1 [][]float64, pa1s1 [][]float64) float64 {
	return MutualInformationBase2(pw2w1) - MutualInformationBase2(pa1s1)
}

// MorphologicalComputationSY quantifies morphological computation as the synergistic information that
// W and A contain about W'. For more details, please read
// TODO Paper reference
func MorphologicalComputationSY(pw2w1a1 [][][]float64, iterations int, eta bool) float64 {
	split := IterativeScaling{}

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
	var bar *pb.ProgressBar

	if eta == true {
		bar = pb.StartNew(iterations)
	}

	for i := 0; i < iterations; i++ {
		split.Iterate()
		if eta == true {
			bar.Increment()
		}
	}

	if eta == true {
		bar.Finish()
	}

	return stat.KullbackLeibler(split.PTarget, split.PEstimate) / math.Log(2)
}

// MorphologicalComputationSyNid quantifies morphological computation as the synergistic
// information that W and A contain about W', excluding the input distribution
// (W,A). For more details, please read
// TODO Paper reference
func MorphologicalComputationSyNid(pw2w1a1 [][][]float64, iterations int) float64 {
	split := IterativeScaling{}

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

	return stat.KullbackLeibler(split.PTarget, split.PEstimate) / math.Log(2)
}

// MorphologicalComputationWp calculates the unique information W -> W'. For more details,
// please see
// Ghazi-Zahedi, Keyan and Langer, Carlotta and Ay, Nihat,
// Morphological Computation: Synergy of Body and Brain, Entropy, 2017
func MorphologicalComputationWp(pw2w1a1 [][][]float64, iterations int, eta bool) float64 {
	return MorphologicalComputationW(pw2w1a1) - MorphologicalComputationSY(pw2w1a1, iterations, eta)
}

// MorphologicalComputationIntrinsicCA [...]
// For more details, please read
// K. Zahedi and N. Ay. Quantifying morphological computation. Entropy, 15(5):1887–1915, 2013.
// http://www.mdpi.com/1099-4300/15/5/1887 (open access)
func MorphologicalComputationIntrinsicCA(ps2s1a1 [][][]float64, sbins int) float64 {
	s2Dim := len(ps2s1a1)
	s1Dim := len(ps2s1a1[0])
	a1Dim := len(ps2s1a1[0][0])
	ps1a1 := Create2D(s1Dim, a1Dim)
	ps2doa1 := Create2D(s1Dim, a1Dim)
	ps2dos1 := Create2D(s1Dim, a1Dim)
	pa1Cs1 := Create2D(a1Dim, s1Dim)
	ps1 := make([]float64, s1Dim, s1Dim)

	for s2 := 0; s2 < s2Dim; s2++ {
		for s1 := 0; s1 < s1Dim; s1++ {
			for a1 := 0; a1 < a1Dim; a1++ {
				ps1a1[s1][a1] += ps2s1a1[s2][s1][a1]
			}
		}
	}

	for s1 := 0; s1 < s1Dim; s1++ {
		for a1 := 0; a1 < a1Dim; a1++ {
			ps1[s1] += ps1a1[s1][a1]
		}
	}

	for s1 := 0; s1 < s1Dim; s1++ {
		for a1 := 0; a1 < a1Dim; a1++ {
			pa1Cs1[a1][s1] = ps1a1[s1][a1] / ps1[s1]
		}
	}

	for s2 := 0; s2 < s2Dim; s2++ {
		for a1 := 0; a1 < a1Dim; a1++ {
			for s1 := 0; s1 < s1Dim; s1++ {
				ps2doa1[s2][a1] += ps1[s1] * ps2s1a1[s2][s1][a1] / ps1a1[s1][a1]
			}
		}
	}

	for s2 := 0; s2 < s2Dim; s2++ {
		for s1 := 0; s1 < s1Dim; s1++ {
			for a1 := 0; a1 < a1Dim; a1++ {
				ps2dos1[s2][s1] += pa1Cs1[a1][s1] * ps2doa1[s2][a1]
			}
		}
	}

	r := 0.0

	for s2 := 0; s2 < s2Dim; s2++ {
		for s1 := 0; s1 < s1Dim; s1++ {
			for a1 := 0; a1 < a1Dim; a1++ {
				r += ps1a1[s1][a1] * ps2doa1[s2][a1] * math.Log2(ps2doa1[s2][a1]/ps2dos1[s2][s1])
			}
		}
	}

	return 1.0 - r/math.Log2(float64(sbins))
}

// MorphologicalComputationIntrinsicCW [...]
// For more details, please read
// K. Zahedi and N. Ay. Quantifying morphological computation. Entropy, 15(5):1887–1915, 2013.
// http://www.mdpi.com/1099-4300/15/5/1887 (open access)
func MorphologicalComputationIntrinsicCW(ps2s1a1 [][][]float64) (r float64) {
	s2Dim := len(ps2s1a1)
	s1Dim := len(ps2s1a1[0])
	a1Dim := len(ps2s1a1[0][0])
	ps2Cs1 := Create2D(s2Dim, s1Dim)
	ps2Cs1Hat := Create2D(s2Dim, s1Dim)
	ps2Cs1a1 := Create3D(s2Dim, s1Dim, a1Dim)
	pa1Cs1 := Create2D(s2Dim, a1Dim)
	ps1a1 := Create2D(s1Dim, a1Dim)
	ps1 := make([]float64, s1Dim, s1Dim)
	pa1 := make([]float64, a1Dim, a1Dim)

	for s2 := 0; s2 < s2Dim; s2++ {
		for s1 := 0; s1 < s1Dim; s1++ {
			for a1 := 0; a1 < a1Dim; a1++ {
				ps1a1[s1][a1] += ps2s1a1[s2][s1][a1]
			}
		}
	}

	for s1 := 0; s1 < s1Dim; s1++ {
		for a1 := 0; a1 < a1Dim; a1++ {
			ps1[s1] += ps1a1[s1][a1]
			pa1[a1] += ps1a1[s1][a1]
		}
	}

	for s1 := 0; s1 < s1Dim; s1++ {
		for a1 := 0; a1 < a1Dim; a1++ {
			pa1Cs1[a1][s1] = ps1a1[s1][a1] / ps1[s1]
		}
	}

	for s2 := 0; s2 < s2Dim; s2++ {
		for s1 := 0; s1 < s1Dim; s1++ {
			for a1 := 0; a1 < a1Dim; a1++ {
				ps2Cs1a1[s2][s1][a1] = ps2s1a1[s2][s1][a1] / ps1a1[s1][a1]
			}
		}
	}

	for s2 := 0; s2 < s2Dim; s2++ {
		for s1 := 0; s1 < s1Dim; s1++ {
			for a1 := 0; a1 < a1Dim; a1++ {
				ps2Cs1[s2][s1] += ps2Cs1a1[s2][s1][a1] * pa1Cs1[a1][s1]
				for s3 := 0; s3 < s1Dim; s3++ {
					ps2Cs1Hat[s2][s1] += pa1Cs1[a1][s1] * ps2Cs1a1[s2][s3][a1] * pa1Cs1[a1][s3] * ps1[s3] / pa1[a1]
				}
			}
		}
	}

	for s2 := 0; s2 < s2Dim; s2++ {
		for s1 := 0; s1 < s1Dim; s1++ {
			r += ps2Cs1[s2][s1] * math.Log2(ps2Cs1[s2][s1]/ps2Cs1Hat[s2][s1])
		}
	}

	return r
}
