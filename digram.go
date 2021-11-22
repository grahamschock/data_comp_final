package main
import (
	"fmt"
	"sort"
	"math/rand"
	"strconv"
	"os"
	"strings"
)

type Pair struct {
	Key string
	Value int
}



type PairList []Pair //to sort

func (p PairList) Len() int           { return len(p) }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
//Reads a two character input

//Assume the dictionary is 256
var dict [256]string
var dict_map map[string]int

//1st 94 is printable ASCII
//2nd 162 is the frequently most used pairs of characters
func genASCIIDict() {
	for i := 32; i < 127; i++ {
		dict[(i - 32)] = string(i)
	}
}

func top162CommonPairs(source []byte) (topK [162]string){
	m := make(map[string]int) //index = pair, value = num times occured
	for i := 0; i < len(source); i += 2 {
		m[string(source[i:i+2])]++
	}

	pairs := make(PairList, len(m))

	i := 0
	for k, v := range m {
		pairs[i] = Pair{k, v}
		i++
	}

	sort.Sort(sort.Reverse(pairs))
	for i := 0; i < len(pairs); i++ {
		topK[i] = pairs[i].Key
	}
	return topK
}

//3 steps for encoding (semi-static)
//read a 2 char input
//if in dictionary, encode
//if not in dictionary encode first char then add second char to digram. repeat
func encode(source []byte) (string) {
	var builder strings.Builder
	for i := 0; i < len(source); {
		if (digramInDict(source[i:i+2])) {
			fmt.Fprintf(&builder, "%s", encodeDigram(source[i:i+2]))
			i += 2
		} else { //the digram is not in dict
			fmt.Fprintf(&builder, "%s", encodeSingle(source[i]))
			i++
		}
	}
	return builder.String()
}

func digramInDict(digram []byte) (bool){
	_, present := dict_map[string(digram)]
	return present
}

func encodeDigram(digram []byte) (string) {
	val := int64(dict_map[string(digram)])
	return strconv.FormatInt(val, 2)
}

func encodeSingle(single byte) (string) {
	val := int64(dict_map[string(single)])
	return strconv.FormatInt(val, 2)
}

func randomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func generateMap() {
	dict_map = make(map[string]int)
	for i := 0; i < len(dict); i++ {
		dict_map[dict[i]] = i
	}
}

func main() {
	genASCIIDict()
	source, err := os.ReadFile("data/sources.100MB")
	if err != nil {
		panic(err)
	}
	a := top162CommonPairs(source)
	for i := 94; i < 256; i++ {
		dict[i] = a[i - 94]
	}
	generateMap()

	bitstring := encode(source)
	fmt.Printf("Compression ratio %v\n", ((float64(len(source) * 8)) / float64(len(bitstring))))
}
