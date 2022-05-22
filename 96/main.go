package main

import "fmt"

func main() {
	fmt.Println(numTrees(3))
}

var m = map[int]int{}

func numTrees(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	if v, ok := m[n]; ok {
		return v
	}
	count := 0
	for i := 1; i <= n; i++ {
		leftNum := numTrees(i - 1)
		rightNum := numTrees(n - i)
		count += leftNum * rightNum
	}
	m[n] = count
	return count
}
