package main

import (
	"github.com/kzahedi/goent/dh"
	"github.com/kzahedi/goent/discrete"
)

////////////////////////////////////////////////////////////
// averaged functions
////////////////////////////////////////////////////////////

func entropyMLBCContinuous(p goentParameters) (r float64) {
	panic("Continuous EntropyMLBC is not available yet")
}

func entropyMLBCDiscrete(p goentParameters) (r float64) {
	data := dh.ReadData(p.Input)
	xd := discretise1D(data, p.XIndices, p.XBins)
	r = discrete.EntropyMLBCBase2(xd)
	return
}

////////////////////////////////////////////////////////////
// state-dependent functions
////////////////////////////////////////////////////////////

func entropyMLBCContinuousState(p goentParameters) (r float64) {
	panic("Continuous-state EntropyMLBC is not available yet")
}

func entropyMLBCDiscreteState(p goentParameters) (r float64) {
	panic("Discrete-state EntropyMLBC is not available yet")
}

func entropyMLBC(p goentParameters) (r float64) {
	if p.UseContinuous == true && p.UseStateDependent == false {
		r = entropyMLBCContinuous(p)
	} else if p.UseContinuous == false && p.UseStateDependent == false {
		r = entropyMLBCDiscrete(p)
	} else if p.UseContinuous == true && p.UseStateDependent == true {
		r = entropyMLBCContinuousState(p)
	} else if p.UseContinuous == false && p.UseStateDependent == true {
		r = entropyMLBCDiscreteState(p)
	}

	return
}
