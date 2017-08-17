package goent

import (
	"math"
)

// H calculates the entropy of a probability distribution.
// It takes the log function as an additional parameter, so that the base
// can be chosen
// H(X) = -\sum_x p(x) lnFunc(p(x))
func H(p []float64, log lnFunc) float64 {
	var r float64
	for _, px := range p {
		if px > 0 {
			r -= px * log(px)
		}
	}
	return r
}

// Entropy calculates the entropy of a probability distribution with base e
// H(X) = -\sum_x p(x) ln(p(x))
func Entropy(p []float64) float64 {
	return H(p, math.Log)
}

// Entropy calculates the entropy of a probability distribution with base 2
// H(X) = -\sum_x p(x) log2(p(x))
func Entropy2(p []float64) float64 {
	return H(p, math.Log2)
}
