package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readLines(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Read Error")
		return nil
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if err == nil {
			// I technically could have just put everything in main here
			// and processed it line by line lol...
			lines = append(lines, scanner.Text())
		}
	}
	return lines
}

func main() {
	passwords := readLines("./passwords.txt")
	validPasswords := 0

	for i := 0; i < len(passwords); i++ {
		fields := strings.Fields(passwords[i])

		// 0th index is min, 1st index
		pos := strings.Split(fields[0], "-")

		pos1, err := strconv.Atoi(pos[0])
		pos2, err1 := strconv.Atoi(pos[1])
		// Position indexes from corporate don't from one
		pos1--
		pos2--

		// Shitty code that should never see the light of day
		if err != nil && err1 != nil {
			fmt.Println("Error")
			break
		}

		// Character
		character := strings.Trim(fields[1], ":")

		// Actual Password Processing
		password := string(fields[2])

		// There's no XOR function :/
		if (string(password[pos1]) == character || string(password[pos2]) == character) && !(string(password[pos1]) == character && string(password[pos2]) == character) {
			validPasswords++
		}
	}

	fmt.Println(validPasswords)
}
