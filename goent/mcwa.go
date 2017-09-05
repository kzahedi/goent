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

func mcwaContinuous(p goentParameters) (r float64) {
	data := dh.ReadData(p.Input)
	w2Indices, w1Indices, a1Indices := createIndices3(p.WIndices, p.WIndices, p.AIndices)
	w2w1a1 := merge3Data(data, p.WIndices, 1, p.WIndices, 0, p.AIndices, 0, false)
	r = continuous.MorphologicalComputationWA1(w2w1a1, w2Indices, w1Indices, a1Indices, p.K, p.UseEta)
	return
}

func mcwaDiscrete(p goentParameters) (r float64) {
	data := dh.ReadData(p.Input)
	w2w1a1 := merge3Data(data, p.WIndices, 1, p.WIndices, 0, p.AIndices, 0, false)
	w2Indices, w1Indices, a1Indices := createIndices3(p.WIndices, p.WIndices, p.AIndices)
	ddata := discretise3D(w2w1a1, w2Indices, p.WBins, w1Indices, p.WBins, a1Indices, p.ABins)
	p3d := discrete.Emperical3D(ddata)
	r = discrete.MorphologicalComputationWA(p3d)
	return
}

////////////////////////////////////////////////////////////
// state-dependent functions
////////////////////////////////////////////////////////////

func mcwaContinuousState(p goentParameters) (r float64) {
	data := dh.ReadData(p.Input)
	w2w1a1 := merge3Data(data, p.WIndices, 1, p.WIndices, 0, p.AIndices, 0, false)
	w2Indices, w1Indices, a1Indices := createIndices3(p.WIndices, p.WIndices, p.AIndices)
	s := cs.MorphologicalComputationWA1(w2w1a1, w2Indices, w1Indices, a1Indices, p.K, p.UseEta)
	writeData(p.Output, s)
	r = average(s)
	return
}

func mcwaDiscreteState(p goentParameters) (r float64) {
	data := dh.ReadData(p.Input)
	w2w1a1 := merge3Data(data, p.WIndices, 1, p.WIndices, 0, p.AIndices, 0, false)
	w2Indices, w1Indices, a1Indices := createIndices3(p.WIndices, p.WIndices, p.AIndices)
	dw2w1a1 := discretise3D(w2w1a1, w2Indices, p.WBins, w1Indices, p.WBins, a1Indices, p.SBins)
	s := ds.MorphologicalComputationWA(dw2w1a1)
	writeData(p.Output, s)
	r = average(s)
	return
}

func mcwa(p goentParameters) (r float64) {
	if p.UseContinuous == true && p.UseStateDependent == false {
		r = mcwaContinuous(p)
	} else if p.UseContinuous == false && p.UseStateDependent == false {
		r = mcwaDiscrete(p)
	} else if p.UseContinuous == true && p.UseStateDependent == true {
		r = mcwaContinuousState(p)
	} else if p.UseContinuous == false && p.UseStateDependent == true {
		r = mcwaDiscreteState(p)
	}
	return
}
