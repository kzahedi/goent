package state

import (
	"math"

	"github.com/kzahedi/goent/discrete"
)

// mi calculates the mutual information for each state with the given lnFunc function
//   MI(X,Y) = \sum_x,y p(x,y) (lnFunc(p(x,y)) - lnFunc(p(x)p(y)))
func mi(data [][]int64, log lnFunc) []float64 {
	pxy := discrete.Emperical2D(data)
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

// MutualInformation calculates the mutual information for each state with base e
//   MI(X,Y) = \sum_x,y p(x,y) (ln(p(x,y)) - ln(p(x)p(y)))
func MutualInformation(data [][]int64) []float64 {
	return mi(data, math.Log)
}

// MutualInformation2 calculates the mutual information with for each state with base 2
//   MI(X,Y) = \sum_x,y p(x,y) (log2(p(x,y)) - log2(p(x)p(y)))
func MutualInformation2(data [][]int64) []float64 {
	return mi(data, math.Log2)
}
