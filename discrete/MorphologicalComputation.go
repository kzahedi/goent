package discrete

import (
	"math"

	stat "gonum.org/v1/gonum/stat"
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
func MorphologicalComputationSY(pw2w1a1 [][][]float64, iterations int) float64 {
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
	// bar := uiprogress.AddBar(iterations).AppendCompleted().PrependElapsed()
	// bar.PrependFunc(func(b *uiprogress.Bar) string {
	// return fmt.Sprintf("%s (%d/%d)", *mc, b.Current(), count)
	// })
	for i := 0; i < iterations; i++ {
		split.Iterate()
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

// MorphologicalComputationP [...]
func MorphologicalComputationP(pw2w1a1 [][][]float64, iterations int) float64 {
	return MorphologicalComputationW(pw2w1a1) - MorphologicalComputationSY(pw2w1a1, iterations)
}
