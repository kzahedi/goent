package discrete

import (
	"math"

	"github.com/kzahedi/goent/sm"
)

// MutualInformation calculates the mutual information with the given lnFunc function
//   I(X,Y) = \sum_x,y p(x,y) (lnFunc(p(x,y)) - lnFunc(p(x)p(y)))
func MutualInformation(pxy [][]float64, log lnFunc) float64 {

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

// MutualInformationBaseE calculates the mutual information with base e
//   I(X,Y) = \sum_x,y p(x,y) (ln(p(x,y)) - ln(p(x)p(y)))
func MutualInformationBaseE(pxy [][]float64) float64 {
	return MutualInformation(pxy, math.Log)
}

// MutualInformationBase2 calculates the mutual information with base 2
//   I(X,Y) = \sum_x,y p(x,y) (log2(p(x,y)) - log2(p(x)p(y)))
func MutualInformationBase2(pxy [][]float64) float64 {
	return MutualInformation(pxy, math.Log2)
}

// MutualInformationSparse calculates the mutual information with the given lnFunc function
//   I(X,Y) = \sum_x,y p(x,y) (lnFunc(p(x,y)) - lnFunc(p(x)p(y)))
func MutualInformationSparse(pxy sm.SparseMatrix, log lnFunc) float64 {

	px := sm.CreateSparseMatrix()
	py := sm.CreateSparseMatrix()

	for _, index := range pxy.Indices {
		x := sm.SparseMatrixIndex{index[0]}
		y := sm.SparseMatrixIndex{index[1]}
		v, _ := pxy.Get(index)

		px.Add(x, v)
		py.Add(y, v)
	}

	mi := 0.0

	for _, index := range pxy.Indices {
		x := sm.SparseMatrixIndex{index[0]}
		y := sm.SparseMatrixIndex{index[1]}
		xyv, _ := pxy.Get(index)
		xv, _ := px.Get(x)
		yv, _ := py.Get(y)

		if xyv > 0.0 && xv > 0.0 && yv > 0.0 {
			mi += xyv * (log(xyv) - log(xv*yv))
		}
	}

	return mi
}

// MutualInformationBaseESparse calculates the mutual information with base e
//   I(X,Y) = \sum_x,y p(x,y) (ln(p(x,y)) - ln(p(x)p(y)))
func MutualInformationBaseESparse(pxy sm.SparseMatrix) float64 {
	return MutualInformationSparse(pxy, math.Log)
}

// MutualInformationBase2 calculates the mutual information with base 2
//   I(X,Y) = \sum_x,y p(x,y) (log2(p(x,y)) - log2(p(x)p(y)))
func MutualInformationBase2Sparse(pxy sm.SparseMatrix) float64 {
	return MutualInformationSparse(pxy, math.Log2)
}
