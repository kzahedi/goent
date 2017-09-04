package main

import (
	"flag"
	"fmt"
)

func main() {
	cfgFile := flag.String("i", "", "input filename (must be a csv)")
	useContPtr := flag.Bool("c", false, "use measures for continuous spaces")
	useStatePtr := flag.Bool("s", false, "use measures for state-dependent analysis")
	basePtr := flag.String("base", "2", "Selection of the log-base to use (if available). Can be either 2 or e.")
	mcPtr := flag.String("m", "MI", "Information theoretic measure. Can be any of the following: H, H_MLBC, ChaoShen, HorvitzThompson, cH, I, MC_W, MC_A, MC_WS, MC_WA, MC_P, MC_SY, MC_SY_NID")

	xIPtr := flag.String("xi", "", "Column indices for X")
	xBinsPtr := flag.String("xb", "", "Number of bins for x")

	yIPtr := flag.String("yi", "", "Column indices for Y")
	yBinsPtr := flag.String("yb", "", "Number of bins for y")

	w2IPtr := flag.String("w2i", "", "Column indices for W'")
	w2BinsPtr := flag.String("w2b", "", "Number of bins for w2")

	w1IPtr := flag.String("w1i", "", "Column indices for W")
	w1BinsPtr := flag.String("w1b", "", "Number of bins for w1")

	a1IPtr := flag.String("a1i", "", "Column indices for A")
	a1BinsPtr := flag.String("a1b", "", "Number of bins for a1")

	s1IPtr := flag.String("s1i", "", "Column indices for S")
	s1BinsPtr := flag.String("s1b", "", "Number of bins for s1")

	flag.Parse()

	parameters := goentParameters{*cfgFile,
		*useContPtr, *useStatePtr, *mcPtr, parseInt(*basePtr),
		parseNumberString(*xIPtr), parseBinsString(*xBinsPtr),
		parseNumberString(*yIPtr), parseBinsString(*yBinsPtr),
		parseNumberString(*w2IPtr), parseBinsString(*w2BinsPtr),
		parseNumberString(*w1IPtr), parseBinsString(*w1BinsPtr),
		parseNumberString(*a1IPtr), parseBinsString(*a1BinsPtr),
		parseNumberString(*s1IPtr), parseBinsString(*s1BinsPtr)}

	fmt.Println(parameters)

	// ok := true

	// if *inputPtr == "" {
	// fmt.Println("Please provide a input file with -i")
	// ok = false
	// }

	// if ok == false {
	// return
	// }

	// fmt.Println(fmt.Sprintf("Reading %s", *inputPtr))

	// data := readData(*inputPtr)

	// r := 0.0

	// if *useContPtr == true {
	// r = continuous(data, parameters)
	// } else {
	// r = discrete(data, parameters)
	// }

	// fmt.Println(fmt.Sprintf("Result %f", r))

}
