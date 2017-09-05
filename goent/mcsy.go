package main

import (
	"github.com/kzahedi/goent/dh"
	"github.com/kzahedi/goent/discrete"
)

////////////////////////////////////////////////////////////
// continuous
////////////////////////////////////////////////////////////

func mcsyContinuous(p goentParameters) (r float64) {
	panic("Continuous MC_SY is not available yet")
	return
}

func mcsyContinuousState(p goentParameters) (r float64) {
	panic("Continuous MC_SY is not available yet")
	return
}

////////////////////////////////////////////////////////////
// discrete
////////////////////////////////////////////////////////////

func mcsyDiscrete(p goentParameters) (r float64) {
	data := dh.ReadData(p.Input)
	w2w1a1 := merge3Data(data, p.WIndices, 1, p.WIndices, 0, p.AIndices, 0, false)
	w2Indices, w1Indices, a1Indices := createIndices3(p.WIndices, p.WIndices, p.AIndices)
	ddata := discretise3D(w2w1a1, w2Indices, p.WBins, w1Indices, p.WBins, a1Indices, p.ABins)
	p3d := discrete.Emperical3D(ddata)
	r = discrete.MorphologicalComputationSY(p3d, p.Iterations, p.UseEta)
	return
}

func mcsyDiscreteState(p goentParameters) (r float64) {
	panic("Discrete-state MC_SY is not available yet")
	// data := dh.ReadData(p.Input)
	// w2w1a1 := merge3Data(data, p.WIndices, 1, p.WIndices, 0, p.AIndices, 0, false)
	// w2Indices, w1Indices, a1Indices := createIndices3(p.WIndices, p.WIndices, p.AIndices)
	// ddata := discretise3D(w2w1a1, w2Indices, p.WBins, w1Indices, p.WBins, a1Indices, p.ABins)
	// s := ds.MorphologicalComputationSY(ddata, p.UseEta)
	// writeData(p.Output, s)
	// r = average(s)
	return
}

////////////////////////////////////////////////////////////
// main
////////////////////////////////////////////////////////////

func mcsy(p goentParameters) (r float64) {
	if p.UseContinuous == true && p.UseStateDependent == false {
		r = mcsyContinuous(p)
	} else if p.UseContinuous == false && p.UseStateDependent == false {
		r = mcsyDiscrete(p)
	} else if p.UseContinuous == true && p.UseStateDependent == true {
		r = mcsyContinuousState(p)
	} else if p.UseContinuous == false && p.UseStateDependent == true {
		r = mcsyDiscreteState(p)
	}
	return
}
