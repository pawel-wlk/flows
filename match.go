package main

import (
	"math/rand"
	"fmt"
	"time"
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

func glpk(graph [][]int)  {
	fmt.Print("param n, integer, >= 2;\n" +
			"/* number of nodes */\n" +
			"\n" +
			"set V, default {1..n};\n" +
			"/* set of nodes */\n" +
			"\n" +
			"set E, within V cross V;\n" +
			"/* set of arcs */\n" +
			"\n" +
			"param a{(i,j) in E}, > 0;\n" +
			"/* a[i,j] is capacity of arc (i,j) */\n" +
			"\n" +
			"param s, symbolic, in V, default 1;\n" +
			"/* source node */\n" +
			"\n" +
			"param t, symbolic, in V, != s, default n;\n" +
			"/* sink node */\n" +
			"\n" +
			"var x{(i,j) in E}, >= 0, <= a[i,j];\n" +
			"/* x[i,j] is elementary flow through arc (i,j) to be found */\n" +
			"\n" +
			"var flow, >= 0;\n" +
			"/* total flow from s to t */\n" +
			"\n" +
			"s.t. node{i in V}:\n" +
			"/* node[i] is conservation constraint for node i */\n" +
			"\n" +
			"	sum{(j,i) in E} x[j,i] + (if i = s then flow)\n" +
			"	/* summary flow into node i through all ingoing arcs */\n" +
			"\n" +
			"	= /* must be equal to */\n" +
			"\n" +
			"	sum{(i,j) in E} x[i,j] + (if i = t then flow);\n" +
			"	/* summary flow from node i through all outgoing arcs */\n" +
			"\n" +
			"maximize obj: flow;\n" +
			"/* objective is to maximize the total flow through the network */\n" +
			"\n" +
			"solve;\n" +
			"\n" +
			"printf{1..56} \"=\"; printf \"\\n\";\n" +
			"printf \"Maximum flow from node %s to node %s is %g\\n\\n\", s, t, flow;\n" +
			"\n" +
			"data;\n\n");
	fmt.Println();
	fmt.Printf("param n := %d;\n", len(graph) + 2)
	fmt.Println();
	fmt.Println("param : E : a :=");
	for j, neighbours := range graph {
		for _, i := range neighbours {
			fmt.Printf("  %d %d %d\n", j + 1, i + 1, 1);
		}
	}
	fmt.Println(";\n");
	fmt.Println("end;");
}

func main() {
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
