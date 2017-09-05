package main

import (
	"github.com/kzahedi/goent/continuous"
	cs "github.com/kzahedi/goent/continuous/state"
	"github.com/kzahedi/goent/dh"
	"github.com/kzahedi/goent/discrete"
	ds "github.com/kzahedi/goent/discrete/state"
)

////////////////////////////////////////////////////////////
// averaged functions
////////////////////////////////////////////////////////////

func mcwsContinuous(p goentParameters) (r float64) {
	data := dh.ReadData(p.Input)
	w2Indices, w1Indices, s1Indices := createIndices3(p.WIndices, p.WIndices, p.SIndices)
	w2w1s1 := merge3Data(data, p.WIndices, 1, p.WIndices, 0, p.SIndices, 0, false)
	r = continuous.MorphologicalComputationWS1(w2w1s1, w2Indices, w1Indices, s1Indices, p.K, p.UseEta)
	return
}

func mcwsDiscrete(p goentParameters) (r float64) {
	data := dh.ReadData(p.Input)
	w2w1s1 := merge3Data(data, p.WIndices, 1, p.WIndices, 0, p.SIndices, 0, false)
	w2Indices, w1Indices, s1Indices := createIndices3(p.WIndices, p.WIndices, p.SIndices)
	ddata := discretise3D(w2w1s1, w2Indices, p.WBins, w1Indices, p.WBins, s1Indices, p.SBins)
	p3d := discrete.Emperical3D(ddata)
	r = discrete.MorphologicalComputationWS(p3d)
	return
}

////////////////////////////////////////////////////////////
// state-dependent functions
////////////////////////////////////////////////////////////

func mcwsContinuousState(p goentParameters) (r float64) {
	data := dh.ReadData(p.Input)
	w2w1s1 := merge3Data(data, p.WIndices, 1, p.WIndices, 0, p.SIndices, 0, false)
	w2Indices, w1Indices, s1Indices := createIndices3(p.WIndices, p.WIndices, p.SIndices)
	s := cs.MorphologicalComputationWS1(w2w1s1, w2Indices, w1Indices, s1Indices, p.K, p.UseEta)
	writeData(p.Output, s)
	r = average(s)
	return
}

func mcwsDiscreteState(p goentParameters) (r float64) {
	data := dh.ReadData(p.Input)
	w2w1s1 := merge3Data(data, p.WIndices, 1, p.WIndices, 0, p.SIndices, 0, false)
	w2Indices, w1Indices, s1Indices := createIndices3(p.WIndices, p.WIndices, p.SIndices)
	dw2w1s1 := discretise3D(w2w1s1, w2Indices, p.WBins, w1Indices, p.WBins, s1Indices, p.SBins)
	s := ds.MorphologicalComputationWS(dw2w1s1)
	writeData(p.Output, s)
	r = average(s)
	return
}

func mcws(p goentParameters) (r float64) {
	if p.UseContinuous == true && p.UseStateDependent == false {
		r = mcwsContinuous(p)
	} else if p.UseContinuous == false && p.UseStateDependent == false {
		r = mcwsDiscrete(p)
	} else if p.UseContinuous == true && p.UseStateDependent == true {
		r = mcwsContinuousState(p)
	} else if p.UseContinuous == false && p.UseStateDependent == true {
		r = mcwsDiscreteState(p)
	}
	return
}
