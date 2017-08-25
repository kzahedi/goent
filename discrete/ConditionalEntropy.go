package discrete

import (
	"math"
)

// cH calculates the conditional entropy of a probability distribution.
// It takes the log function as an additional parameter, so that the base
// can be chosen
// H(X|Y) = -\sum_x p(x,y) LnFunc(p(x,y)/p(y))
func cH(pxy [][]float64, log LnFunc) float64 {
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

// ConditionalEntropy calculates the conditional entropy of a
// probability distribution in nats
// H(X|Y) = -\sum_x p(x,y) ln(p(x,y)/p(y))
func ConditionalEntropy(pxy [][]float64) float64 {
	return cH(pxy, math.Log)
}

// ConditionalEntropy calculates the conditional entropy of a
// probability distribution in bits
// H(X|Y) = -\sum_x p(x,y) log2(p(x,y)/p(y))
func ConditionalEntropy2(pxy [][]float64) float64 {
	return cH(pxy, math.Log2)
}
