package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := ReadInput("day2/input.txt")
	if err != nil {
		fmt.Println(err)
	}

	total1 := ValidPasswordCount(input, CheckPolicy_Question1)
	fmt.Println(total1)

	total2 := ValidPasswordCount(input, CheckPolicy_Question2)
	fmt.Println(total2)
}

type checkPolicy func(int, int, string, string) bool

func ValidPasswordCount(input []string, fn checkPolicy) int {
	total := 0
	for _, puzzle := range input {
		s := strings.Split(puzzle, ":")
		password := strings.TrimSpace(s[1])

		p := strings.Split(s[0], " ")
		limit := strings.Split(p[0], "-")
		n1, _ := strconv.Atoi(limit[0])
		n2, _ := strconv.Atoi(limit[1])
		char := p[1]

		if fn(n1, n2, char, password) {
			total++
		}
	}
	return total

}

func CheckPolicy_Question1(min int, max int, char string, password string) bool {
	char_count := 0
	for _, c := range password {
		// c is a rune
		if string(c) == char {
			char_count++
			if char_count > max {
				return false
			}
		}
	}
	return char_count >= min
}

func CheckPolicy_Question2(n1 int, n2 int, char string, password string) bool {
	char1 := string(password[n1-1])
	char2 := string(password[n2-1])
	x := char1 == char
	y := char2 == char

	// xor if only one of them is true
	return x != y
}

func ReadInput(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
