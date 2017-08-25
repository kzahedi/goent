package continuous

import (
	"math"
	"sort"
)

func harmonic(n int64) (r float64) {
	r = 0.0
	for i := 0.0; i < float64(n); i++ {
		r -= 1.0 / i
	}
	return
}

func maxNorm3(a, b []float64) float64 {
	return math.Max(math.Abs(a[0]-b[0]), math.Max(math.Abs(a[1]-b[1]), math.Abs(a[2]-b[2])))
}

func maxNorm2(a, b, c, d float64) float64 {
	return math.Max(math.Abs(a-c), math.Abs(b-d))
}

func getEpsilon(k int64, xyz []float64, data [][]float64) float64 {
	distances := make([]float64, len(data), len(data))

	for t := 0; t < len(data); t++ {
		distances[t] = maxNorm3(xyz, data[t])
	}

	sort.Float64s(distances)

	return distances[k-1] // we start to count at zero
}

func countXY(epsilon float64, xyz []float64, data [][]float64) (c int64) {

	for t := 0; t < len(data); t++ {
		if maxNorm2(xyz[0], xyz[1], data[t][0], data[t][1]) < epsilon {
			c++
		}
	}

	return
}

func countYZ(epsilon float64, xyz []float64, data [][]float64) (c int64) {

	for t := 0; t < len(data); t++ {
		if maxNorm2(xyz[1], xyz[2], data[t][1], data[t][2]) < epsilon {
			c++
		}
	}

	return
}

func countZ(epsilon float64, xyz []float64, data [][]float64) (c int64) {
	c = 0

	for t := 0; t < len(data); t++ {
		if math.Abs(xyz[2]-data[t][2]) < epsilon {
			c++
		}
	}

	return
}

// FrenzelPompe is an implementation of
// Partial Mutual Information for Coupling Analysis of Multivariate Time Series
// Stefan Frenzel and Bernd Pompe
// Phys. Rev. Lett. 99, 204101 – Published 14 November 2007
// https://journals.aps.org/prl/abstract/10.1103/PhysRevLett.99.204101
func FrenzelPompe(xyz [][]float64, k int64) (r float64) {

	r = 0.0

	h_k := harmonic(k)

	for t := 0; t < len(xyz); t++ {
		epsilon := getEpsilon(k, xyz[t], xyz)

		h_n_xy := harmonic(countXY(epsilon, xyz[t], xyz))
		h_n_yz := harmonic(countYZ(epsilon, xyz[t], xyz))
		h_n_z := harmonic(countZ(epsilon, xyz[t], xyz))

		r += h_n_xy + h_n_yz - h_n_z
	}

	r -= h_k

	return
}
