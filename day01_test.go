package main

// Advent of Code contest
// https://adventofcode.com/2019/day/1
// go test -v

import "testing"

func TestDay01a(t *testing.T) {
	input1 := 12
	input2 := 14
	input3 := 1969
	input4 := 100756

	result := day01a(input1)
	if result != 2 {
		t.Errorf("test for Day01a Failed\nGot:\n%v\nExpected:\n%v", result, input1)
	}
	result = day01a(input2)
	if result != 2 {
		t.Errorf("test for Day01a Failed\nGot:\n%v\nExpected:\n%v", result, input2)
	}
	result = day01a(input3)
	if result != 654 {
		t.Errorf("test for Day01a Failed\nGot:\n%v\nExpected:\n%v", result, input3)
	}
	result = day01a(input4)
	if result != 33583 {
		t.Errorf("test for Day01a Failed\nGot:\n%v\nExpected:\n%v", result, input4)
	}
}

func TestDay01b(t *testing.T) {
	input1 := 14
	input2 := 1969
	input3 := 100756
	result := day01b(input1)
	if result != 2 {
		t.Errorf("test for Day01a Failed\nGot:\n%v\nExpected:\n%v", result, input1)
	}
	result = day01b(input2)
	if result != 966 {
		t.Errorf("test for Day01a Failed\nGot:\n%v\nExpected:\n%v", result, input2)
	}
	result = day01b(input3)
	if result != 50346 {
		t.Errorf("test for Day01a Failed\nGot:\n%v\nExpected:\n%v", result, input3)
	}
}