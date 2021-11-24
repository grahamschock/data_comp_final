package main
import (
	"fmt"
	"sort"
	"os"
	"matchr"
)

//Reads a two character input

//Assume the dictionary is 256
var gdict [256]string
var gdict_map map[string]int

//1st 94 is printable ASCII
//2nd 162 is the frequently most used pairs of characters
func ggenASCIIDict() {
	for i := 32; i < 127; i++ {
		dict[(i - 32)] = string(i)
	}
}

func topKCommonPairs(source []byte, k int) (topK []string){
	m := make(map[string]int) //index = pair, value = num times occured
	for i := 0; i < len(source); i += 2 {
		m[string(source[i:i+2])]++
	}

	topK = make([]string, k)

	pairs := make(PairList, len(m))

	i := 0
	for k, v := range m {
		pairs[i] = Pair{k, v}
		i++
	}

	sort.Sort(sort.Reverse(pairs))
	for i := 0; i < len(pairs) && i < k; i++ {
		topK[i] = pairs[i].Key
	}
	return topK
}

func Digraham(sourcedir string) {
	genASCIIDict()
	source, err := os.ReadFile(sourcedir)
	if err != nil {
		panic(err)
	}
	k := 0
	size := 0
	distance := matchr.DamerauLevenshtein(string(source[1000:2000]), string(source[2000:3000]))
	if distance < 750 {
		k = 128 - 95
		size = 7
	} else if distance < 900 {
		k = 256 - 95
		size = 8
	} else {
		k = 512 - 95
		size = 9
	}

	a := topKCommonPairs(source, k)
	for i := 94; i < 94 + k; i++ {
		dict[i] = a[i - 94]
	}
	generateMap()
	_, di_size := encode(source)
	fmt.Printf("Compression ratio: %v\n", ((float64(len(source) * 8)) / float64(di_size * size)))
	// fmt.Printf("Compression ratio: %v\n", ((float64(len(source) * 8)) / float64(len(bitstring))))
}
