package main

import (
	"flag"
	"fmt"
)

func main() {
	inputPtr := flag.String("i", "", "Input filename (must be a csv). Each column must be normalised between 0 and 1")
	outputPtr := flag.String("o", "", "Output filename (will be a csv)")

	useEta := flag.Bool("eta", false, "use estimated time ahead display")
	useContPtr := flag.Bool("c", false, "use measures for continuous spaces")
	useStatePtr := flag.Bool("s", false, "use measures for state-dependent analysis")
	basePtr := flag.String("base", "2", "Selection of the log-base to use (if available). Can be either 2 or e.")
	kPtr := flag.Int("k", 30, "k, for continuous measures")

	mcPtr := flag.String("m", "MI", "Information theoretic measure. Can be any of the following: H, H_MLBC, ChaoShen, HorvitzThompson, cH, I, cI, MC_W, MC_A, MC_WS, MC_WA, MC_P, MC_SY, MC_SY_NID")

	xIPtr := flag.String("xi", "", "Column indices for X")
	xBinsPtr := flag.String("xb", "", "Number of bins for X")

	yIPtr := flag.String("yi", "", "Column indices for Y")
	yBinsPtr := flag.String("yb", "", "Number of bins for Y")

	zIPtr := flag.String("zi", "", "Column indices for Z")
	zBinsPtr := flag.String("zb", "", "Number of bins for Z")

	w2IPtr := flag.String("w2i", "", "Column indices for W'")
	w2BinsPtr := flag.String("w2b", "", "Number of bins for W2")

	w1IPtr := flag.String("w1i", "", "Column indices for W")
	w1BinsPtr := flag.String("w1b", "", "Number of bins for W1")

	a1IPtr := flag.String("a1i", "", "Column indices for A")
	a1BinsPtr := flag.String("a1b", "", "Number of bins for A1")

	s1IPtr := flag.String("s1i", "", "Column indices for S")
	s1BinsPtr := flag.String("s1b", "", "Number of bins for S1")

	flag.Parse()

	parameters := goentParameters{*inputPtr, *outputPtr,
		*useEta, *useContPtr, *useStatePtr,
		*mcPtr,
		parseInt(*basePtr), int64(*kPtr),
		parseNumberString(*xIPtr), parseBinsString(*xBinsPtr),
		parseNumberString(*yIPtr), parseBinsString(*yBinsPtr),
		parseNumberString(*zIPtr), parseBinsString(*zBinsPtr),
		parseNumberString(*w2IPtr), parseBinsString(*w2BinsPtr),
		parseNumberString(*w1IPtr), parseBinsString(*w1BinsPtr),
		parseNumberString(*a1IPtr), parseBinsString(*a1BinsPtr),
		parseNumberString(*s1IPtr), parseBinsString(*s1BinsPtr)}

	fmt.Println(parameters)

	r := 0.0
	switch parameters.Measure {
	case "cI":
		r = cmi(parameters)
	// case "H":
	// return entropy(p)
	// case "H_MLBC":
	// return entropyMLBC(p)
	// case "ChaoShen":
	// return entropyCS(p)
	// case "HorvitzThompson":
	// return entropyHT(p)
	// case "cH":
	// return conditionalEntropy(p)
	// case "I":
	// return mi(p)
	// case "MC_W":
	// return mcw(p)
	// case "MC_A":
	// return mca(p)
	// case "MC_WS":
	// return mcws(p)
	// case "MC_WA":
	// return mcwa(p)
	// case "MC_P":
	// return mcp(p)
	// case "MC_SY":
	// return mcsy(p)
	// case "MC_SY_NID":
	// return mcsynid(p)
	default:
		panic("unknown measure given")
	}
	fmt.Printf("Result = %f\n", r)
}
