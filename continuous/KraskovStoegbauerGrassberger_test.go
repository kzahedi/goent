package continuous_test

import (
	"fmt"
	"math"
	"math/rand"
	"testing"

	"github.com/kzahedi/goent/continuous"
	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/gonum/stat/distmv"
)

func TestKraskovStoegbauerGrassberger1Independent(t *testing.T) {
	t.Log("Testing KraskovStoegbauerGrassberger1 against independent distribution")
	rand.Seed(2)

	N := 1000

	var xy [][]float64

	xIndex := make([]int, 1, 1)
	xIndex[0] = 0
	yIndex := make([]int, 1, 1)
	yIndex[0] = 1

	for i := 0; i < N; i++ {
		xyd := make([]float64, 2, 2)
		xyd[0] = rand.Float64()
		xyd[1] = rand.Float64()
		xy = append(xy, xyd)
	}

	if r := math.Abs(continuous.KraskovStoegbauerGrassberger1(xy, xIndex, yIndex, 30, false)); r > 0.1 {
		t.Errorf(fmt.Sprintf("Mutual information should be close to be 0.0 but it is %f", r))
	}
}

func TestKraskovStoegbauerGrassberger1Gaussian(t *testing.T) {
	t.Log("Testing KraskovStoegbauerGrassberger1 against Gaussian distribution")
	rand.Seed(1)

	N := 1000
	k := 20

	r := 0.9 // co-variance
	mi := 0.0

	IGauss := -(1.0 / 2.0) * math.Log(1.0-r*r)

	for i := 0; i < 100; i++ {
		mu := []float64{0.0, 0.0}
		sym := mat.NewSymDense(2, []float64{1.0, r, r, 1.0})
		normal, _ := distmv.NewNormal(mu, sym, nil)

		var xy [][]float64

		xIndex := make([]int, 1, 1)
		xIndex[0] = 0
		yIndex := make([]int, 1, 1)
		yIndex[0] = 1

		for i := 0; i < N; i++ {
			xyd := normal.Rand(nil)
			xy = append(xy, xyd)
		}

		mi += math.Abs(continuous.KraskovStoegbauerGrassberger1(xy, xIndex, yIndex, k, false))
	}

	mi /= 100.0

	if math.Abs(mi-IGauss) > 0.1 {
		t.Errorf(fmt.Sprintf("Mutual information should be close %f but it is %f", IGauss, mi))
	}
}

func TestKraskovStoegbauerGrassberger2Independent(t *testing.T) {
	t.Log("Testing KraskovStoegbauerGrassberger1 against Gaussian")
	rand.Seed(1)

	N := 1000

	var xy [][]float64

	xIndex := make([]int, 1, 1)
	xIndex[0] = 0
	yIndex := make([]int, 1, 1)
	yIndex[0] = 1

	for i := 0; i < N; i++ {
		xyd := make([]float64, 2, 2)
		xyd[0] = rand.Float64()
		xyd[1] = rand.Float64()
		xy = append(xy, xyd)
	}

	if r := math.Abs(continuous.KraskovStoegbauerGrassberger2(xy, xIndex, yIndex, 30, false)); r < 0.0001 {
		t.Errorf(fmt.Sprintf("Mutual information should be close to be 0.0 but it is %f", r))
	}
}

func TestKraskovStoegbauerGrassberger2Gaussian(t *testing.T) {
	t.Log("Testing KraskovStoegbauerGrassberger2 against Gaussian distribution")
	rand.Seed(2)

	N := 1000
	k := 30

	r := 0.9 // co-variance

	IGauss := -(1.0 / 2.0) * math.Log(1.0-r*r)

	mi := 0.0

	for i := 0; i < 100; i++ {
		mu := []float64{0.0, 0.0}
		sym := mat.NewSymDense(2, []float64{1.0, r, r, 1.0})
		normal, _ := distmv.NewNormal(mu, sym, nil)

		var xy [][]float64

		xIndex := make([]int, 1, 1)
		xIndex[0] = 0
		yIndex := make([]int, 1, 1)
		yIndex[0] = 1

		for i := 0; i < N; i++ {
			xyd := normal.Rand(nil)
			xy = append(xy, xyd)
		}

		mi += math.Abs(continuous.KraskovStoegbauerGrassberger2(xy, xIndex, yIndex, k, false))
	}

	mi /= 100.0

	if math.Abs(mi-IGauss) > 0.1 {
		t.Errorf(fmt.Sprintf("Mutual information should be close %f but it is %f", IGauss, mi))
	}
}
