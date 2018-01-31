package discrete

// lnFunc log function type
type lnFunc func(float64) float64

// Create2D creates a 2-dimensional slice
func Create2D(xDim, yDim int) [][]float64 {
	r := make([][]float64, xDim)
	for x := 0; x < xDim; x++ {
		r[x] = make([]float64, yDim)
	}
	return r
}

// Create2DInt creates a 2-dimensional slice
func Create2DInt(xDim, yDim int) [][]int {
	r := make([][]int, xDim)
	for x := 0; x < xDim; x++ {
		r[x] = make([]int, yDim)
	}
	return r
}

// Create3D creates a 3-dimensional slice
func Create3D(xDim, yDim, zDim int) [][][]float64 {
	r := make([][][]float64, xDim)
	for x := 0; x < xDim; x++ {
		r[x] = make([][]float64, yDim)
		for y := 0; y < yDim; y++ {
			r[x][y] = make([]float64, zDim)
		}
	}
	return r
}

// Create4D creates a 3-dimensional slice
func Create4D(xDim, yDim, zDim, wDim int) [][][][]float64 {
	r := make([][][][]float64, xDim)
	for x := 0; x < xDim; x++ {
		r[x] = make([][][]float64, yDim)
		for y := 0; y < yDim; y++ {
			r[x][y] = make([][]float64, yDim)
			for z := 0; z < zDim; z++ {
				r[x][y][z] = make([]float64, wDim)
			}
		}
	}
	return r
}

// Create3DInt creates a 3-dimensional slice
func Create3DInt(xDim, yDim, zDim int) [][][]int {
	r := make([][][]int, xDim)
	for x := 0; x < xDim; x++ {
		r[x] = make([][]int, yDim)
		for y := 0; y < yDim; y++ {
			r[x][y] = make([]int, zDim)
		}
	}
	return r
}

// Normalise1D return the normalised matrix
//   a / sum(a)
func Normalise1D(a []float64) []float64 {
	r := make([]float64, len(a), len(a))
	sum := 0.0
	for i := 0; i < len(a); i++ {
		sum += a[i]
	}

	for i := 0; i < len(a); i++ {
		r[i] = a[i] / sum
	}
	return r
}

// Normalise2D return the normalised matrix
//   a / sum(a)
func Normalise2D(a [][]float64) [][]float64 {
	r := Create2D(len(a), len(a[0]))
	sum := 0.0
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			sum += a[i][j]
		}
	}

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			r[i][j] = a[i][j] / sum
		}
	}
	return r
}

// Normalise3D return the normalised matrix
//   a / sum(a)
func Normalise3D(a [][][]float64) [][][]float64 {
	r := Create3D(len(a), len(a[0]), len(a[0][0]))
	sum := 0.0
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			for k := 0; k < len(a[i]); k++ {
				sum += a[i][j][k]
			}
		}
	}

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			for k := 0; k < len(a[i]); k++ {
				r[i][j][k] = a[i][j][k] / sum
			}
		}
	}
	return r
}

// Normalise4D return the normalised matrix
//   a / sum(a)
func Normalise4D(a [][][][]float64) [][][][]float64 {
	r := Create4D(len(a), len(a[0]), len(a[0][0]), len(a[0][0][0]))
	sum := 0.0
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			for k := 0; k < len(a[i]); k++ {
				for l := 0; l < len(a[i]); l++ {
					sum += a[i][j][k][l]
				}
			}
		}
	}

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			for k := 0; k < len(a[i]); k++ {
				for l := 0; l < len(a[i]); l++ {
					r[i][j][k][l] = a[i][j][k][l] / sum
				}
			}
		}
	}
	return r
}
