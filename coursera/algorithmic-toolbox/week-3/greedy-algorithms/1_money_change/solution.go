package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func change(m int64) int64 {
	var coins int64
	pars := []int64{10, 5, 1}
	for _, par := range pars {
		coins += m / par
		m %= par
	}
	return coins
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	_ = scanner.Scan()
	n, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	safe(fmt.Fprintln(os.Stdout, change(n)))
}

func safe(interface{}, error) {}
