package discrete

import "math"

// DiscrestiseVector takes a one-dimensional slice and discretises it.
// Min/max and number of bins must be provided.
func DiscrestiseVector(d []float64, bins int64, min, max float64) []int64 {
	r := make([]int64, len(d), len(d))
	domain := max - min
	for i := 0; i < len(d); i++ {
		r[i] = int64(math.Min(((d[i]-min)/domain)*float64(bins), float64(bins-1)))
	}
	return r
}

// Discrestise takes a two-dimensional slice and discretises it.
// Min/max and number of bins must be provided for each column. The first
// index of the data are the rows and the second index the columns, i.e.,
// d[r][c] is the data point in the r-th row and c-th column
func Discrestise(d [][]float64, bins []int64, min, max []float64) [][]int64 {
	rows := len(d)
	ret := make([][]int64, rows, rows)
	for r := 0; r < rows; r++ {
		cols := len(d[r])
		ret[r] = make([]int64, cols, cols)
		for c := 0; c < cols; c++ {
			domain := max[c] - min[c]
			ret[r][c] = int64(math.Min(((d[r][c]-min[c])/domain)*float64(bins[c]), float64(bins[c]-1)))
		}
	}
	return ret
}

// MakeUnivariate takes a two-dimensional discretised slice and returns
// a one-dimensional representation of it.
func MakeUnivariate(d [][]int64, bins []int64) []int64 {
	rows := len(d)
	cols := len(d[0])
	ret := make([]int64, rows, rows)

	f := int64(1)
	v := int64(0)
	for r := 0; r < rows; r++ {
		v = 0
		f = 1
		for c := 0; c < cols; c++ {
			if c > 0 {
				if bins[c-1] > 0 {
					f *= bins[c-1]
				}
			}
			v += f * d[r][c]
		}
		ret[r] = v
	}

	return ret
}

// Relabel takes a one-dimensional discretised slice and returns
// a one-dimensional representation of it, in which only
// consecutive values, i.e.,
// [1,7,10,15] will be converted into [0,1,2,3]
func Relabel(d []int64) []int64 {
	rows := len(d)
	ret := make([]int64, rows, rows)
	encountered := map[int64]bool{}
	unique := []int64{}

	for i, _ := range d {
		if encountered[d[i]] == false {
			encountered[d[i]] = true
			unique = append(unique, d[i])
		}
	}

	//
	for r := 0; r < rows; r++ {
		for i := 0; i < len(unique); i++ {
			if unique[i] == d[r] {
				ret[r] = int64(i)
				break
			}
		}
	}

	return ret
}
