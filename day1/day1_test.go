package main

import (
	"testing"
)

func Test_find_two_sum(t *testing.T) {
	input := []int{
		1721,
		979,
		366,
		299,
		675,
		1456}

	val1, val2, err := find_two_sum(input, 2020)
	if err != nil {
		t.Errorf("error")
	}
	if val1 != 1721 || val2 != 299 {
		t.Errorf("val1 and val 2 are wrong: %v, %v", val1, val2)
	}
}
func Test_find_three_sum(t *testing.T) {
	input := []int{
		1721,
		979,
		366,
		299,
		675,
		1456}

	val1, val2, val3, err := find_three_sum(input, 2020)
	if err != nil {
		t.Errorf("error")
	}
	if val1 != 979 || val2 != 366 || val3 != 675 {
		t.Errorf("val1 and val 3 are wrong: %v, %v %v", val1, val2, val3)
	}
}
