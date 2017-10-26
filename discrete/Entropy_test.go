package discrete_test

import (
	"fmt"
	"math"
	"math/rand"
	"testing"

	"github.com/kzahedi/goent/discrete"
)

func TestEntropyBase2(t *testing.T) {
	t.Log("Testing Entropy")
	p1 := []float64{0.5, 0.5, 0.5, 0.5}

	if r := discrete.EntropyBase2(p1); r != 2.0 {
		t.Errorf(fmt.Sprintf("Entropy of four state uniform distribution should be 2.0 but it is %f", r))
	}

	p2 := []float64{1.0, 0.0, 0.0, 0.0}

	if r := discrete.EntropyBase2(p2); r != 0.0 {
		t.Errorf(fmt.Sprintf("Entropy of deterministic distribution should be 0.0 but it is %f", r))
	}
}

func TestEntropyBaseE(t *testing.T) {
	t.Log("Testing Entropy")
	p1 := []float64{0.5, 0.5, 0.5, 0.5}

	if r := discrete.EntropyBaseE(p1); math.Abs(r-1.386294) > 0.00001 {
		t.Errorf(fmt.Sprintf("Entropy of four state uniform distribution should be 1.386294 but it is %f", r))
	}

	p2 := []float64{1.0, 0.0, 0.0, 0.0}

	if r := discrete.EntropyBaseE(p2); r != 0.0 {
		t.Errorf(fmt.Sprintf("Entropy of deterministic distribution should be 0.0 but it is %f", r))
	}
}

func TestEntropyChaoShenBaseE(t *testing.T) {
	t.Log("Testing Chao-Shen Entropy")
	r := 0.0
	for i := 0; i < 100; i++ {
		h := make([]int, 5000, 5000)
		for j := 0; j < 5000; j++ {
			h[j] = int(rand.Int63n(100))
		}
		r += discrete.EntropyChaoShenBaseE(h)
	}

	r /= 100.0

	if math.Abs(r-4.595091) > 0.1 {
		t.Errorf("Entropy should be 4.595091 and not %f", r)
	}

}

func TestEntropyChaoShenBase2(t *testing.T) {
	t.Log("Testing Chao-Shen Entropy")
	r := 0.0
	for i := 0; i < 100; i++ {
		h := make([]int, 5000, 5000)
		for j := 0; j < 5000; j++ {
			h[j] = int(rand.Int63n(100))
		}
		r += discrete.EntropyChaoShenBase2(h)
	}

	r /= 100.0

	if math.Abs(r-6.629297) > 0.1 {
		t.Errorf("Entropy should be 6.629297 and not %f", r)
	}

}

func TestEntropyMLBCBaseE(t *testing.T) {
	t.Log("Testing Maximum Likelihood Bias Corrected")
	r := 0.0
	for i := 0; i < 100; i++ {
		h := make([]int, 5000, 5000)
		for j := 0; j < 5000; j++ {
			h[j] = int(rand.Int63n(100))
		}
		r += discrete.EntropyMLBCBaseE(h)
	}

	r /= 100.0

	if math.Abs(r-4.604982) > 0.1 {
		t.Errorf("Entropy should be 4.604982 and not %f", r)
	}
}

func TestEntropyMLBCBase2(t *testing.T) {
	t.Log("Testing Maximum Likelihood Bias Corrected")
	r := 0.0
	for i := 0; i < 100; i++ {
		h := make([]int, 5000, 5000)
		for j := 0; j < 5000; j++ {
			h[j] = int(rand.Int63n(100))
		}
		r += discrete.EntropyMLBCBase2(h)
	}

	r /= 100.0

	if math.Abs(r-6.639633) > 0.1 {
		t.Errorf("Entropy should be 6.639633 and not %f", r)
	}
}

func TestEntropyHorvitzThompsonBase2(t *testing.T) {
	t.Log("Testing Horvitz-Thompson Base 2")
	r := 0.0
	for i := 0; i < 100; i++ {
		h := make([]int, 5000, 5000)
		for j := 0; j < 5000; j++ {
			h[j] = int(rand.Int63n(100))
		}
		r += discrete.EntropyHorvitzThompsonBase2(h)
	}

	r /= 100.0

	if math.Abs(r-4.595013) > 0.1 {
		t.Errorf("Entropy should be 4.595013 and not %f", r)
	}
}

func TestEntropyHorvitzThompsonBaseE(t *testing.T) {
	t.Log("Testing Horvitz-Thompson Base 2")
	r := 0.0
	for i := 0; i < 100; i++ {
		h := make([]int, 5000, 5000)
		for j := 0; j < 5000; j++ {
			h[j] = int(rand.Int63n(100))
		}
		r += discrete.EntropyHorvitzThompsonBaseE(h)
	}

	r /= 100.0

	if math.Abs(r-4.595032) > 0.1 {
		t.Errorf("Entropy should be 4.595032 and not %f", r)
	}
}
