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

var update = flag.Bool("test", false, "use test input")

var testCases = `4 4 0
0 4097
4096 8193
8192 12289
12288 16385
0
4096
42
131313`

func main() {
	var input io.Reader = os.Stdin
	if flag.Parse(); *update {
		input = strings.NewReader(testCases)
	}
	scan := bufio.NewScanner(input)
	scan.Split(bufio.ScanWords)
	for scan.Scan() {
		var (
			m, _ = strconv.Atoi(read(scan))
			q, _ = strconv.Atoi(read(scan))
			r, _ = strconv.Atoi(read(scan))
		)

		memory := make(map[uint64]uint64, m)
		for i := 0; i < m; i++ {
			paddr, _ := strconv.ParseUint(read(scan), 10, 64)
			value, _ := strconv.ParseUint(read(scan), 10, 64)
			memory[paddr] = value
		}

		logical := make([]uint64, 0, q)
		for i := 0; i < q; i++ {
			addr, _ := strconv.ParseUint(read(scan), 10, 64)
			logical = append(logical, addr)
		}

		for i, addr := range logical {
			var paddr uint64
			pml4e, pdpte, pde, pte, offset := split(addr)

			paddr = uint64(r)
			for _, row := range []uint16{pml4e, pdpte, pde, pte} {
				paddr += uint64(row * 8)
				val, ok := memory[paddr]
				if !ok {
					paddr = 0
					break
				}
				paddr, ok = chain(val)
				if !ok {
					paddr = 0
					break
				}
			}
			if paddr != 0 {
				paddr += uint64(offset)
			}
			logical[i] = paddr
		}

		for _, addr := range logical {
			if addr == 0 {
				fmt.Println("fault")
				continue
			}
			fmt.Println(addr)
		}
	}
}

func read(scan *bufio.Scanner) string {
	defer scan.Scan()
	return scan.Text()
}

func split(addr uint64) (pml4e, pdpte, pde, pte, offset uint16) {
	pml4e = uint16((addr >> 39) & 511)
	pdpte = uint16((addr >> 30) & 511)
	pde = uint16((addr >> 21) & 511)
	pte = uint16((addr >> 12) & 511)
	offset = uint16(addr & 4095)
	return
}

func chain(addr uint64) (naddr uint64, valid bool) {
	if addr&1 == 1 {
		return addr - 1, true
	}
	return 0, false
}
