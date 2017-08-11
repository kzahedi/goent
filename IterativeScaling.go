package goent

import (
	"math"
)

type IterativeScalingData struct {
	P_target         []float64
	Features         map[string][]int64
	Nr_of_iterations int64
	Error_treshold   float64
	Alphabet         [][]int64
	Nr_of_states     int64
	Nr_of_variables  int64
}

// IterativeScalingE implements the iterative scaling algorithm as described in
// I. Csiszar. i-divergence geometry of probability distributions and minimization
// problems. Ann. Probab., 3(1):146–158, 02 1975.
// Input is a probability distribution, a feature set, and an error threshold.
// The output is the maximum entropy estimation of p given the feature set
// func IterativeScalingE(p []float64, features map[string][]int64, err float64) []float64 {
// p_est := make([]float64, len(p), len(p))
// for k, v := range features {
// fmt.Println("Key: ", k, " value: ", v)
// }
// fmt.Println(len(features))
// return p_est
// }

// IterativeScaling implements the iterative scaling algorithm as described in
// I. Csiszar. i-divergence geometry of probability distributions and minimization
// problems. Ann. Probab., 3(1):146–158, 02 1975.
// Input is a probability distribution, a feature set, and a number of iterations.
// The output is the maximum entropy estimation of p given the feature set
func IterativeScaling(data *IterativeScalingData) {
}

func IterativeScalingIterate(data *IterativeScalingData) {
}

func IterativeScalingCreateAlphabet(data *IterativeScalingData) {
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

func NewIterativeScalingData() *IterativeScalingData {
	return &IterativeScalingData{P_target: nil,
		Features:         nil,
		Nr_of_iterations: 0,
		Error_treshold:   0.0,
		Alphabet:         nil,
		Nr_of_states:     0,
		Nr_of_variables:  0}
}
