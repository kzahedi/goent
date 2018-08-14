package discrete_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/kzahedi/goent/discrete"
)

func TestMutualInformation(t *testing.T) {
	t.Log("Testing Mutual Information")
	p1 := mat64.NewMatrix(4, 4, [][]float64{
		{1.0 / 16.0, 1.0 / 16.0, 1.0 / 16.0, 1.0 / 16.0},
		{1.0 / 16.0, 1.0 / 16.0, 1.0 / 16.0, 1.0 / 16.0},
		{1.0 / 16.0, 1.0 / 16.0, 1.0 / 16.0, 1.0 / 16.0},
		{1.0 / 16.0, 1.0 / 16.0, 1.0 / 16.0, 1.0 / 16.0}})

	if r := discrete.MutualInformation(p1); r != 0.0 {
		t.Errorf(fmt.Sprintf("Mutual information (base e) of uniform distribution must be 0.0 (4 states) but it is %f", r))
	}

	p2 := mat64.NewMatrix(4, 4, [][]float64{
		{1.0 / 4.0, 0.0, 0.0, 0.0},
		{0.0, 1.0 / 4.0, 0.0, 0.0},
		{0.0, 0.0, 1.0 / 4.0, 0.0},
		{0.0, 0.0, 0.0, 1.0 / 4.0}})

	if r := discrete.MutualInformation(p2); math.Abs(r-1.386) > 0.001 {
		t.Errorf(fmt.Sprintf("Mutual information of deterministic distribution must be 1.386 but it is %f", r))
	}
}
