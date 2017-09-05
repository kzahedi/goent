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

func miContinuous(p goentParameters) (r float64) {
	data := dh.ReadData(p.Input)
	r = continuous.KraskovStoegbauerGrassberger1(data, p.XIndices, p.YIndices, p.K, p.UseEta)
	return
}

func miDiscrete(p goentParameters) (r float64) {
	data := dh.ReadData(p.Input)
	ddata := discretise2D(data, p.XIndices, p.XBins, p.YIndices, p.YBins)

	p2d := discrete.Emperical2D(ddata)
	r = discrete.MutualInformationBase2(p2d)
	return
}

////////////////////////////////////////////////////////////
// state-dependent functions
////////////////////////////////////////////////////////////

func miContinuousState(p goentParameters) (r float64) {
	data := dh.ReadData(p.Input)
	s := cs.KraskovStoegbauerGrassberger1(data, p.XIndices, p.YIndices, p.K, p.UseEta)
	writeData(p.Output, s)
	r = average(s)
	return
}

func miDiscreteState(p goentParameters) (r float64) {
	data := dh.ReadData(p.Input)
	ddata := discretise2D(data, p.XIndices, p.XBins, p.YIndices, p.YBins)
	s := ds.MutualInformationBase2(ddata)
	writeData(p.Output, s)
	r = average(s)
	return
}

func mi(p goentParameters) (r float64) {
	if p.UseContinuous == true && p.UseStateDependent == false {
		r = miContinuous(p)
	} else if p.UseContinuous == false && p.UseStateDependent == false {
		r = miDiscrete(p)
	} else if p.UseContinuous == true && p.UseStateDependent == true {
		r = miContinuousState(p)
	} else if p.UseContinuous == false && p.UseStateDependent == true {
		r = miDiscreteState(p)
	}
	return
}
