package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatalf("could not open file: %s", err)
	}
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

	numberPairs := make(map[int64]int64)
	for _, number := range inputNumbers {
		numberPairs[number] = (2020 - number)
	}

	for _, number := range numberPairs {
		if numberPairs[number]+number == 2020 {
			log.Printf("The winning number is: %d", numberPairs[number]*number)
			break
		}
	}
}
