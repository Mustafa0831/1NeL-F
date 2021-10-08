package main

import "fmt"

func main() {
	arrays := []int{1, 2, 3, 4, 5}
	result := leftRotation(arrays, 2)
	fmt.Println(result)
}

// func leftRotation(arrays []int, n int) []int {
// 	var temp []int
// 	for n >= 1 {
// 		temp = arrays[1:]
// 		temp = append(temp, arrays[0])
// 		arrays = temp
// 		n--
// 	}
// 	return arrays
// }

func leftRotation(arr []int, num int) []int {
	temp := []int{}
	for num >= 1 {
		temp = arr[1:]
		temp = append(temp, arr[0])
		arr = temp
		num--
	}
	return arr
}
