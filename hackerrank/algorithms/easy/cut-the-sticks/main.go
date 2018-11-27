package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the cutTheSticks function below.
func cutTheSticks(arr []int32) []int32 {
	sticks := make([]int32, 0, len(arr))
	count := func(nn []int32) int32 {
		var i int32
		for _, n := range nn {
			if n != 0 {
				i++
			}
		}
		return i
	}
	min := func(nn []int32) int32 {
		m := int32(1001)
		for _, n := range nn {
			if n != 0 && n < m {
				m = n
			}
		}
		return m
	}
	cut := func(m int32, nn []int32) {
		for i := range nn {
			if nn[i] != 0 {
				nn[i] -= m
			}
		}
	}
	for i := int32(len(arr)); i > 0; i = count(arr) {
		sticks = append(sticks, i)
		cut(min(arr), arr)
	}
	return sticks
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	result := cutTheSticks(arr)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%d", resultItem)

		if i != len(result)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

	writer.Flush()
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
