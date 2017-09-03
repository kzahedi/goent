package main

import (
	"flag"
	"fmt"
)

func main() {
	input := flag.String("i", "", "input filename (must be a csv)")
	useCont := flag.Bool("c", false, "use measures for continuous spaces")
	useState := flag.Bool("s", false, "use measures for state-dependent analysis")
	// bins := flag.String("b", "", "Number of bins. One values must be given for each column of the input file")

	flag.Parse()

	ok := true

	if *input == "" {
		fmt.Println("Please provide a input file with -i")
		ok = false
	}

	if ok == false {
		return
	}

	fmt.Println(fmt.Sprintf("Reading %s", *input))

	data := readData(*input)

	r := 0.0

	if *useCont == true {
		r = continuous(data, *useState)
	} else {
		r = discrete(data, *useState)
	}

	fmt.Println(fmt.Sprintf("Result %f", r))
}
