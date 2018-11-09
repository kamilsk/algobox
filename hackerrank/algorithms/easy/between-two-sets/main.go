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
 * Complete the getTotalX function below.
 */
func getTotalX(a []int32, b []int32) int32 {
	// greatest common divisor
	fnGCD := func(x, y int32) int32 {
		if y == 0 {
			return x
		}
		for x%y != 0 {
			x, y = y, x%y
		}
		return y
	}

	// least common multiple
	fnLCM := func(x, y int32) int32 {
		return (x * y) / fnGCD(x, y)
	}

	reduce := func(fn func(x, y int32) int32, list []int32) int32 {
		process := make([]int32, len(list))
		copy(process, list)
		for len(process) > 1 {
			process[1] = fn(process[0], process[1])
			process = process[1:]
		}
		return process[0]
	}

	var counter int32
	lcm, gcd := reduce(fnLCM, a), reduce(fnGCD, b)
	step := lcm
	for step <= gcd {
		if gcd%step == 0 {
			counter++
		}
		step += lcm
	}
	return counter
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	outputFile, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer outputFile.Close()

	writer := bufio.NewWriterSize(outputFile, 1024*1024)

	nm := strings.Split(readLine(reader), " ")

	nTemp, err := strconv.ParseInt(nm[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	mTemp, err := strconv.ParseInt(nm[1], 10, 64)
	checkError(err)
	m := int32(mTemp)

	aTemp := strings.Split(readLine(reader), " ")

	var a []int32

	for aItr := 0; aItr < int(n); aItr++ {
		aItemTemp, err := strconv.ParseInt(aTemp[aItr], 10, 64)
		checkError(err)
		aItem := int32(aItemTemp)
		a = append(a, aItem)
	}

	bTemp := strings.Split(readLine(reader), " ")

	var b []int32

	for bItr := 0; bItr < int(m); bItr++ {
		bItemTemp, err := strconv.ParseInt(bTemp[bItr], 10, 64)
		checkError(err)
		bItem := int32(bItemTemp)
		b = append(b, bItem)
	}

	total := getTotalX(a, b)

	fmt.Fprintf(writer, "%d\n", total)

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
