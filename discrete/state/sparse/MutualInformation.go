package sparse

import (
	"math"

	"github.com/kzahedi/goent/sm"

	"github.com/kzahedi/goent/discrete"
)

type lnFunc func(float64) float64

// MutualInformation calculates the mutual information for each state with the given lnFunc function
//   I(X,Y) = \sum_x,y p(x,y) (lnFunc(p(x,y)) - lnFunc(p(x)p(y)))
func MutualInformation(data [][]int, log lnFunc) []float64 {
	pxy := discrete.Empirical2DSparse(data)
	r := make([]float64, len(data), len(data))

	px := sm.CreateSparseMatrix()
	py := sm.CreateSparseMatrix()

	for _, index := range pxy.Indices {
		x := sm.SparseMatrixIndex{index[0]}
		y := sm.SparseMatrixIndex{index[1]}
		v, _ := pxy.Get(index)

		px.Add(x, v)
		py.Add(y, v)
	}

	for i := 0; i < len(data); i++ {
		xd := data[i][0]
		yd := data[i][1]

		xi := sm.SparseMatrixIndex{xd}
		yi := sm.SparseMatrixIndex{yd}
		xyi := sm.SparseMatrixIndex{xd, yd}

		xy, _ := pxy.Get(xyi)
		x, _ := px.Get(xi)
		y, _ := py.Get(yi)

		if x > 0.0 && y > 0.0 && xy > 0.0 {
			r[i] = log(xy - log(x*y))
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
