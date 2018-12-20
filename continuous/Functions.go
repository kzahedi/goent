package continuous

import (
	"fmt"
	"math"
)

// Distance calculates the distance between to vectors,
// given the set of indices
func Distance(a, b []float64, indices []int) float64 {
	d := 0.0
	for _, v := range indices {
		d += (a[v] - b[v]) * (a[v] - b[v])
	}
	return math.Sqrt(d)
}

// Harmonic calculates the harmonic according to
// A. Kraskov, H. Stoegbauer, and P. Grassberger.
// Estimating mutual information. Phys. Rev. E, 69:066138, Jun 2004.
func Harmonic(n int) (r float64) {
	if n == 0 {
		return
	}

	r = -0.5772156649
	if n > 0 {
		for i := 2.0; i <= float64(n); i++ {
			r -= 1.0 / i
		}
	}

	return
}

// Normalise can be used to normalise the data before passing it
// to FrenzelPompe or KraskovStoegbauerGrassberger1/2.
// This function calls NormaliseByDomain
func Normalise(data [][]float64, verbose bool) ([][]float64, []float64, []float64) {

	min := make([]float64, len(data[0]), len(data[0]))
	max := make([]float64, len(data[0]), len(data[0]))

	for column := range data[0] {
		min[column] = data[0][column]
		max[column] = data[0][column]
	}

	for column := range data[0] {
		for row := range data {
			if min[column] > data[row][column] {
				min[column] = data[row][column]
			}
			if max[column] < data[row][column] {
				max[column] = data[row][column]
			}
		}
	}

	return NormaliseByDomain(data, min, max, verbose), min, max
}

// NormaliseByDomain can be used to normalise the data before passing it
// to FrenzelPompe or KraskovStoegbauerGrassberger1/2
// It takes the data and the minimum and maximum values per column
func NormaliseByDomain(data [][]float64, min, max []float64, verbose bool) [][]float64 {

	if verbose == true {
		minStr := ""
		maxStr := ""
		for i := range min {
			minStr = fmt.Sprintf("%s %f", minStr, min[i])
			maxStr = fmt.Sprintf("%s %f", maxStr, max[i])
		}
	}

	r := make([][]float64, len(data), len(data))

	for row := range data {
		r[row] = make([]float64, len(data[0]), len(data[0]))
		for column := range data[0] {
			if math.Abs(min[column]-max[column]) > 0.000001 {
				value := data[row][column]
				if value > max[column] {
					value = max[column]
				}
				if value < min[column] {
					value = min[column]
				}
				r[row][column] = (value - min[column]) / (max[column] - min[column])
			}
		}
	}

	return r
}
