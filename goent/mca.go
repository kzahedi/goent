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

func mcaContinuous(p goentParameters) (r float64) {
	data := dh.ReadData(p.Input)
	w2a1w1 := merge3Data(data, p.WIndices, 1, p.AIndices, 0, p.WIndices, 0)
	w2Indices, w1Indices, a1Indices := createW2W1A1(p.WIndices, p.AIndices)
	r = continuous.MorphologicalComputationA(w2a1w1, w2Indices, a1Indices, w1Indices, p.K, p.UseEta)
	return
}

func mcaDiscrete(p goentParameters) (r float64) {
	data := dh.ReadData(p.Input)
	w2a1w1 := merge3Data(data, p.WIndices, 1, p.AIndices, 0, p.WIndices, 0)
	w2Indices, w1Indices, a1Indices := createW2W1A1(p.WIndices, p.AIndices)
	ddata := discretise3D(w2a1w1, w2Indices, p.WBins, a1Indices, p.ABins, w1Indices, p.WBins)
	p3d := discrete.Emperical3D(ddata)
	r = discrete.MorphologicalComputationA(p3d)
	return
}

////////////////////////////////////////////////////////////
// state-dependent functions
////////////////////////////////////////////////////////////

func mcaContinuousState(p goentParameters) (r float64) {
	data := dh.ReadData(p.Input)
	w2a1w1 := merge3Data(data, p.WIndices, 1, p.AIndices, 0, p.WIndices, 0)
	w2Indices, w1Indices, a1Indices := createW2W1A1(p.WIndices, p.AIndices)
	s := cs.MorphologicalComputationA(w2a1w1, w2Indices, a1Indices, w1Indices, p.K, p.UseEta)
	writeData(p.Output, s)
	r = average(s)
	return
}

func mcaDiscreteState(p goentParameters) (r float64) {
	data := dh.ReadData(p.Input)
	w2a1w1 := merge3Data(data, p.WIndices, 1, p.AIndices, 0, p.WIndices, 0)
	w2Indices, w1Indices, a1Indices := createW2W1A1(p.WIndices, p.AIndices)
	dw2a1w1 := discretise3D(w2a1w1, w2Indices, p.WBins, a1Indices, p.ABins, w1Indices, p.WBins)
	s := ds.MorphologicalComputationA(dw2a1w1)
	writeData(p.Output, s)
	r = average(s)
	return
}

func mca(p goentParameters) (r float64) {
	if p.UseContinuous == true && p.UseStateDependent == false {
		r = mcaContinuous(p)
	} else if p.UseContinuous == false && p.UseStateDependent == false {
		r = mcaDiscrete(p)
	} else if p.UseContinuous == true && p.UseStateDependent == true {
		r = mcaContinuousState(p)
	} else if p.UseContinuous == false && p.UseStateDependent == true {
		r = mcaDiscreteState(p)
	}
	return
}
