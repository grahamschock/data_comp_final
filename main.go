package main

import (
	"fmt"
	"flag"
	"time"
)

func main() {
	algoPtr := flag.String("algo", "digram", "Algo to run")
	sourceTypePtr := flag.String("source", "english", "Source type to run")
	sourceSizePtr := flag.String("size", "50", "Size of source (mb)")
	flag.Parse()

	sourcedir := "data/" + *sourceTypePtr + "." + *sourceSizePtr + "MB"
	fmt.Println(sourcedir)
	start := time.Now()
	if (*algoPtr == "digram") {
		Digram(sourcedir)
	} else if (*algoPtr == "lz77") {
		Lz77(sourcedir)
	} else if (*algoPtr == "lz78") {
		Lz78(sourcedir)
	} else if (*algoPtr == "repair") {
		Repair(sourcedir)
	} else {
		Digraham(sourcedir)
	}
	fmt.Println(time.Since(start))
}
