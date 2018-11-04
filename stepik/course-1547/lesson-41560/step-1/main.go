package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	_ = scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	numbers := make([]int, 0, n)
	for i := 0; i < n && scanner.Scan(); i++ {
		num, _ := strconv.Atoi(scanner.Text())
		numbers = append(numbers, num)
	}
	swaps := (&heap{}).build(numbers)
	fmt.Println(len(swaps))
	for _, swap := range swaps {
		fmt.Println(strings.Join(swap, " "))
	}
}

type heap struct {
	data []int
}

func (heap *heap) build(numbers []int) [][]string {
	swaps := make([][]string, 0, len(numbers))
	heap.data = make([]int, len(numbers))
	copy(heap.data, numbers)
	return swaps
}
