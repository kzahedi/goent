package continuous

// Harmonic as defined in Frenzel & Pompe, 2007
//   h_x = - \sum_n n^{-1}
func Harmonic(n int64) (r float64) {
	// Harmonic(1) = -C, see A. Kraskov, H. Stoeogbauer, and P. Grassberger.
	// Estimating mutual information. Phys. Rev. E, 69:066138, Jun 2004.
	r = -0.5772156649
	if n > 0 {
		for i := 2.0; i < float64(n); i++ {
			r -= 1.0 / i
		}
	}
	return
}
