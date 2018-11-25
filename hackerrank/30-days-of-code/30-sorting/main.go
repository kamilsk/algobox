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
	nn := make([]int, 0, n)
	for i := 0; i < cap(nn) && scanner.Scan(); i++ {
		n, _ = strconv.Atoi(scanner.Text())
		nn = append(nn, n)
	}
	fmt.Printf("Array is sorted in %d swaps.\n", bubbleSort(nn))
	fmt.Printf("First Element: %d\n", nn[0])
	fmt.Printf("Last Element: %d\n", nn[len(nn)-1])
}

func bubbleSort(nn []int) int {
	var swaps int

	swap := -1
	for swap != 0 {
		swap = 0
		for i := 0; i < len(nn)-1; i++ {
			if nn[i] > nn[i+1] {
				nn[i], nn[i+1] = nn[i+1], nn[i]
				swap++
			}
		}
		swaps += swap
	}

	return swaps
}
