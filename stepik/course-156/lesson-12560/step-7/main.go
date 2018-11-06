package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	_ = scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	fruit := make([]weight, 0, n)
	for i := 0; i < n && scanner.Scan(); i++ {
		w, _ := strconv.Atoi(scanner.Text())
		fruit = append(fruit, weight(w))
	}
	_ = scanner.Scan()
	k, _ := strconv.Atoi(scanner.Text())
	fmt.Println(greedy{weight(k), new(basket).with(n).insert(fruit...)}.eat())
}

type weight int

type basket struct {
	fruit []weight
	head  int
}

func (basket *basket) with(capacity int) *basket {
	basket.fruit, basket.head = make([]weight, 0, capacity), 0
	return basket
}

func (basket *basket) insert(fruit ...weight) *basket {
	// assert(len(fruit) <= cap(basket.fruit) - len(basket.fruit))
	basket.fruit = append(basket.fruit, fruit...)
	return basket.sort()
}

func (basket *basket) empty() bool {
	return len(basket.fruit) == 0 || basket.fruit[0] == 0
}

func (basket *basket) sort() *basket {
	// build min heap
	for i := len(basket.fruit) / 2; i >= 0; i-- {
		siftDown(basket.fruit, i)
	}
	// sort by min heap
	process := basket.fruit
	for last := len(basket.fruit) - 1; last > 0; last-- {
		process[0], process[last] = process[last], process[0]
		process = process[:last]
		siftDown(process, 0)
	}
	// remove stumps
	for i := len(basket.fruit) - 1; i >= 0; i-- {
		if basket.fruit[i] > 0 {
			break
		}
		basket.fruit = basket.fruit[:i]
	}
	return basket
}

func (basket *basket) take() weight {
	if basket.head < len(basket.fruit) {
		fruit := basket.fruit[basket.head]
		basket.head++
		return fruit
	}
	return 0
}

func (basket *basket) back(fruit ...weight) {
	// assert(basket.head - len(fruit) == 0)
	basket.head = 0
	copy(basket.fruit, fruit)
	basket.sort()
}

type greedy struct {
	force  weight
	basket interface {
		empty() bool
		take() weight
		back(fruit ...weight)
	}
}

func (greedy greedy) eat() int {
	var steps int
	for !greedy.basket.empty() {
		steps++
		portion, w := make([]weight, 0, 4), weight(0)
		for w < greedy.force {
			// assert(f <= greedy.force)
			f := greedy.basket.take()
			if f == 0 {
				break
			}
			w += f
			portion = append(portion, f)
		}
		w = greedy.force
		for i := range portion {
			w -= portion[i]
			if w < 0 {
				break
			}
			if portion[i] == 1 {
				portion[i] = 0
				continue
			}
			portion[i] /= 2
		}
		greedy.basket.back(portion...)
	}
	return steps
}

func siftDown(data []weight, i int) {
	min := i
	l := left(i)
	if l < len(data) && data[l] < data[min] {
		min = l
	}
	r := right(i)
	if r < len(data) && data[r] < data[min] {
		min = r
	}
	if i != min {
		data[i], data[min] = data[min], data[i]
		siftDown(data, min)
	}
}

func left(i int) int {
	return 2*i + 1
}

func right(i int) int {
	return 2 * (i + 1)
}
