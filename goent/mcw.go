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

func mcwContinuous(p goentParameters) (r float64) {
	data := dh.ReadData(p.Input)
	w2w1a1 := merge3Data(data, p.WIndices, 1, p.WIndices, 0, p.AIndices, 0, false)
	w2Indices, w1Indices, a1Indices := createIndices3(p.WIndices, p.WIndices, p.AIndices)
	r = continuous.MorphologicalComputationW(w2w1a1, w2Indices, w1Indices, a1Indices, p.K, p.UseEta)
	return
}

func mcwDiscrete(p goentParameters) (r float64) {
	data := dh.ReadData(p.Input)
	w2w1a1 := merge3Data(data, p.WIndices, 1, p.WIndices, 0, p.AIndices, 0, false)
	w2Indices, w1Indices, a1Indices := createIndices3(p.WIndices, p.WIndices, p.AIndices)
	ddata := discretise3D(w2w1a1, w2Indices, p.WBins, w1Indices, p.WBins, a1Indices, p.ABins)
	p3d := discrete.Emperical3D(ddata)
	r = discrete.MorphologicalComputationW(p3d)
	return
}

////////////////////////////////////////////////////////////
// state-dependent functions
////////////////////////////////////////////////////////////

func mcwContinuousState(p goentParameters) (r float64) {
	data := dh.ReadData(p.Input)
	w2w1a1 := merge3Data(data, p.WIndices, 1, p.WIndices, 0, p.AIndices, 0, false)
	w2Indices, w1Indices, a1Indices := createIndices3(p.WIndices, p.WIndices, p.AIndices)
	s := cs.MorphologicalComputationW(w2w1a1, w2Indices, w1Indices, a1Indices, p.K, p.UseEta)
	writeData(p.Output, s)
	r = average(s)
	return
}

func mcwDiscreteState(p goentParameters) (r float64) {
	data := dh.ReadData(p.Input)
	w2w1a1 := merge3Data(data, p.WIndices, 1, p.WIndices, 0, p.AIndices, 0, false)
	w2Indices, w1Indices, a1Indices := createIndices3(p.WIndices, p.WIndices, p.AIndices)
	dw2w1a1 := discretise3D(w2w1a1, w2Indices, p.WBins, w1Indices, p.WBins, a1Indices, p.ABins)
	s := ds.MorphologicalComputationW(dw2w1a1)
	writeData(p.Output, s)
	r = average(s)
	return
}

func mcw(p goentParameters) (r float64) {
	if p.UseContinuous == true && p.UseStateDependent == false {
		r = mcwContinuous(p)
	} else if p.UseContinuous == false && p.UseStateDependent == false {
		r = mcwDiscrete(p)
	} else if p.UseContinuous == true && p.UseStateDependent == true {
		r = mcwContinuousState(p)
	} else if p.UseContinuous == false && p.UseStateDependent == true {
		r = mcwDiscreteState(p)
	}
	return
}
