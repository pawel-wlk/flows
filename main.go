package main

import (
	"fmt"
	"./graphs"
)

func main() {
	k := 16
	graph := graphs.NewHyperCube(k)
	fmt.Println(graphs.Karp(graph, 0, (1 << uint(k))-1, k))
}
