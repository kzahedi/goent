package main

import (
	"fmt"
	"os"

	"github.com/kzahedi/goent/dh"
	pb "gopkg.in/cheggaaa/pb.v1"
)

func discretise1D(data [][]float64, indices, bins []int) []int {

	x := dh.ExtractColumns(data, indices)

	minx := make([]float64, len(x), len(x))
	maxx := make([]float64, len(x), len(x))

	for i := 0; i < len(x); i++ {
		maxx[i] = 1.0
	}

	xd := dh.Discrestise(x, bins, minx, maxx)
	return dh.Relabel(dh.MakeUnivariate(xd, bins))
}

func discretise2D(data [][]float64, xIndices, xBins, yIndices, yBins []int) [][]int {

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

	r := make([][]int, len(data), len(data))
	for i := 0; i < len(data); i++ {
		d := make([]int, 2, 2)
		d[0] = xuv[i]
		d[1] = yuv[i]
		r[i] = d
	}
	return r
}

func discretise3D(data [][]float64, xIndices, xBins, yIndices, yBins, zIndices, zBins []int) [][]int {
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

	r := make([][]int, len(data), len(data))
	for i := 0; i < len(data); i++ {
		d := make([]int, 3, 3)
		d[0] = xuv[i]
		d[1] = yuv[i]
		d[2] = zuv[i]
		r[i] = d
	}
	return r
}

func discretise4D(data [][]float64, xIndices, xBins, yIndices, yBins, zIndices, zBins, wIndices, wBins []int) [][]int {
	x := dh.ExtractColumns(data, xIndices)
	y := dh.ExtractColumns(data, yIndices)
	z := dh.ExtractColumns(data, zIndices)
	w := dh.ExtractColumns(data, wIndices)

	minx := make([]float64, len(xIndices), len(xIndices))
	maxx := make([]float64, len(xIndices), len(xIndices))

	miny := make([]float64, len(yIndices), len(yIndices))
	maxy := make([]float64, len(yIndices), len(yIndices))

	minz := make([]float64, len(zIndices), len(zIndices))
	maxz := make([]float64, len(zIndices), len(zIndices))

	minw := make([]float64, len(wIndices), len(wIndices))
	maxw := make([]float64, len(wIndices), len(wIndices))

	for i := 0; i < len(xIndices); i++ {
		maxx[i] = 1.0
	}
	for i := 0; i < len(yIndices); i++ {
		maxy[i] = 1.0
	}
	for i := 0; i < len(zIndices); i++ {
		maxz[i] = 1.0
	}
	for i := 0; i < len(wIndices); i++ {
		maxw[i] = 1.0
	}

	xd := dh.Discrestise(x, xBins, minx, maxx)
	yd := dh.Discrestise(y, yBins, miny, maxy)
	zd := dh.Discrestise(z, zBins, minz, maxz)
	wd := dh.Discrestise(w, wBins, minw, maxw)

	xuv := dh.Relabel(dh.MakeUnivariate(xd, xBins))
	yuv := dh.Relabel(dh.MakeUnivariate(yd, yBins))
	zuv := dh.Relabel(dh.MakeUnivariate(zd, zBins))
	wuv := dh.Relabel(dh.MakeUnivariate(wd, wBins))

	r := make([][]int, len(data), len(data))
	for i := 0; i < len(data); i++ {
		d := make([]int, 4, 4)
		d[0] = xuv[i]
		d[1] = yuv[i]
		d[2] = zuv[i]
		d[3] = wuv[i]
		r[i] = d
	}
	return r
}

func max2(a, b int) int {
	r := a
	if b > r {
		r = b
	}
	return r
}

func max3(a, b, c int) int {
	r := a
	if b > r {
		r = b
	}
	if c > r {
		r = c
	}
	return r
}

func max4(a, b, c, d int) int {
	r := a
	if b > r {
		r = b
	}
	if c > r {
		r = c
	}
	if d > r {
		r = d
	}
	return r
}

func merge3Data(data [][]float64,
	xIndices []int, xOffset int,
	yIndices []int, yOffset int,
	zIndices []int, zOffset int,
	eta bool) [][]float64 {
	maxOffset := max3(xOffset, yOffset, zOffset)
	N := len(data) - maxOffset
	r := make([][]float64, N, N)

	var bar *pb.ProgressBar

	if eta == true {
		fmt.Println("Extracting W', W, A from data")
		bar = pb.StartNew(N)
	}

	for i := 0; i < N; i++ {
		var d []float64
		for _, x := range xIndices {
			d = append(d, data[i+xOffset][x])
		}
		for _, y := range yIndices {
			d = append(d, data[i+yOffset][y])
		}
		for _, z := range zIndices {
			d = append(d, data[i+zOffset][z])
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

func merge2Data(data [][]float64,
	xIndices []int, xOffset int,
	yIndices []int, yOffset int,
	eta bool) [][]float64 {
	maxOffset := max2(xOffset, yOffset)
	N := len(data) - maxOffset
	r := make([][]float64, N, N)

	var bar *pb.ProgressBar

	if eta == true {
		fmt.Println("Extracting W', W, A from data")
		bar = pb.StartNew(N)
	}

	for i := 0; i < N; i++ {
		var d []float64
		for _, x := range xIndices {
			d = append(d, data[i+xOffset][x])
		}
		for _, y := range yIndices {
			d = append(d, data[i+yOffset][y])
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

func merge4Data(data [][]float64,
	xIndices []int, xOffset int,
	yIndices []int, yOffset int,
	zIndices []int, zOffset int,
	wIndices []int, wOffset int,
	eta bool) [][]float64 {
	maxOffset := max4(xOffset, yOffset, zOffset, wOffset)
	N := len(data) - maxOffset
	r := make([][]float64, N, N)

	var bar *pb.ProgressBar

	if eta == true {
		fmt.Println("Extracting W', W, A from data")
		bar = pb.StartNew(N)
	}

	for i := 0; i < N; i++ {
		var d []float64
		for _, x := range xIndices {
			d = append(d, data[i+xOffset][x])
		}
		for _, y := range yIndices {
			d = append(d, data[i+yOffset][y])
		}
		for _, z := range zIndices {
			d = append(d, data[i+zOffset][z])
		}
		for _, w := range wIndices {
			d = append(d, data[i+wOffset][w])
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

func createIndices3(XIndices, YIndices, ZIndices []int) ([]int, []int, []int) {
	xIndices := make([]int, len(XIndices), len(XIndices))
	yIndices := make([]int, len(YIndices), len(YIndices))
	zIndices := make([]int, len(ZIndices), len(ZIndices))
	index := 0
	for i := range XIndices {
		xIndices[i] = index
		index++
	}
	for i := range YIndices {
		yIndices[i] = index
		index++
	}
	for i := range ZIndices {
		zIndices[i] = index
		index++
	}
	return xIndices, yIndices, zIndices
}

func createIndices4(XIndices, YIndices, ZIndices, WIndices []int) ([]int, []int, []int, []int) {
	xIndices := make([]int, len(XIndices), len(XIndices))
	yIndices := make([]int, len(YIndices), len(YIndices))
	zIndices := make([]int, len(ZIndices), len(ZIndices))
	wIndices := make([]int, len(WIndices), len(WIndices))
	index := 0
	for i := range XIndices {
		xIndices[i] = index
		index++
	}
	for i := range YIndices {
		yIndices[i] = index
		index++
	}
	for i := range ZIndices {
		zIndices[i] = index
		index++
	}
	for i := range WIndices {
		wIndices[i] = index
		index++
	}
	return xIndices, yIndices, zIndices, wIndices
}
