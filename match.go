package main

import (
	"math/rand"
	"fmt"
	"time"
	"os"
	"strconv"
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

func glpk(graph [][]int, k int, n int)  {
	fmt.Printf("param n, integer, >= 2;\n\nset V, default {1..n};\n\nset E, within V cross V;\n\nparam a{(i,j) in E}, > 0;\n\nparam s, symbolic, in V, default 1;\n\nparam t, symbolic, in V, != s, default n;\n\nvar x{(i,j) in E}, >= 0, <= a[i,j];\n\nvar flow, >= 0;\n\ns.t. node{i in V}:\n\n	 sum{(j,i) in E} x[j,i] + (if i = s then flow)\n\n	 =\n\n	 sum{(i,j) in E} x[i,j] + (if i = t then flow);\n\nmaximize obj: flow;\n\nsolve;\n\nprintf{1..56} \"=\"; printf \"\\n\";\nprintf \"Maximum flow from node %%s to node %%s is %%g\\n\\n\", s, t, flow;\n\ndata;\n");
	fmt.Printf("param n := %d;\n\n", (1 << uint(k+1)) + 2)
	fmt.Printf("param : E : a :=\n")
	for i := 0; i < 1 << uint(k); i++ {
		fmt.Printf("\t1 %d 1\n", i+2)
	}
	for i := 0; i < 1 << uint(k); i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("\t%d %d 1\n", i+2, (1 << uint(k)) + graph[i][j] + 2);
		}
	}
	for i := 0; i < 1 << uint(k); i++ {
		fmt.Printf("\t%d %d 1\n", (1 << uint(k)) + i + 2, (1 << uint(k+1)) + 2);
	}
	fmt.Printf(";\nend;\n");
}

func main() {
	if len(os.Args) >= 4 && os.Args[1] == "--glpk" {
		k, _ := strconv.Atoi(os.Args[2])
		i, _ := strconv.Atoi(os.Args[3])
		glpk(generate(k, i), k, i)
		return
	}
	reps := 100
	fmt.Println("k, i, matches, time")
	for k := 3; k<=10; k++ {
		for i := 1; i<=k; i++ {
			for rep := 0; rep<reps; rep++ {
				g := generate(k,i)
				startTime := time.Now()
				matches := maxBPM(g)
				fmt.Printf("%d, %d, %d, %d\n", k, i, matches, time.Since(startTime))
			}
		}
	}


	//k := 10
	//i := 10

	//g := generate(k,i)
	//glpk(g)
}
