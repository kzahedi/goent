package continuous_test

import (
	"fmt"
	"math"
	"math/rand"
	"testing"

	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/gonum/stat/distmv"

	"github.com/kzahedi/goent/continuous"
)

func TestFrenzelPompe(t *testing.T) {
	t.Log("Testing FrenzelPompe against independent distribution")
	rand.Seed(2)

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

	if r := math.Abs(continuous.FrenzelPompe(xyz, xIndex, yIndex, zIndex, 30, false)); r > 0.1 {
		t.Errorf(fmt.Sprintf("Conditional Mutual information should be close to be 0.0 but it is %f", r))
	}
}

func TestFrenzelPompeGaussian(t *testing.T) {
	t.Log("Testing FrenzelPompe against Gaussian distribution")
	rand.Seed(1)

	N := 1000
	k := 30

	r := 0.9 // co-variance

	cmi := 0.0
	cmiTrue := 0.0

	mu := []float64{0.0, 0.0, 0.0}
	sym := mat.NewSymDense(3, []float64{
		1.0, r, r,
		r, 1.0, r,
		r, r, 1.0})
	normal, _ := distmv.NewNormal(mu, sym, nil)

	// I(X,Y|Z) = h(X,Z) + h(Y,Z) - h(Z) - h(X,Y,Z)
	// h(X) = d/2 (1 + ln 2pi) + 1/2 ln detC
	det1 := mat.Det(mat.NewSymDense(1, []float64{1.0}))
	det2 := mat.Det(mat.NewSymDense(2, []float64{1.0, r, r, 1.0}))
	det3 := mat.Det(sym)
	hXZ := (2.0/2.0)*(1.0+math.Log(2.0*math.Pi)) + (1.0/2.0)*math.Log(det2)
	hYZ := (2.0/2.0)*(1.0+math.Log(2.0*math.Pi)) + (1.0/2.0)*math.Log(det2)
	hZ := (1.0/2.0)*(1.0+math.Log(2.0*math.Pi)) + (1.0/2.0)*math.Log(det1)
	hXYZ := (3.0/2.0)*(1.0+math.Log(2.0*math.Pi)) + (1.0/2.0)*math.Log(det3)
	CMIGauss := hXZ + hYZ - hZ - hXYZ

	for i := 0; i < 100; i++ {

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

		cmi += math.Abs(continuous.FrenzelPompe(xyz, xIndex, yIndex, zIndex, k, false))
		cmiTrue += math.Abs(continuous.FrenzelPompe(xyz, xIndex, yIndex, zIndex, k, true))
	}

	cmi /= 100.0
	cmiTrue /= 100.0

	if math.Abs(cmi-CMIGauss) > 0.1 {
		t.Errorf(fmt.Sprintf("Conditional Mutual Information should be close %f but it is %f", CMIGauss, cmi))
	}

	if math.Abs(cmi-cmiTrue) > 0.1 {
		t.Errorf(fmt.Sprintf("Frenzel-Pompe with and without ETA should be the same but %f != %f", cmi, cmiTrue))
	}

}
