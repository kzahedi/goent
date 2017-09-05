package main

import (
	"fmt"
	"os"

	"github.com/kzahedi/goent/dh"
)

func discretise1D(data [][]float64, indices, bins []int64) []int64 {

	x := dh.ExtractColumns(data, indices)

	minx := make([]float64, len(x), len(x))
	maxx := make([]float64, len(x), len(x))

	for i := 0; i < len(x); i++ {
		maxx[i] = 1.0
	}

	xd := dh.Discrestise(x, bins, minx, maxx)
	return dh.Relabel(dh.MakeUnivariate(xd, bins))
}

func discretise2D(data [][]float64, xIndices, xBins, yIndices, yBins []int64) [][]int64 {

	x := dh.ExtractColumns(data, xIndices)
	y := dh.ExtractColumns(data, yIndices)

	minx := make([]float64, len(x), len(x))
	maxx := make([]float64, len(x), len(x))

	miny := make([]float64, len(y), len(y))
	maxy := make([]float64, len(y), len(y))

	for i := 0; i < len(x); i++ {
		maxx[i] = 1.0
	}
	for i := 0; i < len(y); i++ {
		maxy[i] = 1.0
	}

	xd := dh.Discrestise(x, yBins, minx, maxx)
	yd := dh.Discrestise(y, yBins, miny, maxy)
	xuv := dh.Relabel(dh.MakeUnivariate(xd, xBins))
	yuv := dh.Relabel(dh.MakeUnivariate(yd, yBins))

	r := make([][]int64, len(data), len(data))
	for i := 0; i < len(data); i++ {
		d := make([]int64, 2, 2)
		d[0] = xuv[i]
		d[1] = yuv[i]
		r[i] = d
	}
	return r
}

func discretise3D(data [][]float64, xIndices, xBins, yIndices, yBins, zIndices, zBins []int64) [][]int64 {
	x := dh.ExtractColumns(data, xIndices)
	y := dh.ExtractColumns(data, yIndices)
	z := dh.ExtractColumns(data, zIndices)

	minx := make([]float64, len(xIndices), len(xIndices))
	maxx := make([]float64, len(xIndices), len(xIndices))

	miny := make([]float64, len(yIndices), len(yIndices))
	maxy := make([]float64, len(yIndices), len(yIndices))

	minz := make([]float64, len(zIndices), len(zIndices))
	maxz := make([]float64, len(zIndices), len(zIndices))

	for i := 0; i < len(xIndices); i++ {
		maxx[i] = 1.0
	}
	for i := 0; i < len(yIndices); i++ {
		maxy[i] = 1.0
	}
	for i := 0; i < len(zIndices); i++ {
		maxz[i] = 1.0
	}

	xd := dh.Discrestise(x, xBins, minx, maxx)
	yd := dh.Discrestise(y, yBins, miny, maxy)
	zd := dh.Discrestise(z, zBins, minz, maxz)

	xuv := dh.Relabel(dh.MakeUnivariate(xd, xBins))
	yuv := dh.Relabel(dh.MakeUnivariate(yd, yBins))
	zuv := dh.Relabel(dh.MakeUnivariate(zd, zBins))

	r := make([][]int64, len(data), len(data))
	for i := 0; i < len(data); i++ {
		d := make([]int64, 3, 3)
		d[0] = xuv[i]
		d[1] = yuv[i]
		d[2] = zuv[i]
		r[i] = d
	}
	return r
}

func max3(a, b, c int64) int64 {
	r := a
	if b > r {
		r = b
	}
	if c > r {
		r = c
	}
	return r
}

func merge3Data(data [][]float64,
	xIndices []int64, xOffset int64,
	yIndices []int64, yOffset int64,
	zIndices []int64, zOffset int64) [][]float64 {
	maxOffset := max3(xOffset, yOffset, zOffset)
	N := len(data) - int(maxOffset)
	r := make([][]float64, N, N)
	for i := 0; i < N; i++ {
		var d []float64
		for _, x := range xIndices {
			d = append(d, data[i+int(xOffset)][x])
		}
		for _, y := range yIndices {
			d = append(d, data[i+int(yOffset)][y])
		}
		for _, z := range zIndices {
			d = append(d, data[i+int(zOffset)][z])
		}
		r[i] = d
	}
	return r
}

func writeData(filename string, data []float64) {
	file, _ := os.Create(filename)
	defer file.Close()

	for _, value := range data {
		file.WriteString(fmt.Sprintf("%f\n", value))
	}
}

func average(data []float64) (r float64) {
	for _, v := range data {
		r += v
	}
	r /= float64(len(data))
	return
}

func createW2W1A1(WIndices, AIndices []int64) ([]int64, []int64, []int64) {
	w2Indices := make([]int64, len(WIndices), len(WIndices))
	w1Indices := make([]int64, len(WIndices), len(WIndices))
	a1Indices := make([]int64, len(AIndices), len(AIndices))
	index := 0
	for i := range WIndices {
		w2Indices[i] = int64(index)
		index++
	}
	for i := range WIndices {
		w1Indices[i] = int64(index)
		index++
	}
	for i := range AIndices {
		a1Indices[i] = int64(index)
		index++
	}
	return w2Indices, w1Indices, a1Indices
}
