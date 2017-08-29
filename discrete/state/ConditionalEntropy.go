package state

import (
	"math"
)

// ConditionalEntropy calculates the conditional entropy of a probability distribution.
// It takes the log function as an additional parameter, so that the base
// can be chosen
// H(X|Y) = -\sum_x p(x,y) lnFunc(p(x,y)/p(y))
func ConditionalEntropy(pxy [][]float64, log lnFunc) float64 {
	var r float64
	xDim := len(pxy)
	yDim := len(pxy[0])
	py := 0.0
	for y := 0; y < yDim; y++ {
		py = 0.0
		for x := 0; x < xDim; x++ {
			py += pxy[x][y]
		}
		if py > 0.0 {
			for x := 0; x < xDim; x++ {
				if pxy[x][y] > 0.0 {
					r -= pxy[x][y] * (log(pxy[x][y]) - log(py))
				}
			}
		}
	}
	return r
}

// ConditionalEntropyBaseE calculates the conditional entropy of a
// probability distribution in nats
// H(X|Y) = -\sum_x p(x,y) ln(p(x,y)/p(y))
func ConditionalEntropyBaseE(pxy [][]float64) float64 {
	return ConditionalEntropy(pxy, math.Log)
}

// ConditionalEntropyBase2 calculates the conditional entropy of a
// probability distribution in bits
// H(X|Y) = -\sum_x p(x,y) log2(p(x,y)/p(y))
func ConditionalEntropyBase(pxy [][]float64) float64 {
	return ConditionalEntropy(pxy, math.Log2)
}
