package state

import (
	"math"

	"github.com/kzahedi/goent/discrete"
)

// ConditionalMutualInformation calculates the conditional mutual information with the given lnFunc
// function for each (x_t,y_t,z_t)
// CMI(X_t,Y_t|Z_t) = (lnFunc(p(x,y|z)) - lnFunc(p(x|z)p(y|z)))
func ConditionalMutualInformation(xyz [][]int64, ln lnFunc) []float64 {

	pxyz := discrete.Emperical3D(xyz)
	r := make([]float64, len(xyz), len(xyz))

	xDim := len(pxyz)
	yDim := len(pxyz[0])
	zDim := len(pxyz[0][0])

	pxyCz := discrete.Create3D(xDim, yDim, zDim)
	pxCz := discrete.Create2D(xDim, zDim)
	pyCz := discrete.Create2D(yDim, zDim)

	pz := make([]float64, zDim)

	for x := 0; x < xDim; x++ {
		for y := 0; y < yDim; y++ {
			for z := 0; z < zDim; z++ {
				pz[z] += pxyz[x][y][z]
				pxCz[x][z] += pxyz[x][y][z]
				pyCz[y][z] += pxyz[x][y][z]
			}
		}
	}

	for x := 0; x < xDim; x++ {
		for z := 0; z < zDim; z++ {
			pxCz[x][z] /= pz[z]
		}
	}

	for y := 0; y < yDim; y++ {
		for z := 0; z < zDim; z++ {
			pyCz[y][z] /= pz[z]
		}
	}

	for x := 0; x < xDim; x++ {
		for y := 0; y < yDim; y++ {
			for z := 0; z < zDim; z++ {
				pxyCz[x][y][z] = pxyz[x][y][z] / pz[z]
			}
		}
	}

	for i := 0; i < len(xyz); i++ {
		x := xyz[i][0]
		y := xyz[i][1]
		z := xyz[i][2]
		if pxyz[x][y][z] > 0.0 && pxyCz[x][y][z] > 0.0 && pxCz[x][z] > 0.0 && pyCz[y][z] > 0.0 {
			r[i] = ln(pxyCz[x][y][z]) - ln(pxCz[x][z]*pyCz[y][z])
		}
	}

	return r
}

// ConditionalMutualInformationBaseE calculates the conditional mutual information with base e
// CMI(X,Y|Z) = \sum_x,y, p(x,y,z) (ln(p(x,y|z)) - ln(p(x|z)p(y|z)))
func ConditionalMutualInformationBaseE(xyz [][]int64) []float64 {
	return ConditionalMutualInformation(xyz, math.Log)
}

// ConditionalMutualInformationBase2 calculates the conditional mutual information with base 2
// CMI(X,Y|Z) = \sum_x,y, p(x,y,z) (log2(p(x,y|z)) - log2(p(x|z)p(y|z)))
func ConditionalMutualInformationBase2(xyz [][]int64) []float64 {
	return ConditionalMutualInformation(xyz, math.Log2)
}
