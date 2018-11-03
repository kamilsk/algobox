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
	parents := make([]int, 0, n)
	for i := 0; i < n && scanner.Scan(); i++ {
		parent, _ := strconv.Atoi(scanner.Text())
		parents = append(parents, parent)
	}
	fmt.Println(height(parents))
}

func height(parents []int) int {
	var max int
	depths, root := map[int]int{}, -1
	for i, parent := range parents {
		depth, until := 1, root

		for parent != root {
			if depths[parent] > 0 {
				until = parent
				depth += depths[parent]
				break
			}
			parent = parents[parent]
			depth++
		}

		if depths[i] < depth {
			depths[i] = depth
		}
		if max < depth {
			max = depth
		}

		parent = parents[i]
		for parent != until {
			depth--
			if depths[parent] < depth {
				depths[parent] = depth
			}
			parent = parents[parent]
		}
	}
	return max
}
