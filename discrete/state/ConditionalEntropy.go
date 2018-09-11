package state

import (
	"math"

	"github.com/kzahedi/goent/discrete"
)

// ConditionalEntropy calculates the conditional entropy of a probability distribution.
// It takes the log function as an additional parameter, so that the base
// can be chosen
//   H(X|Y) = -\sum_x p(x,y) lnFunc(p(x,y)/p(y))
func ConditionalEntropy(data [][]int, log lnFunc) []float64 {
	pxy := discrete.Empirical2D(data)
	xDim := len(pxy)
	yDim := len(pxy[0])

	r := make([]float64, len(data), len(data))
	py := make([]float64, yDim, yDim)

	for y := 0; y < yDim; y++ {
		py[y] = 0.0
		for x := 0; x < xDim; x++ {
			py[y] += pxy[x][y]
		}
	}

	for i := range data {
		x := data[i][0]
		y := data[i][1]
		if py[y] > 0.0 && pxy[x][y] > 0.0 {
			r[i] = -(log(pxy[x][y]) - log(py[y]))
		}
	}
	return r
}

// ConditionalEntropyBaseE calculates the conditional entropy of a
// probability distribution in nats
//   H(X|Y) = -\sum_x p(x,y) ln(p(x,y)/p(y))
func ConditionalEntropyBaseE(data [][]int) []float64 {
	return ConditionalEntropy(data, math.Log)
}

// ConditionalEntropyBase2 calculates the conditional entropy of a
// probability distribution in bits
//   H(X|Y) = -\sum_x p(x,y) log2(p(x,y)/p(y))
func ConditionalEntropyBase2(data [][]int) []float64 {
	return ConditionalEntropy(data, math.Log2)
}
