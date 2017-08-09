package goent

import "math"

func DiscrestiseVector(d []float64, bins int64, min, max float64) []int64 {
	r := make([]int64, len(d), len(d))
	domain := max - min
	for i := 0; i < len(d); i++ {
		r[i] = int64(math.Min(((d[i]-min)/domain)*float64(bins), float64(bins-1)))
	}
	return r
}
