package main

// Advent of Code contest
// https://adventofcode.com/2019/day/1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func day01a(n int) int {
	return (n/3)-2
}

func day01b(n int) (ans int) {
	for {
		res := day01a(n)
		if res > 0{
			ans += res
		} else {
			break
		}
		n = res
	}
	return ans
}

func main() {
	const FILENAME = "./_advent/day01input.txt"

	f, err := os.Open(FILENAME)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	count1 := 0
	count2 := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		num, _ := strconv.Atoi(line)

		// Part1 - divide by three, round down, and subtract 2.
		count1 += day01a(num)

		// Part2 - do the same until its negative
		count2 += day01b(num)
	}

	fmt.Println("Answer 1: ", count1)
	fmt.Println("Answer 2: ", count2)
}