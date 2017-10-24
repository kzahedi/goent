package discrete_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/kzahedi/goent/discrete"
	stat "gonum.org/v1/gonum/stat"
)

// alphabet         [][]int

func TestIterativeScalingAND(t *testing.T) {
	split := discrete.IterativeScaling{}

	split.NrOfVariables = 3
	split.NrOfStates = make([]int, 3, 3)
	split.NrOfStates[0] = 2
	split.NrOfStates[1] = 2
	split.NrOfStates[2] = 2

	split.PTarget = make([]float64, 8, 8)
	split.PTarget[0] = 0.5
	split.PTarget[4] = 0.25
	split.PTarget[6] = 0.25

	split.Features = make(map[string][]int)
	split.Features["X,Z"] = []int{0, 2}
	split.Features["Y,Z"] = []int{1, 2}
	split.Features["X,Y"] = []int{0, 1}

	split.CreateAlphabet()
	split.Init()
	for i := 0; i < 100; i++ {
		split.Iterate()
	}

	r := stat.KullbackLeibler(split.PTarget, split.PEstimate) / math.Log(2)

	if r > 0.0001 {
		t.Errorf(fmt.Sprintf("AND should be 0 but it is %f", r))
	}
}

func TestIterativeScalingOR(t *testing.T) {
	split := discrete.IterativeScaling{}

	split.NrOfVariables = 3
	split.NrOfStates = make([]int, 3, 3)
	split.NrOfStates[0] = 2
	split.NrOfStates[1] = 2
	split.NrOfStates[2] = 2

	split.PTarget = make([]float64, 8, 8)
	split.PTarget[0] = 0.25
	split.PTarget[4] = 0.25
	split.PTarget[6] = 0.5

	split.Features = make(map[string][]int)
	split.Features["X,Z"] = []int{0, 2}
	split.Features["Y,Z"] = []int{1, 2}
	split.Features["X,Y"] = []int{0, 1}

	split.CreateAlphabet()
	split.Init()
	for i := 0; i < 100; i++ {
		split.Iterate()
	}

	r := stat.KullbackLeibler(split.PTarget, split.PEstimate) / math.Log(2)

	if r > 0.0001 {
		t.Errorf(fmt.Sprintf("AND should be 0 but it is %f", r))
	}
}

func TestIterativeScalingXOR(t *testing.T) {
	split := discrete.IterativeScaling{}

	split.NrOfVariables = 3
	split.NrOfStates = make([]int, 3, 3)
	split.NrOfStates[0] = 2
	split.NrOfStates[1] = 2
	split.NrOfStates[2] = 2

	split.PTarget = make([]float64, 8, 8)
	split.PTarget[0] = 0.25
	split.PTarget[3] = 0.25
	split.PTarget[5] = 0.25
	split.PTarget[6] = 0.25

	split.Features = make(map[string][]int)
	split.Features["X,Z"] = []int{0, 2}
	split.Features["Y,Z"] = []int{1, 2}
	// split.Features["X,Y"] = []int{0, 1}

	split.CreateAlphabet()
	split.Init()
	for i := 0; i < 100; i++ {
		split.Iterate()
	}

	r := stat.KullbackLeibler(split.PTarget, split.PEstimate) / math.Log(2)

	if math.Abs(r-1.0) > 0.00001 {
		t.Errorf(fmt.Sprintf("XOR should be 1 but it is %f", r))
	}
}

func TestConstructor(t *testing.T) {
	isPtr := discrete.NewIterativeScaling()

	if *isPtr.Alphabet != nil {
		t.Errorf("Alphabet should be nil")
	}

	if *isPtr.CurrentFeatureIndex != -1 {
		t.Errorf("CurrentFeatureIndex should be -1")
	}

	if *isPtr.CurrentIteration != 0 {
		t.Errorf("CurrentIteration should be 0")
	}

	if *isPtr.ErrorThreshold != 0.0 {
		t.Errorf("ErrorThreshold should be 0.0")
	}

	if *isPtr.Features != nil {
		t.Errorf("Features should be nil")
	}

	if *isPtr.LastKLStep != -1.0 {
		t.Errorf("LastKLStep should be -1.0")
	}

	if *isPtr.NrOfIterations != 0 {
		t.Errorf("NrOfIterations should be 0")
	}

	if *isPtr.NrOfStates != nil {
		t.Errorf("NrOfStates should be nil")
	}

	if *isPtr.NrOfVariables != 0 {
		t.Errorf("NrOfVariables should be 0")
	}

	if *isPtr.PEstimate != nil {
		t.Errorf("PEstimate should be nil")
	}

	if *isPtr.PTarget != nil {
		t.Errorf("PTarget should be nil")
	}

	if *isPtr.Keys != nil {
		t.Errorf("Keys should be nil")
	}
}
