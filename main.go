package main

import "encoding/csv"
import "os"
import "io"
import "fmt"

func main() {
	var fp *os.File
	var err error

	fp, err = os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	reader := csv.NewReader(fp)
	reader.LazyQuotes = true

	rows := 0
	columns := 0
	nonzeroCount := 0
	for ; ; rows++ {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		columns = len(record)
		for _, v := range record {
			if v != "0" {
				nonzeroCount++
			}
		}
	}

	sparsity := 1.0 - float64(nonzeroCount)/float64(rows*columns)
	fmt.Printf("input sparsity ratio: %g\n", sparsity)
}
