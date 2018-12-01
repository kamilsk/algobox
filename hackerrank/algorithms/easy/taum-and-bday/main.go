package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the taumBday function below.
func taumBday(b int64, w int64, bc int64, wc int64, z int64) int64 {
	if bc-wc > z {
		return (w+b)*wc + b*z
	}
	if wc-bc > z {
		return (w+b)*bc + w*z
	}
	return w*wc + b*bc
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		bw := strings.Split(readLine(reader), " ")

		b, err := strconv.ParseInt(bw[0], 10, 64)
		checkError(err)

		w, err := strconv.ParseInt(bw[1], 10, 64)
		checkError(err)

		bcWcz := strings.Split(readLine(reader), " ")

		bc, err := strconv.ParseInt(bcWcz[0], 10, 64)
		checkError(err)

		wc, err := strconv.ParseInt(bcWcz[1], 10, 64)
		checkError(err)

		z, err := strconv.ParseInt(bcWcz[2], 10, 64)
		checkError(err)

		result := taumBday(b, w, bc, wc, z)

		fmt.Fprintf(writer, "%d\n", result)
	}

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
