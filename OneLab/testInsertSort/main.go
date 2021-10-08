package main

import "fmt"

// func main() {
// 	var scores = []int{12, 32, 43, 5, 4, 6}
// 	sort(scores)
// 	for i := 0; i < len(scores); i++ {
// 		fmt.Println(scores[i])
// 	}
// }

// func sort(arr []int) {
// 	for i := 0; i < len(arr); i++ {
// 		var insertElement = arr[i]
// 		var insertPosition = i
// 		for j := insertPosition - 1; j >= 0; j-- {
// 			if insertElement < arr[j] {
// 				arr[j+1] = arr[j]
// 				insertPosition--
// 			}
// 		}
// 		arr[insertPosition] = insertElement

// 	}
// }

func main() {
	var scores= []int{12,23,43,5,4,6}
	sort(scores)
	for i:=0; i <len(scores); i++{
		fmt.Println(scores[i])
	}
}

func sort(arr []int) {
	for i :=0; i<len(arr); i++{
		var insertElement= arr[i]
		var insertPosition= i
		for j:=insertPosition-1; j>=0; j--{
			if insertElement <arr[j]{
				arr[j+1]= arr[j]
				insertPosition--
			}
		}
		arr[insertPosition]=insertElement
	}
}