package goent

import (
	"math"
)

// H calculates the entropy of a probability distribution.
// It takes the log function as an additional parameter, so that the base
// can be chosen
// H(X) = -\sum_x p(x) lnFunc(p(x))
func H(p []float64, log lnFunc) float64 {
	var r float64
	for _, px := range p {
		if px > 0 {
			r -= px * log(px)
		}
	}
	return r
}

// Entropy calculates the entropy of a probability distribution with base e
// H(X) = -\sum_x p(x) ln(p(x))
func Entropy(p []float64) float64 {
	return H(p, math.Log)
}

// Entropy calculates the entropy of a probability distribution with base 2
// H(X) = -\sum_x p(x) log2(p(x))
func Entropy2(p []float64) float64 {
	return H(p, math.Log2)
}

// these are next
// # implemented from [1] (see below)
// function entropy_MLBC(data::Vector{Int64}, base::Number)
// p = fe1p(data)
// n = float(size(data)[1])
// S = float(size(p)[1])
// H = -sum([ p[x] > ϵ ? (p[x] * log(base, p[x])) : 0 for x=1:size(p)[1]])
// return H + (S-1) / (2.0 * n)
// end

// # implemented from [1] (see below)
// function entropy_HT(data::Vector{Int64}, base::Number)
// p = fe1p(data)
// n = size(data)[1]
// return -sum([ p[x] > ϵ ? ((p[x] * log(base, p[x])) / (1.0 - ((1.0 - p[x])^n))) : 0 for x=1:size(p)[1]])
// end

// function entropy_CS(data::Vector{Int64}, base::Number)
// m = maximum(data)
// n  = size(data)[1]
// c  = counts(data, 1:m)
// c = c ./ n
// # just to get rid of the numerical inaccuracies and make sure its a probability distribution
// s = sum(c)
// p = c ./ s
// C = 1.0 - float(sum(c .== 1)) / float(n)
// p = p .* C
// return -sum([ p[x] > ϵ ? ((p[x] * log(base, p[x])) / (1.0 - ((1.0 - p[x])^l))) : 0 for x=1:size(p)[1]])
// end
