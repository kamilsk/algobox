package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	_ = scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	total := 2 * n
	input := make([]int, total)
	for i := 0; i < total && scanner.Scan(); i++ {
		input[i], _ = strconv.Atoi(scanner.Text())
	}
	fmt.Println(compare(input[:n], input[n:]))
}

func compare(first []int, second []int) (i0, j0 int) {
	// assert(len(first) == len(second))
	maxI, maxSum := i0, first[i0]+second[j0]
	for i, size := 0, len(first); i < size; i++ {
		if first[i] > first[maxI] {
			maxI = i
		}
		if sum := first[maxI] + second[i]; sum > maxSum {
			i0, j0, maxSum = maxI, i, sum
		}
	}
	return i0, j0
}
