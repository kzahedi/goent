package goent_test

import (
	"testing"

	"github.com/kzahedi/goent"
)

func TestEntropy(t *testing.T) {
	t.Log("Testing Entropy")
	p1 := []float64{0.5, 0.5, 0.5, 0.5}

	if r := goent.Entropy2(p1); r != 2.0 {
		t.Errorf("Entropy of four state uniform distribution should be 2.0 but it is ", r)
	}

	p2 := []float64{1.0, 0.0, 0.0, 0.0}

	if r := goent.Entropy2(p2); r != 0.0 {
		t.Errorf("Entropy of deterministic distribution should be 0.0 but it is ", r)
	}
}
