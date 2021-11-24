package main

import (
	"os"
	"fmt"
	// "unsafe"
)

type Triple78 struct{
	Offset int
	LastChar byte
}

var dict78 = make(map[string]bool)

//attempts to reuse similar words
func encode_lz78(source []byte) {
	var EncodingArr []Triple78
	var search []byte
	search = append(search, source[0])
	for currPos := 0; currPos < len(source); {
		if !(dict78[string(search)]) {
			//add to dict
			dict78[string(search)] = true

			//send encoding
			newEntry := Triple78{(currPos + 1) - len(search), search[len(search) - 1]}
			EncodingArr = append(EncodingArr, newEntry)
			search = []byte{source[currPos]}
		} else {
			search = append(search, source[currPos])
		}
		currPos++
	}
}

func Lz78(sourcedir string) {
	source, err := os.ReadFile(sourcedir)
	if err != nil {
		panic(err)
	}
	encode_lz78(source)
}
