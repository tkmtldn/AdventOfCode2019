//package main
//
//import (
//	"bufio"
//	"fmt"
//	"log"
//	"math"
//	"os"
//	"strconv"
//	"strings"
//)
//
//const FILENAME = "./_advent/day03/day03input.txt"
//const MAX_DISTANCE  = 99999
//const MAX_STEPS  = 99999
//
//func day03_create(data []string) map[string]int{
//	// Creating and returning a map of visited 2D points
//	x_start := 0
//	x_finish := 0
//	y_start := 0
//	y_finish := 0
//	step := 0
//
//	visited := map[string]int{}
//
//	for _, v := range data {
//		pointer := v[0]
//		move, _ := strconv.Atoi(v[1:])
//		switch pointer {
//		case 'L':
//			x_finish -= move
//			for i:= x_start; i>x_finish; i-- {
//				add := day03_convert(i, y_start)
//				visited[add] = step
//				step += 1
//				x_start = x_finish
//			}
//		case 'R':
//			x_finish += move
//			for i:= x_start; i<x_finish; i++{
//				add := day03_convert(i, y_start)
//				visited[add] = step
//				step += 1
//				x_start = x_finish
//			}
//		case 'U':
//			y_finish += move
//			for i:= y_start; i<y_finish; i++{
//				add := day03_convert(x_start, i)
//				visited[add] = step
//				step += 1
//				y_start = y_finish
//			}
//		case 'D':
//			y_finish -= move
//			for i:= y_start; i>y_finish; i--{
//				add := day03_convert(x_start, i)
//				visited[add] = step
//				step += 1
//				y_start = y_finish
//			}
//		}
//	}
//	return visited
//}
//
//func day03_convert(i1, i2 int) string {
//	// creating a key for visited map
//	return strconv.Itoa(i1) + "_" + strconv.Itoa(i2)
//}
//
//func day03_compare(v1, v2 map[string]int) map[string]int{
//	// returning a map where intersections are keys and number of steps are values
//	ans := map[string]int{}
//
//	for k1, v1 := range v1 {
//		for k2, v2 := range v2 {
//			if k1 == k2 {
//				ans[k1] = v1+v2
//			}
//		}
//	}
//	return ans
//}
//
//func day03_calculate(a map[string]int, min_manhattan_distance int, min_steps int) (int, int){
//
//	for k, v := range a{
//		coord := strings.Split(k, "_")
//		x, _ := strconv.Atoi(coord[0])
//		y, _ := strconv.Atoi(coord[1])
//		distance := math.Abs(float64(x)) + math.Abs(float64(y))
//		if distance != 0 && int(distance) < min_manhattan_distance {
//			min_manhattan_distance = int(distance)
//		}
//		if v !=0 && v < min_steps {
//			min_steps = v
//		}
//	}
//	return min_manhattan_distance, min_steps
//}
//
//func day03_calculateData(data1, data2 string) ([]string, []string) {
//	// Split data into slices
//	d1 := strings.Split(data1, ",")
//	d2 := strings.Split(data2, ",")
//	return d1, d2
//}
//
//func main() {
//
//	f, err := os.Open(FILENAME)
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer f.Close()
//
//	var lines []string
//	scanner := bufio.NewScanner(f)
//	for scanner.Scan() {
//		lines = append(lines, scanner.Text())
//	}
//
//	data1 := lines[0]
//	data2 := lines[1]
//
//	d1, d2 := day03_calculateData(data1, data2)
//
//	visited_1 := day03_create(d1)
//	visited_2 := day03_create(d2)
//
//	count1, count2 := day03_calculate(day03_compare(visited_1, visited_2), MAX_DISTANCE, MAX_STEPS)
//
//	fmt.Println("Answer 1:", count1)
//	fmt.Println("Answer 2:", count2)
//	// 399, 15678
//}