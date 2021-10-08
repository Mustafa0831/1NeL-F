package main

import "fmt"

func main() {
	score :=[]int{90,23,53,56,75}
	sort(score)
	// for i:=0;i<len(score); i++{

	// 	fmt.Print(score[i], " ")

	// }
	fmt.Println(score)
}

func sort(arr []int) []int {
	for i:=0; i <len(arr); i++{
		for j:=0; j<len(arr)-i-1; j++{
			if arr[j]> arr[j+1]{
				var temp = arr[j]
				arr[j]=arr[j+1]
				arr[j+1]=temp
			}
		}
	}
	return arr
}