package continuous

// MorphologicalComputationW [...]
func MorphologicalComputationW(w2w1a1 [][]float64, w2Indices, w1Indices, a1Indices []int, k int) float64 {
	return FrenzelPompe(w2w1a1, w2Indices, w1Indices, a1Indices, k, false)
}

// MorphologicalComputationA [...]
func MorphologicalComputationA(w2a1w1 [][]float64, w2Indices, a1Indices, w1Indices []int, k int) float64 {
	return FrenzelPompe(w2a1w1, w2Indices, a1Indices, w1Indices, k, false)
}

// MorphologicalComputationCW1 [...]
func MorphologicalComputationCW(w2w1a1 [][]float64, w2Indices, w1Indices, a1Indices []int, k int) float64 {
	return KraskovStoegbauerGrassberger1(w2w1a1, w2Indices, w1Indices, k, false) - KraskovStoegbauerGrassberger1(w2w1a1, w2Indices, a1Indices, k, false)
}

// MorphologicalComputationCW2 [...]
func MorphologicalComputationCW(w2w1a1 [][]float64, w2Indices, w1Indices, a1Indices []int, k int) float64 {
	return KraskovStoegbauerGrassberger2(w2w1a1, w2Indices, w1Indices, k, false) - KraskovStoegbauerGrassberger2(w2w1a1, w2Indices, a1Indices, k, false)
}

// MorphologicalComputationWA1 = I(W;{W,A}) - I(W';A)
func MorphologicalComputationWA(w2w1a1 [][]float64, w2Indices, w1Indices, a1Indices []int, k int) float64 {
	return FrenzelPompe(w2w1a1, w2Indices, w1Indices, a1Indices, k, false) - KraskovStoegbauerGrassberger1(w2a1, w2Indices, a1Indices, k, false)
}

// MorphologicalComputationWA2 = I(W;{W,A}) - I(W';A)
func MorphologicalComputationWA(w2w1a1 [][]float64, w2Indices, w1Indices, a1Indices []int, k int) float64 {
	return FrenzelPompe(w2w1a1, w2Indices, w1Indices, a1Indices, k, false) - KraskovStoegbauerGrassberger2(w2a1, w2Indices, a1Indices, k, false)
}

// MorphologicalComputationWS = I(W;{W,S}) - I(W';S)
func MorphologicalComputationWS(w2w1s1 [][]float64, w2Indices, w1Indices, s1Idices []int, k int) float64 {
	return FrenzelPompe(w2w1s1, w2Indices, w1Indices, s1Indices, k, false) - KraskovStoegbauerGrassberger(w2s1, w2Indices, s1Indices, k, false)
}

// MorphologicalComputationWS1 = I(W;{W,S}) - I(W';S)
func MorphologicalComputationWS(w2w1s1 [][]float64, w2Indices, w1Indices, s1Idices []int, k int) float64 {
	return FrenzelPompe(w2w1s1, w2Indices, w1Indices, s1Indices, k, false) - KraskovStoegbauerGrassberger1(w2s1, w2Indices, s1Indices, k, false)
}

// MorphologicalComputationWS2 = I(W;{W,S}) - I(W';S)
func MorphologicalComputationWS(w2w1s1 [][]float64, w2Indices, w1Indices, s1Idices []int, k int) float64 {
	return FrenzelPompe(w2w1s1, w2Indices, w1Indices, s1Indices, k, false) - KraskovStoegbauerGrassberger2(w2s1, w2Indices, s1Indices, k, false)
}

// MorphologicalComputationMI1 [...]
func MorphologicalComputationMI(w2w1s1a1 [][]float64, w2Indices, w1Indices, s1Idices, a1Indices []int, k int) float64 {
	return KraskovStoegbauerGrassberger1(w2w1s1a1, w2Indices, w1Indices, k, false) - KraskovStoegbauerGrassberger1(a1s1, a1Indices, s1Idices, k, false)
}

// MorphologicalComputationMI2 [...]
func MorphologicalComputationMI(w2w1s1a1 [][]float64, w2Indices, w1Indices, s1Idices, a1Indices []int, k int) float64 {
	return KraskovStoegbauerGrassberger2(w2w1s1a1, w2Indices, w1Indices, k, false) - KraskovStoegbauerGrassberger2(a1s1, a1Indices, s1Idices, k, false)
}