package main

import (
	"testing"
)

// Advent of Code contest
// https://adventofcode.com/2019/day/3
// go test -v

func TestDay03_1(t *testing.T) {
	input_x := "R8,U5,L5,D3"
	input_y := "U7,R6,D4,L4"

	d1x, d1y := day03_calculateData(input_x, input_y)
	res1, res2 := day03_calculate(day03_compare(day03_create(d1x), day03_create(d1y)), MAX_DISTANCE, MAX_STEPS)
	if res1 != 6{
		t.Errorf("test for Day03 Failed\nGot:\n%v,\nExpected:\n6", res1)
	}
	if res2 != 30{
		t.Errorf("test for Day03 Failed\nGot:\n%v,\nExpected:\n30", res2)
	}
}

func TestDay03_2(t *testing.T) {
	input_x := "R75,D30,R83,U83,L12,D49,R71,U7,L72"
	input_y := "U62,R66,U55,R34,D71,R55,D58,R83"

	d1x, d1y := day03_calculateData(input_x, input_y)
	res1, res2 := day03_calculate(day03_compare(day03_create(d1x), day03_create(d1y)), MAX_DISTANCE, MAX_STEPS)
	if res1 != 159{
		t.Errorf("test for Day03 Failed\nGot:\n%v,\nExpected:\n159", res1)
	}
	if res2 != 610{
		t.Errorf("test for Day03 Failed\nGot:\n%v,\nExpected:\n610", res2)
	}
}

func TestDay03_3(t *testing.T) {
	input_x := "R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51"
	input_y := "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7"

	d1x, d1y := day03_calculateData(input_x, input_y)
	res1, res2 := day03_calculate(day03_compare(day03_create(d1x), day03_create(d1y)), MAX_DISTANCE, MAX_STEPS)
	if res1 != 135{
		t.Errorf("test for Day03 Failed\nGot:\n%v,\nExpected:\n135", res1)
	}
	if res2 != 410{
		t.Errorf("test for Day03 Failed\nGot:\n%v,\nExpected:\n410", res2)
	}
}