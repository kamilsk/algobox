package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func gcd(a, b int64) int64 {
	if b == 0 {
		return a
	}

	if a < b {
		return gcd(a, b%a)
	}
	return gcd(b, a%b)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	_ = scanner.Scan()
	a, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	_ = scanner.Scan()
	b, _ := strconv.ParseInt(scanner.Text(), 10, 64)

	safe(fmt.Fprintln(os.Stdout, gcd(a, b)))
}

func safe(interface{}, error) {}
