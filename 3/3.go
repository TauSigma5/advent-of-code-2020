package main

import (
	"bufio"
	"fmt"
	"os"
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
			// Can't do it line by line here... or can I?
			lines = append(lines, scanner.Text())
		}
	}
	return lines
}

func treesHit(forestMap []string, slopeX int, slopeY int) int {
	treesHit := 0
	x := 0
	width := len(forestMap[0])

	for i := 0; i < len(forestMap); i += slopeY {
		if string(forestMap[i][x]) == "#" {
			treesHit++
		}

		x = (x + slopeX) % width
	}

	return treesHit
}

func main() {
	forestMap := readLines("./map.txt")

	fmt.Println(treesHit(forestMap, 1, 1) * treesHit(forestMap, 3, 1) * treesHit(forestMap, 5, 1) * treesHit(forestMap, 7, 1) * treesHit(forestMap, 1, 2))
}
