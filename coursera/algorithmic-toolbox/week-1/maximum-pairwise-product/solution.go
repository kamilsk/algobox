package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func maxPairwiseProduct(numbers []int64) int64 {
	var (
		local int64
		idx1, idx2 int
	)
	local, idx1 = 0, -1
	for k, v := range numbers {
		if v >= local {
			local, idx1 = v, k
		}
	}
	local, idx2 = 0, -1
	for k, v := range numbers {
		if k != idx1 && v >= local {
			local, idx2 = v, k
		}
	}
	return numbers[idx1] * numbers[idx2]
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	_ = scanner.Scan()
	n, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	numbers := make([]int64, 0, n)
	for n > 0 && scanner.Scan() {
		n--
		number, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		numbers = append(numbers, number)
	}

	safe(fmt.Fprintln(os.Stdout, maxPairwiseProduct(numbers)))
}

func safe(interface{}, error) {}
