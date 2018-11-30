package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	_ = scanner.Scan()
	T, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < T && scanner.Scan(); i++ {
		n, _ := strconv.Atoi(scanner.Text())
		if simple(n) {
			fmt.Println("Prime")
			continue
		}
		fmt.Println("Not prime")
	}
}

func simple(n int) bool {
	if n == 1 {
		return false
	}
	for i, limit := 2, int(math.Sqrt(float64(n))); i <= limit; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
