package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func readLines() string {
	data, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		fmt.Println("Error")
	}
	return string(data)
}

func main() {
	inputFile := readLines()
	validPassports := 0

	passports := strings.Split(inputFile, "\n\n")
	// var passports []passport

	for i := 0; i < len(passports); i++ {
		if checkValidity(string(passports[i])) {
			validPassports++
		}
	}

	fmt.Println(validPassports)
}

func checkValidity(passport string) bool {
	fields := strings.Fields(passport)
	validFields := 0

	for i := 0; i < len(fields); i++ {
		if strings.Contains(string(fields[i]), "byr") {
			if checkByr(fields[i]) {
				validFields++
			}
		} else if strings.Contains(string(fields[i]), "iyr") {
			if checkIyr(fields[i]) {
				validFields++
			}
		} else if strings.Contains(string(fields[i]), "eyr") {
			if checkEyr(fields[i]) {
				validFields++
			}
		} else if strings.Contains(string(fields[i]), "hgt") {
			if checkHgt(fields[i]) {
				validFields++
			}
		} else if strings.Contains(string(fields[i]), "hcl") {
			if checkHcl(fields[i]) {
				validFields++
			}
		} else if strings.Contains(string(fields[i]), "ecl") {
			if checkEcl(fields[i]) {
				validFields++
			}
		} else if strings.Contains(string(fields[i]), "pid") {
			if checkPid(fields[i]) {
				validFields++
			}
		}
	}

	fmt.Println(validFields)

	if validFields == 7 {
		return true
	}

	return false
}

func kvConv(input string) []string {
	return strings.Split(input, ":")
}

// good 1
func checkByr(input string) bool {
	kv := kvConv(input)
	value, err := strconv.Atoi(kv[1])

	if err != nil {
		return false
	}

	if value <= 2002 && 1920 <= value {
		return true
	}

	return false
}

// good 1
func checkIyr(input string) bool {
	kv := kvConv(input)
	value, err := strconv.Atoi(kv[1])

	if err != nil {
		return false
	}

	if value <= 2020 && 2010 <= value {
		return true
	}

	return false
}

func checkEyr(input string) bool {
	kv := kvConv(input)
	value, err := strconv.Atoi(kv[1])

	if err != nil {
		return false
	}

	if value >= 2020 && 2030 >= value {
		return true
	}
	return false
}

// good
func checkHgt(input string) bool {
	kv := kvConv(input)
	heightUnits := string(kv[1])

	if strings.Contains(heightUnits, "in") {
		height, err := strconv.Atoi(strings.Trim(heightUnits, "in"))

		if err != nil {
			return false
		}

		if height <= 76 && height >= 59 {
			return true
		}
	} else {
		height, err := strconv.Atoi(strings.Trim(heightUnits, "cm"))

		if err != nil {
			return false
		}

		if height <= 193 && height >= 150 {
			return true
		}
	}

	return false
}

//good
func checkHcl(input string) bool {
	kv := kvConv(input)
	value := string(kv[1])

	if string(value[0]) == "#" && len(value) == 7 && !(strings.ContainsAny(value, "ghijklmnopqrstuvwxyz")) {
		return true
	}

	return false
}

// good 1
func checkEcl(input string) bool {
	kv := kvConv(input)
	value := string(kv[1])

	var re = regexp.MustCompile(`amb|blu|brn|gry|grn|hzl|oth`)

	if re.MatchString(value) {
		return true
	}

	return false
}

// good
func checkPid(input string) bool {
	kv := kvConv(input)
	value := string(kv[1])

	if len(value) == 9 && !(strings.Contains(value, "abcdefghijklmnopqrstuvwxyz")) {
		return true
	}

	return false
}

/*
type passport struct {
	byr int
	iyr int
	eyr int
	hgt string
	hcl string
	ecl string
	pid string
}
*/

/*
func split(r rune) bool {
	return r == '\n' || r == ' '
}


Mistakes have been made, this does not work and requires
wayyyy too much effort to get working
*/
