package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		fmt.Println(strings.Join(split(n), " "))
	}
}

func split(n int) []string {
	result := make([]string, 0, 64)

	for n > 1 {
		if simple(n) {
			result = append(result, strconv.Itoa(n))
			break
		}
		for i, limit := 2, int(math.Sqrt(float64(n))); i <= limit; i++ {
			if n%i != 0 || !simple(i) {
				continue
			}
			result = append(result, strconv.Itoa(i))
			n /= i
			break
		}
	}

	return result
}

func simple(n int) bool {
	for i, limit := 2, int(math.Sqrt(float64(n))); i <= limit; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
