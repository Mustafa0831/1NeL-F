package main

import "fmt"
func main(){
	scores := []int{12,322,54,3,6}
	fmt.Println(minValue(scores))
}

func minValue(arr []int)int{
	minIndex:=0
	for j:=1; j<len(arr); j++{
		if arr[minIndex]>arr[j]{
			minIndex=j
		}
	}
	return arr[minIndex]
}