package main

import (
	"math/rand"
	"fmt"
)

func search(arr []int, searched int) bool {
	for _, el := range arr {
		if el == searched {
			return true
		}
	}
	return false
}

func generate(size int, neighbours int) [][]int {
	g := make([][]int, 1<<uint(size))

	for i:=0; i < 1<<uint(size); i++ {
		g[i] = make([]int, neighbours)

		for j:=0; j<neighbours; j++ {
			n := rand.Intn(1<<uint(size))
			for search(g[i], n) {
				n = rand.Intn(1<<uint(size))
			}

			g[i][j] = n
		}
	}

	return g
}

func bpm(bpGraph [][]int, u int, seen []bool, matchR []int) bool {
	for v := range bpGraph {
		if search(bpGraph[u], v) && !seen[v] {
			seen[v] = true

			if matchR[v] < 0 || bpm(bpGraph, matchR[v], seen, matchR) {
				matchR[v] = u
				return true
			}
		}
	}

	return false
}

func maxBPM(bpGraph [][]int) int {
	matchR := make([]int, len(bpGraph))

	for i := range matchR {
		matchR[i] = -1
	}

	result := 0

	for u := range bpGraph {
		seen := make([]bool, len(bpGraph))
		for i := range seen {
			seen[i] = false
		}

		if bpm(bpGraph, u, seen, matchR) {
			result++
		}
	}

	return result
}

func main() {
	reps := 100
	fmt.Println("k, i, matches")
	for k := 3; k<=10; k++ {
		for i := 1; i<=k; i++ {
			for rep := 0; rep<reps; rep++ {
				g := generate(k,i)
				fmt.Printf("%d, %d, %d\n", k, i, maxBPM(g))
			}
		}
	}
}
