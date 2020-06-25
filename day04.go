package main

// Advent of Code contest
// https://adventofcode.com/2019/day/4

import (
	"fmt"
	"strconv"
	"strings"
)

func countLargerGroup(in string) int {
	if strings.Count(in, "0") == 2 || strings.Count(in, "1") == 2 ||
		strings.Count(in, "2") == 2 || strings.Count(in, "3") == 2 ||
		strings.Count(in, "4") == 2 || strings.Count(in, "5") == 2 ||
		strings.Count(in, "6") == 2 || strings.Count(in, "7") == 2 ||
		strings.Count(in, "8") == 2 || strings.Count(in, "9") == 2 {
		return 1
	}
	return 0
}

func main()  {
	input_start := 278384
	input_finish := 824795
	result1 := 0
	result2 := 0

	for i:=input_start; i<=input_finish; i++ {
		str := strconv.Itoa(i)
		if (str[0] == str[1] || str[1] == str[2] || str[2] == str[3] || str[3] == str[4] || str[4] == str[5]) &&
			(str[0] <= str[1] && str[1] <= str[2] &&  str[2] <= str[3] &&  str[3] <= str[4] &&  str[4] <= str[5]) {
			result1 += 1
			result2 += countLargerGroup(str)
		}
	}

	fmt.Println("Answer 1:", result1)
	fmt.Println("Answer 2:", result2)

}