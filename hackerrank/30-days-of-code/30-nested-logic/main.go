package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	_ = scanner.Scan()
	d1, m1, y1 := convert(strings.Split(scanner.Text(), " "))
	_ = scanner.Scan()
	d2, m2, y2 := convert(strings.Split(scanner.Text(), " "))
	fmt.Println(fine(d1, m1, y1, d2, m2, y2))
}

func convert(raw []string) (d, m, y int) {
	d, _ = strconv.Atoi(raw[0])
	m, _ = strconv.Atoi(raw[1])
	y, _ = strconv.Atoi(raw[2])
	return
}

// see https://www.hackerrank.com/challenges/library-fine/problem
func fine(d1, m1, y1, d2, m2, y2 int) int {
	returned := time.Date(y1, time.Month(m1), d1, 0, 0, 0, 0, time.Local)
	due := time.Date(y2, time.Month(m2), d2, 0, 0, 0, 0, time.Local)
	if due.Sub(returned) > 0 {
		return 0
	}
	if y2 == y1 {
		if m2 == m1 {
			return 15 * (d1 - d2)
		}
		return 500 * (m1 - m2)
	}
	return 10000
}
