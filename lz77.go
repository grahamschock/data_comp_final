package main

import (
	"fmt"
	"os"
	"strings"
	"bytes"
)

type Triple struct{
	Offset int
	Length int
	FirstChar byte
}

//attempts to reuse similar words
func encode(source []byte) (string) {
	var builder strings.Builder
	searchBuffer := source[0:1024]
	lookAhead := source[1024:2048]
	var EncodingArr []Triple

	for currPos := 0; currPos < len(source) - 2048; currPos++ {
		toEncode := containsSearch(searchBuffer, lookAhead)
		fmt.Println(toEncode)
		searchBuffer = source[0 + currPos : 1024 + currPos]
		lookAhead = source[1024 + currPos : 2048 + currPos]
	}
}

func containsSearch(searchBuffer []byte, lookAhead []byte) (toEncode Triple) {
	offset := 0
	length := 0
	firstChar := lookAhead[0]
	coutner_ahead := 0
	for i := 0; i < len(searchBuffer); i++ {
		counter_search := i
		for searchBuffer[counter_search] == lookAhead[counter_ahead] {
			offset = i
			length++
			counter_search++
			counter_ahead++
			firstChar = lookAhead[counter_ahead]
		}
	}
	toEncode.Offset = offset
	toEncode.Length = length
	toEncode.FirstChar = firstChar
	return toEncode

}

func lz77(sourcedir string) {
	source, err := os.ReadFile(sourcedir)
	if err != nil {
		panic(err)
	}
}
