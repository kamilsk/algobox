package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

var test = flag.Bool("test", false, "use test input")

var testCases = `3
9 7 3`

func main() {
	var input io.Reader = os.Stdin
	if flag.Parse(); *test {
		input = strings.NewReader(testCases)
	}
	scan := bufio.NewScanner(input)
	scan.Split(bufio.ScanWords)
	for scan.Scan() {
		tasks, _ := strconv.Atoi(scan.Text())
		durations := make(taskByDuration, tasks)
		for i := 0; i < tasks && scan.Scan(); i++ {
			duration, _ := strconv.Atoi(scan.Text())
			durations[i][0] = i
			durations[i][1] = duration
		}
		sort.Sort(durations)
		index := make([]string, 0, len(durations))
		for _, duration := range durations {
			index = append(index, strconv.Itoa(duration[0]))
		}
		fmt.Println(strings.Join(index, " "))
	}
}

type taskByDuration [][2]int

func (l taskByDuration) Len() int {
	return len(l)
}

func (l taskByDuration) Less(i, j int) bool {
	return l[i][1] < l[j][1]
}

func (l taskByDuration) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}
