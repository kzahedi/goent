package goent_test

import (
	"fmt"
	"testing"

	"github.com/kzahedi/goent"
)

// alphabet         [][]int64

func TestIterativeScalingIAND(t *testing.T) {
	data := goent.IterativeScalingData{}

	data.Nr_of_variables = 3
	data.Nr_of_states = 3

	// data.P_target = make([]float64, 8, 8)
	// data.P_target[0] = 0.5
	// data.P_target[4] = 0.25
	// data.P_target[6] = 0.25

	// data.Features = make(map[string][]int64)
	// data.Features["X,Z"] = []int64{1, 2, 5, 6}
	// data.Features["Y,Z"] = []int64{3, 4, 5, 6}
	// data.Features["X,Y"] = []int64{1, 2, 3, 4}

	data.Nr_of_iterations = 10
	data.Error_treshold = 0.0

	goent.IterativeScalingCreateAlphabet(&data)

	// 0 0 0
	// 0 1 0
	// 1 0 0
	// 1 1 1

	goent.IterativeScaling(&data)
	fmt.Println(data)
}
