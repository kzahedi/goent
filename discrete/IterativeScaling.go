package discrete

import (
	"math"
	"sort"

	"gonum.org/v1/gonum/stat"
)

// IterativeScaling contains the configuration and data required
// for the iterative scaling algorithm
type IterativeScaling struct {
	PTarget             []float64
	PEstimate           []float64
	Features            map[string][]int
	NrOfIterations      int
	ErrorThreshold      float64
	Alphabet            [][]int
	NrOfStates          []int
	NrOfVariables       int
	CurrentFeatureIndex int
	CurrentIteration    int
	LastKLStep          float64
	Keys                []string
}

// NewIterativeScaling Creates a new struct
func NewIterativeScaling() *IterativeScaling {
	return &IterativeScaling{
		Alphabet:            nil,
		CurrentFeatureIndex: -1,
		CurrentIteration:    0,
		ErrorThreshold:      0.0,
		Features:            nil,
		LastKLStep:          -1.0,
		NrOfIterations:      0,
		NrOfStates:          nil,
		NrOfVariables:       0,
		PEstimate:           nil,
		PTarget:             nil,
		Keys:                nil,
	}
}

// Init extract the feature names for faster processing
func (data *IterativeScaling) Init() {
	data.Keys = make([]string, 0, len(data.Features))
	for k := range data.Features {
		data.Keys = append(data.Keys, k)
	}
	for _, k := range data.Keys {
		v := data.Features[k]
		sort.Ints(v)
		data.Features[k] = v
	}

	data.CurrentFeatureIndex = -1
	data.PEstimate = make([]float64, len(data.PTarget), len(data.PTarget))
	for i := range data.PTarget {
		data.PEstimate[i] = 1.0 / float64(len(data.PTarget))
	}
}

// CreateAlphabet creates the alphabet given NrOfStates and NrOfVariables
func (data *IterativeScaling) CreateAlphabet() {

	n := 1
	for i := 0; i < data.NrOfVariables; i++ {
		n = n * data.NrOfStates[i]
	}
	data.Alphabet = make([][]int, n, n)

	for i := 0; i < n; i++ {
		data.Alphabet[i] = make([]int, data.NrOfVariables, data.NrOfVariables)
	}

	nrvi := data.NrOfVariables
	nrsf := 0.0

	for i := 0; i < n; i++ {
		for j := 0; j < nrvi; j++ {
			nrsf = float64(data.NrOfStates[j])
			b := int(math.Pow(nrsf, float64(j)))
			w := int((i / b) % data.NrOfStates[j])
			// data.Alphabet[i][nrvi-j-1] = w
			data.Alphabet[i][j] = w
		}
	}
}

// Iterate implements the iterative scaling algorithm as described in
// I. Csiszar. i-divergence geometry of probability distributions and minimization
// problems. Ann. Probab., 3(1):146–158, 02 1975.
// Input is a probability distribution, a feature set, and a number of iterations.
// The output is the maximum entropy estimation of p given the feature set
// p_est^(n+1)(x) = p_target(x_a) * p_est^(n)(x_{without a}|x_a)
// where a is cycled through the list of features
// We calculate it in the following form:
// p_est^(n+1)(x) =  p_est^(n)(x) * p_target(x_a) / p_est(x_a)
func (data *IterativeScaling) Iterate() {
	data.CurrentFeatureIndex++
	data.CurrentFeatureIndex = data.CurrentFeatureIndex % len(data.Features)
	pCopy := make([]float64, len(data.PEstimate), len(data.PEstimate))
	copy(pCopy, data.PEstimate) // for step with calculation with Kullback-Leibler
	f := data.Features[data.Keys[data.CurrentFeatureIndex]]

	a := 0.0
	b := 0.0
	var indices []int
	for i := range data.PEstimate {
		indices = getAlphabetIndices(i, f, &data.Alphabet)
		b = calculateMarginal(data.PEstimate, indices)
		if b > 0 {
			a = calculateMarginal(data.PTarget, indices)
			data.PEstimate[i] = data.PEstimate[i] * a / b
		}
	}
	data.LastKLStep = stat.KullbackLeibler(pCopy, data.PEstimate)
}

func calculateMarginal(p []float64, indices []int) float64 {
	sum := 0.0
	if len(indices) > 0 {
		for _, v := range indices {
			sum += p[v]
		}
	}
	return sum
}

func getAlphabetIndices(index int, feature []int, alphabet *[][]int) []int {
	a := (*alphabet)[index]
	var indices []int
	for i, v := range *alphabet {
		if checkFeatureAlphabet(feature, a, v) == true {
			indices = append(indices, i)
		}
	}
	return indices
}

func checkFeatureAlphabet(feature, values, alphabet []int) bool {
	for _, v := range feature {
		if values[v] != alphabet[v] {
			return false
		}
	}
	return true
}

// CalculateMarginalProbability calculates the marginal probability p(x)
func (data *IterativeScaling) CalculateMarginalProbability(feature []int) float64 {
	var indices []int
	sum := 0.0
	for i := range data.PEstimate {
		indices = getAlphabetIndices(i, feature, &data.Alphabet)
		sum += calculateMarginal(data.PEstimate, indices)
	}
	return sum
}
