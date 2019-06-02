package main

import (
	"fmt"
	"os"
	"./graphs"
)

func main() {
	if len(os.Args) >= 2 && os.Args[1] == "--glpk" {
		graphs.Glpk(graphs.NewHyperCube(4))
		return
	}
	fmt.Println("k, flow, paths, time")
	for k := 1; k <= 16; k++  {
		for i:=0; i<100; i++ {
			fmt.Printf("%d,", k)
			graph := graphs.NewHyperCube(k)
			graphs.Karp(graph, 0, (1 << uint(k))-1, k)
		}
	}

	//k:=4
	//graph := graphs.NewHyperCube(k)
	//graphs.Karp(graph, 0, (1 << uint(k))-1, k)
}
