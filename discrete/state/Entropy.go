package state

import (
	"math"

	"github.com/kzahedi/goent/discrete"
)

// Entropy calculates the entropy of a probability distribution.
// It takes the log function as an additional parameter, so that the base
// can be chosen
// H(X) = -\sum_x p(x) lnFunc(p(x))
func Entropy(data []int64, ln lnFunc) []float64 {
	r := make([]float64, len(data), len(data))
	p := discrete.Emperical1D(data)
	for i := 0; i < len(data); i++ {
		x := data[i]
		if p[x] > 0 {
			r[i] = ln(p[x])
		}
	}
	return r
}

// EntropyBaseE calculates the entropy of a probability distribution with base e
// H(X) = -\sum_x p(x) ln(p(x))
func EntropyBaseE(data []int64) []float64 {
	return Entropy(data, math.Log)
}

// EntropyBase2 calculates the entropy of a probability distribution with base 2
// H(X) = -\sum_x p(x) log2(p(x))
func EntropyBase2(data []int64) []float64 {
	return Entropy(data, math.Log2)
}
