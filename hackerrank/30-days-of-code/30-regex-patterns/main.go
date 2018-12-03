package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	NTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	N := int32(NTemp)

	type name string
	type email string

	db := map[email]name{}
	for NItr := 0; NItr < int(N); NItr++ {
		firstNameEmailID := strings.Split(readLine(reader), " ")
		db[email(firstNameEmailID[1])] = name(firstNameEmailID[0])
	}

	result := make([]string, 0, len(db))
	for emailID, firstName := range db {
		if strings.HasSuffix(string(emailID), "@gmail.com") {
			result = append(result, string(firstName))
		}
	}

	sort.Strings(result)
	for _, output := range result {
		fmt.Println(output)
	}
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
