package main

import (
	"github.com/kzahedi/goent/dh"
	"github.com/kzahedi/goent/discrete"
)

////////////////////////////////////////////////////////////
// continuous
////////////////////////////////////////////////////////////

func mcpContinuous(p goentParameters) (r float64) {
	panic("Continuous MC_P is not available yet")
}

func mcpContinuousState(p goentParameters) (r float64) {
	panic("Continuous MC_P is not available yet")
}

////////////////////////////////////////////////////////////
// discrete
////////////////////////////////////////////////////////////

func mcpDiscrete(p goentParameters) (r float64) {
	data := dh.ReadData(p.Input)
	w2w1a1 := merge3Data(data, p.WIndices, 1, p.WIndices, 0, p.AIndices, 0, false)
	w2Indices, w1Indices, a1Indices := createIndices3(p.WIndices, p.WIndices, p.AIndices)
	ddata := discretise3D(w2w1a1, w2Indices, p.WBins, w1Indices, p.WBins, a1Indices, p.ABins)
	p3d := discrete.Emperical3D(ddata)
	r = discrete.MorphologicalComputationP(p3d, p.Iterations, p.UseEta)
	return
}

func mcpDiscreteState(p goentParameters) (r float64) {
	panic("Discrete-state MC_P is not available yet")
	// data := dh.ReadData(p.Input)
	// w2w1a1 := merge3Data(data, p.WIndices, 1, p.WIndices, 0, p.AIndices, 0, false)
	// w2Indices, w1Indices, a1Indices := createIndices3(p.WIndices, p.WIndices, p.AIndices)
	// ddata := discretise3D(w2w1a1, w2Indices, p.WBins, w1Indices, p.WBins, a1Indices, p.ABins)
	// s := ds.MorphologicalComputationP(ddata, p.UseEta)
	// writeData(p.Output, s)
	// r = average(s)
}

////////////////////////////////////////////////////////////
// main
////////////////////////////////////////////////////////////

func mcp(p goentParameters) (r float64) {
	if p.UseContinuous == true && p.UseStateDependent == false {
		r = mcpContinuous(p)
	} else if p.UseContinuous == false && p.UseStateDependent == false {
		r = mcpDiscrete(p)
	} else if p.UseContinuous == true && p.UseStateDependent == true {
		r = mcpContinuousState(p)
	} else if p.UseContinuous == false && p.UseStateDependent == true {
		r = mcpDiscreteState(p)
	}
	return
}
