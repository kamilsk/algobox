package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	var arr [][]int32
	for i := 0; i < 6; i++ {
		arrRowTemp := strings.Split(readLine(reader), " ")

		var arrRow []int32
		for _, arrRowItem := range arrRowTemp {
			arrItemTemp, err := strconv.ParseInt(arrRowItem, 10, 64)
			checkError(err)
			arrItem := int32(arrItemTemp)
			arrRow = append(arrRow, arrItem)
		}

		if len(arrRow) != int(6) {
			panic("Bad input")
		}

		arr = append(arr, arrRow)
	}

	var max int32 = -9*7 - 1
	for i, floor := 0, 4; i < floor; i++ {
		for j, left := 0, 4; j < left; j++ {
			var sum int32
			sum += arr[i+0][j+0] + arr[i+0][j+1] + arr[i+0][j+2]
			sum += 0000000000000 + arr[i+1][j+1] + 0000000000000
			sum += arr[i+2][j+0] + arr[i+2][j+1] + arr[i+2][j+2]
			if max < sum {
				max = sum
			}
		}
	}
	fmt.Println(max)
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
