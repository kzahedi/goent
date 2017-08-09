package goent

func MC_W(pw2w1a1 [][][]float64) float64 {
	return ConditionalMutualInformation2(pw2w1a1)
}

func MC_A(pw2a1w1 [][][]float64) float64 {
	return ConditionalMutualInformation2(pw2a1w1)
}

func MC_MI(pw2w1 [][]float64, ps1a1 [][]float64) float64 {
	return MutualInformation2(pw2w1) - MutualInformation2(ps1a1)
}
