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
		minMax := strings.Split(fields[0], "-")

		min, err := strconv.Atoi(minMax[0])

		max, err1 := strconv.Atoi(minMax[1])

		if err != nil && err1 != nil {
			fmt.Println("Error")
			break
		}

		// Character
		character := strings.Trim(fields[1], ":")

		// Actual Password Processing
		appearances := strings.Count(fields[2], character)

		// fmt.Println(min, max, character, appearances)

		if appearances >= min && appearances <= max {
			validPasswords++
		}
	}

	fmt.Println(validPasswords)
}
