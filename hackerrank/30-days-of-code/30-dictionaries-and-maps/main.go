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
	scanner.Split(bufio.ScanLines)
	_ = scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	phonebook := make(map[string]string)
	for i := 0; i < n && scanner.Scan(); i++ {
		entry := strings.Split(scanner.Text(), " ")
		name, phone := entry[0], entry[1]
		phonebook[name] = phone
	}
	for scanner.Scan() {
		name := scanner.Text()
		phone, found := phonebook[name]
		if !found {
			fmt.Println("Not found")
			continue
		}
		fmt.Printf("%s=%s\n", name, phone)
	}
}
