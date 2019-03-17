package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func fibonacciLastDigit(n int64) int64 {
	if n <= 1 {
		return n
	}

	var prev, curr int64 = 0, 1
	for i := int64(0); i < n-1; i++ {
		prev, curr = curr, prev+curr%10
	}

	return curr % 10
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	_ = scanner.Scan()
	n, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	safe(fmt.Fprintln(os.Stdout, fibonacciLastDigit(n)))
}

func safe(interface{}, error) {}
