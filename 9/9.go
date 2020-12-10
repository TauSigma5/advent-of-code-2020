package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Global variable since this should be immutable
var preambleLength int = 25

func main() {
	file, err := os.Open("./input.txt")

	if err != nil {
		panic("File Read Error")
	}

	var numbers []int
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if err == nil {
			num, err1 := strconv.Atoi(scanner.Text())

			if err1 != nil {
				panic("rip")
			}
			numbers = append(numbers, num)
		}
	}

	// Starts with the first n values
	segment := preambleLength

	for i := preambleLength; i < len(numbers); i++ {
		if !inPreamble(numbers, segment, numbers[segment]) {
			break
		}
		segment++
	}

	fmt.Println(contigNum(numbers, numbers[segment]))
}

func inPreamble(numbers []int, segment int, number int) bool {

	for i := segment - preambleLength; i < segment; i++ {
		for j := segment - preambleLength; j < segment; j++ {
			// fmt.Println(numbers[i] + numbers[j])
			if numbers[i]+numbers[j] == number {
				return true
			}
		}
	}

	return false
}

func contigNum(numbers []int, target int) int {
	// Minimum of two continuous numbers
	numNums := 2
	val := 0
	endVal := 0
	var minMax []int

	for true {
		// Outter loop, checks 2, then 3 then four etc and stops when a number is found
		for i := 0; i < len(numbers)-(numNums-1); i++ {
			// middle loop, goes through the array
			val = 0
			for j := 0; j < numNums; j++ {
				// Goes through the next nums
				val += numbers[i+j]
				endVal = j
			}

			if val == target {
				for k := i; k < i+endVal; k++ {
					minMax = append(minMax, numbers[k])
				}
				min, max := getMinMax(minMax)

				return min + max
			}
		}
		numNums++
	}

	return 0
}

func getMinMax(input []int) (int, int) {
	min := input[0]
	max := 0

	for _, value := range input {
		if value < min {
			min = value
		} else if max < value {
			max = value
		}
	}

	return min, max
}
