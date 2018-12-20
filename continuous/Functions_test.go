package continuous

import (
	"fmt"
	"math"
	"math/rand"
	"testing"

	"github.com/kzahedi/goent/discrete"
)

func TestDistanceZero(t *testing.T) {
	a := []float64{1.0, 2.0, 3.0}
	b := []float64{1.0, 2.0, 3.0}

	if diff := Distance(a, b, []int{0, 1, 2}); diff > 0.0 {
		t.Errorf(fmt.Sprintf("Distance should be zero but it is %f", diff))
	}
}

func TestDistanceNotZero(t *testing.T) {
	a := []float64{1.0, 2.0, 3.0}
	b := []float64{2.0, 4.0, 6.0}

	dist1 := math.Sqrt(1.0 + 4.0 + 9.0)
	dist2 := Distance(a, b, []int{0, 1, 2})

	if math.Abs(dist1-dist2) > 0.00001 {
		t.Errorf(fmt.Sprintf("Distance should be %f but it is %f", dist1, dist2))
	}
}

func TestHarmonic0(t *testing.T) {
	h := Harmonic(0)

	if h != 0.0 {
		t.Errorf(fmt.Sprintf("Harmonic(0) should be 0.0 but it is %f", h))
	}
}

func TestHarmonic1(t *testing.T) {
	h := Harmonic(1)
	r := -0.5772156649

	if math.Abs(h-r) > 0.00001 {
		t.Errorf(fmt.Sprintf("Harmonic(1) should be %f but it is %f", r, h))
	}

}

func TestHarmonic5(t *testing.T) {
	h := Harmonic(5)
	r := -1.0/2.0 - 1.0/3.0 - 1.0/4.0 - 1.0/5.0 - 0.5772156649

	if math.Abs(h-r) > 0.00001 {
		t.Errorf(fmt.Sprintf("Harmonic(5) should be %f but it is %f", r, h))
	}

	h2 := Harmonic(0)
	if h2 != 0.0 {
		t.Errorf(fmt.Sprintf("Harmonic(0) should be 0.0 but it is %f", h2))
	}
}

func TestNormaliseZero(t *testing.T) {
	data := discrete.Create2D(10, 5)

	nData, _, _ := Normalise(data, false)

	if len(nData) != len(data) {
		t.Errorf(fmt.Sprintf("Normalise should not change the number of rows from %d to %d", len(data), len(nData)))
	}

	for row := range data {
		for column := range data[row] {
			if len(nData[row]) != 5 {
				t.Errorf(fmt.Sprintf("Normalise should not change the number of columns from 5 to %d", len(nData[row])))
			}
			if math.Abs(nData[row][column]) > 0.00000001 {
				t.Errorf(fmt.Sprintf("Normalised data [%d,%d] should be zero and not %f", row, column, nData[row][column]))
			}
		}
	}

	for i := range data {
		for j := 0; j < 5; j++ {
			data[i][j] = rand.Float64()*10.0 - 5.0
		}
	}

	nData, _, _ = Normalise(data, false)

	for row := range data {
		for column := range data[row] {
			if nData[row][column] < 0.0 {
				t.Errorf("Normalise should not produce values smaller than 0, but we have %f", nData[row][column])
			}
			if nData[row][column] > 1.0 {
				t.Errorf("Normalise should not produce values larger than 1, but we have %f", nData[row][column])
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

	nData, _, _ := Normalise(data, false)

	for row := range data {
		for column := range data[row] {
			if math.Abs(nData[row][column]-(data[row][column]/10.0+0.5)) > 0.00000001 {
				t.Errorf(fmt.Sprintf("Normalised data [%d,%d] should be %f and not %f",
					row, column, (data[row][column]/10.0 + 0.5), nData[row][column]))
			}
		}
	}
}
