package main

import (
	"fmt"
	"math/rand"
)

type Kmer struct {
	// Do I use list, tuple or map?
}

func generate_kmer_seq() string {
	var base [4]string
	base[0] = "A"
	base[1] = "C"
	base[2] = "G"
	base[3] = "T"

	rand.New(rand.NewSource(42))

	var res string = ""

	for i := 0; i < 20; i++ {
		res += base[rand.Int()%4]
	}
	return res
}

func One_mer()

func main() {
	fmt.Println("Kmer Sequence Generator")
	fmt.Println(generate_kmer_seq())
}
