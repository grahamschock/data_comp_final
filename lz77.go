package main

import (
	"fmt"
	"os"
	// "strings"
)

type Triple struct{
	Offset int
	Length int
	FirstChar byte
}

var max_len = 0

//attempts to reuse similar words
func encode_lz77(source []byte) {
	// var builder strings.Builder
	searchBuffer := source[0:2048]
	lookAhead := source[2049:4096]
	var EncodingArr []Triple

	for currPos := 0; currPos < len(source) - 5000; {
		toEncode := containsSearch(searchBuffer, lookAhead)
		currPos += (toEncode.Length + 1)
		EncodingArr = append(EncodingArr, toEncode)
		searchBuffer = source[0 + currPos : 2048 + currPos]
		lookAhead = source[2049 + currPos : 4096 + currPos]
	}
	fmt.Printf("Compression ratio: %v\n", ((float64(len(source))) / float64(len(EncodingArr) * 2)))
}

func containsSearch(searchBuffer []byte, lookAhead []byte) (toEncode Triple) {
	offset := 0
	length := 0
	max_length := 0
	max_offset := 0
	var max_char byte
	firstChar := lookAhead[0]
	for i := 0; i < len(searchBuffer); i++ {
		counter_search := i
		counter_ahead := 0
		length = 0
		for searchBuffer[counter_search] == lookAhead[counter_ahead] && (counter_search + 1 < len(searchBuffer)) {
			offset = i
			length++
			counter_search++
			counter_ahead++
			firstChar = lookAhead[counter_ahead]
			if length > max_length {
				max_length = length
				max_char = firstChar
				max_offset = offset
			}

		}
	}
	if(max_length > max_len) {
		max_len = max_length
	}
	toEncode.Offset = max_offset
	toEncode.Length = max_length
	toEncode.FirstChar = max_char
	return toEncode

}

func Lz77(sourcedir string) {
	source, err := os.ReadFile(sourcedir)
	if err != nil {
		panic(err)
	}
	encode_lz77(source)
}
