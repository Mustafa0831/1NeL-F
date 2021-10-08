package main

import (
	"fmt"
)

/*
 * Complete the 'closestNumbers' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func closestNumbers(arr []int32) []int32 {
	// Write your code here
	var minVal int32
	var res []int32
	sort(arr, 0, len(arr)-1)
	minVal = arr[1] - arr[0]
	for i := 2; i < len(arr); i++ {
		if minVal > arr[i]-arr[i-1] {
			minVal = arr[i] - arr[i-1]
		}
	}
    fmt.Println(minVal)
	for i := 1; i < len(arr); i++ {
		if minVal == arr[i]-arr[i-1] {
			res = append(res, arr[i-1])
			res = append(res, arr[i])
		}
    }
	return res
}

func sort(arr []int32, low, high int) {
	if low > high {
		return
	}
	pivot := quickSort(arr, low, high)
	sort(arr, low, pivot-1)
	sort(arr, pivot+1, high)
}

func quickSort(arr []int32, low int, high int) int {
	pivot := arr[high]
	for i := low; i < high; i++ {
		if arr[i] < pivot {
			arr[i], arr[low] = arr[low], arr[i]
			low++
		}
	}
	arr[low], arr[high] = arr[high], arr[low]
	return low
}

func main() {
	values := []int32{5, 2, 3, 4,1}
	fmt.Println(closestNumbers(values))
	//     reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

	//     stdout, err := os.Create(("OUTPUT_PATH.txt"))
	//     // checkError(err)

	//     defer stdout.Close()

	//     writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

	//     nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	//     checkError(err)
	//     n := int32(nTemp)

	//     arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	//     var arr []int32

	//     for i := 0; i < int(n); i++ {
	//         arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
	//         checkError(err)
	//         arrItem := int32(arrItemTemp)
	//         arr = append(arr, arrItem)
	//     }

	//     result := closestNumbers(arr)

	//     for i, resultItem := range result {
	//         fmt.Fprintf(writer, "%d", resultItem)

	//         if i != len(result) - 1 {
	//             fmt.Fprintf(writer, " ")
	//         }
	//     }

	//     fmt.Fprintf(writer, "\n")

	//     writer.Flush()
	// }

	// func readLine(reader *bufio.Reader) string {
	//     str, _, err := reader.ReadLine()
	//     if err == io.EOF {
	//         return ""
	//     }

	//     return strings.TrimRight(string(str), "\r\n")
	// }

	// func checkError(err error) {
	//     if err != nil {
	//         panic(err)
	//     }
}
