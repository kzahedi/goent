package discrete

import (
	"math"

	"gonum.org/v1/gonum/mat"
)

// ConditionalEntropy calculates the conditional entropy of a probability distribution.
// It takes the log function as an additional parameter, so that the base
// can be chosen
//   H(X|Y) = -\sum_x p(x,y) log2(p(x,y)/p(y))
// Results are given in nats
func ConditionalEntropy(pxy mat.Matrix) float64 {
	var r float64
	var py float64

	xDim, yDim := pxy.Dims()

	for y := 0; y < yDim; y++ {
		py = 0.0
		for x := 0; x < xDim; x++ {
			py += pxy.At(x, y)
		}
		if py > 0.0 {
			for x := 0; x < xDim; x++ {
				if pxy.At(x, y) > 0.0 {
					r -= pxy.At(x, y) * (math.Log(pxy.At(x, y)) - math.Log(py))
				}
			}
		}
	}
	return r
}
