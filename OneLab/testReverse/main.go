package main

import "fmt"

// func main() {
// 	var scores= []int{12,23,43,5,3}
// 	reverse(scores)
// 	for i:=0; i<len(scores); i++{
// 		fmt.Print(scores[i]," ")
// 	}
// 	fmt.Println()
// }

// func reverse(arr []int) {
// 	var middle = len(arr)/2
// 	for i:=0; i<middle; i++{
// 		arr[i], arr[len(arr)-i-1]= arr[len(arr)-i-1], arr[i]
// 	}
// }

func main() {
	var scores= []int{12,23,54,56,78}
	reverse(scores)
	for i:=0; i<len(scores); i++{
		fmt.Print(scores[i], " ")
	}
	fmt.Println()
}

func reverse(arr []int) {
	var middle= len(arr)/2
	for i:=0; i<middle; i++{
		arr[i], arr[len(arr)-i-1]= arr[len(arr)-i-1], arr[i]
	}
}