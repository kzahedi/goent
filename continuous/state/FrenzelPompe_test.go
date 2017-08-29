package state_test

import (
	"fmt"
	"math"
	"math/rand"
	"testing"

	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/gonum/stat/distmv"

	"github.com/kzahedi/goent/continuous"
	"github.com/kzahedi/goent/continuous/state"
)

func TestFrenzelPompe(t *testing.T) {
	t.Log("Testing FrenzelPompe against independent distribution")
	rand.Seed(1)

	N := 1000

	var xyz [][]float64

	xIndex := make([]int, 1, 1)
	xIndex[0] = 0
	yIndex := make([]int, 1, 1)
	yIndex[0] = 1
	zIndex := make([]int, 1, 1)
	zIndex[0] = 2

	for i := 0; i < N; i++ {
		xyzd := make([]float64, 3, 3)
		xyzd[0] = rand.Float64()
		xyzd[1] = rand.Float64()
		xyzd[2] = rand.Float64()
		xyz = append(xyz, xyzd)
	}

	rr := continuous.FrenzelPompe(xyz, xIndex, yIndex, zIndex, 5, false)
	s := state.FrenzelPompe(xyz, xIndex, yIndex, zIndex, 5, false)
	q := 0.0

	for _, v := range s {
		q += v
	}

	if d := math.Abs(q - rr); d > 0.0001 {
		t.Errorf(fmt.Sprintf("Sum over states should be equal to averaged, but the difference is %f (%f %f)", d, q, rr))
	}
}

func TestFrenzelPompeGaussian(t *testing.T) {
	t.Log("Testing FrenzelPompe against Gaussian distribution")
	rand.Seed(1)

	N := 1000
	k := 30

	r := 0.9 // co-variance

	mu := []float64{0.0, 0.0, 0.0}
	sym := mat.NewSymDense(3, []float64{
		1.0, r, r,
		r, 1.0, r,
		r, r, 1.0})
	normal, _ := distmv.NewNormal(mu, sym, nil)

	var xyz [][]float64

	xIndex := make([]int, 1, 1)
	xIndex[0] = 0
	yIndex := make([]int, 1, 1)
	yIndex[0] = 1
	zIndex := make([]int, 1, 1)
	zIndex[0] = 2

	for i := 0; i < N; i++ {
		xyzd := normal.Rand(nil)
		xyz = append(xyz, xyzd)
	}

	rr := continuous.FrenzelPompe(xyz, xIndex, yIndex, zIndex, k, false)
	s := state.FrenzelPompe(xyz, xIndex, yIndex, zIndex, k, false)
	q := 0.0

	for _, v := range s {
		q += v
	}

	if d := math.Abs(q - rr); d > 0.0001 {
		t.Errorf(fmt.Sprintf("Sum over states should be equal to averaged, but the difference is %f (%f vs. %f)", d, q, rr))
	}
}
