package goent

func Create2D(xDim, yDim int) [][]float64 {
	r := make([][]float64, xDim)
	for x := 0; x < xDim; x++ {
		r[x] = make([]float64, yDim)
	}
	return r
}

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
