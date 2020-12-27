// squares (.) and trees (#)
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input, err := ReadInput("day3/input.txt")
	if err != nil {
		fmt.Println(err)
	}

	Q1 := TreeCount(input, 3, 1)
	fmt.Println(Q1)

	tmp1 := TreeCount(input, 1, 1)
	tmp2 := TreeCount(input, 3, 1)
	tmp3 := TreeCount(input, 5, 1)
	tmp4 := TreeCount(input, 7, 1)
	tmp5 := TreeCount(input, 1, 2)
	Q2 := tmp1 * tmp2 * tmp3 * tmp4 * tmp5
	fmt.Println(Q2)
}

// counts the tree while traversing the slope
func TreeCount(input []string, right int, down int) int {
	treeCounter := 0
	i, j := 0, 0

	for i < len(input) {
		curPoint := input[i][j]
		if isTree(string(curPoint)) {
			treeCounter++
		}
		j += right
		j = j % len(input[0]) // the same pattern repeats to the right many times
		i += down
	}

	return treeCounter
}

func isTree(char string) bool {
	return char == "#"
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
