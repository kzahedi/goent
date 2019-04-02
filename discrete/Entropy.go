package discrete

import (
	"math"

	"github.com/kzahedi/goent/sm"
)

// Entropy calculates the entropy of a probability distribution.
// It takes the log function as an additional parameter, so that the base
// can be chosen:
//   H(X) = -\sum_x p(x) lnFunc(p(x))
func Entropy(p []float64, ln lnFunc) float64 {
	var r float64
	for _, px := range p {
		if px > 0 {
			r -= px * ln(px)
		}
	}
	return r
}

// EntropyBaseE calculates the entropy of a probability distribution with base e
//   H(X) = -\sum_x p(x) ln(p(x))
func EntropyBaseE(p []float64) float64 {
	return Entropy(p, math.Log)
}

// EntropyBase2 calculates the entropy of a probability distribution with base 2
//   H(X) = -\sum_x p(x) log2(p(x))
func EntropyBase2(p []float64) float64 {
	return Entropy(p, math.Log2)
}

// EntropyMLBC is maximum likelihood estimator with bias correction
// It takes discretised data and the log
// function as input. Implemented from
// A. Chao and T.-J. Shen. Nonparametric estimation of shannon’s
// index of diversity when there are unseen species in sample.
// Environmental and Ecological Statistics, 10(4):429–443, 2003.
func EntropyMLBC(data []int, ln lnFunc) float64 {
	p := Empirical1D(data)
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

// EntropyMLBCBaseE is maximum likelihood estimator with bias correction
// It takes discretised data as input and
// returns the entropy in nats.
// Implemented from
// A. Chao and T.-J. Shen. Nonparametric estimation of shannon’s
// index of diversity when there are unseen species in sample.
// Environmental and Ecological Statistics, 10(4):429–443, 2003.
func EntropyMLBCBaseE(data []int) float64 {
	return EntropyMLBC(data, math.Log)
}

// EntropyMLBCBase2 is maximum likelihood estimator with bias correction
// It takes discretised data as input and
// returns the entropy in bits.
// Implemented from
// A. Chao and T.-J. Shen. Nonparametric estimation of shannon’s
// index of diversity when there are unseen species in sample.
// Environmental and Ecological Statistics, 10(4):429–443, 2003.
func EntropyMLBCBase2(data []int) float64 {
	return EntropyMLBC(data, math.Log2)
}

// EntropyHorvitzThompson is the Horvitz-Thompson entropy estimator.
// It takes discretised data and log function as input.
// Implemented from
// A. Chao and T.-J. Shen. Nonparametric estimation of shannon’s
// index of diversity when there are unseen species in sample.
// Environmental and Ecological Statistics, 10(4):429–443, 2003.
func EntropyHorvitzThompson(data []int, ln lnFunc) float64 {
	p := Empirical1D(data)
	n := float64(len(data))
	r := 0.0

	for _, v := range p {
		if v > 0.0 {
			N := v * ln(v)
			D := 1.0 - math.Pow(1.0-v, n)
			r -= N / D
		}
	}

	return r
}

// EntropyHorvitzThompsonBaseE is the Horvitz-Thompson entropy estimator.
// It takes discretised data as input and
// return the entropy in nats.
// Implemented from
// A. Chao and T.-J. Shen. Nonparametric estimation of shannon’s
// index of diversity when there are unseen species in sample.
// Environmental and Ecological Statistics, 10(4):429–443, 2003.
func EntropyHorvitzThompsonBaseE(data []int) float64 {
	return EntropyHorvitzThompson(data, math.Log)
}

// EntropyHorvitzThompsonBase2 is the Horvitz-Thompson entropy estimator.
// It takes discretised data as input and
// return the entropy in bits.
// Implemented from
// A. Chao and T.-J. Shen. Nonparametric estimation of shannon’s
// index of diversity when there are unseen species in sample.
// Environmental and Ecological Statistics, 10(4):429–443, 2003.
func EntropyHorvitzThompsonBase2(data []int) float64 {
	return EntropyHorvitzThompson(data, math.Log)
}

// EntropyChaoShen is the Chao-Shen entropy estimator. It take discretised data
// and the log-function as input
// Implemented from
// A. Chao and T.-J. Shen. Nonparametric estimation of shannon’s
// index of diversity when there are unseen species in sample.
// Environmental and Ecological Statistics, 10(4):429–443, 2003.
func EntropyChaoShen(data []int, ln lnFunc) float64 {
	n := float64(len(data))
	nrOfSingletons := 0.0
	histogram := map[int]float64{}
	for _, v := range data {
		histogram[v] += 1.0
	}

	p := make([]float64, len(histogram), len(histogram))

	var keys []int
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

	var z float64
	var r float64

	for i := range p {
		if p[i] > 0.0 {
			z = math.Pow((1.0 - p[i]), n)
			z = (1.0 - z)
			r -= p[i] * ln(p[i]) / z
		}
	}

	return r
}

// EntropyChaoShenBaseE is the Chao-Shen entropy estimator. It take discretised data
// and return nats.
// Implemented from
// A. Chao and T.-J. Shen. Nonparametric estimation of shannon’s
// index of diversity when there are unseen species in sample.
// Environmental and Ecological Statistics, 10(4):429–443, 2003.
func EntropyChaoShenBaseE(data []int) float64 {
	return EntropyChaoShen(data, math.Log)
}

// EntropyChaoShenBase2 is the Chao-Shen entropy estimator. It take discretised data
// and return bits.
// Implemented from
// A. Chao and T.-J. Shen. Nonparametric estimation of shannon’s
// index of diversity when there are unseen species in sample.
// Environmental and Ecological Statistics, 10(4):429–443, 2003.
func EntropyChaoShenBase2(data []int) float64 {
	return EntropyChaoShen(data, math.Log2)
}

// EntropySparse calculates the entropy of a probability distribution.
// It takes the log function as an additional parameter, so that the base
// can be chosen:
//   H(X) = -\sum_x p(x) lnFunc(p(x))
func EntropySparse(p sm.SparseMatrix, ln lnFunc) float64 {
	var r float64
	for _, x := range p.Values {
		if x > 0 {
			r -= x * ln(x)
		}
	}
	return r
}

// EntropyBaseESparse calculates the entropy of a probability distribution with base e
//   H(X) = -\sum_x p(x) ln(p(x))
func EntropyBaseESparse(p sm.SparseMatrix) float64 {
	return EntropySparse(p, math.Log)
}

// EntropyBase2Sparse calculates the entropy of a probability distribution with base 2
//   H(X) = -\sum_x p(x) log2(p(x))
func EntropyBase2Sparse(p sm.SparseMatrix) float64 {
	return EntropySparse(p, math.Log2)
}
