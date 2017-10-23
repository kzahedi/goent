package main

import (
	"github.com/kzahedi/goent/dh"
	"github.com/kzahedi/goent/discrete"
)

////////////////////////////////////////////////////////////
// averaged functions
////////////////////////////////////////////////////////////

func entropyHorvitzThompsonContinuous(p goentParameters) (r float64) {
	panic("Continuous EntropyHorvitzThompson is not available yet")
}

func entropyHorvitzThompsonDiscrete(p goentParameters) (r float64) {
	data := dh.ReadData(p.Input)
	xd := discretise1D(data, p.XIndices, p.XBins)
	r = discrete.EntropyHorvitzThompsonBase2(xd)
	return
}

////////////////////////////////////////////////////////////
// state-dependent functions
////////////////////////////////////////////////////////////

func entropyHorvitzThompsonContinuousState(p goentParameters) (r float64) {
	panic("Continuous-state EntropyHorvitzThompson is not available yet")
}

func entropyHorvitzThompsonDiscreteState(p goentParameters) (r float64) {
	panic("Discrete-state EntropyHorvitzThompson is not available yet")
}

func entropyHorvitzThompson(p goentParameters) (r float64) {
	if p.UseContinuous == true && p.UseStateDependent == false {
		r = entropyHorvitzThompsonContinuous(p)
	} else if p.UseContinuous == false && p.UseStateDependent == false {
		r = entropyHorvitzThompsonDiscrete(p)
	} else if p.UseContinuous == true && p.UseStateDependent == true {
		r = entropyHorvitzThompsonContinuousState(p)
	} else if p.UseContinuous == false && p.UseStateDependent == true {
		r = entropyHorvitzThompsonDiscreteState(p)
	}

	return
}
