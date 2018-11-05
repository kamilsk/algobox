package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"unicode/utf8"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		fmt.Println(fix(scanner.Text()))
	}
}

type stack []rune

func (s *stack) push(char rune) {
	*s = append(*s, char)
}

func (s *stack) pop() (rune, bool) {
	var r rune
	size := len(*s)
	if size == 0 {
		return r, false
	}
	r, *s = (*s)[size-1], (*s)[:size-1]
	return r, true
}

func fix(input string) string {
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
	var (
		size   = utf8.RuneCountInString(input)
		buffer = bytes.NewBuffer(make([]byte, 0, 2*size))
	)
	front, back := make(stack, 0, size), make(stack, 0, size)
	for _, char := range input {
		bracket := brackets[char]
		if bracket.opening {
			back.push(char)
			continue
		}
		if len(back) == 0 {
			front.push(char)
			continue
		}
		prev, _ := back.pop()
		if bracket.reverse != prev {
			return "IMPOSSIBLE"
		}
	}
	for len(front) > 0 {
		char, _ := front.pop()
		buffer.WriteRune(brackets[char].reverse)
	}
	buffer.WriteString(input)
	for len(back) > 0 {
		char, _ := back.pop()
		buffer.WriteRune(brackets[char].reverse)
	}
	return buffer.String()
}
