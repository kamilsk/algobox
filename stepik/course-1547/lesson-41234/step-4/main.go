package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	opMax  = "max"
	opPop  = "pop"
	opPush = "push "
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	_ = scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	commands := make([]string, 0, n)
	for i := 0; i < n && scanner.Scan(); i++ {
		commands = append(commands, scanner.Text())
	}
	for _, max := range New(n).process(commands) {
		fmt.Println(max)
	}
}

// New returns a new instance of stack with maximums.
func New(capacity int) *stack {
	return &stack{data: make([]int, 0, capacity), max: make([]int, 0, capacity)}
}

type stack struct {
	data []int
	max  []int
}

func (stack *stack) process(commands []string) []int {
	maximums := make([]int, 0, len(commands)/10)
	for _, command := range commands {
		command = strings.ToLower(command)
		switch {
		case command == opMax:
			max, found := stack.maximum()
			if found {
				maximums = append(maximums, max)
			}
		case command == opPop:
			_, _ = stack.pop()
		case strings.HasPrefix(command, opPush):
			el, _ := strconv.Atoi(strings.TrimPrefix(command, opPush))
			stack.push(el)
		}
	}
	return maximums
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
