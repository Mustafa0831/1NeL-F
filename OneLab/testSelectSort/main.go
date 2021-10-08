package main

import "fmt"

func main(){
	var scores= []int{13,15,25,71,63}
	sort(scores)
	fmt.Println(scores)
	for i:=0; i<len(scores); i++{
		fmt.Println(scores[i])
	}
}

func sort(arr []int) {
	minIndex :=0
	for i:=0; i<len(arr)-1; i++{
		minIndex=i
		var minValue= arr[minIndex]
		for j:=i; j<len(arr)-1; j++{
			if minValue >arr[j+1]{
				minValue= arr[j+1]
				minIndex=j+1
			}
		}
		if i!=minIndex{
			arr[i], arr[minIndex]=arr[minIndex], arr[i]
		}
	}
}