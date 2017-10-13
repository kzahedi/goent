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
	Base              int
	K                 int
	Iterations        int
	XIndices          []int
	XBins             []int
	YIndices          []int
	YBins             []int
	ZIndices          []int
	ZBins             []int
	WIndices          []int
	WBins             []int
	AIndices          []int
	ABins             []int
	SIndices          []int
	SBins             []int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parseNumberString(str string) []int {
	var r []int
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

func parseBinsString(str string) []int {
	var r []int
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
			for i := 0; i < n; i++ {
				r = append(r, v)
			}
		}
	}

	return r
}

func parseInt(str string) int {
	n, err := strconv.ParseInt(str, 10, 64)
	check(err)
	return n
}
