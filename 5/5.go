package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")

	if err != nil {
		fmt.Println("Error")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if err == nil {
			row, col, id := findSeat(scanner.Text())
		}
	}

}

func findSeat(binaryInput string) (int, int, int) {
	instructions := strings.Split(binaryInput, "")
	// Define parameters of BSP
	topBottom := []int{0, 127}
	center := 63

	// I would do this recursively but brain go woosh (hehe get the pun?)
	for i := 0; i < len(instructions); i++ {
		currentInstruction := instructions[i]
		if currentInstruction == "F" || currentInstruction == "B" {
			if currentInstruction == "F" {
				topBottom[1] = center
			} else {
				topBottom[0] = center
			}
		} else {
			// Since we are guarenteed of conformity to the standard
			// I think I can get away with using just this
		}
	}
}
