package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func sumOfTwoDigits(first, second int64) int64 {
	return first + second
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	_ = scanner.Scan()
	a, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	_ = scanner.Scan()
	b, _ := strconv.ParseInt(scanner.Text(), 10, 64)

	safe(fmt.Fprintln(os.Stdout, sumOfTwoDigits(a, b)))
}

func safe(interface{}, error) {}
