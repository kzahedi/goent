package state

import (
	"math"
	"sort"

	"github.com/kzahedi/goent/continuous"
	pb "gopkg.in/cheggaaa/pb.v1"
)

// FrenzelPompe is an implementation of
// S. Frenzel and B. Pompe.
// Partial mutual information for coupling analysis of multivariate time series.
// Phys. Rev. Lett., 99:204101, Nov 2007.
// The function assumes that the data xyz is normalised column-wise
func FrenzelPompe(xyz [][]float64, xIndices, yIndices, zIndices []int, k int, eta bool) []float64 {

	T := len(xyz)
	r := make([]float64, T, T)

	hk := continuous.Harmonic(k - 1)

	var bar *pb.ProgressBar

	if eta == true {
		bar = pb.StartNew(T)
	}

	for t, v := range xyz {
		epsilon := fpGetEpsilon(k, v, xyz, xIndices, yIndices, zIndices)

		cNxz := fpCount2(epsilon, v, xyz, xIndices, zIndices)
		hNxz := continuous.Harmonic(cNxz)

		cNyz := fpCount2(epsilon, v, xyz, yIndices, zIndices)
		hNyz := continuous.Harmonic(cNyz)

		cNz := fpCount1(epsilon, v, xyz, zIndices)
		hNz := continuous.Harmonic(cNz)

		r[t] = (hNxz + hNyz - hNz - hk)

		if eta == true {
			bar.Increment()
		}
	}

	if eta == true {
		bar.Finish()
	}

	return r

}

// fpMaxNorm3 computes the max-norm of two 3-dimensional vectors
//   maxnorm(a,b) = max( |a[0] - b[0]|, |a[1] - b[1]|, |a[2] - b[2]|)
func fpMaxNorm3(a, b []float64, xIndices, yIndices, zIndices []int) float64 {
	xDist := continuous.Distance(a, b, xIndices)
	yDist := continuous.Distance(a, b, yIndices)
	zDist := continuous.Distance(a, b, zIndices)
	return math.Max(xDist, math.Max(yDist, zDist))
}

// fpGetEpsilon calculate epsilon_k(t) as defined by Frenzel & Pompe, 2007
// epsilon_k(t) is the continuous.Distance of the k-th nearest neighbour. The function
// takes k, the point from which the continuous.Distance is calculated (xyz), and the
// data from which the k-th nearest neighbour should be determined
func fpGetEpsilon(k int, xyz []float64, data [][]float64, xIndices, yIndices, zIndices []int) float64 {
	distances := make([]float64, len(data), len(data))

	for t := 0; t < len(data); t++ {
		distances[t] = fpMaxNorm3(xyz, data[t], xIndices, yIndices, zIndices)
	}

	sort.Float64s(distances)

	return distances[k] // we start to count at zero, but the first one is xyz[t] vs. xyz[t]
}

// fpCount2 count the number of points for which the x and y coordinate is
// closer than epsilon, where the continuous.Distance is measured by the max-norm
func fpCount2(epsilon float64, xyz []float64, data [][]float64, xIndices, yIndices []int) (c int) {

	c = -1 // because we will also count xyz[t] vs. xyz[t]
	for t := 0; t < len(data); t++ {
		if fpMaxNorm2(xyz, data[t], xIndices, yIndices) < epsilon {
			c++
		}
	}

	return
}

// fpMaxNorm2 computes the max-norm of two 3-dimensional vectors
//   maxnorm(a,b) = max( |a[0] - b[0]|, |a[1] - b[1]|)
func fpMaxNorm2(a, b []float64, xIndices, yIndices []int) float64 {
	xDist := continuous.Distance(a, b, xIndices)
	yDist := continuous.Distance(a, b, yIndices)
	return math.Max(xDist, yDist)
}

// fpCount1 count the number of points for which the z coordinate is
// closer than epsilon
func fpCount1(epsilon float64, xyz []float64, data [][]float64, zIndices []int) (c int) {
	c = -1 // because we will also count xyz[t] vs. xyz[t]
	for t := 0; t < len(data); t++ {
		if continuous.Distance(xyz, data[t], zIndices) < epsilon {
			c++
		}
	}
	return
}
