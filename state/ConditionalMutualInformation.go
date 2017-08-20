package state

import (
	"math"

	"github.com/kzahedi/goent"
)

// CMI calculates the conditional mutual information with the given LnFunc
// function for each (x_t,y_t,z_t)
// CMI(X_t,Y_t|Z_t) = (LnFunc(p(x,y|z)) - LnFunc(p(x|z)p(y|z)))
func CMI(xyz [][]int64, ln goent.LnFunc) []float64 {

	pxyz := goent.Emperical3D(xyz)
	r := make([]float64, len(xyz), len(xyz))

	xDim := len(pxyz)
	yDim := len(pxyz[0])
	zDim := len(pxyz[0][0])

	pxy_c_z := goent.Create3D(xDim, yDim, zDim)
	px_c_z := goent.Create2D(xDim, zDim)
	py_c_z := goent.Create2D(yDim, zDim)

	pz := make([]float64, zDim)

	for x := 0; x < xDim; x++ {
		for y := 0; y < yDim; y++ {
			for z := 0; z < zDim; z++ {
				pz[z] += pxyz[x][y][z]
				px_c_z[x][z] += pxyz[x][y][z]
				py_c_z[y][z] += pxyz[x][y][z]
			}
		}
	}

	for x := 0; x < xDim; x++ {
		for z := 0; z < zDim; z++ {
			px_c_z[x][z] /= pz[z]
		}
	}

	for y := 0; y < yDim; y++ {
		for z := 0; z < zDim; z++ {
			py_c_z[y][z] /= pz[z]
		}
	}

	for x := 0; x < xDim; x++ {
		for y := 0; y < yDim; y++ {
			for z := 0; z < zDim; z++ {
				pxy_c_z[x][y][z] = pxyz[x][y][z] / pz[z]
			}
		}
	}

	for i := 0; i < len(xyz); i++ {
		x := xyz[i][0]
		y := xyz[i][1]
		z := xyz[i][2]
		if pxyz[x][y][z] > 0.0 && pxy_c_z[x][y][z] > 0.0 && px_c_z[x][z] > 0.0 && py_c_z[y][z] > 0.0 {
			r[i] = ln(pxy_c_z[x][y][z]) - ln(px_c_z[x][z]*py_c_z[y][z])
		}
	}

	return r
}

// CMI calculates the conditional mutual information with base e
// CMI(X,Y|Z) = \sum_x,y, p(x,y,z) (ln(p(x,y|z)) - ln(p(x|z)p(y|z)))
func ConditionalMutualInformation(xyz [][]int64) []float64 {
	return CMI(xyz, math.Log)
}

// CMI calculates the conditional mutual information with base 2
// CMI(X,Y|Z) = \sum_x,y, p(x,y,z) (log2(p(x,y|z)) - log2(p(x|z)p(y|z)))
func ConditionalMutualInformation2(xyz [][]int64) []float64 {
	return CMI(xyz, math.Log2)
}
