package goent

import (
	"math"
)

type lnFunc func(float64) float64

// H calculates the entropy of a probability distribution.
// It takes the log function as an additional parameter, so that the base
// can be chosen
// H(X) = -\sum_x p(x) lnFunc(p(x))
func H(p []float64, log lnFunc) float64 {
	var r float64
	for _, px := range p {
		if px != 0 {
			r -= px * log(px)
		}
	}
	return r
}

// Entropy calculates the entropy of a probability distribution with base e
// H(X) = -\sum_x p(x) ln(p(x))
func Entropy(p []float64) float64 {
	return H(p, math.Log)
}

// Entropy calculates the entropy of a probability distribution with base 2
// H(X) = -\sum_x p(x) log2(p(x))
func Entropy2(p []float64) float64 {
	return H(p, math.Log2)
}

// MI calculates the mutual information with the given lnFunc function
// MI(X,Y) = \sum_x,y p(x,y) (lnFunc(p(x,y)) - lnFunc(p(x)p(y)))
func MI(pxy [][]float64, log lnFunc) float64 {

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
		for y := 0; y < yDim; y++ {
			if px[x] > 0.0 && py[y] > 0.0 && pxy[x][y] > 0.0 {
				mi = pxy[x][y] * (log(pxy[x][y]) - log(px[x]*py[y]))
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

// CMI calculates the conditional mutual information with the given lnFunc function
// CMI(X,Y|Z) = \sum_x,y, p(x,y,z) (lnFunc(p(x,y|z)) - lnFunc(p(x|z)p(y|z)))
func CMI(pxyz [][][]float64, ln lnFunc) float64 {

	xDim := len(pxyz)
	yDim := len(pxyz[0])
	zDim := len(pxyz[0][0])

	pxy_c_z := Create3D(xDim, yDim, zDim)
	px_c_z := Create2D(xDim, zDim)
	py_c_z := Create2D(yDim, zDim)
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

	r := 0.0
	for x := 0; x < xDim; x++ {
		for y := 0; y < yDim; y++ {
			for z := 0; z < zDim; z++ {
				if pxyz[x][y][z] > 0.0 && pxy_c_z[x][y][z] > 0.0 && px_c_z[x][z] > 0.0 && py_c_z[y][z] > 0.0 {
					r += pxyz[x][y][z] * (ln(pxy_c_z[x][y][z]) - ln(px_c_z[x][z]*py_c_z[y][z]))
				}
			}
		}
	}

	return r
}

// CMI calculates the conditional mutual information with base e
// CMI(X,Y|Z) = \sum_x,y, p(x,y,z) (ln(p(x,y|z)) - ln(p(x|z)p(y|z)))
func ConditionalMutualInformation(pxyz [][][]float64) float64 {
	return CMI(pxyz, math.Log)
}

// CMI calculates the conditional mutual information with base 2
// CMI(X,Y|Z) = \sum_x,y, p(x,y,z) (log2(p(x,y|z)) - log2(p(x|z)p(y|z)))
func ConditionalMutualInformation2(pxyz [][][]float64) float64 {
	return CMI(pxyz, math.Log2)
}
