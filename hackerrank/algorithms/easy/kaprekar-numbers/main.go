package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the kaprekarNumbers function below.
func kaprekarNumbers(p int32, q int32) {
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	pTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	p := int32(pTemp)

	qTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	q := int32(qTemp)

	kaprekarNumbers(p, q)
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
