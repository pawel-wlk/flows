package main

import (
	"fmt"
	"./graphs"
)

func main() {
	fmt.Println("k, flow, paths, time")
	for k := 1; k <= 16; k++  {
		for i:=0; i<100; i++ {
			fmt.Printf("%d,", k)
			graph := graphs.NewHyperCube(k)
			graphs.Karp(graph, 0, (1 << uint(k))-1, k)
		}
	}
}
