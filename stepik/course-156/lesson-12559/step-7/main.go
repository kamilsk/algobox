package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var test = flag.Bool("test", false, "use test input")

var testCases = `5
1 44
3 50
2 44
2 50
2 -1

7
4 44
4 -1
3 22
4 22
2 -1
1 11
4 11`

func main() {
	var input io.Reader = os.Stdin
	if flag.Parse(); *test {
		input = strings.NewReader(testCases)
	}
	scan := bufio.NewScanner(input)
	scan.Split(bufio.ScanWords)

	stepik := true // hack for stepik.org because it provides invalid input at test #6

	for stepik && scan.Scan() {
		size, _ := strconv.Atoi(scan.Text())
		success := new(bool)
		for i, deq := 0, new(DEQ); i < size; i++ {
			operation, element := read(scan)
			switch operation {
			case PushFront:
				deq.enqueue(element)
			case PopFront:
				*success = deq.assert(deq.dequeue, element)
				if !*success {
					rewind(scan, size-i)
					i = size
				}
			case PushBack:
				deq.push(element)
			case PopBack:
				*success = deq.assert(deq.pop, element)
				if !*success {
					rewind(scan, size-i)
					i = size
				}
			}
		}
		stepik = *test
		if success != nil {
			if *success {
				fmt.Println("YES")
				continue
			}
			fmt.Println("NO")
		}
	}
}

// Operations above the DEQ.
const (
	PushFront = 1 + iota
	PopFront
	PushBack
	PopBack
)

func read(scan *bufio.Scanner) (int, int) {
	scan.Scan()
	first, _ := strconv.Atoi(scan.Text())
	scan.Scan()
	second, _ := strconv.Atoi(scan.Text())
	return first, second
}

func rewind(scan *bufio.Scanner, count int) {
	for i := 0; i < count; i++ {
		_, _ = read(scan)
	}
}

// DEQ is a double-ended queue based on slice.
type DEQ []int

func (d *DEQ) assert(operation func() int, expected int) bool {
	return operation() == expected
}

func (d *DEQ) enqueue(element int) {
	*d = append([]int{element}, *d...)
}

func (d *DEQ) dequeue() (element int) {
	if len(*d) == 0 {
		return -1
	}
	element, *d = (*d)[0], (*d)[1:]
	return
}

func (d *DEQ) push(element int) {
	*d = append(*d, element)
}

func (d *DEQ) pop() (element int) {
	last := len(*d) - 1
	if last < 0 {
		return -1
	}
	element, *d = (*d)[last], (*d)[:last]
	return
}
