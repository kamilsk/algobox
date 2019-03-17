package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func fibonacciNumber(n int64) int64 {
	memo := append(make([]int64, 0, n), 0, 1)

	for i := int64(2); i < n+1; i++ {
		memo = append(memo, memo[i-1]+memo[i-2])
	}

	return memo[n]
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	_ = scanner.Scan()
	n, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	safe(fmt.Fprintln(os.Stdout, fibonacciNumber(n)))
}

func safe(interface{}, error) {}
