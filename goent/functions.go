package main

import (
	"fmt"
	"os"

	"github.com/kzahedi/goent/dh"
	pb "gopkg.in/cheggaaa/pb.v1"
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
	zIndices []int64, zOffset int64,
	eta bool) [][]float64 {
	maxOffset := max3(xOffset, yOffset, zOffset)
	N := len(data) - int(maxOffset)
	r := make([][]float64, N, N)

	var bar *pb.ProgressBar

	if eta == true {
		fmt.Println("Extracting W', W, A from data")
		bar = pb.StartNew(N)
	}

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
		if eta == true {
			bar.Increment()
		}
	}

	if eta == true {
		bar.Increment()
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

func createIndices3(XIndices, YIndices, ZIndices []int64) ([]int64, []int64, []int64) {
	xIndices := make([]int64, len(XIndices), len(XIndices))
	yIndices := make([]int64, len(YIndices), len(YIndices))
	zIndices := make([]int64, len(ZIndices), len(ZIndices))
	index := 0
	for i := range XIndices {
		xIndices[i] = int64(index)
		index++
	}
	for i := range YIndices {
		yIndices[i] = int64(index)
		index++
	}
	for i := range ZIndices {
		zIndices[i] = int64(index)
		index++
	}
	return xIndices, yIndices, zIndices
}
