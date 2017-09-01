package discrete

import "math"

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
