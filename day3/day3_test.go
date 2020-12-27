package main

import (
	"testing"
)

func TestValidPasswordCount_Question1(t *testing.T) {
	input := []string{
		"..##.......",
		"#...#...#..",
		".#....#..#.",
		"..#.#...#.#",
		".#...##..#.",
		"..#.##.....",
		".#.#.#....#",
		".#........#",
		"#.##...#...",
		"#...##....#",
		".#..#...#.#",
	}

	total := TreeCount(input, 3, 1)

	if total != 7 {
		t.Errorf("total trees encountered are: %v", total)
	}
}

func TestValidPasswordCount_Question2(t *testing.T) {
	input := []string{
		"..##.......",
		"#...#...#..",
		".#....#..#.",
		"..#.#...#.#",
		".#...##..#.",
		"..#.##.....",
		".#.#.#....#",
		".#........#",
		"#.##...#...",
		"#...##....#",
		".#..#...#.#",
	}

	total := TreeCount(input, 1, 1)

	if total != 2 {
		t.Errorf("total trees encountered are: %v", total)
	}
	total = TreeCount(input, 5, 1)

	if total != 3 {
		t.Errorf("total trees encountered are: %v", total)
	}
	total = TreeCount(input, 7, 1)

	if total != 4 {
		t.Errorf("total trees encountered are: %v", total)
	}
	total = TreeCount(input, 1, 2)

	if total != 2 {
		t.Errorf("total trees encountered are: %v", total)
	}
}
