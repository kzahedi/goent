package sparse

import (
	"math"

	"github.com/kzahedi/goent/sm"
)

// ConditionalMutualInformation calculates the conditional mutual information with the given lnFunc function
//   I(X,Y|Z) = \sum_x,y, p(x,y,z) (lnFunc(p(x,y|z)) - lnFunc(p(x|z)p(y|z)))
func ConditionalMutualInformation(pxyz sm.SparseMatrix, ln lnFunc) float64 {

	pxyCz := sm.CreateSparseMatrix()
	pxCz := sm.CreateSparseMatrix()
	pyCz := sm.CreateSparseMatrix()
	pz := sm.CreateSparseMatrix()

	for _, index := range pxyz.Indices {
		x := index[0]
		y := index[1]
		z := index[2]
		v, _ := pxyz.Get(index)

		zi := sm.SparseMatrixIndex{z}
		xzi := sm.SparseMatrixIndex{x, z}
		yzi := sm.SparseMatrixIndex{y, z}

		pz.Add(zi, v)
		pxyCz.Add(index, v)
		pxCz.Add(xzi, v)
		pyCz.Add(yzi, v)
	}

	for _, index := range pxyCz.Indices {
		zi := sm.SparseMatrixIndex{index[2]}
		v, _ := pz.Get(zi)
		pxyCz.Mul(zi, 1.0/v)
	}

	for _, index := range pxCz.Indices {
		zi := sm.SparseMatrixIndex{index[2]}
		v, _ := pz.Get(zi)
		pxCz.Mul(zi, 1.0/v)
	}

	for _, index := range pyCz.Indices {
		zi := sm.SparseMatrixIndex{index[2]}
		v, _ := pz.Get(zi)
		pyCz.Mul(zi, 1.0/v)
	}

	r := 0.0
	for _, index := range pxyz.Indices {
		xi := index[0]
		yi := index[1]
		zi := index[2]

		xyz, _ := pxyz.Get(index)
		xyCz, _ := pxyCz.Get(index)
		xCz, _ := pxCz.Get(sm.SparseMatrixIndex{xi, zi})
		yCz, _ := pxCz.Get(sm.SparseMatrixIndex{yi, zi})

		if xyz > 0.0 && xyCz > 0.0 && xCz > 0.0 && yCz > 0.0 {
			r += xyz * (ln(xyCz) - ln(xCz*yCz))
		}
	}

	return r
}

// ConditionalMutualInformationBaseE calculates the conditional mutual information with base e
// I(X,Y|Z) = \sum_x,y, p(x,y,z) (ln(p(x,y|z)) - ln(p(x|z)p(y|z)))
func ConditionalMutualInformationBaseE(pxyz sm.SparseMatrix) float64 {
	return ConditionalMutualInformation(pxyz, math.Log)
}

// ConditionalMutualInformationBase2 calculates the conditional mutual information with base 2
//   I(X,Y|Z) = \sum_x,y, p(x,y,z) (log2(p(x,y|z)) - log2(p(x|z)p(y|z)))
func ConditionalMutualInformationBase2(pxyz sm.SparseMatrix) float64 {
	return ConditionalMutualInformation(pxyz, math.Log2)
}
