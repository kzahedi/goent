package state_test

import (
	"fmt"
	"math"
	"math/rand"
	"testing"

	"github.com/kzahedi/goent/continuous"
	"github.com/kzahedi/goent/continuous/state"

	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/gonum/stat/distmv"
)

func TestKraskovStoegbauerGrassberger1Independent(t *testing.T) {
	t.Log("Testing KraskovStoegbauerGrassberger1 against independent distribution")
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

	r1 := continuous.KraskovStoegbauerGrassberger1(xy, xIndex, yIndex, 30, false)
	r2 := state.KraskovStoegbauerGrassberger1(xy, xIndex, yIndex, 30, false)
	r3 := 0.0

	for _, v := range r2 {
		r3 += v
	}

	if d := math.Abs(r1 - r3); d > 0.0001 {
		t.Errorf(fmt.Sprintf("Summed over states should be the same as averaged but the difference is %f (%f %f)", d, r1, r3))
	}
}

func TestKraskovStoegbauerGrassberger1Gaussian(t *testing.T) {
	t.Log("Testing KraskovStoegbauerGrassberger1 against Gaussian distribution")
	rand.Seed(1)

	N := 1000
	k := 20

	r := 0.9 // co-variance

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

	r1 := continuous.KraskovStoegbauerGrassberger1(xy, xIndex, yIndex, k, false)
	r2 := state.KraskovStoegbauerGrassberger1(xy, xIndex, yIndex, k, false)
	r3 := 0.0

	for _, v := range r2 {
		r3 += v
	}

	if d := math.Abs(r1 - r3); d > 0.0001 {
		t.Errorf(fmt.Sprintf("Summed over states should be the same as averaged but the difference is %f (%f %f)", d, r1, r3))
	}
}

func TestKraskovStoegbauerGrassberger2Independent(t *testing.T) {
	t.Log("Testing KraskovStoegbauerGrassberger2 against independent distribution")
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

	r1 := continuous.KraskovStoegbauerGrassberger2(xy, xIndex, yIndex, 30, false)
	r2 := state.KraskovStoegbauerGrassberger2(xy, xIndex, yIndex, 30, false)
	r3 := 0.0

	for _, v := range r2 {
		r3 += v
	}

	if d := math.Abs(r1 - r3); d > 0.0001 {
		t.Errorf(fmt.Sprintf("Summed over states should be the same as averaged but the difference is %f (%f %f)", d, r1, r3))
	}
}

func TestKraskovStoegbauerGrassberger2Gaussian(t *testing.T) {
	t.Log("Testing KraskovStoegbauerGrassberger2 against Gaussian distribution")
	rand.Seed(1)

	N := 1000
	k := 20

	r := 0.9 // co-variance

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

	r1 := continuous.KraskovStoegbauerGrassberger2(xy, xIndex, yIndex, k, false)
	r2 := state.KraskovStoegbauerGrassberger2(xy, xIndex, yIndex, k, false)
	r3 := 0.0

	for _, v := range r2 {
		r3 += v
	}

	if d := math.Abs(r1 - r3); d > 0.0001 {
		t.Errorf(fmt.Sprintf("Summed over states should be the same as averaged but the difference is %f (%f %f)", d, r1, r3))
	}
}
