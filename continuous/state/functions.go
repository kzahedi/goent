package state

import (
	"math"
)

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
