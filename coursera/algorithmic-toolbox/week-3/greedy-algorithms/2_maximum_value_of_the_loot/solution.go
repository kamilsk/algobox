package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Loot struct{ value, weight int64 }

func (loot Loot) Score() float64 { return float64(loot.value / loot.weight) }

type Loots []Loot

func (loots Loots) Len() int           { return len(loots) }
func (loots Loots) Less(i, j int) bool { return loots[i].Score() < loots[j].Score() }
func (loots Loots) Swap(i, j int)      { loots[i], loots[j] = loots[j], loots[i] }

func optimalValue(capacity int64, loots []Loot) float64 {
	var optimal float64

	sort.Sort(Loots(loots))
	for capacity > 0 && len(loots) > 0 {
		var best Loot
		best, loots = loots[0], loots[1:]
		weight := best.weight
		if weight > capacity {
			weight = capacity
		}
		optimal += best.Score() * float64(weight)
		capacity -= weight
	}

	return optimal
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	_ = scanner.Scan()
	n, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	_ = scanner.Scan()
	capacity, _ := strconv.ParseInt(scanner.Text(), 10, 64)

	loots := make([]Loot, 0, capacity)
	var i int64
	for ; i < n && scanner.Scan(); n++ {
		value, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		_ = scanner.Scan()
		weight, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		loots = append(loots, Loot{value, weight})
	}
	safe(fmt.Fprintf(os.Stdout, "%.4f\n", optimalValue(capacity, loots)))
}

func safe(interface{}, error) {}
