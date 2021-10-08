package main

import "fmt"

func main()  {
	var scores= []int32{50, 65, 99, 87, 74, 63, 76, 100, 92}
	sort(scores)
	for i:=0; i<len(scores); i++{
		fmt.Println(scores[i])
	}
	fmt.Println(scores)
}

func sort(arr []int32){
	if len(arr)>0{
		quickSort(arr, 0, len(arr)-1)
	}
}

func quickSort(arr []int32, low int, high int){
	if low >high {
		return
	}
	var i = low
	var j=high
	var threshold= arr[low]
	for {
		if i>=j {
			break
		}
		for {
			if i>j || arr[j]<=threshold{
				break
			}
			j--
		}
		if i<j{
			arr[i]= arr[j]
			i++
		}
		for {
			if i>=j ||arr[i]>threshold{
				break
			}
			i++
		}
		if i<j{
			arr[j]=arr[i]
			j--
		}
	}
	arr[i]=threshold
	quickSort(arr, low, i-1)
	quickSort(arr,i+1, high)
}