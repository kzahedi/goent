package goent_test

import (
	"math"
	"testing"

	"github.com/kzahedi/goent"
)

func TestMIasEntropies(t *testing.T) {
	t.Log("Testing Mutual Information as Entropy minus Conditional Entropy")
	p1 := [][]float64{
		{1.0 / 16.0, 1.0 / 16.0, 1.0 / 16.0, 1.0 / 16.0},
		{1.0 / 16.0, 1.0 / 16.0, 1.0 / 16.0, 1.0 / 16.0},
		{1.0 / 16.0, 1.0 / 16.0, 1.0 / 16.0, 1.0 / 16.0},
		{1.0 / 16.0, 1.0 / 16.0, 1.0 / 16.0, 1.0 / 16.0}}

	px := []float64{1.0 / 4.0, 1.0 / 4.0, 1.0 / 4.0, 1.0 / 4.0}

	mi1 := goent.MutualInformation2(p1)
	ch1 := goent.ConditionalEntropy2(p1)
	h1 := goent.Entropy2(px)
	diff1 := mi1 - (h1 - ch1)

	if math.Abs(diff1) > 0.0001 {
		t.Errorf(" I(X;Y) = H(X) - H(X|Y) but the difference is %f, MI: %f, cH: %f, H:%f", math.Abs(diff1), mi1, ch1, h1)
	}

	p2 := [][]float64{
		{1.0 / 4.0, 0.0, 0.0, 0.0},
		{0.0, 1.0 / 4.0, 0.0, 0.0},
		{0.0, 0.0, 1.0 / 4.0, 0.0},
		{0.0, 0.0, 0.0, 1.0 / 4.0}}

	mi2 := goent.MutualInformation2(p2)  // I(X;Y) = H(X) - H(X|Y)
	ch2 := goent.ConditionalEntropy2(p2) // H(X|Y)
	h2 := goent.Entropy2(px)             // H(X)
	diff2 := mi2 - (h2 - ch2)

	if math.Abs(diff2) > 0.0001 {
		t.Errorf(" I(X;Y) = H(X) - H(X|Y) but the difference is %f, MI: %f, cH: %f, H:%f", math.Abs(diff2), mi2, ch2, h2)
	}

}
