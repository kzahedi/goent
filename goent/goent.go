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

	mcPtr := flag.String("m", "MI", "Information theoretic measure. Can be any of the following: Entropy, EntropyMLBC, EntropyChaoShen, Entropy_HorvitzThompson, ConditionalEntropy, MutualInformation, ConditionalMutualInformation, MC_W, MC_A, MC_WS, MC_WA, MC_P, MC_SY, MC_SY_NID")

	xIPtr := flag.String("xi", "", "Column indices for X")
	xBinsPtr := flag.String("xb", "", "Number of bins for X")

	yIPtr := flag.String("yi", "", "Column indices for Y")
	yBinsPtr := flag.String("yb", "", "Number of bins for Y")

	zIPtr := flag.String("zi", "", "Column indices for Z")
	zBinsPtr := flag.String("zb", "", "Number of bins for Z")

	wIPtr := flag.String("wi", "", "Column indices for W")
	wBinsPtr := flag.String("wb", "", "Number of bins for W")

	aIPtr := flag.String("ai", "", "Column indices for A")
	aBinsPtr := flag.String("ab", "", "Number of bins for A")

	sIPtr := flag.String("si", "", "Column indices for S")
	sBinsPtr := flag.String("sb", "", "Number of bins for S")

	flag.Parse()

	parameters := goentParameters{*inputPtr, *outputPtr,
		*useEta, *useContPtr, *useStatePtr,
		*mcPtr,
		parseInt(*basePtr), int64(*kPtr),
		parseNumberString(*xIPtr), parseBinsString(*xBinsPtr),
		parseNumberString(*yIPtr), parseBinsString(*yBinsPtr),
		parseNumberString(*zIPtr), parseBinsString(*zBinsPtr),
		parseNumberString(*wIPtr), parseBinsString(*wBinsPtr),
		parseNumberString(*aIPtr), parseBinsString(*aBinsPtr),
		parseNumberString(*sIPtr), parseBinsString(*sBinsPtr)}

	fmt.Println(parameters)

	r := 0.0
	switch parameters.Measure {
	case "ConditionalMutualInformation":
		r = conditionalMutualInformation(parameters)
	case "Entropy":
		r = entropy(parameters)
	case "MutualInformation":
		r = mi(parameters)
	case "EntropyMLBC":
		r = entropyMLBC(parameters)
	case "EntropyHorvitzThompson":
		r = entropyHorvitzThompson(parameters)
	case "EntropyChaoShen":
		r = entropyChaoShen(parameters)
	case "ConditionalEntropy":
		r = conditionalEntropy(parameters)
	case "MC_W":
		r = mcw(parameters)
	case "MC_A":
		r = mca(parameters)
	// case "MC_WS":
	// r = mcws(parameters)
	// case "MC_WA":
	// r = mcwa(parameters)
	// case "MC_P":
	// r = mcp(parameters)
	// case "MC_SY":
	// r = mcsy(parameters)
	// case "MC_SY_NID":
	// r = mcsynid(parameters)
	default:
		panic("unknown measure given")
	}
	fmt.Printf("Result = %f\n", r)
}
