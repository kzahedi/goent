package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

func getvalues(str string) []float64 {
	var r []float64
	values := strings.Split(str, ":")

	start, err := strconv.ParseFloat(values[0], 64)
	if err != nil {
		panic(err)
	}
	end := start
	delta := start - end + 1.0

	if len(values) == 3 {
		delta, err = strconv.ParseFloat(values[1], 64)
		if err != nil {
			panic(err)
		}
		end, err = strconv.ParseFloat(values[2], 64)
		if err != nil {
			panic(err)
		}
	}

	for v := start; v <= end; v += delta {
		r = append(r, v)
	}

	return r
}

func main() {

	muStr := flag.String("mu", "0:0.1:1", "mu values. can take list (1,2,3) to range with delta (0:0.1:1.0)")
	phiStr := flag.String("phi", "0:0.1:1", "phi values. can take list (1,2,3) to range with delta (0:0.1:1.0)")
	psiStr := flag.String("psi", "0:0.1:1", "phi values. can take list (1,2,3) to range with delta (0:0.1:1.0)")
	zetaStr := flag.String("zeta", "0:0.1:1", "phi values. can take list (1,2,3) to range with delta (0:0.1:1.0)")
	tauStr := flag.String("tau", "0:0.1:1", "phi values. can take list (1,2,3) to range with delta (0:0.1:1.0)")
	mc := flag.String("mc", "MC_W", "quantification to use: MC_W (soon: MC_A, MC_MI, MC_SY, MC_SY_NIS, MC_SY_GIS, MC_SY_SCGIS)")
	verbose := flag.Bool("v", false, "verbose")
	output := flag.String("o", "out.csv", "output file. default out.csv")

	flag.Parse()

	mu := getvalues(*muStr)
	phi := getvalues(*phiStr)
	psi := getvalues(*psiStr)
	zeta := getvalues(*zetaStr)
	tau := getvalues(*tauStr)

	if *verbose == true {
		fmt.Println("mu:", *muStr)
		fmt.Println("mu:", mu)
		fmt.Println("phi:", *phiStr)
		fmt.Println("phi:", phi)
		fmt.Println("psi:", *psiStr)
		fmt.Println("psi:", psi)
		fmt.Println("zeta:", *zetaStr)
		fmt.Println("zeta:", zeta)
		fmt.Println("tau:", *tauStr)
		fmt.Println("tau:", tau)
		fmt.Println("mc:", *mc)
		fmt.Println("out:", *output)
	}

}
