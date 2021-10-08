package main

import "fmt"

func main() {
	var scores = []int32{50, 65, 99, 87, 74, 63, 76, 100, 92}
	sort(scores, 0, len(scores)-1)
	for i := 0; i < len(scores); i++ {
		fmt.Println(scores[i])
	}
	fmt.Println(scores)
}

func sort(arr []int32, low, high int) {
	if low > high {
		return
	}
	pivot:=quickSort(arr, low, high)
	sort(arr, low,pivot-1)
	sort(arr, pivot+1, high)
}

func quickSort(arr []int32, low int, high int) int {
	pivot := arr[high]
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[j], arr[low] = arr[low], arr[j]
			low++
		}
	}
	arr[low], arr[high]= arr[high], arr[low]
	return low
}
