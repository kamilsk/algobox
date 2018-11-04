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
	plan := (&scheduler{n}).schedule(tasks)
	for _, process := range plan {
		fmt.Println(strings.Join(process, " "))
	}
}

type scheduler struct {
	cpus int
}

func (scheduler *scheduler) schedule(tasks []int) [][]string {
	return nil
}
