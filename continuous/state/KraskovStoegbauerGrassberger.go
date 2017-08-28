package state

import (
	"math"
	"sort"

	"github.com/kzahedi/goent/continuous"

	pb "gopkg.in/cheggaaa/pb.v1"
)

// KraskovStoegbauerGrassberger1 is an implementation of the first
// algorithm presented in
// A. Kraskov, H. Stoegbauer, and P. Grassberger.
// Estimating mutual information. Phys. Rev. E, 69:066138, Jun 2004.
func KraskovStoegbauerGrassberger1(xy [][]float64, xIndices, yIndices []int64, k int64) (r float64) {

	r = 0.0

	h_k := continuous.Harmonic(k)
	h_N := continuous.Harmonic(len(xy))

	bar := pb.StartNew(len(xy))
	for t := 0; t < len(xy); t++ {
		epsilon := getEpsilon(k, xy[t], xy, xIndices, yIndices)

		c_n_x := count(epsilon, xy[t], xy, xIndices)
		h_n_x := continuous.Harmonic(c_n_x + 1)
		// fmt.Println(fmt.Sprintf("c_n_xy %d h_n_xy %f", c_n_xy, h_n_xy))

		c_n_y := countY(epsilon, xy[t], xy, yIndices)
		h_n_y := continuous.Harmonic(c_n_y + 1)
		// fmt.Println(fmt.Sprintf("c_n_yz %d h_n_yz %f", c_n_yz, h_n_yz))

		r -= h_n_x + h_n_z

		bar.Increment()
	}

	bar.FinishPrint("Finished")

	r /= float64(len(xy))

	r += h_k + h_N

	return
}

// KraskovStoegbauerGrassberger2 is an implementation of the second
// algorithm presented in
// A. Kraskov, H. Stoegbauer, and P. Grassberger.
// Estimating mutual information. Phys. Rev. E, 69:066138, Jun 2004.
func KraskovStoegbauerGrassberger2(xy [][]float64, xIndices, yIndices []int64, k int64) (r float64) {

	r = 0.0

	h_k := continuous.Harmonic(k)
	h_N := continuous.Harmonic(len(xy))

	bar := pb.StartNew(len(xy))
	for t := 0; t < len(xy); t++ {
		epsilon := getEpsilon(k, xy[t], xy, xIndices, yIndices)

		c_n_x := count(epsilon, xy[t], xy, xIndices)
		h_n_x := continuous.Harmonic(c_n_x)

		c_n_y := countY(epsilon, xy[t], xy, yIndices)
		h_n_y := continuous.Harmonic(c_n_y)

		r -= h_n_x + h_n_z

		bar.Increment()
	}

	bar.FinishPrint("Finished")

	r /= float64(len(xy))

	r += h_k + h_N - 1.0/float64(k)

	return
}

// getEpsilon calculate epsilon_k(t) as defined by Frenzel & Pompe, 2007
// epsilon_k(t) is the distance of the k-th nearest neighbour. The function
// takes k, the point from which the distance is calculated (xyz), and the
// data from which the k-th nearest neighbour should be determined
func getEpsilon(k int64, xy []float64, data [][]float64, xIndices, yIndices []int64) float64 {
	distances := make([]float64, len(data), len(data))

	for t := 0; t < len(data); t++ {
		distances[t] = maxNorm2(xy, data[t], xIndices, yIndices)
	}

	sort.Float64s(distances)

	return distances[k-1] // we start to count at zero
}

func maxNorm2(a, b []float64, xIndices, yIndices []int64) float64 {
	xDist := distance(a, b, xIndices)
	yDist := distance(a, b, yIndices)
	return math.Max(xDist, yDist)
}
