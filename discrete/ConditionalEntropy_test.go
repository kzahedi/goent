package discrete_test

import (
	"math"
	"testing"

	"github.com/kzahedi/goent/discrete"
	"gonum.org/v1/gonum/mat"
)

func TestConditionalEntropy(t *testing.T) {
	t.Log("Testing Conditional Entropy")
	p1 := mat.NewDense(4, 4, []float64{
		1.0 / 16.0, 1.0 / 16.0, 1.0 / 16.0, 1.0 / 16.0,
		1.0 / 16.0, 1.0 / 16.0, 1.0 / 16.0, 1.0 / 16.0,
		1.0 / 16.0, 1.0 / 16.0, 1.0 / 16.0, 1.0 / 16.0,
		1.0 / 16.0, 1.0 / 16.0, 1.0 / 16.0, 1.0 / 16.0})

	if r := discrete.ConditionalEntropy(p1); math.Abs(r-1.386294) > 0.0001 {
		t.Errorf("Conditional entropy (Base E) of uniform distribution must be 1.386294 (4 states) but it is %f", r)
	}

	p2 := mat.NewDense(4, 4, []float64{
		1.0 / 4.0, 0.0, 0.0, 0.0,
		0.0, 1.0 / 4.0, 0.0, 0.0,
		0.0, 0.0, 1.0 / 4.0, 0.0,
		0.0, 0.0, 0.0, 1.0 / 4.0})

	if r := discrete.ConditionalEntropy(p2); r != 0.0 {
		t.Errorf("Conditional entropy of deterministic distribution must be 0.0 but it is %f", r)
	}

}
