package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	// I realized that it doesn't make a lot of sense
	// to use a bufio for my purposes.
	inputFile, err := ioutil.ReadFile("./input.txt")

	if err != nil {
		panic("File read error")
	}

	input := string(inputFile)
	groups := strings.Split(input, "\n\n")

	totalCount := 0

	for i := 0; i < len(groups); i++ {
		// Add 1 since you have to remember to count the last person,
		// whose \n got eaten by the string.Split() monster.
		numPeople := strings.Count(groups[i], "\n") + 1

		for j := 0; j < 26; j++ {
			// small brain time
			if strings.Count(groups[i], string(rune(j+97))) == numPeople {
				totalCount++
			}
		}
	}

	fmt.Println(totalCount)
}

/* For part 1 of the problem
func removeDuplicate(input string) string {
	output := strings.Builder{}

	for i := 0; i < len(input); i++ {
		if !strings.Contains(output.String(), string(input[i])) {
			output.WriteString(string(input[i]))
		}
	}

	return output.String()
}
*/
