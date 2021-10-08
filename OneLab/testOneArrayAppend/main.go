package main

import "fmt"

// func main() {
// 	var scores = []int{90, 23, 43, 54, 35}
// 	var tempArray = make([]int, len(scores)+1)
// 	for i := 0; i < len(scores); i++ {
// 		tempArray[i] = scores[i]
// 	}
// 	tempArray[len(scores)] = 75

// 	fmt.Println(tempArray)
// 	scores = tempArray
// 	fmt.Println(len(scores))
// 	for i := 0; i < len(scores); i++ {
// 		fmt.Print(scores[i], " ")
// 	}
// }

func main() {
	score := []int{12,23,54,63,645}
	tempArray:=make([]int, len(score)+1)
	for i:=0; i<len(score); i++{
		tempArray[i]= score[i]
	}
	tempArray[len(score)]=75
	score= tempArray
	for i:=0; i<len(score); i++ {
		fmt.Print(score[i], )
	}
}