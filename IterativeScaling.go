package goent

import (
	"math"
	"sort"

	stat "gonum.org/v1/gonum/stat"
)

type IterativeScaling struct {
	P_target              []float64
	P_estimate            []float64
	Features              map[string][]int
	Nr_of_iterations      int
	Error_treshold        float64
	Alphabet              [][]int
	Nr_of_states          []int
	Nr_of_variables       int
	Current_feature_index int
	Current_iteration     int
	Last_KL_step          float64
	Keys                  []string
}

// NewIterativeScaling Creates a new struct
func NewIterativeScaling() *IterativeScaling {
	return &IterativeScaling{
		Alphabet:              nil,
		Current_feature_index: -1,
		Current_iteration:     0,
		Error_treshold:        0.0,
		Features:              nil,
		Last_KL_step:          -1.0,
		Nr_of_iterations:      0,
		Nr_of_states:          nil,
		Nr_of_variables:       0,
		P_estimate:            nil,
		P_target:              nil,
		Keys:                  nil,
	}
}

// Init extract the feature names for faster processing
func (data *IterativeScaling) Init() {
	data.Keys = make([]string, 0, len(data.Features))
	for k, _ := range data.Features {
		data.Keys = append(data.Keys, k)
	}
	for _, k := range data.Keys {
		v := data.Features[k]
		sort.Ints(v)
		data.Features[k] = v
	}

	data.Current_feature_index = -1
	data.P_estimate = make([]float64, len(data.P_target), len(data.P_target))
	for i, _ := range data.P_target {
		data.P_estimate[i] = 1.0 / float64(len(data.P_target))
	}
}

// CreateAlphabet creates the alphabet given Nr_of_states and Nr_of_variables
func (data *IterativeScaling) CreateAlphabet() {

	n := 1
	for i := 0; i < data.Nr_of_variables; i++ {
		n = n * data.Nr_of_states[i]
	}
	data.Alphabet = make([][]int, n, n)

	for i := 0; i < n; i++ {
		data.Alphabet[i] = make([]int, data.Nr_of_variables, data.Nr_of_variables)
	}

	nrvi := data.Nr_of_variables
	nrsf := 0.0

	for i := 0; i < n; i++ {
		for j := 0; j < nrvi; j++ {
			nrsf = float64(data.Nr_of_states[j])
			b := int(math.Pow(nrsf, float64(j)))
			w := int((i / b) % data.Nr_of_states[j])
			data.Alphabet[i][nrvi-j-1] = w
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
	data.Current_feature_index++
	data.Current_feature_index = data.Current_feature_index % len(data.Features)
	p_copy := make([]float64, len(data.P_estimate), len(data.P_estimate))
	copy(p_copy, data.P_estimate) // for step with calculation with Kullback-Leibler
	f := data.Features[data.Keys[data.Current_feature_index]]

	a := 0.0
	b := 0.0
	var indices []int
	for i, _ := range data.P_estimate {
		indices = Get_alphabet_indices(i, f, &data.Alphabet)
		b = calculate_marginal(data.P_estimate, indices)
		if b > 0 {
			a = calculate_marginal(data.P_target, indices)
			data.P_estimate[i] = data.P_estimate[i] * a / b
		}
	}
	data.Last_KL_step = stat.KullbackLeibler(p_copy, data.P_estimate)
}

func calculate_marginal(p []float64, indices []int) float64 {
	sum := 0.0
	if len(indices) > 0 {
		for _, v := range indices {
			sum += p[v]
		}
	}
	return sum
}

func Get_alphabet_indices(index int, feature []int, alphabet *[][]int) []int {
	a := (*alphabet)[index]
	var indices []int
	for i, v := range *alphabet {
		if Check_feature_alphabet(feature, a, v) == true {
			indices = append(indices, i)
		}
	}
	return indices
}

func Check_feature_alphabet(feature, values, alphabet []int) bool {
	for _, v := range feature {
		if values[v] != alphabet[v] {
			return false
		}
	}
	return true
}

func (data *IterativeScaling) CalculateMarginalProbability(feature []int) float64 {
	var indices []int
	sum := 0.0
	for i, _ := range data.P_estimate {
		indices = Get_alphabet_indices(i, feature, &data.Alphabet)
		sum += calculate_marginal(data.P_estimate, indices)
	}
	return sum
}
