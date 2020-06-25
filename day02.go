package main

import (
	"fmt"
	"log"
)

// Advent of Code contest
// https://adventofcode.com/2019/day/2

const SECOND_GOAL = 19690720

func twoIntcode(data[]int) {

    for idx:=0; idx<len(data); idx+=4{
		switch data[idx] {
		case 99:
			return
		case 1:
			arg1 := data[idx+1]
			arg2 := data[idx+2]
			resultIdx := data[idx+3]
			data[resultIdx] = data[arg1] + data[arg2]
		case 2:
			arg1 := data[idx+1]
			arg2 := data[idx+2]
			resultIdx := data[idx+3]
			data[resultIdx] = data[arg1] * data[arg2]
		default:
			log.Fatalf("Position error %d: %d", idx, data[idx])
		}
    }
}

func twoCalculate(data []int, noun, verb int) int {
	newList := make([]int, len(data))
	copy(newList, data)
	newList[1], newList[2] = noun, verb
	twoIntcode(newList)
	return newList[0]
}

//func readIntoIntSlice(s string) []int {
//	var n []int
//	for _, f := range strings.Fields(s) {
//		i, err := strconv.Atoi(f)
//		if err == nil {
//			n = append(n, i)
//		}
//	}
//	return n
//}


func main(){

	//f, err := os.Open("./_advent/day02input.txt")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer f.Close()

	dataset := []int{1,0,0,3,1,1,2,3,1,3,4,3,1,5,0,3,2,10,1,19,2,19,6,23,2,13,23,27,1,9,27,31,2,31,9,35,1,6,35,39,2,10,39,43,1,5,43,47,1,5,47,51,2,51,6,55,2,10,55,59,1,59,9,63,2,13,63,67,1,10,67,71,1,71,5,75,1,75,6,79,1,10,79,83,1,5,83,87,1,5,87,91,2,91,6,95,2,6,95,99,2,10,99,103,1,103,5,107,1,2,107,111,1,6,111,0,99,2,14,0,0}

	fmt.Println("Answer 1: ", twoCalculate(dataset, 12, 2))

	for noun := 0 ; noun < 100; noun ++ {
		for verb := 0; verb < 100; verb++ {
			result := twoCalculate(dataset, noun, verb)
			if result == SECOND_GOAL {
				fmt.Println("Answer 2: ", noun*100 + verb)
				return
			}
		}
	}
}