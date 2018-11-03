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
	size, _ := strconv.Atoi(scanner.Text())
	_ = scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	traffic := make([][2]int, 0, n)
	for i := 0; i < n && scanner.Scan(); i++ {
		arrival, _ := strconv.Atoi(scanner.Text())
		_ = scanner.Scan()
		duration, _ := strconv.Atoi(scanner.Text())
		traffic = append(traffic, [2]int{arrival, duration})
	}
	for _, ts := range process(traffic, size) {
		fmt.Println(ts)
	}
}

func process(traffic [][2]int, capacity int) []int {
	if len(traffic) == 0 {
		return nil
	}

	result := make([]int, 0, len(traffic))
	fl := &flow{cap: capacity}

	for len(traffic) > 0 {
		p := packet{arrival: traffic[0][0], duration: traffic[0][1]}
		traffic = traffic[1:]
		fl.push(&p)
		result = append(result, p.start)
	}

	return result
}

type packet struct {
	arrival, duration int
	start, end        int
	next              *packet
}

type flow struct {
	cap, size int
	head      *packet
	tail      *packet
}

func (f *flow) push(p *packet) {
	f.shrink(p.arrival)
	if f.head == nil {
		p.start, p.end = p.arrival, p.arrival+p.duration
		f.head, f.tail = p, p
		f.size++
		return
	}
	if f.size < f.cap {
		p.start, p.end = f.tail.end, f.tail.end+p.duration
		f.tail.next, f.tail = p, p
		f.size++
		return
	}
	p.start, p.end = -1, f.tail.end
}

func (f *flow) shrink(arrival int) {
	for f.head != nil && f.head.end <= arrival {
		f.head = f.head.next
		f.size--
	}
}
