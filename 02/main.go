package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/02", "relative file path for input data file")

// PasswordEntry contains an unmarshaled line of input that look like this:
// 4-15 k: stjkjvvxrmwdpkwsjqvc
type PasswordEntry struct {
	MinConstrain int
	MaxConstrain int
	Character    string
	Password     string
}

// ValidPasswordsA stores the number of passwords that are valid according to task A password policy
var ValidPasswordsA int

// ValidPasswordsB stores the number of passwords that are valid according to task B password policy
var ValidPasswordsB int

func main() {
	flag.Parse()
	passwordsToValidate := ParseContents(*inputFile)
	for _, pe := range passwordsToValidate {
		validA := ValidatePasswordA(pe)
		if validA {
			ValidPasswordsA++
		}
		validB := ValidatePasswordB(pe)
		if validB {
			ValidPasswordsB++
		}
	}
	log.Printf("The number of valid passwords in task A is: %d", ValidPasswordsA)
	log.Printf("The number of valid passwords in task B is: %d", ValidPasswordsB)
}

// ValidatePasswordA checks if a password is valid according to part A of the task:
// does in have a specific character a specific number of times.
// If it does, true is returned
func ValidatePasswordA(pe PasswordEntry) bool {
	count := strings.Count(pe.Password, pe.Character)
	if count >= pe.MinConstrain && count <= pe.MaxConstrain {
		return true
	}
	return false
}

// ValidatePasswordB checks if a password is valid according to part B of the task:
// the Character has to be present on exactly one of the positions defined
// in pe.MinConstrain/pe.MaxConstrain.
// If it is, the function returns true
func ValidatePasswordB(pe PasswordEntry) bool {
	firstChar := string(string(pe.Password)[pe.MinConstrain-1])
	secondChar := string(string(pe.Password)[pe.MaxConstrain-1])
	if (firstChar == pe.Character && secondChar != pe.Character) || (firstChar != pe.Character && secondChar == pe.Character) {
		return true
	}
	return false
}

// ParseContents loads an input file and returns an array of integers.
func ParseContents(inputFile string) []PasswordEntry {
	file := openFile(inputFile)
	defer file.Close()

	var passwordEntries []PasswordEntry

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var passwordEntry PasswordEntry
		fmt.Sscanf(
			scanner.Text(),
			"%d-%d %1s: %s",
			&passwordEntry.MinConstrain,
			&passwordEntry.MaxConstrain,
			&passwordEntry.Character,
			&passwordEntry.Password,
		)
		passwordEntries = append(passwordEntries, passwordEntry)
	}
	return passwordEntries
}

func openFile(inputFile string) *os.File {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatalf("could not open file: %s", err)
		return nil
	}
	return file
}
