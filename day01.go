package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func day01_2(n int) (ans int) {
	for {
		res := (n/3)-2
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
		count1 += ((num/3)-2)

		// Part2 - do the same until its negative
		count2 += day01_2(num)
	}

	fmt.Println("Answer 1: ", count1)
	fmt.Println("Answer 2: ", count2)
}