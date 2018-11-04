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
	for i := int(len(heap.data) / 2); i >= 0; i-- {
		swaps = append(swaps, heap.siftDown(i)...)
	}
	return swaps
}

func (heap *heap) siftDown(i int) [][]string {
	max := i
	l := left(i)
	if l < len(heap.data) && heap.data[l] < heap.data[max] {
		max = l
	}
	r := right(i)
	if r < len(heap.data) && heap.data[r] < heap.data[max] {
		max = r
	}
	if i != max {
		heap.data[i], heap.data[max] = heap.data[max], heap.data[i]
		swap := make([][]string, 0, 4)
		swap = append(swap, []string{strconv.FormatInt(int64(i), 10), strconv.FormatInt(int64(max), 10)})
		return append(swap, heap.siftDown(max)...)
	}
	return nil
}

func left(i int) int {
	return 2*i + 1
}

func right(i int) int {
	return 2 * (i + 1)
}
