package goent

import "math"

// MI calculates the mutual information with the given LnFunc function
// MI(X,Y) = \sum_x,y p(x,y) (LnFunc(p(x,y)) - LnFunc(p(x)p(y)))
func MI(pxy [][]float64, log LnFunc) float64 {

	xDim := len(pxy)
	yDim := len(pxy[0])

	px := make([]float64, xDim)
	py := make([]float64, yDim)

	for x := 0; x < xDim; x++ {
		for y := 0; y < yDim; y++ {
			px[x] += pxy[x][y]
		}
	}

	for x := 0; x < xDim; x++ {
		for y := 0; y < yDim; y++ {
			py[y] += pxy[x][y]
		}
	}

	mi := 0.0

	for x := 0; x < xDim; x++ {
		if px[x] > 0.0 {
			for y := 0; y < yDim; y++ {
				if py[y] > 0.0 && pxy[x][y] > 0.0 {
					mi += pxy[x][y] * (log(pxy[x][y]) - log(px[x]*py[y]))
				}
			}
		}
	}
	return mi
}

// MutualInformation calculates the mutual information with base e
// MI(X,Y) = \sum_x,y p(x,y) (ln(p(x,y)) - ln(p(x)p(y)))
func MutualInformation(pxy [][]float64) float64 {
	return MI(pxy, math.Log)
}

// MutualInformation calculates the mutual information with base 2
// MI(X,Y) = \sum_x,y p(x,y) (log2(p(x,y)) - log2(p(x)p(y)))
func MutualInformation2(pxy [][]float64) float64 {
	return MI(pxy, math.Log2)
}
