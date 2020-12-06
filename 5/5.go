package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")

	if err != nil {
		panic("File Read Error")
	}

	var ids []int

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if err == nil {
			id := findSeat(scanner.Text())
			ids = append(ids, id)
		}
	}

	sort.Ints(ids)

	// fmt.Println(ids)

	// This is probably stupid but the easiest way I can think of
	for i := 1; i < len(ids); i++ {
		if ids[i]-1 != ids[i-1] {
			fmt.Println(ids[i])
			break
		}
	}
}

func findSeat(binaryInput string) int {
	instructions := strings.Split(binaryInput, "")
	// Define parameters of BSP
	frontBack := []int{0, 127}
	leftRight := []int{0, 7}

	// fmt.Println(frontBack)

	// I would do this recursively but brain go woosh (hehe get the pun?)
	for i := 0; i < len(instructions); i++ {
		currentInstruction := instructions[i]
		if currentInstruction == "F" || currentInstruction == "B" {
			if currentInstruction == "F" {
				frontBack[1] = frontBack[1] - int(math.Round(128/math.Pow(2, float64(i+1))))
				// fmt.Println(frontBack[1])
			} else {
				frontBack[0] = frontBack[0] + int(math.Round(128/math.Pow(2, float64(i+1))))
				// fmt.Println(frontBack[0])
			}
			// fmt.Println(frontBack)
		} else {
			// Since we are guarenteed of conformity to the standard
			// I think I can get away with using just this
			if currentInstruction == "L" {
				leftRight[1] = leftRight[1] - int(math.Round(8/math.Pow(2, float64(i-6))))
			} else {
				leftRight[0] = leftRight[0] + int(math.Round(8/math.Pow(2, float64(i-6))))
			}
		}
	}

	// Make sure that the values are correct or else throw panic
	if frontBack[0] != frontBack[1] {
		panic("Front back error" + " " + fmt.Sprint(frontBack[0]) + " " + fmt.Sprint(frontBack[1]))
	} else if leftRight[0] != leftRight[1] {
		panic("Left right error" + " " + fmt.Sprint(leftRight[0]) + " " + fmt.Sprint(leftRight[1]))
	}

	row := frontBack[0]
	col := leftRight[0]

	return row*8 + col
}
