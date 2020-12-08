package main

import (
	"testing"
)

func Test_ValidPasswordCountQuestion1(t *testing.T) {
	input := []string{
		"1-3 a: abcde",
		"1-3 b: cdefg",
		"2-9 c: ccccccccc"}

	total := ValidPasswordCount(input, CheckPolicy_Question1)

	if total != 2 {
		t.Errorf("total counted valid passowrds is: %v", total)
	}
}

func Test_ValidPasswordCountQuestion2(t *testing.T) {
	input := []string{
		"1-3 a: abcde",
		"1-3 b: cdefg",
		"2-9 c: ccccccccc"}

	total := ValidPasswordCount(input, CheckPolicy_Question2)

	if total != 1 {
		t.Errorf("total counted valid passowrds is: %v", total)
	}
}
