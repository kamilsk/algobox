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
	nn := make([]int, 0, 512)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		nn = append(nn, n)
	}
	fmt.Println(strings.Join(convert(insertion(nn)), " "))
}

func selection(array []int) []int {
	for i, size := 0, len(array); i < size-1; i++ {
		min := i
		for j := i + 1; j < size; j++ {
			if array[j] < array[min] {
				min = j
			}
		}
		if min != i {
			array[i], array[min] = array[min], array[i]
		}
	}
	return array
}

func insertion(array []int) []int {
	for i, j, size := 1, 0, len(array); i < size; i++ {
		n := array[i]
		for j = i - 1; j >= 0 && n < array[j]; j-- {
			array[j+1] = array[j]
		}
		array[j+1] = n
	}
	return array
}

func bubble(array []int) []int {
	return array
}

func convert(array []int) []string {
	converted := make([]string, len(array))
	for i, n := range array {
		converted[i] = strconv.Itoa(n)
	}
	return converted
}
