package main

import "fmt"

func main() {
	scores:= []int{12,433,56,32,56}
	fmt.Println(len(scores))
	fmt.Println(maxValue(scores))
}
func maxValue(arr []int)int {
	for i:=0; i<len(arr)-1; i++{
		if arr[i]>arr[i+1]{
			var temp = arr[i]
			arr[i]=arr[i+1]
			arr[i+1]= temp
			// arr[i],arr[i+1]= arr[i+1], arr[i]
		}
	}
	return arr[len(arr)-1]
}