package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	_ = scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < n && scanner.Scan(); i++ {
		str := scanner.Text()
		for j, ch := range str {
			if j%2 == 0 {
				fmt.Print(string(ch))
			}
		}
		fmt.Print(" ")
		for j, ch := range str {
			if j%2 != 0 {
				fmt.Print(string(ch))
			}
		}
		fmt.Println()
	}
}
