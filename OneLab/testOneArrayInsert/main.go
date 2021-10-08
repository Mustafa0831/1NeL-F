package main

import "fmt"

// func main() {
// 	var scores = []int{23,234,54,6,3,2}
// 	var tempArray = make([]int, len(scores)+1)
// 	insert(scores, tempArray, 76,2)
// 	scores =tempArray
// 	for i:=0; i<len(scores); i++{
// 		fmt.Print(scores[i], " ")
// 	}
// }

// func insert(arr []int, temp []int, score, insertIndex int) {
// 	for i:=0; i<len(arr); i++{
// 		if i <insertIndex {
// 			temp[i]= arr[i]
// 		}else {
// 			temp[i+1]= arr[i]
// 		}
// 	}
// 	temp[insertIndex]=score
// }

func main() {
	var score= []int{12,12,43,5,4,3}
	var temp = [7]int{}
	var tempArray= make([]int, len(score)+1)
	fmt.Println(temp)
	fmt.Println(tempArray)
	insert(score, tempArray, 77, 2)
	score =tempArray
	for i:=0; i<len(score); i++ {
		// fmt.Println(score[i])
	}
}

func insert(arr,temp []int, score, index int){
	for i:=0; i<len(arr); i++{
		if i <index {
			temp[i]= arr[i]
		}else {
			temp[i+1]= arr[i]
		}
	}
	temp[index]= score
}