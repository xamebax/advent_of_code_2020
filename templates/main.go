package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"strconv"
)

var inputFile = flag.String("inputFile", "inputs/1", "relative file path for input data file")

func main() {
	flag.Parse()
	inputData := ParseContents(*inputFile)
	log.Printf("%v", inputData)
}

// ParseContents loads an input file and returns an array of integers.
func ParseContents(inputFile string) []int64 {
	file := openFile(inputFile)
	defer file.Close()

	var inputData []int64

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		entry, err := strconv.ParseInt(scanner.Text(), 0, 64)
		if err != nil {
			log.Fatalf("could not parse entry: %s", err)
		}
		inputData = append(inputData, entry)
	}

	return inputData
}

func openFile(inputFile string) *os.File {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatalf("could not open file: %s", err)
		return nil
	}
	return file
}
