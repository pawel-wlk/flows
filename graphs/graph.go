package graphs

import (
	"math/rand"
	"fmt"

	"../binary"
)

func NewHyperCube(k int) [][]int {
	cube := make([][]int, binary.PowOf2(k))
	for i:=0; i<binary.PowOf2(k); i++ {
		cube[i] = make([]int, k)
		uWeight := binary.HammingWeight(i)
		l := max(uWeight+1, k-uWeight)
		for j:=0; j<k; j++ {
			if (1 << uint(j)) & i == 0 {
				cube[i][j] = rand.Intn(binary.PowOf2(l)) + 1
			} else {
				cube[i][j] = 0
			}
		}
	}

	return cube
}

func max(values ...int) int {
	var m int
	for i, e := range values {
		if i==0 || e > m {
			m = e }
	}

	return m
}

func Glpk(graph [][]int) {
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
			"data;\n")
	fmt.Println()
	fmt.Printf("param n := %d;\n", len(graph))
	fmt.Println()
	fmt.Println("param : E : a :=")
	for i, neighbours := range graph {
		for j := range neighbours {
			if graph[i][j] != 0 {
				s := i + (1 << uint(j)) + 1
				fmt.Printf("  %d %d %d\n", i + 1, s, graph[i][j])
			}
		}
	}
	fmt.Println(";\n")
	fmt.Println("end;");
}
