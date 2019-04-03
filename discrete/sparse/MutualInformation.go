package sparse

import (
	"math"

	"github.com/kzahedi/goent/sm"
)

type lnFunc func(float64) float64

// MutualInformation calculates the mutual information with the given lnFunc function
//   I(X,Y) = \sum_x,y p(x,y) (lnFunc(p(x,y)) - lnFunc(p(x)p(y)))
func MutualInformation(pxy sm.SparseMatrix, log lnFunc) float64 {

	px := sm.CreateSparseMatrix()
	py := sm.CreateSparseMatrix()

	for _, index := range pxy.Indices {
		x := sm.SparseMatrixIndex{index[0]}
		y := sm.SparseMatrixIndex{index[1]}
		v, _ := pxy.Get(index)

		px.Add(x, v)
		py.Add(y, v)
	}

	var r float64

	for _, index := range pxy.Indices {
		x := sm.SparseMatrixIndex{index[0]}
		y := sm.SparseMatrixIndex{index[1]}
		xyv, _ := pxy.Get(index)
		xv, _ := px.Get(x)
		yv, _ := py.Get(y)

		if xyv > 0.0 && xv > 0.0 && yv > 0.0 {
			r += xyv * (log(xyv) - log(xv*yv))
		}
	}

	return r
}

// MutualInformationBaseE calculates the mutual information with base e
//   I(X,Y) = \sum_x,y p(x,y) (ln(p(x,y)) - ln(p(x)p(y)))
func MutualInformationBaseE(pxy sm.SparseMatrix) float64 {
	return MutualInformation(pxy, math.Log)
}

// MutualInformationBase2 calculates the mutual information with base 2
//   I(X,Y) = \sum_x,y p(x,y) (log2(p(x,y)) - log2(p(x)p(y)))
func MutualInformationBase2(pxy sm.SparseMatrix) float64 {
	return MutualInformation(pxy, math.Log2)
}
