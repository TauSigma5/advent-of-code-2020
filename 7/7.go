package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var bagTrees []bag

func cleanString(input string) string {
	// Removes words like bag, bags and all punctuation
	// as well as leading and trailing spaces

	input = strings.ReplaceAll(input, ",", "")
	input = strings.ReplaceAll(input, ".", "")
	input = strings.ReplaceAll(input, "bags", "")
	input = strings.ReplaceAll(input, "bag", "")
	input = strings.Trim(input, " ")

	// fmt.Println(cleanString)

	return input
}

func generateRules() [][]string {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Read Error")
		return nil
	}
	defer file.Close()

	var rules [][]string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if err == nil {
			line := scanner.Text()
			// Split the line first into the containing bag and the
			// contained bags within the containing bag
			var cleanedRule []string

			rule := strings.Split(line, " contain ")
			cleanedRule = append(cleanedRule, cleanString(string(rule[0])))

			// Then split the contained bags into individual elements and
			// have them as colors + numbers
			for _, containedBag := range strings.Split(string(rule[1]), ",") {
				cleanedRule = append(cleanedRule, cleanString(containedBag))
			}

			rules = append(rules, cleanedRule)
		}
	}
	return rules
}

func main() {
	rules := generateRules()
	

}


func newTree(color string, parent string, children []string) bag {

	var children []bag

	for i := 0; i < len(children); i++ {
		childColors := strings.Split(children[i], " ")

	}

	newBag := bag {
		color: traits[0]
		children
	}
}

type bag struct {
	// Bag struct
	color    string
	children *[]bag
	parent   *bag
}
