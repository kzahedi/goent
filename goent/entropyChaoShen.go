package main

import (
	"github.com/kzahedi/goent/dh"
	"github.com/kzahedi/goent/discrete"
)

////////////////////////////////////////////////////////////
// averaged functions
////////////////////////////////////////////////////////////

func entropyChaoShenContinuous(p goentParameters) (r float64) {
	panic("Continuous EntropyChaoShen is not available yet")
}

func entropyChaoShenDiscrete(p goentParameters) (r float64) {
	data := dh.ReadData(p.Input)
	xd := discretise1D(data, p.XIndices, p.XBins)
	r = discrete.EntropyChaoShenBase2(xd)
	return
}

////////////////////////////////////////////////////////////
// state-dependent functions
////////////////////////////////////////////////////////////

func entropyChaoShenContinuousState(p goentParameters) (r float64) {
	panic("Continuous-state EntropyChaoShen is not available yet")
}

func entropyChaoShenDiscreteState(p goentParameters) (r float64) {
	panic("Discrete-state EntropyChaoShen is not available yet")
}

func entropyChaoShen(p goentParameters) (r float64) {
	if p.UseContinuous == true && p.UseStateDependent == false {
		r = entropyChaoShenContinuous(p)
	} else if p.UseContinuous == false && p.UseStateDependent == false {
		r = entropyChaoShenDiscrete(p)
	} else if p.UseContinuous == true && p.UseStateDependent == true {
		r = entropyChaoShenContinuousState(p)
	} else if p.UseContinuous == false && p.UseStateDependent == true {
		r = entropyChaoShenDiscreteState(p)
	}

	return
}
