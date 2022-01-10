package main

import (
	"log"
	"os"
	"sort"
	"bytes"
	"io/ioutil"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	if len(os.Args) != 3 {
		log.Fatalf("Usage: %v inputfile outputfile\n", os.Args[0])
	}

	log.Printf("Sorting %s to %s\n", os.Args[1], os.Args[2])
	inputFile := os.Args[1]
	outputFile := os.Args[2]

	Sort(inputFile, outputFile)
}

func Sort(input string, output string) {
	data, err := ioutil.ReadFile(input)
	check(err)

	records := make([][]byte, len(data) / 100)
	for i := range records {
		records[i] = make([]byte, 100)
	}

	for i := 0; i < len(data) / 100; i++ {
		records[i] = data[i * 100: (i + 1) * 100]
	}

	// sort according to key
	sort.Slice(records, func(i, j int) bool {
		return bytes.Compare(records[i][:10], records[j][:10]) < 0
	})

	// write to the output file
	ioutil.WriteFile(output, bytes.Join(records, []byte("")), 0666)
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}