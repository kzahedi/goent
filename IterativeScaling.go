package goent

import (
	"fmt"
	"math"
)

type IterativeScaling struct {
	P_target              []float64
	P_estimate            []float64
	Features              map[string][]int64
	Nr_of_iterations      int64
	Error_treshold        float64
	Alphabet              [][]int64
	Nr_of_states          int64
	Nr_of_variables       int64
	Current_feature_index int64
	Keys                  []string
}

// NewIterativeScaling Creates a new struct
func NewIterativeScaling() *IterativeScaling {
	return &IterativeScaling{P_target: nil,
		P_estimate:            nil,
		Features:              nil,
		Nr_of_iterations:      0,
		Error_treshold:        0.0,
		Alphabet:              nil,
		Nr_of_states:          0,
		Nr_of_variables:       0,
		Current_feature_index: -1,
		Keys: nil}
}

// Init extract the feature names for faster processing
func (data *IterativeScaling) Init() {
	data.Keys = make([]string, 0, len(data.Features))
	for k, _ := range data.Features {
		data.Keys = append(data.Keys, k)
	}
	data.Current_feature_index = -1
	data.P_estimate = make([]float64, len(data.P_target), len(data.P_target))
	for i, _ := range data.P_target {
		data.P_estimate[i] = 1.0 / float64(len(data.P_target))
	}
}

// CreateAlphabet creates the alphabet given Nr_of_states and Nr_of_variables
func (data *IterativeScaling) CreateAlphabet() {

	n := int(math.Pow(float64(data.Nr_of_states), float64(data.Nr_of_variables)))
	data.Alphabet = make([][]int64, n, n)

	for i := 0; i < n; i++ {
		data.Alphabet[i] = make([]int64, data.Nr_of_variables, data.Nr_of_variables)
	}

	nrsi := int(data.Nr_of_states)
	nrvi := int(data.Nr_of_variables)
	nrsf := float64(data.Nr_of_states)

	for i := 0; i < n; i++ {
		for j := 0; j < nrvi; j++ {
			b := int(math.Pow(nrsf, float64(j)))
			w := int64((i / b) % nrsi)
			data.Alphabet[i][nrvi-j-1] = w
		}
	}
}

// Iterate implements the iterative scaling algorithm as described in
// I. Csiszar. i-divergence geometry of probability distributions and minimization
// problems. Ann. Probab., 3(1):146–158, 02 1975.
// Input is a probability distribution, a feature set, and a number of iterations.
// The output is the maximum entropy estimation of p given the feature set
func (data *IterativeScaling) Iterate() {
	data.Current_feature_index++
	feature := data.Features[data.Keys[data.Current_feature_index]]
	fmt.Println(feature)
}
