package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'miniMaxSum' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func reduce(arr []int32, fnc func(prev int64, curr int64) int64) int64 {
	var reduced int64 = 0
	for _, el := range arr {
		reduced = fnc(reduced, int64(el))
	}
	return reduced
}

func miniMaxSum(arr []int32) {
	// Write your code here

	reduced := reduce(arr, func(prev int64, curr int64) int64 {
		return prev + curr
	})

	var min int64 = int64(arr[0])
	var max int64 = 0

	for _, el := range arr {
		if int64(el) < min {
			min = int64(el)
		}

		if int64(el) > max {
			max = int64(el)
		}
	}

	fmt.Printf("%d %d", reduced-max, reduced-min)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < 5; i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	miniMaxSum(arr)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
