package main

import "fmt"

/*
假设有一个很长的花坛，一部分地块种植了花，另一部分却没有。可是，花不能种植在相邻的地块上，它们会争夺水源，两者都会死去。

给你一个整数数组flowerbed 表示花坛，由若干 0 和 1 组成，其中 0 表示没种植花，1 表示种植了花。另有一个数n ，能否在不打破种植规则的情况下种入n朵花？能则返回 true ，不能则返回 false。

输入：flowerbed = [1,0,0,0,1], n = 1
输出：true

输入：flowerbed = [1,0,0,0,1], n = 2
输出：false
1 <= flowerbed.length <= 2 * 104
flowerbed[i] 为 0 或 1
flowerbed 中不存在相邻的两朵花
0 <= n <= flowerbed.length
 */

func canPlaceFlowers(flowerbed []int, n int) bool {
	var num int
	numFlower := len(flowerbed)
	for i := 0; i < numFlower; i++ {
		if flowerbed[i] == 1 {
			continue
		}
		before := i - 1
		after := i + 1
		isCanPlaceFlower := true
		if before >= 0 && flowerbed[before] == 1 {
			isCanPlaceFlower = false
			continue
		}
		if after < numFlower && flowerbed[after] == 1 {
			isCanPlaceFlower = false
			continue
		}
		if isCanPlaceFlower {
			num++
			flowerbed[i] = 1
		}
	}
	if num >= n {
		return true
	} else {
		return false
	}
}

func main()  {
	fmt.Println(canPlaceFlowers([]int{1,0,0,0,0,1}, 2))
	//fmt.Println(canPlaceFlowers([]int{1,0,0,0,1,0,0}, 3))
}