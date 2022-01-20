package main

import "fmt"

/*
n 个孩子站成一排。给你一个整数数组 ratings 表示每个孩子的评分。

你需要按照以下要求，给这些孩子分发糖果：

每个孩子至少分配到 1 个糖果。
相邻两个孩子评分更高的孩子会获得更多的糖果。
请你给每个孩子分发糖果，计算并返回需要准备的 最少糖果数目 。
case 1:
输入：ratings = [1,0,2]
输出：5
解释：你可以分别给第一个、第二个、第三个孩子分发 2、1、2 颗糖果。
case 2:
输入：ratings = [1,2,2]
输出：4
解释：你可以分别给第一个、第二个、第三个孩子分发 1、2、1 颗糖果。
     第三个孩子只得到 1 颗糖果，这满足题面中的两个条件。
 */

func candy(ratings []int) int {
	num := len(ratings)
	candyList := make([]int, len(ratings))
	candyList[0] = 1
	for i := 1; i < num; i++ {
		if ratings[i] > ratings[i - 1] {
			candyList[i] = candyList[i - 1] + 1
		}else {
			candyList[i] = 1
		}
	}
	for j := num -2; j >= 0; j-- {
		if ratings[j] > ratings[j+1] {
			if candyList[j+1] + 1 > candyList[j] {
				candyList[j] = candyList[j+1] + 1
			}
		}
	}
	sum := 0
	for _, v := range candyList {
		sum += v
	}
	return sum
}

func main()  {
	fmt.Println(candy([]int{1, 0, 2}))
	fmt.Println(candy([]int{1, 2, 2}))
}
