package main

import "fmt"

func main() {
	var scores = []int{30, 40, 70, 90, 101}
	var searchValue = 40
	var position = binarySearch(scores, searchValue)
	fmt.Print("position: ", searchValue, position)
	fmt.Print("\n-------\n")
	searchValue = 90
	position = binarySearch(scores, searchValue)
	fmt.Print("position: ", searchValue, position, "\n")
}

func binarySearch(arr []int, searchValue int)  int{
	var leftPointer = 0
	var rightPointer = len(arr)-1
	for leftPointer<=rightPointer{
		var midPointer=(leftPointer+rightPointer)/2
		var midValue=arr[midPointer]
		if midValue== searchValue{
			return midPointer
		}else if midValue <searchValue{
			leftPointer=midPointer+1
		}else {
			rightPointer= midPointer-1
		}
	}
	return -1
}
