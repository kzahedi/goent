package main

import (
	"github.com/kzahedi/goent/dh"
	"github.com/kzahedi/goent/discrete"
)

////////////////////////////////////////////////////////////
// averaged functions
////////////////////////////////////////////////////////////

func conditionalEntropyContinuous(p goentParameters) (r float64) {
	panic("Continuous Entropy is not available yet")
}

func conditionalEntropyDiscrete(p goentParameters) (r float64) {
	data := dh.ReadData(p.Input)

	x := dh.ExtractColumns(data, p.XIndices)
	y := dh.ExtractColumns(data, p.YIndices)

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

	xd := dh.Discrestise(x, p.XBins, minx, maxx)
	yd := dh.Discrestise(y, p.YBins, miny, maxy)

	xuv := dh.Relabel(dh.MakeUnivariate(xd, p.XBins))
	yuv := dh.Relabel(dh.MakeUnivariate(yd, p.YBins))

	ddata := make([][]int, len(data), len(data))
	for i := 0; i < len(data); i++ {
		d := make([]int, 2, 2)
		d[0] = xuv[i]
		d[1] = yuv[i]
		ddata[i] = d
	}

	p2d := discrete.Emperical2D(ddata)
	r = discrete.ConditionalEntropyBase2(p2d)
	return
}

////////////////////////////////////////////////////////////
// state-dependent functions
////////////////////////////////////////////////////////////

func conditionalEntropyContinuousState(p goentParameters) (r float64) {
	panic("Continuous Entropy is not available yet")
}

func conditionalEntropyDiscreteState(p goentParameters) (r float64) {
	panic("Continuous Entropy discrete state is not available yet")
}

func conditionalEntropy(p goentParameters) (r float64) {
	if p.UseContinuous == true && p.UseStateDependent == false {
		r = conditionalEntropyContinuous(p)
	} else if p.UseContinuous == false && p.UseStateDependent == false {
		r = conditionalEntropyDiscrete(p)
	} else if p.UseContinuous == true && p.UseStateDependent == true {
		r = conditionalEntropyContinuousState(p)
	} else if p.UseContinuous == false && p.UseStateDependent == true {
		r = conditionalEntropyDiscreteState(p)
	}

	return
}
