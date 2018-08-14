package discrete

import "math"

// ConditionalMutualInformation calculates the conditional mutual information
//   I(X,Y|Z) = \sum_x,y, p(x,y,z) (log2(p(x,y|z)) - log2(p(x|z)p(y|z)))
// Results are given in nats
func ConditionalMutualInformation(pxyz [][][]float64) float64 {

	xDim := len(pxyz)
	yDim := len(pxyz[0])
	zDim := len(pxyz[0][0])

	pxyCz := Create3D(xDim, yDim, zDim)
	pxCz := mat64.NewMatrix(xDim, zDim, nil)
	pyCz := mat64.NewMatrix(yDim, zDim, nil)

	pz := mat.NewVecDense(zDim, [])

	for x := 0; x < xDim; x++ {
		for y := 0; y < yDim; y++ {
			for z := 0; z < zDim; z++ {
				pz.Set(z, pz.At(z) + pxyz[x][y][z])
				pxCz.Set(x, z, pxCz.At(x, z)+pxyz[x][y][z])
				pyCz.Set(y, z, pyCz.At(y, z)+pxyz[x][y][z])
			}
		}
	}

	for x := 0; x < xDim; x++ {
		for z := 0; z < zDim; z++ {
			pxCz.Set(x, z, pxCz.At(x, z)/pz[z])
		}
	}

	for y := 0; y < yDim; y++ {
		for z := 0; z < zDim; z++ {
			pyCz.Set(y, z, pyCz.At(y, z)/pz[z])
		}
	}

	for x := 0; x < xDim; x++ {
		for y := 0; y < yDim; y++ {
			for z := 0; z < zDim; z++ {
				pxyCz[x][y][z] = pxyz[x][y][z] / pz.At(z)
			}
		}
	}

	r := 0.0
	for x := 0; x < xDim; x++ {
		for y := 0; y < yDim; y++ {
			for z := 0; z < zDim; z++ {
				if pxyz[x][y][z] > 0.0 && pxyCz[x][y][z] > 0.0 && pxCz.At(x,z) > 0.0 && pyCz.At(y,z) > 0.0 {
					r += pxyz[x][y][z] * (math.Log(pxyCz[x][y][z]) - math.Log(pxCz.At(x,z)*pyCz.At(y,z)))
				}
			}
		}
	}
	return r
}
