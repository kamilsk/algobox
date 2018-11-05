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
	sorted := make([]int, n)
	for i := 0; i < n && scanner.Scan(); i++ {
		sorted[i], _ = strconv.Atoi(scanner.Text())
	}
	_ = scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())
	find := make([]int, m)
	for i := 0; i < m && scanner.Scan(); i++ {
		find[i], _ = strconv.Atoi(scanner.Text())
	}
	fmt.Println(strings.Join(search(sorted, find), " "))
}

func search(where []int, what []int) []string {
	min := func(int1, int2 int) int {
		if int1 < int2 {
			return int1
		}
		return int2
	}

	binary := func(slice []int, what int) (nearest int) {
		left, right := 0, len(slice)-1
		for left+1 < right {
			middle := left + (right-left)/2
			if what < slice[middle] {
				right = middle
			} else {
				left = middle
			}
		}
		if what-slice[left] <= slice[right]-what {
			return left
		}
		return right
	}

	result := make([]string, 0, len(what))
	for len(what) > 0 {
		start, end, power := 0, len(where)-1, uint(2)
		for start < end {
			last := min(1<<power, end)
			if what[0] < where[last] {
				end = start + binary(where[start:last+1], what[0])
				break
			}
			power++
			start = last
		}
		result = append(result, strconv.Itoa(end))
		what = what[1:]
	}
	return result
}
