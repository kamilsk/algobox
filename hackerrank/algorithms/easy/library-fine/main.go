package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

// Complete the libraryFine function below.
func libraryFine(d1 int32, m1 int32, y1 int32, d2 int32, m2 int32, y2 int32) int32 {
	returned := time.Date(int(y1), time.Month(m1), int(d1), 0, 0, 0, 0, time.Local)
	due := time.Date(int(y2), time.Month(m2), int(d2), 0, 0, 0, 0, time.Local)
	if due.Sub(returned) > 0 {
		return 0
	}
	if y2 == y1 {
		if m2 == m1 {
			return 15 * (d1 - d2)
		}
		return 500 * (m1 - m2)
	}
	return 10000
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	d1M1Y1 := strings.Split(readLine(reader), " ")

	d1Temp, err := strconv.ParseInt(d1M1Y1[0], 10, 64)
	checkError(err)
	d1 := int32(d1Temp)

	m1Temp, err := strconv.ParseInt(d1M1Y1[1], 10, 64)
	checkError(err)
	m1 := int32(m1Temp)

	y1Temp, err := strconv.ParseInt(d1M1Y1[2], 10, 64)
	checkError(err)
	y1 := int32(y1Temp)

	d2M2Y2 := strings.Split(readLine(reader), " ")

	d2Temp, err := strconv.ParseInt(d2M2Y2[0], 10, 64)
	checkError(err)
	d2 := int32(d2Temp)

	m2Temp, err := strconv.ParseInt(d2M2Y2[1], 10, 64)
	checkError(err)
	m2 := int32(m2Temp)

	y2Temp, err := strconv.ParseInt(d2M2Y2[2], 10, 64)
	checkError(err)
	y2 := int32(y2Temp)

	result := libraryFine(d1, m1, y1, d2, m2, y2)

	fmt.Fprintf(writer, "%d\n", result)

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
