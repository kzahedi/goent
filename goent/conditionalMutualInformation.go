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

func conditionalMutualInformationContinuous(p goentParameters) (r float64) {
	data := dh.ReadData(p.Input)
	r = continuous.FrenzelPompe(data, p.XIndices, p.YIndices, p.ZIndices, p.K, p.UseEta)
	return
}

func conditionalMutualInformationDiscrete(p goentParameters) (r float64) {
	data := dh.ReadData(p.Input)
	ddata := discretise3D(data, p.XIndices, p.XBins, p.YIndices, p.YBins, p.ZIndices, p.ZBins)
	p3d := discrete.Emperical3D(ddata)
	r = discrete.ConditionalMutualInformationBase2(p3d)
	return
}

////////////////////////////////////////////////////////////
// state-dependent functions
////////////////////////////////////////////////////////////

func conditionalMutualInformationContinuousState(p goentParameters) (r float64) {
	data := dh.ReadData(p.Input)
	s := cs.FrenzelPompe(data, p.XIndices, p.YIndices, p.ZIndices, p.K, p.UseEta)
	writeData(p.Output, s)
	r = average(s)
	return
}

func conditionalMutualInformationDiscreteState(p goentParameters) (r float64) {
	data := dh.ReadData(p.Input)
	ddata := discretise3D(data, p.XIndices, p.XBins, p.YIndices, p.YBins, p.ZIndices, p.ZBins)
	s := ds.ConditionalMutualInformationBase2(ddata)
	writeData(p.Output, s)
	r = average(s)
	return
}

func conditionalMutualInformation(p goentParameters) (r float64) {
	if p.UseContinuous == true && p.UseStateDependent == false {
		r = conditionalMutualInformationContinuous(p)
	} else if p.UseContinuous == false && p.UseStateDependent == false {
		r = conditionalMutualInformationDiscrete(p)
	} else if p.UseContinuous == true && p.UseStateDependent == true {
		r = conditionalMutualInformationContinuousState(p)
	} else if p.UseContinuous == false && p.UseStateDependent == true {
		r = conditionalMutualInformationDiscreteState(p)
	}
	return
}
