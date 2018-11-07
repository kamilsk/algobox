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
	nn := make([]int, 0, 2048)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		nn = append(nn, n)
	}
	fmt.Println(strings.Join(convert(insertion(nn)), " "))
}

func insertion(array []int) []int {
	if len(array) < 2 {
		return array
	}
	var j int
	for i, size := 1, len(array); i < size; i++ {
		n := array[i]
		for j = i - 1; j >= 0 && n < array[j]; j-- {
			array[j+1] = array[j]
		}
		array[j+1] = n
	}
	return array
}

func convert(array []int) []string {
	converted := make([]string, len(array))
	for i, n := range array {
		converted[i] = strconv.Itoa(n)
	}
	return converted
}
