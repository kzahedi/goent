package main

import (
	"github.com/kzahedi/goent/continuous"
	"github.com/kzahedi/goent/dh"
	"github.com/kzahedi/goent/discrete"
)

func cmiContinuous(p goentParameters) (r float64) {
	data := dh.ReadData(p.Input)
	r = continuous.FrenzelPompe(data, p.XIndices, p.YIndices, p.ZIndices, p.K, p.UseEta)
	return
}

func cmiDiscrete(p goentParameters) (r float64) {
	data := dh.ReadData(p.Input)

	x := dh.ExtractColumns(data, p.XIndices)
	y := dh.ExtractColumns(data, p.YIndices)
	z := dh.ExtractColumns(data, p.ZIndices)

	minx := make([]float64, len(x), len(x))
	maxx := make([]float64, len(x), len(x))
	miny := make([]float64, len(y), len(y))
	maxy := make([]float64, len(y), len(y))
	minz := make([]float64, len(z), len(z))
	maxz := make([]float64, len(z), len(z))

	for i := 0; i < len(x); i++ {
		maxx[i] = 1.0
	}
	for i := 0; i < len(y); i++ {
		maxy[i] = 1.0
	}
	for i := 0; i < len(z); i++ {
		maxz[i] = 1.0
	}

	xd := dh.Discrestise(x, p.XBins, minx, maxx)
	yd := dh.Discrestise(y, p.YBins, miny, maxy)
	zd := dh.Discrestise(z, p.ZBins, minz, maxz)

	xuv := dh.Relabel(dh.MakeUnivariate(xd, p.XBins))
	yuv := dh.Relabel(dh.MakeUnivariate(yd, p.YBins))
	zuv := dh.Relabel(dh.MakeUnivariate(zd, p.ZBins))

	ddata := make([][]int64, len(data), len(data))
	for i := 0; i < len(data); i++ {
		d := make([]int64, 3, 3)
		d[0] = xuv[i]
		d[1] = yuv[i]
		d[2] = zuv[i]
		ddata[i] = d
	}
	p3d := discrete.Emperical3D(ddata)
	r = discrete.ConditionalMutualInformationBase2(p3d)
	return
}

func cmi(p goentParameters) (r float64) {
	if p.UseContinuous {
		r = cmiContinuous(p)
	} else {
		r = cmiDiscrete(p)
	}
	return
}
