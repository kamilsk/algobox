package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode/utf8"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	_ = scanner.Scan()
	if mistake := check(scanner.Text()); mistake != -1 {
		fmt.Println(mistake)
		return
	}
	fmt.Println("Success")
}

func check(input string) int {
	type bracket struct {
		i int
		v rune
	}

	brackets := map[rune]struct {
		opening bool
		reverse rune
	}{
		'[': {true, ']'},
		'{': {true, '}'},
		'(': {true, ')'},
		']': {false, '['},
		'}': {false, '{'},
		')': {false, '('},
	}

	result, stack := -1, make([]bracket, 0, utf8.RuneCountInString(input))
	pop := func() bracket {
		var br bracket
		if last := len(stack) - 1; last > -1 {
			br, stack = stack[last], stack[:last]
		}
		return br
	}

	for i, r := range input {
		i++
		if br, is := brackets[r]; is {
			if br.opening {
				stack = append(stack, bracket{i, r})
				continue
			}
			if pop().v != br.reverse {
				return i
			}
		}
	}
	if len(stack) > 0 {
		return pop().i
	}

	return result
}
