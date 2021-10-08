package main

import (
	"fmt"
)

/*
 * Complete the 'countPairs' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY numbers
 *  2. INTEGER k
 */
func countPairs(numbers []int32, k int32) int32 {
	dict := make(map[int32]bool)
	for _, v := range numbers {
		dict[v] = true
		// fmt.Println(dict[v])
		// fmt.Println(dict)
	}
	// fmt.Println(dict)
	var res int32
	for key := range dict {
		if dict[key+k] {
			res++
		}
		// fmt.Println(dict[k+key], " ")
		// fmt.Println(res)
	}
	// fmt.Println(dict)
	// fmt.Println(dict[k])
	return res
}

func main() {
	numbers := []int32{5, 1, 3, 4, 2}
	fmt.Println(countPairs(numbers,2))
	// 	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	// 	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	// 	checkError(err)

	// 	defer stdout.Close()

	// 	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	// 	numbersCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	// 	checkError(err)

	// 	var numbers []int32

	// 	for i := 0; i < int(numbersCount); i++ {
	// 		numbersItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	// 		checkError(err)
	// 		numbersItem := int32(numbersItemTemp)
	// 		numbers = append(numbers, numbersItem)
	// 	}

	// 	kTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	// 	checkError(err)
	// 	k := int32(kTemp)

	// 	result := countPairs(numbers, k)

	// 	fmt.Fprintf(writer, "%d\n", result)

	// 	writer.Flush()
	// }

	// func readLine(reader *bufio.Rfunc countPairs(numbers []int32, k int32) int32 {

	// 	}

	// 	return strings.TrimRight(string(str), "\r\n")
	// }

	// func checkError(err error) {
	// 	if err != nil {
	// 		panic(err)
	// 	}
}
