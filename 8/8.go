package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	inputFile, err := ioutil.ReadFile("./input.txt")

	if err != nil {
		panic("File read error")
	}

	instructions := strings.Split(string(inputFile), "\n")

	// Get a list of values to test.
	for i, value := range instructions {
		if strings.Contains(value, "jmp") || strings.Contains(value, "nop") {
			var changedInstruction []string

			for i := range instructions {
				changedInstruction = append(changedInstruction, instructions[i])
			}
			// fmt.Println(changedInstruction)

			val, success := compute(changeValue(changedInstruction, i))
			if success {
				fmt.Println(val)
			}

		}
	}
}

func changeValue(input []string, valChange int) []string {
	value := input[valChange]

	if strings.Contains(value, "jmp") {
		input[valChange] = strings.ReplaceAll(value, "jmp", "nop")
	} else {
		input[valChange] = strings.ReplaceAll(value, "nop", "jmp")
	}

	// fmt.Println(input)

	return input
}

func contains(input []int, value int) bool {
	for _, number := range input {
		if number == value {
			return true
		}
	}

	return false
}

func compute(instructions []string) (int, bool) {
	// For the accumulator
	globalValue := 0
	currentInstruction := 0
	var visitedInstructions []int

	for !contains(visitedInstructions, currentInstruction) {
		instruction := strings.Split(instructions[currentInstruction], " ")
		operator := instruction[1]

		visitedInstructions = append(visitedInstructions, currentInstruction)

		if instruction[0] == "acc" {
			if strings.Contains(operator, "+") {
				number, err := strconv.Atoi(strings.ReplaceAll(operator, "+", ""))
				if err != nil {
					panic("Something went wrong")
				}
				globalValue += number
			} else {
				number, err := strconv.Atoi(strings.ReplaceAll(operator, "-", ""))
				if err != nil {
					panic("Something went wrong")
				}
				globalValue -= number
			}
		} else if instruction[0] == "jmp" {
			if strings.Contains(operator, "+") {
				number, err := strconv.Atoi(strings.ReplaceAll(operator, "+", ""))
				if err != nil {
					panic("Something went wrong")
				}
				currentInstruction += (number - 1)
			} else {
				number, err := strconv.Atoi(strings.ReplaceAll(operator, "-", ""))
				if err != nil {
					panic("Something went wrong")
				}
				currentInstruction -= (number + 1)
			}
		}

		currentInstruction++

		if currentInstruction == len(instructions) {
			return globalValue, true
		}
	}

	return globalValue, false
}
