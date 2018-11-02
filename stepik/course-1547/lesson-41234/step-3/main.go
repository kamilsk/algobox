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
	count, _ := strconv.Atoi(scanner.Text())
	traffic := make([][2]int, 0, count)
	for i := 0; scanner.Scan() && i < count; i++ {
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

	type packet struct {
		start, end int
		buffer     int
	}
	result := make([]int, 0, len(traffic))
	flow := make([]packet, 0, len(traffic))

	for len(traffic) > 0 {
		arrival, duration := traffic[0][0], traffic[0][1]
		traffic = traffic[1:]

		if len(flow) == 0 {
			flow = append(flow, packet{arrival, arrival + duration, 1})
			continue
		}

		head := flow[len(flow)-1]
		if arrival >= head.end {
			flow = append(flow, packet{arrival, arrival + duration, head.buffer})
			continue
		}

		if head.buffer < capacity {
			flow = append(flow, packet{head.end, head.end + duration, head.buffer + 1})
			continue
		}
		flow = append(flow, packet{-1, head.end, head.buffer})
	}

	for len(flow) > 0 {
		var tail packet
		tail, flow = flow[0], flow[1:]
		result = append(result, tail.start)
	}

	return result
}
