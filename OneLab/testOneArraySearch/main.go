package main

import "fmt"

func main() {
	var scores= []int{12,23,34,45,65}
	fmt.Print("Enter the value: ")
	var value int
	fmt.Scan(&value)
	var isSearch bool
	for i:=0; i<len(scores); i++{
		if value== scores[i]{
			isSearch= true
			fmt.Print("Correct value: ", value, " index: ", i)
			break
		}
	}
	if !isSearch{
		fmt.Print("Not Found !", value)
		fmt.Println()
	}
}