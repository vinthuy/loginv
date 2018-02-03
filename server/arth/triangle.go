package main

import (
	"fmt"
	"math"
)

// 7
// 3 8
// 8 1 0
// 2 7 4 4
// 4 5 2 6 5
// 1. max(r,j) = max(max(r+1,j),max(r+1,j+1)) + d(r,j)
// 2. max(r,j) = d(r,j)
func main() {
	recycleArth()
}

var (
	N int = 5
	d     = [5][5]int{{7}, {3, 8}, {8, 1, 0}, {2, 7, 4, 4}, {4, 5, 2, 6, 5}}
)

func recycleArth() {
	// fmt.Println(max(0, 0))
	// max2()
	max3()
}

var cache [5][5]int

func max(r int, j int) int {
	var result int
	if cache[r][j] != 0 {
		fmt.Println("hit cache", r, j)
		return cache[r][j]
	}
	if r == N-1 && j <= r {
		result = d[r][j]
	} else if r < N-1 && j <= r {
		a1 := float64(max(r+1, j))
		a2 := float64(max(r+1, j+1))
		result = int(math.Max(a1, a2)) + d[r][j]
	}
	cache[r][j] = result
	return result
}

func maxInt(r int, j int) int {
	if r >= j {
		return r
	}
	return j

}

// 7
// 3 8
// 8 1 0       20,12,10,0,0
// 2 7 4 4     7,12,10,10
// 4 5 2 6 5   4,5,2,6,5
func max2() {
	var max1 [5][5]int
	for r := N - 1; r >= 0; r-- {
		for j := 0; j <= r; j++ {
			if r == N-1 {
				max1[r][j] = d[r][j]
			} else {
				max1[r][j] = d[r][j] + maxInt(max1[r+1][j+1], max1[r+1][j])
			}

		}
	}
	fmt.Println(max1[0][0])
}

func max3() {
	var max1 [5]int
	for r := N - 1; r >= 0; r-- {
		for j := 0; j <= r; j++ {
			if r == N-1 {
				max1[j] = d[r][j]
			} else {
				max1[j] = d[r][j] + maxInt(max1[j], max1[j+1])
			}
		}
		fmt.Println(max1)
	}
	fmt.Println(max1[0])

}
