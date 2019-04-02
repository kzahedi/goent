package sparse

import (
	"math"

	"github.com/kzahedi/goent/discrete"
	"github.com/kzahedi/goent/sm"
)

// ConditionalMutualInformation calculates the conditional
// mutual information with the given lnFunc function for each (x_t,y_t,z_t)
//   I(X_t,Y_t|Z_t) = (lnFunc(p(x,y|z)) - lnFunc(p(x|z)p(y|z)))
func ConditionalMutualInformation(xyz [][]int, ln lnFunc) []float64 {

	pxyz := discrete.Empirical3DSparse(xyz)
	r := make([]float64, len(xyz), len(xyz))

	pxyCz := sm.CreateSparseMatrix()
	pxCz := sm.CreateSparseMatrix()
	pyCz := sm.CreateSparseMatrix()
	pz := sm.CreateSparseMatrix()

	for _, index := range pxyz.Indices {
		x := index[0]
		y := index[1]
		z := index[2]

		zi := sm.SparseMatrixIndex{z}
		xzi := sm.SparseMatrixIndex{x, z}
		yzi := sm.SparseMatrixIndex{y, z}

		v, _ := pxyz.Get(index)

		pz.Add(zi, v)
		pxCz.Add(xzi, v)
		pyCz.Add(yzi, v)
		pxyCz.Add(index, v)
	}

	for _, index := range pxCz.Indices {
		zi := sm.SparseMatrixIndex{index[1]}
		v, _ := pz.Get(zi)
		pxCz.Mul(index, 1.0/v)
	}

	for _, index := range pyCz.Indices {
		zi := sm.SparseMatrixIndex{index[1]}
		v, _ := pz.Get(zi)
		pyCz.Mul(index, 1.0/v)
	}

	for _, index := range pxyCz.Indices {
		zi := sm.SparseMatrixIndex{index[1]}
		v, _ := pz.Get(zi)
		pxyCz.Mul(index, 1.0/v)
	}

	for i, d := range xyz {
		x := d[0]
		y := d[1]
		z := d[2]

		xyzi := sm.SparseMatrixIndex{x, y, z}
		xzi := sm.SparseMatrixIndex{x, z}
		yzi := sm.SparseMatrixIndex{y, z}

		xyCz, _ := pxyCz.Get(xyzi)
		xCz, _ := pxCz.Get(xzi)
		yCz, _ := pxCz.Get(yzi)

		if xyCz > 0.0 && xCz > 0.0 && yCz > 0.0 {
			r[i] = ln(xyCz - ln(xCz*yCz))
		}
	}

	return r
}

// ConditionalMutualInformationBaseE calculates the conditional
// mutual information with base e
//   I(X,Y|Z) = \sum_x,y, p(x,y,z) (ln(p(x,y|z)) - ln(p(x|z)p(y|z)))
func ConditionalMutualInformationBaseE(xyz [][]int) []float64 {
	return ConditionalMutualInformation(xyz, math.Log)
}

// ConditionalMutualInformationBase2 calculates the conditional
// mutual information with base 2
//   I(X,Y|Z) = \sum_x,y, p(x,y,z) (log2(p(x,y|z)) - log2(p(x|z)p(y|z)))
func ConditionalMutualInformationBase2(xyz [][]int) []float64 {
	return ConditionalMutualInformation(xyz, math.Log2)
}
