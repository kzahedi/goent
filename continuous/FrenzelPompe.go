package continuous

import (
	"math"
	"sort"

	pb "gopkg.in/cheggaaa/pb.v1"
)

// FrenzelPompe is an implementation of
// S. Frenzel and B. Pompe.
// Partial mutual information for coupling analysis of multivariate time series.
// Phys. Rev. Lett., 99:204101, Nov 2007.
func FrenzelPompe(xyz [][]float64, xIndices, yIndices, zIndices []int, k int, eta bool) (r float64) {

	r = 0.0

	hk := Harmonic(k - 1)

	var bar *pb.ProgressBar

	if eta == true {
		bar = pb.StartNew(len(xyz))
	}
	for t := 0; t < len(xyz); t++ {
		epsilon := fpGetEpsilon(k, xyz[t], xyz, xIndices, yIndices, zIndices)

		cNxy := fpCountXY(epsilon, xyz[t], xyz, xIndices, yIndices)
		hNxy := Harmonic(cNxy)
		// fmt.Println(fmt.Sprintf("cNxy %d hNxy %f", cNxy, hNxy))

		cNyz := fpCountYZ(epsilon, xyz[t], xyz, yIndices, zIndices)
		hNyz := Harmonic(cNyz)
		// fmt.Println(fmt.Sprintf("cNyz %d hNyz %f", cNyz, hNyz))

		cNz := fpCountZ(epsilon, xyz[t], xyz, zIndices)
		hNz := Harmonic(cNz)
		// fmt.Println(fmt.Sprintf("cNz %d hNz %f", cNz, hNz))

		r += hNxy + hNyz - hNz

		if eta == true {
			bar.Increment()
		}
	}

	if eta == true {
		bar.FinishPrint("Finished")
	}

	r /= float64(len(xyz))

	r -= hk

	return
}

// fpMaxNorm3 computes the max-norm of two 3-dimensional vectors
//   maxnorm(a,b) = max( |a[0] - b[0]|, |a[1] - b[1]|, |a[2] - b[2]|)
func fpMaxNorm3(a, b []float64, xIndices, yIndices, zIndices []int) float64 {
	xDist := Distance(a, b, xIndices)
	yDist := Distance(a, b, yIndices)
	zDist := Distance(a, b, zIndices)
	return math.Max(xDist, math.Max(yDist, zDist))
}

// fpGetEpsilon calculate epsilon_k(t) as defined by Frenzel & Pompe, 2007
// epsilon_k(t) is the Distance of the k-th nearest neighbour. The function
// takes k, the point from which the Distance is calculated (xyz), and the
// data from which the k-th nearest neighbour should be determined
func fpGetEpsilon(k int, xyz []float64, data [][]float64, xIndices, yIndices, zIndices []int) float64 {
	Distances := make([]float64, len(data), len(data))

	for t := 0; t < len(data); t++ {
		Distances[t] = fpMaxNorm3(xyz, data[t], xIndices, yIndices, zIndices)
	}

	sort.Float64s(Distances)

	return Distances[k-1] // we start to count at zero
}

// fpCountXY count the number of points for which the x and y coordinate is
// closer than epsilon, where the Distance is measured by the max-norm
func fpCountXY(epsilon float64, xyz []float64, data [][]float64, xIndices, yIndices []int) (c int) {

	for t := 0; t < len(data); t++ {
		if fpMaxNorm2(xyz, data[t], xIndices, yIndices) < epsilon {
			c++
		}
	}

	return
}

// fpCountYZ count the number of points for which the y and z coordinate is
// closer than epsilon, where the Distance is measured by the max-norm
func fpCountYZ(epsilon float64, xyz []float64, data [][]float64, yIndices, zIndices []int) (c int) {

	for t := 0; t < len(data); t++ {
		if fpMaxNorm2(xyz, data[t], yIndices, zIndices) < epsilon {
			c++
		}
	}

	return
}

func fpMaxNorm2(a, b []float64, xIndices, yIndices []int) float64 {
	xDist := Distance(a, b, xIndices)
	yDist := Distance(a, b, yIndices)
	return math.Max(xDist, yDist)
}

// fpCountZ count the number of points for which the z coordinate is
// closer than epsilon
func fpCountZ(epsilon float64, xyz []float64, data [][]float64, zIndices []int) (c int) {
	for t := 0; t < len(data); t++ {
		if Distance(xyz, data[t], zIndices) < epsilon {
			c++
		}
	}
	return
}
