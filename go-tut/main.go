package main

import "fmt"

func main() {
	var x = []int{}
	var y = make([]int, len(x))

	for i := 1; i <= 10; i++ {
		x = append(x, i)
		i++
	}

	y = make ([]int, len(x))
	copy(y, x)

	fmt.Println("y slice is : ", y)
	fmt.Println("x slice is : ", x)


}

// package main

// import "fmt"

// func main (){
// 	//for -> only construct in go for looping

// 	var x int

// 	fmt.Println("Type the No. : ")
// 	fmt.Scan(&x)

// 	y :=(2*x)-1

// 	for i:= 1;i<=x; i++{
// 		for j:=1 ; j<=y; j++{
// 			if (i%2 !=0 && j%2 == 0 && j>=x-i+1 && j<=x+i-1) {
// 				fmt.Print("*")
// 			}else if(i%2 ==0 && j%2 != 0 && j>=x-i+1 && j<=x+i-1){
// 				fmt.Print("*")
// 			}else{
// 				fmt.Print(" ")
// 			}
// 		}

// 		fmt.Println()

// 	}
// }

// // package main

// // import "fmt"

// // func main(){

// // 	//if not infering then you have to assing type
// // 	// var name string = "go land"

// // 	//infer if you directly do it
// // 	var name = "goland"
// // 	const age = 30

// // 	//shorthand syntax
// // 	// x:=2
// // 	fmt.Println(age)

// // 	fmt.Println(name)
// // }