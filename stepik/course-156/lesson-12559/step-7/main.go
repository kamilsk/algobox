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
	for i, d := 0, new(deq); i < n && scanner.Scan(); i++ {
		code, _ := strconv.Atoi(scanner.Text())
		_ = scanner.Scan()
		arg, _ := strconv.Atoi(scanner.Text())
		if !d.do(code, arg) {
			fmt.Println("NO")
			return
		}
	}
	fmt.Println("YES")
}

const (
	pushFront = 1 + iota
	popFront
	pushBack
	popBack
)

type deq []int

func (deq *deq) assert(operation func() int, expected int) bool {
	return operation() == expected
}

func (deq *deq) enqueue(element int) {
	*deq = append([]int{element}, *deq...)
}

func (deq *deq) dequeue() (element int) {
	if len(*deq) == 0 {
		return -1
	}
	element, *deq = (*deq)[0], (*deq)[1:]
	return
}

func (deq *deq) push(element int) {
	*deq = append(*deq, element)
}

func (deq *deq) pop() (element int) {
	last := len(*deq) - 1
	if last < 0 {
		return -1
	}
	element, *deq = (*deq)[last], (*deq)[:last]
	return
}

func (deq *deq) do(code, arg int) bool {
	switch code {
	case pushFront:
		deq.enqueue(arg)
	case popFront:
		if !deq.assert(deq.dequeue, arg) {
			return false
		}
	case pushBack:
		deq.push(arg)
	case popBack:
		if !deq.assert(deq.pop, arg) {
			return false
		}
	}
	return true
}

func (deq *deq) process(operations [][2]int) bool {
	for _, operation := range operations {
		if !deq.do(operation[0], operation[1]) {
			return false
		}
	}
	return true
}
