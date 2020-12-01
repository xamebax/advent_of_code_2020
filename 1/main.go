package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	findTwoEntriesThatSumUpTo2020()
	findThreeEntriesThatSumUpTo2020()
}

func openFile(filePath string) *os.File {
	file, err := os.Open("input")
	if err != nil {
		log.Fatalf("could not open file: %s", err)
		return nil
	}
	return file
}

func findThreeEntriesThatSumUpTo2020() {
	file := openFile("input")
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

func findTwoEntriesThatSumUpTo2020() {
	file := openFile("input")
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
