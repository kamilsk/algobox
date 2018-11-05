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
	_ = scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())
	tasks := make([]int, 0, m)
	for i := 0; i < m && scanner.Scan(); i++ {
		duration, _ := strconv.Atoi(scanner.Text())
		tasks = append(tasks, duration)
	}
	for _, process := range new(scheduler).with(n).schedule(tasks) {
		fmt.Println(strings.Join(process, " "))
	}
}

type scheduler struct {
	cpus int
}

func (scheduler *scheduler) schedule(tasks []int) [][]string {
	system := heap{bus: make([]cpu, scheduler.cpus)}
	for i := 0; i < scheduler.cpus; i++ {
		system.bus[i].id = i
	}
	result := make([][]string, 0, len(tasks))
	for len(tasks) > 0 {
		result, tasks = append(result, system.run(tasks[0])), tasks[1:]
	}
	return result
}

func (scheduler *scheduler) with(cpus int) *scheduler {
	scheduler.cpus = cpus
	return scheduler
}

type cpu struct {
	id    int
	timer int
}

type heap struct {
	bus []cpu
}

func (heap *heap) run(duration int) []string {
	current := heap.bus[0]
	heap.bus[0].timer += duration
	siftDown(heap.bus, 0)
	return []string{strconv.Itoa(current.id), strconv.Itoa(current.timer)}
}

func siftDown(data []cpu, i int) {
	min := i
	l := left(i)
	if l < len(data) {
		if data[l].timer == data[min].timer && data[l].id < data[min].id {
			min = l
		}
		if data[l].timer < data[min].timer {
			min = l
		}
	}
	r := right(i)
	if r < len(data) {
		if data[r].timer == data[min].timer && data[r].id < data[min].id {
			min = r
		}
		if data[r].timer < data[min].timer {
			min = r
		}
	}
	if i != min {
		data[i], data[min] = data[min], data[i]
		siftDown(data, min)
	}
}

func left(i int) int {
	return 2*i + 1
}

func right(i int) int {
	return 2 * (i + 1)
}
