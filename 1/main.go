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
	findTwoEntriesThatSumUpTo2020(*inputFile)
	findThreeEntriesThatSumUpTo2020(*inputFile)
}

func openFile(filePath string) *os.File {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("could not open file: %s", err)
		return nil
	}
	return file
}

func findThreeEntriesThatSumUpTo2020(inputFile string) {
	file := openFile(inputFile)
	defer file.Close()

	var inputNumbers []int64

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.ParseInt(scanner.Text(), 0, 64)
		if err != nil {
			log.Fatalf("could not parse number: %s", err)
		}
		inputNumbers = append(inputNumbers, number)
	}

MegaLoop:
	for _, firstNumber := range inputNumbers {
		for _, secondNumber := range inputNumbers {
			for _, thirdNumber := range inputNumbers {
				if firstNumber+secondNumber+thirdNumber == 2020 {
					log.Printf("The product of three entries that sum up to 2020 is: %d", firstNumber*secondNumber*thirdNumber)
					break MegaLoop
				}
			}
		}
	}
}

func findTwoEntriesThatSumUpTo2020(inputFile string) {
	file := openFile(inputFile)
	defer file.Close()

	numberPairs := make(map[int64]int64)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.ParseInt(scanner.Text(), 0, 64)
		if err != nil {
			log.Fatalf("could not parse number: %s", err)
		}
		numberPairs[number] = (2020 - number)
	}

	for _, number := range numberPairs {
		if numberPairs[number]+number == 2020 {
			log.Printf("The product of two entries that sum up to 2020 is: %d", numberPairs[number]*number)
			break
		}
	}
}
