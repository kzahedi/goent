package main

import (
	"strconv"
	"strings"
)

type goentParameters struct {
	Input             string
	Output            string
	UseEta            bool
	UseContinuous     bool
	UseStateDependent bool
	Measure           string
	Base              int64
	K                 int64
	XIndices          []int64
	XBins             []int64
	YIndices          []int64
	YBins             []int64
	ZIndices          []int64
	ZBins             []int64
	W2Indices         []int64
	W2Bins            []int64
	W1Indices         []int64
	W1Bins            []int64
	A1Indices         []int64
	A1Bins            []int64
	S1Indices         []int64
	S1Bins            []int64
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parseNumberString(str string) []int64 {
	var r []int64
	if str == "" {
		return r
	}

	comma_values := strings.Split(str, ",")

	for _, c := range comma_values {
		values := strings.Split(c, ":")
		if len(values) == 1 {
			v, err := strconv.ParseInt(values[0], 10, 64)
			check(err)
			r = append(r, v)
		} else {
			start, err := strconv.ParseInt(values[0], 10, 64)
			check(err)
			end, err := strconv.ParseInt(values[1], 10, 64)
			check(err)
			for n := start; n <= end; n++ {
				r = append(r, n)
			}
		}
	}

	return r
}

func parseBinsString(str string) []int64 {
	var r []int64
	if str == "" {
		return r
	}

	comma_values := strings.Split(str, ",")

	for _, c := range comma_values {
		values := strings.Split(c, "x")
		if len(values) == 1 {
			v, err := strconv.ParseInt(values[0], 10, 64)
			check(err)
			r = append(r, v)
		} else {
			n, err := strconv.ParseInt(values[0], 10, 64)
			check(err)
			v, err := strconv.ParseInt(values[1], 10, 64)
			check(err)
			for i := 0; i < int(n); i++ {
				r = append(r, v)
			}
		}
	}

	return r
}

func parseInt(str string) int64 {
	n, err := strconv.ParseInt(str, 10, 64)
	check(err)
	return n
}
