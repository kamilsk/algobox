package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

var update = flag.Bool("test", false, "use test input")

var testCases = `75 2 109`

func main() {
	var input io.Reader = os.Stdin
	if flag.Parse(); *update {
		input = strings.NewReader(testCases)
	}
	scan := bufio.NewScanner(input)
	scan.Split(bufio.ScanWords)
	for scan.Scan() {
		n, _ := strconv.Atoi(scan.Text())
		result := make([]string, 0, 64)
		for n > 1 {
			if simple(n) {
				result = append(result, strconv.Itoa(n))
				break
			}
			for i, limit := 2, int(math.Sqrt(float64(n))); i <= limit; i++ {
				if n%i != 0 || !simple(i) {
					continue
				}
				result = append(result, strconv.Itoa(i))
				n /= i
				break
			}
		}
		fmt.Println(strings.Join(result, " "))
	}
}

func simple(n int) bool {
	for i, limit := 2, int(math.Sqrt(float64(n))); i <= limit; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
