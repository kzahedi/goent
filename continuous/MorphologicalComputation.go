package continuous

// MorphologicalComputationW [...]
func MorphologicalComputationW(w2w1a1 [][]float64, w2Indices, w1Indices, a1Indices []int, k int, eta bool) float64 {
	return FrenzelPompe(w2w1a1, w2Indices, w1Indices, a1Indices, k, eta)
}

// MorphologicalComputationA [...]
func MorphologicalComputationA(w2a1w1 [][]float64, w2Indices, a1Indices, w1Indices []int, k int, eta bool) float64 {
	return FrenzelPompe(w2a1w1, w2Indices, a1Indices, w1Indices, k, eta)
}

// MorphologicalComputationCW1 [...]
func MorphologicalComputationCW1(w2w1a1 [][]float64, w2Indices, w1Indices, a1Indices []int, k int, eta bool) float64 {
	return KraskovStoegbauerGrassberger1(w2w1a1, w2Indices, w1Indices, k, false) - KraskovStoegbauerGrassberger1(w2w1a1, w2Indices, a1Indices, k, eta)
}

// MorphologicalComputationCW2 [...]
func MorphologicalComputationCW2(w2w1a1 [][]float64, w2Indices, w1Indices, a1Indices []int, k int, eta bool) float64 {
	return KraskovStoegbauerGrassberger2(w2w1a1, w2Indices, w1Indices, k, false) - KraskovStoegbauerGrassberger2(w2w1a1, w2Indices, a1Indices, k, eta)
}

// MorphologicalComputationWA1 = I(W;{W,A}) - I(W';A)
func MorphologicalComputationWA1(w2w1a1 [][]float64, w2Indices, w1Indices, a1Indices []int, k int, eta bool) float64 {
	return FrenzelPompe(w2w1a1, w2Indices, w1Indices, a1Indices, k, false) - KraskovStoegbauerGrassberger1(w2w1a1, w2Indices, a1Indices, k, eta)
}

// MorphologicalComputationWA2 = I(W;{W,A}) - I(W';A)
func MorphologicalComputationWA2(w2w1a1 [][]float64, w2Indices, w1Indices, a1Indices []int, k int, eta bool) float64 {
	return FrenzelPompe(w2w1a1, w2Indices, w1Indices, a1Indices, k, false) - KraskovStoegbauerGrassberger2(w2w1a1, w2Indices, a1Indices, k, eta)
}

// MorphologicalComputationWS1 = I(W;{W,S}) - I(W';S)
func MorphologicalComputationWS1(w2w1s1 [][]float64, w2Indices, w1Indices, s1Indices []int, k int, eta bool) float64 {
	return FrenzelPompe(w2w1s1, w2Indices, w1Indices, s1Indices, k, false) - KraskovStoegbauerGrassberger1(w2w1s1, w2Indices, s1Indices, k, eta)
}

// MorphologicalComputationWS2 = I(W;{W,S}) - I(W';S)
func MorphologicalComputationWS2(w2w1s1 [][]float64, w2Indices, w1Indices, s1Indices []int, k int, eta bool) float64 {
	return FrenzelPompe(w2w1s1, w2Indices, w1Indices, s1Indices, k, false) - KraskovStoegbauerGrassberger2(w2w1s1, w2Indices, s1Indices, k, eta)
}

// MorphologicalComputationMI1 [...]
func MorphologicalComputationMI1(w2w1s1a1 [][]float64, w2Indices, w1Indices, s1Indices, a1Indices []int, k int, eta bool) float64 {
	return KraskovStoegbauerGrassberger1(w2w1s1a1, w2Indices, w1Indices, k, false) - KraskovStoegbauerGrassberger1(w2w1s1a1, s1Indices, a1Indices, k, eta)
}

// MorphologicalComputationMI2 [...]
func MorphologicalComputationMI2(w2w1s1a1 [][]float64, w2Indices, w1Indices, s1Indices, a1Indices []int, k int, eta bool) float64 {
	return KraskovStoegbauerGrassberger2(w2w1s1a1, w2Indices, w1Indices, k, false) - KraskovStoegbauerGrassberger2(w2w1s1a1, a1Indices, s1Indices, k, eta)
}
