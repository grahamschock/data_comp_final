package main

import (
	"os"
	"fmt"
	// "unsafe"
)

var current_source = make(map[int]int)

var grammar_rule = make(map[int][2]int)

var all_pairs = make(map[[2]int]int)

var new_rule_index = 1

//attempts to reuse similar words
func translate_to_int(source []byte) {
	for i := 0; i < len(source); i++ {
		current_source[i] = int(source[i] - 32)
		// fmt.Println(current_source[i])
	}
}

func generate_grammar() bool {
	max_pair_occur := 0
	var max_pair [2]int

	fmt.Println(len(current_source))

	for i := 0; i < len(current_source) - 1; i++ {
		first := current_source[i]
		second := current_source[i + 1]
		curr_pair := [2]int{first, second}

		num_occur := all_pairs[curr_pair]
		num_occur += 1
		if num_occur > max_pair_occur {
			max_pair_occur = num_occur
			max_pair[0] = current_source[i] // our grammar is the max_pair
			max_pair[1] = current_source[i + 1]
		}
		all_pairs[curr_pair] = num_occur
	}

	if max_pair_occur > 1 {
		//add the grammar rule
		grammar_rule[new_rule_index] = max_pair
		new_rule_index++
	} else {
		return false
	}

	counter := 0
	new_source := make(map[int]int)
	for index := 0; index < len(current_source); index++{
		first := current_source[index]
		second := current_source[index + 1]
		if first == max_pair[0] && second == max_pair[1] {
			new_source[counter] = new_rule_index - 1
			index++
		} else {
			new_source[counter] = current_source[index]
		}
		counter++
	}

	current_source = new_source
	return true
}

func Repair(sourcedir string) {
	source, err := os.ReadFile(sourcedir)
	if err != nil {
		panic(err)
	}
	translate_to_int(source)

	fmt.Println("Translated to ints")

	for generate_grammar() == true {
		fmt.Println("Generating grammar")
	}
	fmt.Printf("Compression ratio: %v\n", ((float64(len(source))) / (float64(( float64(new_rule_index) * 1.5) + float64((len(current_source)))))))
}
