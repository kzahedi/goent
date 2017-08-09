package goent_test

import (
	"testing"

	"gonum.org/v1/gonum/mat"

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

func TestMutualInformation(t *testing.T) {
	t.Log("Testing Mutual Information")
	p1 := [][]float64{
		{1.0 / 16.0, 1.0 / 16.0, 1.0 / 16.0, 1.0 / 16.0},
		{1.0 / 16.0, 1.0 / 16.0, 1.0 / 16.0, 1.0 / 16.0},
		{1.0 / 16.0, 1.0 / 16.0, 1.0 / 16.0, 1.0 / 16.0},
		{1.0 / 16.0, 1.0 / 16.0, 1.0 / 16.0, 1.0 / 16.0}}

	if r := goent.MutualInformation2(p1); r != 0.0 {
		t.Errorf("Mutual information of uniform distribution must be 0.0 but it is ", r)
	}

	p2 := [][]float64{
		{1.0 / 4.0, 1.0 / 0.0, 1.0 / 0.0, 1.0 / 0.0},
		{1.0 / 0.0, 1.0 / 4.0, 1.0 / 0.0, 1.0 / 0.0},
		{1.0 / 0.0, 1.0 / 0.0, 1.0 / 4.0, 1.0 / 0.0},
		{1.0 / 0.0, 1.0 / 0.0, 1.0 / 0.0, 1.0 / 4.0}}

	if r := goent.MutualInformation2(p2); r != 0.5 {
		t.Errorf("Mutual information of deterministic distribution must be 0.5 but it is ", r)
	}

}

func TestRowSum(t *testing.T) {
	t.Log("Testing RowSum")
	m := mat.NewDense(5, 10, nil)
	var index float64

	for i := 0; i < 5; i++ {
		for j := 0; j < 10; j++ {
			m.Set(i, j, index)
			index++
		}
	}

	rs := goent.RowSum(m)

	if len(rs) != 5 {
		t.Errorf("Wrong number of rows left after RowSum. Should be 5 but it is", len(rs))
	}

	for r := 0; r < 5; r++ {
		sum := (0.0 + 1.0 + 2.0 + 3.0 + 4.0 + 5.0 + 6.0 + 7.0 + 8.0 + 9.0) + float64(r)*100.0
		if rs[r] != sum {
			t.Errorf(string(r+1), "th row sum should be", sum, " but it is ", rs[r])
		}
	}
}

func TestColSum(t *testing.T) {
	t.Log("Testing ColSum")
	m := mat.NewDense(5, 10, nil)
	var index float64

	for i := 0; i < 5; i++ {
		for j := 0; j < 10; j++ {
			m.Set(i, j, index)
			index++
		}
	}

	cs := goent.ColSum(m)

	if len(cs) != 10 {
		t.Errorf("Wrong number of columns left after ColSum. Should be 10 but it is", len(cs))
	}

	for c := 0; c < 10; c++ {
		sum := (0.0 + 10.0 + 20.0 + 30.0 + 40.0) + float64(c)*5
		if cs[c] != sum {
			t.Errorf(string(c+1), "th column sum should be", sum, " but it is ", cs[c])
		}
	}
}
