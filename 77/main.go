package main

import "fmt"

var res [][]int

func combine(n int, k int) [][]int {
	res = make([][]int, 0)
	path := make([]int, 0)
	backtrack(n, 1, k, path)
	return res
}

func backtrack(n, i, k int, path []int) {
	if len(path) == k {
		cpath := make([]int, k)
		copy(cpath, path)
		res = append(res, cpath)
		return
	}

	for j := i; j <= n; j++ {
		path = append(path, j)
		fmt.Printf("递归之前:%v\n", path)
		backtrack(n, j+1, k, path)
		path = path[:len(path)-1]
		fmt.Printf("递归之前:%v\n", path)
	}
}

func main() {
	combine(5, 3)
}
