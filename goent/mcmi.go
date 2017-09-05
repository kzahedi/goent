package main

import (
	"github.com/kzahedi/goent/continuous"
	cs "github.com/kzahedi/goent/continuous/state"
	"github.com/kzahedi/goent/dh"
	"github.com/kzahedi/goent/discrete"
	ds "github.com/kzahedi/goent/discrete/state"
)

////////////////////////////////////////////////////////////
// continuous
////////////////////////////////////////////////////////////

func mcmiContinuous(p goentParameters) (r float64) {
	data := dh.ReadData(p.Input)
	w2Indices, w1Indices, s1Indices, a1Indices := createIndices4(p.WIndices, p.WIndices, p.SIndices, p.AIndices)
	w2w1s1a1 := merge4Data(data, p.WIndices, 1, p.WIndices, 0, p.SIndices, 0, p.AIndices, 0, false)
	r = continuous.MorphologicalComputationMI1(w2w1s1a1, w2Indices, w1Indices, s1Indices, a1Indices, p.K, p.UseEta)
	return
}

func mcmiContinuousState(p goentParameters) (r float64) {
	data := dh.ReadData(p.Input)
	w2Indices, w1Indices, s1Indices, a1Indices := createIndices4(p.WIndices, p.WIndices, p.SIndices, p.AIndices)
	w2w1s1a1 := merge4Data(data, p.WIndices, 1, p.WIndices, 0, p.SIndices, 0, p.AIndices, 0, false)
	s := cs.MorphologicalComputationMI1(w2w1s1a1, w2Indices, w1Indices, s1Indices, a1Indices, p.K, p.UseEta)
	writeData(p.Output, s)
	r = average(s)
	return
}

////////////////////////////////////////////////////////////
// discrete
////////////////////////////////////////////////////////////

func mcmiDiscrete(p goentParameters) (r float64) {
	data := dh.ReadData(p.Input)
	w2Indices, w1Indices, s1Indices, a1Indices := createIndices4(p.WIndices, p.WIndices, p.SIndices, p.AIndices)
	w2w1s1a1 := merge4Data(data, p.WIndices, 1, p.WIndices, 0, p.SIndices, 0, p.AIndices, 0, false)
	dw2w1 := discretise2D(w2w1s1a1, w2Indices, p.WBins, w1Indices, p.WBins)
	ds1a1 := discretise2D(w2w1s1a1, s1Indices, p.SBins, a1Indices, p.ABins)
	pw2w1 := discrete.Emperical2D(dw2w1)
	ps1a1 := discrete.Emperical2D(ds1a1)
	r = discrete.MorphologicalComputationMI(pw2w1, ps1a1)
	return
}

func mcmiDiscreteState(p goentParameters) (r float64) {
	data := dh.ReadData(p.Input)
	w2Indices, w1Indices, s1Indices, a1Indices := createIndices4(p.WIndices, p.WIndices, p.SIndices, p.AIndices)
	w2w1s1a1 := merge4Data(data, p.WIndices, 1, p.WIndices, 0, p.SIndices, 0, p.AIndices, 0, false)
	dw2w1s1a1 := discretise4D(w2w1s1a1, w2Indices, p.WBins, w1Indices, p.WBins, s1Indices, p.SBins, a1Indices, p.ABins)
	s := ds.MorphologicalComputationMI(dw2w1s1a1)
	writeData(p.Output, s)
	r = average(s)
	return
}

////////////////////////////////////////////////////////////
// main
////////////////////////////////////////////////////////////

func mcmi(p goentParameters) (r float64) {
	if p.UseContinuous == true && p.UseStateDependent == false {
		r = mcmiContinuous(p)
	} else if p.UseContinuous == false && p.UseStateDependent == false {
		r = mcmiDiscrete(p)
	} else if p.UseContinuous == true && p.UseStateDependent == true {
		r = mcmiContinuousState(p)
	} else if p.UseContinuous == false && p.UseStateDependent == true {
		r = mcmiDiscreteState(p)
	}
	return
}
