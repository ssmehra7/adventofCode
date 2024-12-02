package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func parseInput(s string, leftArr *[]int, rightArr *[]int){
	lines := strings.Split(s, "\n")
	// fmt.Println((lines))

	for i:=0; i<len(lines); i++{
		temp := strings.Split(lines[i], "   ")
		left,_ := strconv.Atoi(temp[0])
		right,_ := strconv.Atoi(temp[1])
		if len(temp) == 2{
			*leftArr = append(*leftArr, left)
			*rightArr = append(*rightArr, right)
		}

	}

}

func main(){
	leftArr := []int{}
	rightArr := []int{}
	parseInput(inputData, &leftArr, &rightArr)

	sort.Ints(leftArr)
	sort.Ints(rightArr)

	// fmt.Println(leftArr, rightArr)
	sum := 0
	for i:=0 ; i<len(leftArr); i++{
		l := leftArr[i]
		r := rightArr[i]

		if (l>r){
			sum+= l-r
		}else{
			sum+= r-l
		}

	}

	fmt.Println(sum)

}

