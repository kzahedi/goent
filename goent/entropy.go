package main

import (
	"github.com/kzahedi/goent/dh"
	"github.com/kzahedi/goent/discrete"
	ds "github.com/kzahedi/goent/discrete/state"
)

////////////////////////////////////////////////////////////
// averaged functions
////////////////////////////////////////////////////////////

func entropyContinuous(p goentParameters) (r float64) {
	panic("Continuous Entropy is not available yet")
}

func entropyDiscrete(p goentParameters) (r float64) {
	data := dh.ReadData(p.Input)
	xd := discretise1D(data, p.XIndices, p.XBins)
	p1d := discrete.Emperical1D(xd)
	r = discrete.EntropyBase2(p1d)
	return
}

////////////////////////////////////////////////////////////
// state-dependent functions
////////////////////////////////////////////////////////////

func entropyContinuousState(p goentParameters) (r float64) {
	panic("Continuous-state Entropy is not available yet")
}

func entropyDiscreteState(p goentParameters) (r float64) {
	data := dh.ReadData(p.Input)
	xd := discretise1D(data, p.XIndices, p.XBins)
	s := ds.EntropyBase2(xd)
	writeData(p.Output, s)
	r = average(s)
	return
}

func entropy(p goentParameters) (r float64) {
	if p.UseContinuous == true && p.UseStateDependent == false {
		r = entropyContinuous(p)
	} else if p.UseContinuous == false && p.UseStateDependent == false {
		r = entropyDiscrete(p)
	} else if p.UseContinuous == true && p.UseStateDependent == true {
		r = entropyContinuousState(p)
	} else if p.UseContinuous == false && p.UseStateDependent == true {
		r = entropyDiscreteState(p)
	}

	return
}
