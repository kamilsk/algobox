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
	_ = scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())
	fmt.Println(strings.Join(New(m).calculate(numbers), " "))
}

// New returns a new instance of sliding window above two stacks with maximums.
func New(size int) *window {
	return &window{
		in:   &stack{data: make([]int, 0, size), max: make([]int, 0, size)},
		out:  &stack{data: make([]int, 0, size), max: make([]int, 0, size)},
		size: size,
	}
}

type window struct {
	in, out *stack
	size    int
}

func (window *window) calculate(numbers []int) []string {
	if window.size == 0 || len(numbers) == 0 {
		return nil
	}
	capacity := len(numbers) - window.size + 1
	if capacity < 1 {
		capacity = 1
	}
	result := make([]string, 0, capacity)

	// optimization for the edge case
	if window.size == 1 {
		for _, num := range numbers {
			result = append(result, strconv.Itoa(num))
		}
		return result
	}

	var el int

	for i := 0; i < window.size && len(numbers) > 0; i++ {
		el, numbers = numbers[0], numbers[1:]
		window.in.push(el)
	}

	for {
		for !window.in.isEmpty() {
			el, _ = window.in.pop()
			window.out.push(el)
		}

		for !window.out.isEmpty() {
			maxIn, foundIn := window.in.maximum()
			max, _ := window.out.maximum()
			if foundIn && maxIn > max {
				max = maxIn
			}
			result = append(result, strconv.Itoa(max))
			_, _ = window.out.pop()

			if len(numbers) == 0 {
				return result
			}
			el, numbers = numbers[0], numbers[1:]
			window.in.push(el)
		}
	}
}

type stack struct {
	data []int
	max  []int
}

func (stack *stack) isEmpty() bool {
	return len(stack.data) == 0
}

func (stack *stack) push(el int) {
	max, found := stack.maximum()
	if !found || (found && max < el) {
		max = el
	}
	stack.max = append(stack.max, max)
	stack.data = append(stack.data, el)
}

func (stack *stack) pop() (el int, found bool) {
	if len(stack.data) == 0 {
		return
	}
	last := len(stack.data) - 1
	el, stack.data, stack.max = stack.data[last], stack.data[:last], stack.max[:last]
	return el, true
}

func (stack *stack) maximum() (max int, found bool) {
	// stack.data and stack.max must be consistent
	// panic if not (index out of range)
	if len(stack.data) == 0 {
		return
	}
	return stack.max[len(stack.data)-1], true
}
