package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var _ = strconv.Itoa // Ignore this comment. You can still use the package "strconv".

	var i uint64 = 4
	var d float64 = 4.0
	var s string = "HackerRank "

	scanner := bufio.NewScanner(os.Stdin)

	//

	scanner.Split(bufio.ScanLines)

	_ = scanner.Scan()
	si, _ := strconv.ParseUint(scanner.Text(), 10, 64)

	_ = scanner.Scan()
	sd, _ := strconv.ParseFloat(scanner.Text(), 64)

	_ = scanner.Scan()
	ss := scanner.Text()

	fmt.Println(i + si)
	fmt.Printf("%.1f\n", d+sd)
	fmt.Println(s + ss)

	//
}
