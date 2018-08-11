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

var update = flag.Bool("test", false, "use test input")

var testCases = `3
10 20 30
3
9 15 35

10
0 10 20 30 40   50 60 70 80   90
5
-1 11 40 41 92`

func main() {
	var input io.Reader = os.Stdin
	if flag.Parse(); *update {
		input = strings.NewReader(testCases)
	}
	scan := bufio.NewScanner(input)
	scan.Split(bufio.ScanWords)
	for scan.Scan() {
		size, _ := strconv.Atoi(scan.Text())
		sorted := make([]int, size)
		for i := 0; i < size; i++ {
			scan.Scan()
			sorted[i], _ = strconv.Atoi(scan.Text())
		}
		scan.Scan()
		size, _ = strconv.Atoi(scan.Text())
		out := make([]string, size)
		for i := 0; i < size; i++ {
			scan.Scan()
			what, _ := strconv.Atoi(scan.Text())
			out[i] = strconv.Itoa(search(sorted, what))
		}
		fmt.Println(strings.Join(out, " "))
	}
}

func search(where []int, what int) (nearest int) {
	min := func(int1, int2 int) int {
		if int1 < int2 {
			return int1
		}
		return int2
	}

	bsearch := func(slice []int, what int) (nearest int) {
		left, right := 0, len(slice)-1
		for left+1 < right {
			middle := left + (right-left)/2
			if what < slice[middle] {
				right = middle
			} else {
				left = middle
			}
		}
		if what-slice[left] <= slice[right]-what {
			return left
		}
		return right
	}

	lsearch := func(slice []int, what int) (nearest int) {
		for i, end := 0, len(slice); i < end; i++ {
			if what < slice[i] {
				if prev := i - 1; prev > -1 && what-slice[prev] <= slice[i]-what {
					return prev
				}
				return i
			}
		}
		panic("unreachable")
	}

	start, end, power := 0, len(where)-1, uint(2)
	for start < end {
		last := min(1<<power, end)
		if what < where[last] {
			l, b := lsearch(where[start:last+1], what), bsearch(where[start:last+1], what)
			if l != b {
				panic(fmt.Errorf("%d != %d : %+v.find(%d)\n", l, b, where, what))
			}
			return start + b
		}
		power++
		start = last
	}
	return end
}
