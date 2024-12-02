package main

import (
	"fmt"
	"strconv"
	"strings"
)


func parseInput(inputData string, leftArr *[]int, rightMap *map[int]int) {
	lines := strings.Split(inputData, "\n")
	for i := 0; i < len(lines); i++ {
		temp := strings.Split(lines[i],"   ")

		left, _ := strconv.Atoi(temp[0])
		right, _ := strconv.Atoi(temp[1])

		*leftArr = append(*leftArr, left)
		(*rightMap)[right]++

	}

}

func main() {


	leftArr := []int{}
	rightMap := make(map[int]int)
	parseInput(inputData, &leftArr, &rightMap)

	sum := 0
	for i := 0; i < len(leftArr); i++ {
		if count, ok := rightMap[leftArr[i]]; ok {
			sum += leftArr[i] * count
		}else{
			sum+= 0
		}
	}

	fmt.Println(sum)
}