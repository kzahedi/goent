package main

func continuous(data [][]float64, p goentParameters) float64 {
	if p.UseStateDependent == true {
		return 10.0
	}
	return 20.0
}
