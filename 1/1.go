package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readLines(path string) []int {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Read Error")
		return nil
	}
	defer file.Close()

	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text, err := strconv.Atoi(scanner.Text())
		if err == nil {
			lines = append(lines, text)
		}
	}
	return lines
}

func main() {
	input := readLines("./1a.txt")
	exit := false
	number1 := 0
	number2 := 0
	number3 := 0

	for i := 0; i < len(input); i++ {
		number1 = input[i]

		for j := 0; j < len(input); j++ {
			number2 = input[j]

			for k := 0; k < len(input); k++ {
				number3 = input[k]

				if number1+number2+number3 == 2020 {
					exit = true
					break
				}
			}

			if exit == true {
				break
			}
		}

		if exit == true {
			break
		}
	}

	fmt.Println(number1 * number2 * number3)
}
