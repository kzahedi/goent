package continuous

import (
	"math"
	"sort"

	pb "gopkg.in/cheggaaa/pb.v1"
)

// KraskovStoegbauerGrassberger1 is an implementation of the first
// algorithm presented in
// A. Kraskov, H. Stoegbauer, and P. Grassberger.
// Estimating mutual information. Phys. Rev. E, 69:066138, Jun 2004.
func KraskovStoegbauerGrassberger1(xy [][]float64, xIndices, yIndices []int, k int, eta bool) (r float64) {

	r = 0.0

	hk := harmonic(k)       // h(k)
	hN := harmonic(len(xy)) // h(N)

	var bar *pb.ProgressBar

	if eta == true {
		bar = pb.StartNew(len(xy))
	}
	for t := 0; t < len(xy); t++ {
		epsilon := ksgGetEpsilon(k, xy[t], xy, xIndices, yIndices)

		cNx := ksgCount(epsilon, xy[t], xy, xIndices) // N_x
		hNx := harmonic(cNx + 1)                      // h(N_x)

		cNy := ksgCount(epsilon, xy[t], xy, yIndices) // N_y
		hNy := harmonic(cNy + 1)                      // h(N_y)

		r -= hNx + hNy

		if eta == true {
			bar.Increment()
		}
	}

	if eta == true {
		bar.Finish()
	}

	r /= float64(len(xy))

	r += hk + hN

	return
}

// KraskovStoegbauerGrassberger2 is an implementation of the second
// algorithm presented in
// A. Kraskov, H. Stoegbauer, and P. Grassberger.
// Estimating mutual information. Phys. Rev. E, 69:066138, Jun 2004.
func KraskovStoegbauerGrassberger2(xy [][]float64, xIndices, yIndices []int, k int, eta bool) (r float64) {

	r = 0.0

	hk := harmonic(k)
	hN := harmonic(len(xy))

	var bar *pb.ProgressBar

	if eta == true {
		bar = pb.StartNew(len(xy))
	}
	for t := 0; t < len(xy); t++ {
		epsilon := ksgGetEpsilon(k, xy[t], xy, xIndices, yIndices)

		cNx := ksgCount(epsilon, xy[t], xy, xIndices)
		hNx := harmonic(cNx)

		cNy := ksgCount(epsilon, xy[t], xy, yIndices)
		hNy := harmonic(cNy)

		r -= hNx + hNy

		if eta == true {
			bar.Increment()
		}
	}

	if eta == true {
		bar.FinishPrint("Finished")
	}

	r /= float64(len(xy))

	r += hk + hN - 1.0/float64(k)

	return
}

// ksgGetEpsilon calculate epsilon_k(t) as defined by Frenzel & Pompe, 2007
// epsilon_k(t) is the distance of the k-th nearest neighbour. The function
// takes k, the point from which the distance is calculated (xyz), and the
// data from which the k-th nearest neighbour should be determined
func ksgGetEpsilon(k int, xy []float64, data [][]float64, xIndices, yIndices []int) float64 {
	distances := make([]float64, len(data), len(data))

	for t := 0; t < len(data); t++ {
		distances[t] = ksgMaxNorm2(xy, data[t], xIndices, yIndices)
	}

	sort.Float64s(distances)

	return distances[k-1] // we start to ksgCount at zero
}

func ksgMaxNorm2(a, b []float64, xIndices, yIndices []int) float64 {
	xDistance := distance(a, b, xIndices)
	yDistance := distance(a, b, yIndices)
	return math.Max(xDistance, yDistance)
}

// ksgCount count the number of points for which the x or y coordinate is
// closer than epsilon, where the ksgDistance is measured by the max-norm
func ksgCount(epsilon float64, xy []float64, data [][]float64, indices []int) (c int) {

	for t := 0; t < len(data); t++ {
		if distance(xy, data[t], indices) < epsilon {
			c++
		}
	}

	return
}
