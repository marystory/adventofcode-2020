package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input, err := readLines("day1/input.txt")
	if err != nil {
		fmt.Println(err)
	}

	val1, val2, err := find_two_sum(input, 2020)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("The two value are: %v, %v, and their multiplication is %v\n", val1, val2, val1*val2)
	}

	val1, val2, val3, err := find_three_sum(input, 2020)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("The three values are: %v, %v, %v and their multiplication is %v\n", val1, val2, val3, val1*val2*val3)
	}
}

func find_two_sum(input []int, sum int) (int, int, error) {
	m := map[int]bool{}
	for _, num := range input {
		diff := sum - num
		if _, ok := m[diff]; ok {
			return diff, num, nil
		} else {
			m[num] = true
		}
	}
	return 0, 0, fmt.Errorf("did not find match")
}

func find_three_sum(input []int, sum int) (int, int, int, error) {
	for i, num := range input {
		diff := sum - num
		val1, val2, _ := find_two_sum(input[i+1:], diff)
		if val1 != 0 || val2 != 0 {
			return num, val1, val2, nil
		}
	}
	return 0, 0, 0, fmt.Errorf("did not find match")
}


func readLines(path string) ([]int, error) {
	fmt.Println(path)
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		lines = append(lines, num)
	}
	return lines, scanner.Err()
}
