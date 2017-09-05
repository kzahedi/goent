package state

// MorphologicalComputationW [...]
func MorphologicalComputationW(w2w1a1 [][]float64, w2Indices, w1Indices, a1Indices []int64, k int64, eta bool) []float64 {
	return FrenzelPompe(w2w1a1, w2Indices, w1Indices, a1Indices, k, eta)
}

// MorphologicalComputationA [...]
func MorphologicalComputationA(w2a1w1 [][]float64, w2Indices, a1Indices, w1Indices []int64, k int64, eta bool) []float64 {
	return FrenzelPompe(w2a1w1, w2Indices, a1Indices, w1Indices, k, eta)
}

// MorphologicalComputationCW1 [...]
func MorphologicalComputationCW1(w2w1a1 [][]float64, w2Indices, w1Indices, a1Indices []int64, k int64, eta bool) []float64 {
	r1 := KraskovStoegbauerGrassberger1(w2w1a1, w2Indices, w1Indices, k, eta)
	r2 := KraskovStoegbauerGrassberger1(w2w1a1, w2Indices, a1Indices, k, eta)
	r := make([]float64, len(r1), len(r2))
	for i := range r1 {
		r[i] = r1[i] - r2[i]
	}
	return r
}

// MorphologicalComputationCW2 [...]
func MorphologicalComputationCW2(w2w1a1 [][]float64, w2Indices, w1Indices, a1Indices []int64, k int64, eta bool) []float64 {
	r1 := KraskovStoegbauerGrassberger2(w2w1a1, w2Indices, w1Indices, k, eta)
	r2 := KraskovStoegbauerGrassberger2(w2w1a1, w2Indices, a1Indices, k, eta)
	r := make([]float64, len(r1), len(r2))
	for i := range r1 {
		r[i] = r1[i] - r2[i]
	}
	return r
}

// MorphologicalComputationWA1 = I(W;{W,A}) - I(W';A)
func MorphologicalComputationWA1(w2w1a1 [][]float64, w2Indices, w1Indices, a1Indices []int64, k int64, eta bool) []float64 {
	r1 := FrenzelPompe(w2w1a1, w2Indices, w1Indices, a1Indices, k, eta)
	r2 := KraskovStoegbauerGrassberger1(w2w1a1, w2Indices, a1Indices, k, eta)
	r := make([]float64, len(r1), len(r2))
	for i := range r1 {
		r[i] = r1[i] - r2[i]
	}
	return r
}

// MorphologicalComputationWA2 = I(W;{W,A}) - I(W';A)
func MorphologicalComputationWA(w2w1a1 [][]float64, w2Indices, w1Indices, a1Indices []int64, k int64, eta bool) []float64 {
	r1 := FrenzelPompe(w2w1a1, w2Indices, w1Indices, a1Indices, k, eta)
	r2 := KraskovStoegbauerGrassberger2(w2w1a1, w2Indices, a1Indices, k, eta)
	r := make([]float64, len(r1), len(r2))
	for i := range r1 {
		r[i] = r1[i] - r2[i]
	}
	return r
}

// MorphologicalComputationWS1 = I(W;{W,S}) - I(W';S)
func MorphologicalComputationWS1(w2w1s1 [][]float64, w2Indices, w1Indices, s1Indices []int64, k int64, eta bool) []float64 {
	r1 := FrenzelPompe(w2w1s1, w2Indices, w1Indices, s1Indices, k, eta)
	r2 := KraskovStoegbauerGrassberger1(w2w1s1, w2Indices, s1Indices, k, eta)
	r := make([]float64, len(r1), len(r2))
	for i := range r1 {
		r[i] = r1[i] - r2[i]
	}
	return r
}

// MorphologicalComputationWS2 = I(W;{W,S}) - I(W';S)
func MorphologicalComputationWS2(w2w1s1 [][]float64, w2Indices, w1Indices, s1Indices []int64, k int64, eta bool) []float64 {
	r1 := FrenzelPompe(w2w1s1, w2Indices, w1Indices, s1Indices, k, eta)
	r2 := KraskovStoegbauerGrassberger2(w2w1s1, w2Indices, s1Indices, k, eta)
	r := make([]float64, len(r1), len(r2))
	for i := range r1 {
		r[i] = r1[i] - r2[i]
	}
	return r
}

// MorphologicalComputationMI1 [...]
func MorphologicalComputationMI1(w2w1s1a1 [][]float64, w2Indices, w1Indices, s1Indices, a1Indices []int64, k int64, eta bool) []float64 {
	r1 := KraskovStoegbauerGrassberger1(w2w1s1a1, w2Indices, w1Indices, k, eta)
	r2 := KraskovStoegbauerGrassberger1(w2w1s1a1, a1Indices, s1Indices, k, eta)
	r := make([]float64, len(r1), len(r2))
	for i := range r1 {
		r[i] = r1[i] - r2[i]
	}
	return r
}

// MorphologicalComputationMI2 [...]
func MorphologicalComputationMI2(w2w1s1a1 [][]float64, w2Indices, w1Indices, s1Indices, a1Indices []int64, k int64, eta bool) []float64 {
	r1 := KraskovStoegbauerGrassberger2(w2w1s1a1, w2Indices, w1Indices, k, eta)
	r2 := KraskovStoegbauerGrassberger2(w2w1s1a1, a1Indices, s1Indices, k, eta)
	r := make([]float64, len(r1), len(r2))
	for i := range r1 {
		r[i] = r1[i] - r2[i]
	}
	return r
}
