// https://adventofcode.com/2020/day/4
package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var expectedFields = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"} // "cid" is optional

func main() {
	input, err := ReadInput("day4/input.txt")
	if err != nil {
		fmt.Println(err)
	}

	Q1 := ValidPassportCount(input, PassportValid_Question1)
	fmt.Println(Q1)

	Q2 := ValidPassportCount(input, PassportValid_Question2)
	fmt.Println(Q2)
}

type passportPolicy func(string) bool

func ValidPassportCount(input []string, fn passportPolicy) int {
	validPassportCounter := 0
	for _, document := range input {
		if fn(document) {
			validPassportCounter++
		}
	}
	return validPassportCounter
}

func PassportValid_Question1(document string) bool {
	for _, field := range expectedFields {
		if !strings.Contains(document, field) {
			return false
		}
	}
	return true
}

func PassportValid_Question2(document string) bool {
	for _, field := range expectedFields {
		if !strings.Contains(document, field) {
			return false
		}
	}

	var docFields = strings.Fields(document)
	for _, f := range docFields {
		pair := strings.Split(f, ":")
		prop := pair[0]
		val := pair[1]

		switch prop {
		case "byr":
			// byr (Birth Year) - four digits; at least 1920 and at most 2002.
			val_num, _ := strconv.Atoi(val)
			if val_num < 1920 || val_num > 2002 {
				return false
			}
		case "iyr":
			//	iyr (Issue Year) - four digits; at least 2010 and at most 2020.
			val_num, _ := strconv.Atoi(val)
			if val_num < 2010 || val_num > 2020 {
				return false
			}
		case "eyr":
			//	eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
			val_num, _ := strconv.Atoi(val)
			if val_num < 2020 || val_num > 2030 {
				return false
			}
		case "hgt":
			//	hgt (Height) - a number followed by either cm or in:
			//  - If cm, the number must be at least 150 and at most 193.
			//  - If in, the number must be at least 59 and at most 76.
			if strings.HasSuffix(val, "cm") {
				tmp := strings.Split(val, "cm")
				hgt, err := strconv.Atoi(tmp[0])
				if err != nil || hgt < 150 || hgt > 193 {
					return false
				}
			} else if strings.HasSuffix(val, "in") {
				tmp := strings.Split(val, "in")
				hgt, err := strconv.Atoi(tmp[0])
				if err != nil || hgt < 59 || hgt > 76 {
					return false
				}
			} else {
				return false
			}

		case "hcl":
			//hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
			if strings.HasPrefix(val, "#") {
				hcl := strings.Split(val, "#")[1]
				if len(hcl) != 6 {
					return false
				} else {
					for i := 0; i < len(hcl); i++ {
						char := hcl[i]
						if !(('a' <= char && char <= 'f') || ('0' <= char && char <= '9')) {
							return false
						}
					}
				}
			} else {
				return false
			}

		case "ecl":
			// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
			valid_colors := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
			if !contains(valid_colors, val) {
				return false
			}

		case "pid":
			// pid (Passport ID) - a nine-digit number, including leading zeroes.
			if _, err := strconv.Atoi(val); err != nil || len(val) != 9 {
				return false
			}

		}
	}
	return true
}

func contains(slice []string, target string) bool {
	for _, item := range slice {
		if item == target {
			return true
		}
	}
	return false
}

func ReadInput(path string) ([]string, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	fileContent := string(file)
	slicedContent := strings.Split(fileContent, "\n\n") // line break

	return slicedContent, err
}
