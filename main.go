package main

import (
    "fmt"
    "math/rand"
)

type KmerSeq struct{ 
    Kmer  string
    Count int
}

func makeSequence(size int) string {
    seq := ""
    chars := [4]string{"A", "G", "T", "C"}

    for i:= 0; i < size; i++ {
        seq += chars[rand.Intn(3)]
    }
    return seq
}

func findKmers(dna string, k int) []KmerSeq {
    kmerMap := make(map[string]int)

    for i:=0; i<len(dna)-k; i++ {
        kmer := dna[i : i+k]
        kmerMap[kmer]++
    }

    var kmerSeq []KmerSeq

    for kmer, count := range kmerMap {
        kmerSeq = append(kmerSeq, KmerSeq{Kmer: kmer, Count: count})
    }
    return kmerSeq
}

func main() {
    var dna_sequence string = makeSequence(20)

    fmt.Println(dna_sequence)

    fmt.Println()

    IVmers := findKmers(dna_sequence, 4)

    fmt.Printf("Kmers of length %d:\n", 4)

    for _, kmer := range IVmers {
        fmt.Printf("%s: %d\n", kmer.Kmer, kmer.Count)
    }
}
