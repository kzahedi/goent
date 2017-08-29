package discrete

import "math"

// ConditionalMutualInformation calculates the conditional mutual information with the given lnFunc function
// CMI(X,Y|Z) = \sum_x,y, p(x,y,z) (lnFunc(p(x,y|z)) - lnFunc(p(x|z)p(y|z)))
func ConditionalMutualInformation(pxyz [][][]float64, ln lnFunc) float64 {

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

// ConditionalMutualInformationBaseE calculates the conditional mutual information with base e
// CMI(X,Y|Z) = \sum_x,y, p(x,y,z) (ln(p(x,y|z)) - ln(p(x|z)p(y|z)))
func ConditionalMutualInformationBaseE(pxyz [][][]float64) float64 {
	return ConditionalMutualInformation(pxyz, math.Log)
}

// ConditionalMutualInformationBase2 calculates the conditional mutual information with base 2
// CMI(X,Y|Z) = \sum_x,y, p(x,y,z) (log2(p(x,y|z)) - log2(p(x|z)p(y|z)))
func ConditionalMutualInformationBase2(pxyz [][][]float64) float64 {
	return ConditionalMutualInformation(pxyz, math.Log2)
}
