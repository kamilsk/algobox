package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() { sumOfTwoDigits(os.Stdin, os.Stdout) }

func sumOfTwoDigits(reader io.Reader, writer io.Writer) {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)

	_ = scanner.Scan()
	a, _ := strconv.Atoi(scanner.Text())
	_ = scanner.Scan()
	b, _ := strconv.Atoi(scanner.Text())

	safe(fmt.Fprintln(writer, a+b))
}

func safe(interface{}, error) {}
