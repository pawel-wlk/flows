package graphs

import (
	"math"
	"fmt"
	"time"

	"../binary"
)

var paths uint = 0


func bfs(graph [][]int, source int, end int, k int, pi []int) bool {
	paths++
	visited := make([]bool, binary.PowOf2(k))

	for vertex := range graph {
		visited[vertex] = false
		pi[vertex] = -1
	}

	visited[source] = true
	pi[source] = -1

	queue := make([]int, 0)
	queue = append(queue, source)

	for len(queue) != 0 {
		u := queue[0]
		queue = queue[1:]

		for i := 0; i<k; i++ {
			v := u | binary.PowOf2(i)
			if !visited[v] && graph[u][i] > 0 {
				visited[v] = true
				pi[v] = u

				queue = append(queue, v)
			}
		}
	}

	return visited[end]
}

func Karp(graph [][]int, source int, end int, k int) {
	startTime := time.Now()
	paths = 0
	residualGraph := make([][]int, binary.PowOf2(k))

	for u:=0; u<binary.PowOf2(k); u++ {
		residualGraph[u] = make([]int, k)
		for v:=0; v<k; v++ {
			residualGraph[u][v] = graph[u][v]
		}
	}


	path := make([]int, binary.PowOf2(k))
	maxFlow := 0

	for bfs(residualGraph, source, end, k, path) {
		// find minimum weight on path
		pathFlow := math.MaxInt32
		for v := end; v != source; v = path[v] {
			edge := residualGraph[path[v]][binary.Log2(path[v]^v)]
			if edge < pathFlow {
				pathFlow = edge
			}
		}
		for v := end; v != source; v = path[v] {
			residualGraph[path[v]][binary.Log2(path[v]^v)] -= pathFlow
			residualGraph[v][binary.Log2(path[v]^v)] += pathFlow
		}

		maxFlow += pathFlow
	}

	fmt.Printf("%d, %d, %d\n", maxFlow, paths, time.Since(startTime))
}
