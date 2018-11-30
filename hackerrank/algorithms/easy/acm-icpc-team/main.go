package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the acmTeam function below.
func acmTeam(topic []string) []int32 {
	var howMach, howMany int32
	for len(topic) > 1 {
		var compare string
		compare, topic = topic[0], topic[1:]
		for _, t := range topic {
			coverage := int32(0)
			for i := range compare {
				if compare[i] == '1' || t[i] == '1' {
					coverage++
				}
			}
			if coverage > howMach {
				howMach, howMany = coverage, 1
				continue
			}
			if coverage == howMach {
				howMany++
				continue
			}
		}
	}
	return []int32{howMach, howMany}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nm := strings.Split(readLine(reader), " ")

	nTemp, err := strconv.ParseInt(nm[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	mTemp, err := strconv.ParseInt(nm[1], 10, 64)
	checkError(err)
	m := int32(mTemp)

	var topic []string

	for i := 0; i < int(n); i++ {
		topicItem := readLine(reader)
		if int32(len(topicItem)) > m {
			panic("length of input is grater then limit")
		}
		topic = append(topic, topicItem)
	}

	result := acmTeam(topic)

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
