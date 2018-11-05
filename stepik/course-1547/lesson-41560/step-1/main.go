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
	swaps := new(heap).build(numbers)
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
	min := i
	l := left(i)
	if l < len(heap.data) && heap.data[l] < heap.data[min] {
		min = l
	}
	r := right(i)
	if r < len(heap.data) && heap.data[r] < heap.data[min] {
		min = r
	}
	if i != min {
		heap.data[i], heap.data[min] = heap.data[min], heap.data[i]
		swap := make([][]string, 0, 4)
		swap = append(swap, []string{strconv.Itoa(i), strconv.Itoa(min)})
		return append(swap, heap.siftDown(min)...)
	}
	return nil
}

func left(i int) int {
	return 2*i + 1
}

func right(i int) int {
	return 2 * (i + 1)
}
