package discrete

import (
	"math"
)

// entropy calculates the entropy of a probability distribution.
// It takes the log function as an additional parameter, so that the base
// can be chosen
// H(X) = -\sum_x p(x) LnFunc(p(x))
func entropy(p []float64, ln LnFunc) float64 {
	var r float64
	for _, px := range p {
		if px > 0 {
			r -= px * ln(px)
		}
	}
	return r
}

// Entropy calculates the entropy of a probability distribution with base e
//  H(X) = -\sum_x p(x) ln(p(x))
func Entropy(p []float64) float64 {
	return entropy(p, math.Log)
}

// Entropy2 calculates the entropy of a probability distribution with base 2
// H(X) = -\sum_x p(x) log2(p(x))
func Entropy2(p []float64) float64 {
	return entropy(p, math.Log2)
}

// entropyMLBC is maximum likelihood estimator with bias correction
// It takes discretised data and the log
// function as input. Implemented from
// A. Chao and T.-J. Shen. Nonparametric estimation of shannon’s
// index of diversity when there are unseen species in sample.
// Environmental and Ecological Statistics, 10(4):429–443, 2003.
func entropyMLBC(data []int64, ln LnFunc) float64 {
	p := Emperical1D(data)
	n := float64(len(data))
	S := float64(len(p))

	r := 0.0

	for _, v := range p {
		if v > 0.0 {
			r -= v * ln(v)
		}
	}

	return r + (S-1.0)/(2.0*n)

}

// EntropyMLBC is maximum likelihood estimator with bias correction
// It takes discretised data as input and
// returns the entropy in nats.
// Implemented from
// A. Chao and T.-J. Shen. Nonparametric estimation of shannon’s
// index of diversity when there are unseen species in sample.
// Environmental and Ecological Statistics, 10(4):429–443, 2003.
func EntropyMLBC(data []int64) float64 {
	return entropyMLBC(data, math.Log)
}

// EntropyMLBC2 is maximum likelihood estimator with bias correction
// It takes discretised data as input and
// returns the entropy in bits.
// Implemented from
// A. Chao and T.-J. Shen. Nonparametric estimation of shannon’s
// index of diversity when there are unseen species in sample.
// Environmental and Ecological Statistics, 10(4):429–443, 2003.
func EntropyMLBC2(data []int64) float64 {
	return entropyMLBC(data, math.Log2)
}

// entropyHorvitzThompson is the Horvitz-Thompson entropy estimator.
// It takes discretised data and log function as input.
// Implemented from
// A. Chao and T.-J. Shen. Nonparametric estimation of shannon’s
// index of diversity when there are unseen species in sample.
// Environmental and Ecological Statistics, 10(4):429–443, 2003.
func entropyHorvitzThompson(data []int64, ln LnFunc) float64 {
	p := Emperical1D(data)
	n := float64(len(data))
	r := 0.0

	for _, v := range p {
		if v > 0.0 {
			r -= v * ln(v) / (1.0 - (1.0 - math.Pow(v, n)))
		}
	}

	return r
}

// EntropyHorvitzThompson is the Horvitz-Thompson entropy estimator.
// It takes discretised data as input and
// return the entropy in nats.
// Implemented from
// A. Chao and T.-J. Shen. Nonparametric estimation of shannon’s
// index of diversity when there are unseen species in sample.
// Environmental and Ecological Statistics, 10(4):429–443, 2003.
func EntropyHorvitzThompson(data []int64) float64 {
	return entropyHorvitzThompson(data, math.Log)
}

// EntropyHorvitzThompson2 is the Horvitz-Thompson entropy estimator.
// It takes discretised data as input and
// return the entropy in bits.
// Implemented from
// A. Chao and T.-J. Shen. Nonparametric estimation of shannon’s
// index of diversity when there are unseen species in sample.
// Environmental and Ecological Statistics, 10(4):429–443, 2003.
func EntropyHorvitzThompson2(data []int64) float64 {
	return entropyHorvitzThompson(data, math.Log)
}

// entropyChaoShen is the Chao-Shen entropy estimator. It take discretised data
// and the log-function as input
// Implemented from
// A. Chao and T.-J. Shen. Nonparametric estimation of shannon’s
// index of diversity when there are unseen species in sample.
// Environmental and Ecological Statistics, 10(4):429–443, 2003.
func entropyChaoShen(data []int64, ln LnFunc) float64 {
	n := float64(len(data))
	nrOfSingletons := 0.0
	histogram := map[int64]float64{}
	for _, v := range data {
		histogram[v] += 1.0
	}

	p := make([]float64, len(histogram), len(histogram))

	var keys []int64
	for k, v := range histogram {
		keys = append(keys, k)
		if v == 1.0 {
			nrOfSingletons += 1.0
		}
	}

	if nrOfSingletons == n {
		nrOfSingletons -= 1.0
	}

	for i := range histogram {
		p[i] = histogram[keys[i]] / n
	}

	C := 1.0 - nrOfSingletons/n

	for i := range p {
		p[i] = p[i] * C
	}

	z := 0.0
	r := 0.0
	for i := range p {
		if p[i] > 0.0 {
			z = math.Pow((1.0 - p[i]), n)
			z = (1.0 - z)
			r -= p[i] * ln(p[i]) / z
		}
	}

	return r
}

// EntropyChaoShen is the Chao-Shen entropy estimator. It take discretised data
// and return nats.
// Implemented from
// A. Chao and T.-J. Shen. Nonparametric estimation of shannon’s
// index of diversity when there are unseen species in sample.
// Environmental and Ecological Statistics, 10(4):429–443, 2003.
func EntropyChaoShen(data []int64) float64 {
	return entropyChaoShen(data, math.Log)
}

// EntropyChaoShen2 is the Chao-Shen entropy estimator. It take discretised data
// and return bits.
// Implemented from
// A. Chao and T.-J. Shen. Nonparametric estimation of shannon’s
// index of diversity when there are unseen species in sample.
// Environmental and Ecological Statistics, 10(4):429–443, 2003.
func EntropyChaoShen2(data []int64) float64 {
	return entropyChaoShen(data, math.Log2)
}
