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

var testCases = `4
 4  -8 6 0
-10  3 1 1

6
1 1  39 40 1 1
1 30 29 1  1 1`

func main() {
	var input io.Reader = os.Stdin
	if flag.Parse(); *test {
		input = strings.NewReader(testCases)
	}
	scan := bufio.NewScanner(input)
	scan.Split(bufio.ScanWords)
	for scan.Scan() {
		size, _ := strconv.Atoi(scan.Text())
		in := make([]int, 2*size)
		for i, limit := 0, 2*size; i < limit; i++ {
			scan.Scan()
			in[i], _ = strconv.Atoi(scan.Text())
		}
		first, second := in[:size], in[size:]
		maxI, maxJ := 0, 0
		maxF, maxSum := maxI, first[maxI]+second[maxJ]
		for j := 0; j < size; j++ {
			if first[j] > first[maxF] {
				maxF = j
			}
			if sum := first[maxF] + second[j]; sum > maxSum {
				maxI, maxJ, maxSum = maxF, j, sum
			}
		}
		fmt.Println(maxI, maxJ)
	}
}
