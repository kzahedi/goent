package continuous

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
)

func TestDistanceZero(t *testing.T) {
	a := []float64{1.0, 2.0, 3.0}
	b := []float64{1.0, 2.0, 3.0}

	if diff := distance(a, b, []int{0, 1, 2}); diff > 0.0 {
		t.Errorf(fmt.Sprintf("Distance should be zero but it is %f", diff))
	}
}

func TestDistanceNotZero(t *testing.T) {
	a := []float64{1.0, 2.0, 3.0}
	b := []float64{2.0, 4.0, 6.0}

	dist1 := math.Sqrt(1.0 + 4.0 + 9.0)
	dist2 := distance(a, b, []int{0, 1, 2})

	if math.Abs(dist1-dist2) > 0.00001 {
		t.Errorf(fmt.Sprintf("Distance should be %f but it is %f", dist1, dist2))
	}
}

func TestHarmonic1(t *testing.T) {
	h := harmonic(1)
	r := -0.5772156649

	if math.Abs(h-r) > 0.00001 {
		t.Errorf(fmt.Sprintf("Harmonic(1) should be %f but it is %f", r, h))
	}

}

func TestHarmonic5(t *testing.T) {
	h := harmonic(5)
	r := -1.0/2.0 - 1.0/3.0 - 1.0/4.0 - 1.0/5.0 - 0.5772156649

	if math.Abs(h-r) > 0.00001 {
		t.Errorf(fmt.Sprintf("Harmonic(1) should be %f but it is %f", r, h))
	}
}

func TestNormaliseZero(t *testing.T) {
	data := make([][]float64, 10, 10)
	for i := range data {
		data[i] = make([]float64, 5, 5)
	}

	ndata := Normalise(data)

	if len(ndata) != len(data) {
		t.Errorf(fmt.Sprintf("Normalise should not change the number of rows from %d to %d", len(data), len(ndata)))
	}

	for row := range data {
		for column := range data[row] {
			if len(ndata[row]) != 5 {
				t.Errorf(fmt.Sprintf("Normalise should not change the number of columns from 5 to %d", len(ndata[row])))
			}
			if math.Abs(ndata[row][column]) > 0.00000001 {
				t.Errorf(fmt.Sprintf("Normalised data [%d,%d] should be zero and not %f", row, column, ndata[row][column]))
			}
		}
	}
}

func TestNormaliseRandom(t *testing.T) {
	data := make([][]float64, 10, 10)
	for i := range data {
		data[i] = make([]float64, 5, 5)
		for j := range data[i] {
			data[i][j] = 10.0*rand.Float64() - 5.0
		}
	}

	// to make sure that the full range of values is used
	for i := 0; i < 5; i++ {
		data[0][i] = -5.0
		data[1][i] = 5.0
	}

	ndata := Normalise(data)

	for row := range data {
		for column := range data[row] {
			if math.Abs(ndata[row][column]-(data[row][column]/10.0+0.5)) > 0.00000001 {
				t.Errorf(fmt.Sprintf("Normalised data [%d,%d] should be %f and not %f",
					row, column, (data[row][column]/10.0 + 0.5), ndata[row][column]))
			}
		}
	}
}
