package dh

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// DiscrestiseVector takes a one-dimensional slice and discretises it.
// Min/max and number of bins must be provided.
func DiscrestiseVector(d []float64, bins int, min, max float64) []int {
	r := make([]int, len(d), len(d))
	domain := max - min
	for i := 0; i < len(d); i++ {
		r[i] = int(math.Min(((d[i]-min)/domain)*float64(bins), float64(bins-1)))
	}
	return r
}

// Discrestise takes a two-dimensional slice and discretises it.
// Min/max and number of bins must be provided for each column. The first
// index of the data are the rows and the second index the columns, i.e.,
// d[r][c] is the data point in the r-th row and c-th column
func Discrestise(d [][]float64, bins []int, min, max []float64) [][]int {
	rows := len(d)
	ret := make([][]int, rows, rows)
	for r := 0; r < rows; r++ {
		cols := len(d[r])
		ret[r] = make([]int, cols, cols)
		for c := 0; c < cols; c++ {
			domain := max[c] - min[c]
			ret[r][c] = int(math.Min(((d[r][c]-min[c])/domain)*float64(bins[c]), float64(bins[c]-1)))
		}
	}
	return ret
}

// MakeUnivariate takes a two-dimensional discretised slice and returns
// a one-dimensional representation of it.
func MakeUnivariate(d [][]int, bins []int) []int {
	rows := len(d)
	cols := len(d[0])
	ret := make([]int, rows, rows)

	var f int
	var v int

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
func Relabel(d []int) []int {
	rows := len(d)
	ret := make([]int, rows, rows)
	encountered := map[int]bool{}
	unique := []int{}

	for i := range d {
		if encountered[d[i]] == false {
			encountered[d[i]] = true
			unique = append(unique, d[i])
		}
	}

	//
	for r := 0; r < rows; r++ {
		for i := 0; i < len(unique); i++ {
			if unique[i] == d[r] {
				ret[r] = i
				break
			}
		}
	}

	return ret
}

// ReadData [...]
func ReadData(filename string) (r [][]float64) {

	f, _ := os.Open(filename)
	defer f.Close()

	lineCount := 0

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()

		record := strings.Split(line, ",")

		if strings.HasPrefix(record[0], "#") {
			continue
		}

		d := make([]float64, len(record), len(record))
		for i, v := range record {
			s, err := strconv.ParseFloat(v, 64)
			if err != nil {
				fmt.Println(err)
				return
			}
			d[i] = s
		}

		r = append(r, d)

		lineCount++
		fmt.Print(fmt.Sprintf("Line count: %d\r", lineCount))
	}

	fmt.Println(fmt.Sprintf("\nRead %d lines from %s", lineCount, filename))

	return
}

// ExtractColumns [...]
func ExtractColumns(data [][]float64, indices []int) [][]float64 {

	r := make([][]float64, len(data), len(data))

	for i := range data {
		d := make([]float64, len(indices), len(indices))
		for j, w := range indices {
			d[j] = data[i][w]
		}
		r[i] = d
	}
	return r
}
