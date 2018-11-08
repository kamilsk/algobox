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
	scanner.Split(bufio.ScanLines)
	_ = scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	points := make([]point, 0, n)
	for i := 0; i < n && scanner.Scan(); i++ {
		cc := strings.Split(scanner.Text(), " ")
		x, _ := strconv.Atoi(cc[0])
		y, _ := strconv.Atoi(cc[1])
		points = append(points, point{x, y})
	}
	for _, p := range new(heap).sort(points) {
		fmt.Println(strings.Join(convert(p), " "))
	}
}

func convert(p point) []string {
	return []string{strconv.Itoa(p.x), strconv.Itoa(p.y)}
}

type point struct {
	x, y int
}

type heap struct{}

func (heap *heap) sort(points []point) []point {
	heap.build(points)
	process := points
	for len(process) > 1 {
		last := len(process) - 1
		process[0], process[last] = process[last], process[0]
		process = process[:last]
		heap.siftDown(process, 0)
	}
	return points
}

func (heap *heap) build(points []point) {
	for i := len(points) / 2; i >= 0; i-- {
		heap.siftDown(points, i)
	}
}

func (heap *heap) siftDown(points []point, i int) {
	var (
		left = func(i int) int {
			return 2*i + 1
		}
		right = func(i int) int {
			return 2 * (i + 1)
		}
	)
	max := i
	l := left(i)
	if l < len(points) {
		if points[l].x == points[max].x && points[l].y > points[max].y {
			max = l
		}
		if points[l].x > points[max].x {
			max = l
		}
	}
	r := right(i)
	if r < len(points) {
		if points[r].x == points[max].x && points[r].y > points[max].y {
			max = r
		}
		if points[r].x > points[max].x {
			max = r
		}
	}
	if i != max {
		points[i], points[max] = points[max], points[i]
		heap.siftDown(points, max)
	}
}
