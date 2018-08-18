package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode/utf8"
)

var test = flag.Bool("test", false, "use test input")

var testCases = `
}[[([{[]}  -> {}[[([{[]}])]]
{[({[({[(  -> {[({[({[()]})]})]}
)]})]})]}  -> {[({[({[()]})]})]}
{][[[[{}[] -> IMPOSSIBLE`

func main() {
	var input io.Reader = os.Stdin
	if flag.Parse(); *test {
		input = strings.NewReader(testCases)
	}
	scan := bufio.NewScanner(input)
	scan.Split(bufio.ScanWords)
	for scan.Scan() {
		in, expected := read(scan)
		result := fix(in)
		if *test && result != expected {
			panic(fmt.Errorf("%q != %q : `%s`.fix()", result, expected, in))
		}
		fmt.Println(result)
	}
}

var (
	opening = map[rune]rune{'}': '{', ']': '[', ')': '('}
	closing = map[rune]rune{'{': '}', '[': ']', '(': ')'}
)

type stack []rune

func (s *stack) push(r rune) {
	*s = append(*s, r)
}

func (s *stack) pop() (rune, bool) {
	var r rune
	size := len(*s)
	if size < 1 {
		return r, false
	}
	r, *s = (*s)[size-1], (*s)[:size-1]
	return r, true
}

func read(scan *bufio.Scanner) (in, out string) {
	in = scan.Text()
	if *test {
		scan.Scan()
		scan.Scan()
		out = scan.Text()
	}
	return
}

func fix(in string) string {
	buf, size := bytes.NewBufferString(in), utf8.RuneCountInString(in)
	front, back := make(stack, 0, size), make(stack, 0, size)
	for _, bracket := range in {
		if _, is := closing[bracket]; is {
			back.push(bracket)
			continue
		}
		op := opening[bracket]
		prev, exists := back.pop()
		if !exists {
			front.push(bracket)
		} else if op != prev {
			return "IMPOSSIBLE"
		}
	}
	for {
		op, has := back.pop()
		if !has {
			break
		}
		buf.WriteRune(closing[op])
	}
	prefix := bytes.NewBuffer(nil)
	for {
		cl, has := front.pop()
		if !has {
			break
		}
		prefix.WriteRune(opening[cl])
	}
	return prefix.String() + buf.String()
}
