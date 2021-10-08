package main

import "fmt"

// func main() {
// 	var score= []int{21,43,2,54,23}//[0]-21, [1]-43
// 	fmt.Print("Enter the index you want to delete: ")
// 	var index int
// 	fmt.Scan(&index)

// 	var tempArray = make([]int, len(score)-1)
// 	for i:=0; i<len(score); i++{
// 		if i <index {
// 			tempArray[i]= score[i]
// 		}
// 		if i >index {
// 			tempArray[i-1]= score[i]
// 		}
// 	}
// 	score = tempArray
// 	for i :=0; i <len(score); i++{
// 		fmt.Print(score[i],",")
// 	}
// }

func main() {
	var score= []int{12,32,43,54,64}
	var tempArray= make([]int, len(score)-1)
	fmt.Print("Enter the index you want to delete : ")
	var index int
	fmt.Scan(&index)
	if len(score)-1<index || index <0 {
		fmt.Println("\n"+"Please enter that contain index of score")
		return
	}
	for i :=0; i <len(score); i++{
		if i <index {
			tempArray[i]= score[i]
		}
		if i >index {
			tempArray[i-1]= score[i]
		}
	}
	score = tempArray
	for i:=0; i<len(score);i++{
		fmt.Println(score[i])
	}
}