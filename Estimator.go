package goent

import "fmt"

// Emperical1D is an empirical estimator for a one-dimensional
// probability distribution
func Emperical1D(d []int64) []float64 {
	max := int64(0)
	for _, v := range d {
		if v > max {
			max = v
		}
	}

	max += 1

	p := make([]float64, max, max)

	for _, v := range d {
		p[v] += 1.0
	}

	l := float64(len(d))
	for i, _ := range p {
		p[i] /= l
	}

	return p
}

// Emperical2D is an empirical estimator for a two-dimensional
// probability distribution
func Emperical2D(d [][]int64) [][]float64 {
	max := make([]int64, 2, 2)
	rows := len(d)
	for r := 0; r < rows; r++ {
		for c := 0; c < 2; c++ {
			if d[r][c] > max[c] {
				max[c] = d[r][c]
			}
		}
	}

	max[0] += 1
	max[1] += 1

	p := make([][]float64, max[0], max[0])
	for m := 0; m < int(max[0]); m++ {
		p[m] = make([]float64, max[1], max[1])
	}

	for r := 0; r < rows; r++ {
		p[d[r][0]][d[r][1]] += 1.0
	}

	fmt.Println(p)

	l := float64(len(d))
	for r := 0; r < int(max[0]); r++ {
		for c := 0; c < int(max[1]); c++ {
			p[r][c] /= l
		}
	}

	return p
}

// Emperical3D is an empirical estimator for a three-dimensional
// probability distribution
func Emperical3D(d [][]int64) [][][]float64 {
	max := make([]int64, 3, 3)
	rows := len(d)
	for r := 0; r < rows; r++ {
		for c := 0; c < 3; c++ {
			if d[r][c] > max[c] {
				max[c] = d[r][c]
			}
		}
	}

	max[0] += 1
	max[1] += 1
	max[2] += 1

	p := make([][][]float64, max[0], max[0])
	for m := 0; m < int(max[0]); m++ {
		p[m] = make([][]float64, max[1], max[1])
		for n := 0; n < int(max[1]); n++ {
			p[m][n] = make([]float64, max[2], max[2])
		}
	}

	fmt.Println(p)

	for r := 0; r < rows; r++ {
		p[d[r][0]][d[r][1]][d[r][2]] += 1.0
	}

	fmt.Println(p)

	l := float64(len(d))
	for a := 0; a < int(max[0]); a++ {
		for b := 0; b < int(max[1]); b++ {
			for c := 0; c < int(max[2]); c++ {
				p[a][b][c] /= l
			}
		}
	}

	return p
}
