package main

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Advent of Code contest
// https://adventofcode.com/2019/day/3


func TestDayo3(t *testing.T) {

	tests := []struct {
		input_x string
		input_y string
		wantA   int
		wantB   int
	}{
		{
			input_x: "R8,U5,L5,D3",
			input_y: "U7,R6,D4,L4",
			wantA:   6,
			wantB:   30,
		},
		{
			input_x: "R75,D30,R83,U83,L12,D49,R71,U7,L72",
			input_y: "U62,R66,U55,R34,D71,R55,D58,R83",
			wantA:   159,
			wantB:   610,
		},
		{
			input_x: "R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51",
			input_y: "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7",
			wantA:   135,
			wantB:   410,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprint("test_", i), func(t *testing.T) {

			d1x, d1y := threePrepareData(test.input_x, test.input_y)
			res1, res2 := threeCalculateResults(threeCompare(threeVisited2D(d1x), threeVisited2D(d1y)), MAX_DISTANCE, MAX_STEPS)

			assert.Equal(t, test.wantA, res1)
			assert.Equal(t, test.wantB, res2)
		})
	}
}