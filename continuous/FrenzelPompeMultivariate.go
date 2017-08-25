package continuous

import (
	"math"
	"sort"

	pb "gopkg.in/cheggaaa/pb.v1"
)

// maxNorm3 computes the max-norm of two 3-dimensional vectors
//   maxnorm(a,b) = max( |a[0] - b[0]|, |a[1] - b[1]|, |a[2] - b[2]|)
func maxNorm3(a, b []float64, xIndices, yIndices, zIndices []int64) float64 {
	xDist := distance(a, b, xIndices)
	yDist := distance(a, b, yIndices)
	zDist := distance(a, b, zIndices)
	return math.Max(xDist, math.Max(yDist, zDist))
}

// getEpsilon calculate epsilon_k(t) as defined by Frenzel & Pompe, 2007
// epsilon_k(t) is the distance of the k-th nearest neighbour. The function
// takes k, the point from which the distance is calculated (xyz), and the
// data from which the k-th nearest neighbour should be determined
func getEpsilon(k int64, xyz []float64, data [][]float64, xIndices, yIndices, zIndices []int64) float64 {
	distances := make([]float64, len(data), len(data))

	for t := 0; t < len(data); t++ {
		distances[t] = maxNorm3(xyz, data[t], xIndices, yIndices, zIndices)
	}

	sort.Float64s(distances)

	return distances[k-1] // we start to count at zero
}

// countXY count the number of points for which the x and y coordinate is
// closer than epsilon, where the distance is measured by the max-norm
func countXY(epsilon float64, xyz []float64, data [][]float64, xIndices, yIndices []int64) (c int64) {

	for t := 0; t < len(data); t++ {
		if maxNorm2(xyz, data[t], xIndices, yIndices) < epsilon {
			c++
		}
	}

	return
}

// countYZ count the number of points for which the y and z coordinate is
// closer than epsilon, where the distance is measured by the max-norm
func countYZ(epsilon float64, xyz []float64, data [][]float64, yIndices, zIndices []int64) (c int64) {

	for t := 0; t < len(data); t++ {
		if maxNorm2(xyz, data[t], yIndices, zIndices) < epsilon {
			c++
		}
	}

	return
}

func maxNorm2(a, b []float64, xIndices, yIndices []int64) float64 {
	xDist := distance(a, b, xIndices)
	yDist := distance(a, b, yIndices)
	return math.Max(xDist, yDist)
}

// countZ count the number of points for which the z coordinate is
// closer than epsilon
func countZ(epsilon float64, xyz []float64, data [][]float64, zIndices []int64) (c int64) {
	for t := 0; t < len(data); t++ {
		if distance(xyz, data[t], zIndices) < epsilon {
			c++
		}
	}
	return
}

func distance(a, b []float64, indices []int64) float64 {
	d := 0.0
	for _, v := range indices {
		d += (a[v] - b[v]) * (a[v] - b[v])
	}
	return math.Sqrt(d)
}

// FrenzelPompeMultivariate is an implementation of
// Partial Mutual Information for Coupling Analysis of Multivariate Time Series
// Stefan Frenzel and Bernd Pompe
// Phys. Rev. Lett. 99, 204101 – Published 14 November 2007
// https://journals.aps.org/prl/abstract/10.1103/PhysRevLett.99.204101
func FrenzelPompeMultivariate(xyz [][]float64, xIndices, yIndices, zIndices []int64, k int64) (r float64) {

	r = 0.0

	h_k := Harmonic(k - 1)

	bar := pb.StartNew(len(xyz))
	for t := 0; t < len(xyz); t++ {
		epsilon := getEpsilon(k, xyz[t], xyz, xIndices, yIndices, zIndices)

		c_n_xy := countXY(epsilon, xyz[t], xyz, xIndices, yIndices)
		h_n_xy := Harmonic(c_n_xy)
		// fmt.Println(fmt.Sprintf("c_n_xy %d h_n_xy %f", c_n_xy, h_n_xy))

		c_n_yz := countYZ(epsilon, xyz[t], xyz, yIndices, zIndices)
		h_n_yz := Harmonic(c_n_yz)
		// fmt.Println(fmt.Sprintf("c_n_yz %d h_n_yz %f", c_n_yz, h_n_yz))

		c_n_z := countZ(epsilon, xyz[t], xyz, zIndices)
		h_n_z := Harmonic(c_n_z)
		// fmt.Println(fmt.Sprintf("c_n_z %d h_n_z %f", c_n_z, h_n_z))

		r += h_n_xy + h_n_yz - h_n_z

		bar.Increment()
	}

	bar.FinishPrint("Finished")

	r -= h_k

	r /= float64(len(xyz))

	return
}
