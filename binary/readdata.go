package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func readData(filename string) (r [][]float64) {

	f, _ := os.Open(filename)
	defer f.Close()
	reader := csv.NewReader(bufio.NewReader(f))

	lineCount := 0

	record, err := reader.Read()
	for err != io.EOF {

		if strings.HasPrefix(record[0], "#") {
			record, err = reader.Read()
			continue
		}

		d := make([]float64, len(record), len(record))
		for i, v := range record {
			s, err := strconv.ParseFloat(v, 64)
			if err != nil {
				fmt.Println(err)
				return
			}
			d[i] = s
		}

		r = append(r, d)

		lineCount++
		fmt.Print(fmt.Sprintf("Line count: %d\r", lineCount))

		record, err = reader.Read()
	}

	fmt.Println(fmt.Sprintf("\nRead %d lines from %s", lineCount, filename))

	return
}
