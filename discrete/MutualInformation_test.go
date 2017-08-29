package discrete_test

import (
	"fmt"
	"testing"

	"github.com/kzahedi/goent/discrete"
)

func TestMutualInformation(t *testing.T) {
	t.Log("Testing Mutual Information")
	p1 := [][]float64{
		{1.0 / 16.0, 1.0 / 16.0, 1.0 / 16.0, 1.0 / 16.0},
		{1.0 / 16.0, 1.0 / 16.0, 1.0 / 16.0, 1.0 / 16.0},
		{1.0 / 16.0, 1.0 / 16.0, 1.0 / 16.0, 1.0 / 16.0},
		{1.0 / 16.0, 1.0 / 16.0, 1.0 / 16.0, 1.0 / 16.0}}

	if r := discrete.MutualInformationBase2(p1); r != 0.0 {
		t.Errorf(fmt.Sprintf("Mutual information of uniform distribution must be 0.0 (4 states) but it is %f", r))
	}

	p2 := [][]float64{
		{1.0 / 4.0, 0.0, 0.0, 0.0},
		{0.0, 1.0 / 4.0, 0.0, 0.0},
		{0.0, 0.0, 1.0 / 4.0, 0.0},
		{0.0, 0.0, 0.0, 1.0 / 4.0}}

	if r := discrete.MutualInformationBase2(p2); r != 2.0 {
		t.Errorf(fmt.Sprintf("Mutual information of deterministic distribution must be 0.5 but it is %f", r))
	}

}
