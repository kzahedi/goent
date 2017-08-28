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

	N := 1000

	var xy [][]float64

	xIndex := make([]int64, 1, 1)
	xIndex[0] = 0
	yIndex := make([]int64, 1, 1)
	yIndex[0] = 1

	for i := 0; i < N; i++ {
		xyd := make([]float64, 2, 2)
		xyd[0] = rand.Float64()
		xyd[1] = rand.Float64()
		xy = append(xy, xyd)
	}

	if r := math.Abs(continuous.KraskovStoegbauerGrassberger1(xy, xIndex, yIndex, 5)); r < 0.0001 {
		t.Errorf(fmt.Sprintf("Mutual information should be close to be 0.0 but it is %f", r))
	}
}

func TestKraskovStoegbauerGrassberger1Gaussian(t *testing.T) {
	t.Log("Testing KraskovStoegbauerGrassberger1 against Gaussian distribution")

	N := 10000

	r := 0.9 // co-variance

	IGauss := -(1.0 / 2.0) * math.Log(1.0-r*r)

	mu := []float64{0.0, 0.0}
	sym := mat.NewSymDense(2, []float64{1.0, r, r, 1.0})
	normal, _ := distmv.NewNormal(mu, sym, nil)

	var xy [][]float64

	xIndex := make([]int64, 1, 1)
	xIndex[0] = 0
	yIndex := make([]int64, 1, 1)
	yIndex[0] = 1

	for i := 0; i < N; i++ {
		xyd := normal.Rand(nil)
		xy = append(xy, xyd)
	}

	mi := math.Abs(continuous.KraskovStoegbauerGrassberger1(xy, xIndex, yIndex, 400))

	if math.Abs(mi-IGauss) > 0.01 {
		t.Errorf(fmt.Sprintf("Mutual information should be close %f but it is %f", IGauss, mi))
	}
}

func TestKraskovStoegbauerGrassberger2Independent(t *testing.T) {
	t.Log("Testing KraskovStoegbauerGrassberger1 against Gaussian")

	N := 1000

	var xy [][]float64

	xIndex := make([]int64, 1, 1)
	xIndex[0] = 0
	yIndex := make([]int64, 1, 1)
	yIndex[0] = 1

	for i := 0; i < N; i++ {
		xyd := make([]float64, 2, 2)
		xyd[0] = rand.Float64()
		xyd[1] = rand.Float64()
		xy = append(xy, xyd)
	}

	if r := math.Abs(continuous.KraskovStoegbauerGrassberger2(xy, xIndex, yIndex, 5)); r < 0.0001 {
		t.Errorf(fmt.Sprintf("Mutual information should be close to be 0.0 but it is %f", r))
	}
}

func TestKraskovStoegbauerGrassberger2Gaussian(t *testing.T) {
	t.Log("Testing KraskovStoegbauerGrassberger2 against Gaussian distribution")

	N := 10000

	r := 0.9 // co-variance

	IGauss := -(1.0 / 2.0) * math.Log(1.0-r*r)

	mu := []float64{0.0, 0.0}
	sym := mat.NewSymDense(2, []float64{1.0, r, r, 1.0})
	normal, _ := distmv.NewNormal(mu, sym, nil)

	var xy [][]float64

	xIndex := make([]int64, 1, 1)
	xIndex[0] = 0
	yIndex := make([]int64, 1, 1)
	yIndex[0] = 1

	for i := 0; i < N; i++ {
		xyd := normal.Rand(nil)
		xy = append(xy, xyd)
	}

	mi := math.Abs(continuous.KraskovStoegbauerGrassberger2(xy, xIndex, yIndex, 400))

	if math.Abs(mi-IGauss) > 0.01 {
		t.Errorf(fmt.Sprintf("Mutual information should be close %f but it is %f", IGauss, mi))
	}
}
