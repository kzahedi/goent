package state_test

import (
	"fmt"
	"testing"

	"github.com/kzahedi/goent/discrete"
)

func TestConditionalEntropy(t *testing.T) {
	t.Log("Testing Conditional Entropy")
	p1 := [][]float64{
		{1.0 / 16.0, 1.0 / 16.0, 1.0 / 16.0, 1.0 / 16.0},
		{1.0 / 16.0, 1.0 / 16.0, 1.0 / 16.0, 1.0 / 16.0},
		{1.0 / 16.0, 1.0 / 16.0, 1.0 / 16.0, 1.0 / 16.0},
		{1.0 / 16.0, 1.0 / 16.0, 1.0 / 16.0, 1.0 / 16.0}}

	if r := discrete.ConditionalEntropy2(p1); r != 2.0 {
		t.Errorf(fmt.Sprintf("Conditional entropy information of uniform distribution must be 2.0 (4 states) but it is %f", r))
	}

	p2 := [][]float64{
		{1.0 / 4.0, 0.0, 0.0, 0.0},
		{0.0, 1.0 / 4.0, 0.0, 0.0},
		{0.0, 0.0, 1.0 / 4.0, 0.0},
		{0.0, 0.0, 0.0, 1.0 / 4.0}}

	if r := discrete.ConditionalEntropy2(p2); r != 0.0 {
		t.Errorf(fmt.Sprintf("Conditional entropy of deterministic distribution must be 0.0 but it is %f", r))
	}

}
