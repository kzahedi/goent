package state

import (
	"math"

	"github.com/kzahedi/goent/discrete"
)

// MutualInformation calculates the mutual information for each state with the given lnFunc function
//   I(X,Y) = \sum_x,y p(x,y) (lnFunc(p(x,y)) - lnFunc(p(x)p(y)))
func MutualInformation(data [][]int, log lnFunc) []float64 {
	pxy := discrete.Empirical2D(data)
	r := make([]float64, len(data), len(data))

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

	for i := 0; i < len(data); i++ {
		x := data[i][0]
		y := data[i][1]
		if px[x] > 0.0 && py[y] > 0.0 && pxy[x][y] > 0.0 {
			r[i] = log(pxy[x][y]) - log(px[x]*py[y])
		}
	}
	return r
}

// MutualInformationBaseE calculates the mutual information for each state with base e
//   I(X,Y) = \sum_x,y p(x,y) (ln(p(x,y)) - ln(p(x)p(y)))
func MutualInformationBaseE(data [][]int) []float64 {
	return MutualInformation(data, math.Log)
}

// MutualInformationBase2 calculates the mutual information with for each state with base 2
//   I(X,Y) = \sum_x,y p(x,y) (log2(p(x,y)) - log2(p(x)p(y)))
func MutualInformationBase2(data [][]int) []float64 {
	return MutualInformation(data, math.Log2)
}
