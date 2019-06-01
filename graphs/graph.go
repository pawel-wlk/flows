package graphs

import (
	"math/rand"

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
