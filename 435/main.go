package main

import (
	"fmt"
	"sort"
)

/*
给定一个区间的集合，找到需要移除区间的最小数量，使剩余区间互不重叠。

注意:

可以认为区间的终点总是大于它的起点。
区间 [1,2] 和 [2,3] 的边界相互“接触”，但没有相互重叠。

case1:
输入: [ [1,2], [2,3], [3,4], [1,3] ]
输出: 1
解释: 移除 [1,3] 后，剩下的区间没有重叠。

case2:
输入: [ [1,2], [1,2], [1,2] ]
输出: 2
解释: 你需要移除两个 [1,2] 来使剩下的区间没有重叠。

case3:
输入: [ [1,2], [2,3] ]
输出: 0
解释: 你不需要移除任何区间，因为它们已经是无重叠的了。
 */
func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	currentMaxEnd := intervals[0][1]
	remove_num := 0
	for i := 1; i < len(intervals); i++ {
		maxEnd := intervals[i][0]
		if maxEnd < currentMaxEnd {
			remove_num++
		} else {
			currentMaxEnd = intervals[i][1]
		}
	}
	return remove_num
}

func main()  {
	a := make([][]int, 0)
	//a = append(a, []int{1,2}, []int{2,3}, []int{3,4}, []int{4,5})
	fmt.Println(eraseOverlapIntervals(a))
}