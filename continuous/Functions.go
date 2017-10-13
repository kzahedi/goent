package continuous

import "math"

func distance(a, b []float64, indices []int) float64 {
	d := 0.0
	for _, v := range indices {
		d += (a[v] - b[v]) * (a[v] - b[v])
	}
	return math.Sqrt(d)
}

func harmonic(n int) (r float64) {
	// harmonic(1) = -C, see A. Kraskov, H. Stoeogbauer, and P. Grassberger.
	// Estimating mutual information. Phys. Rev. E, 69:066138, Jun 2004.
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
// to FrenzelPompe or KraskovStoegbauerGrassberger1/2
func Normalise(data [][]float64) [][]float64 {

	min := make([]float64, len(data[0]), len(data[0]))
	max := make([]float64, len(data[0]), len(data[0]))

	r := make([][]float64, len(data), len(data))

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

	for row := range data {
		r[row] = make([]float64, len(data[0]), len(data[0]))
		for column := range data[0] {
			if math.Abs(min[column]-max[column]) > 0.000001 {
				r[row][column] = (data[row][column] - min[column]) / (max[column] - min[column])
			}
		}
	}

	return r
}
