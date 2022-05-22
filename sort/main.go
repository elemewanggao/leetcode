package main

import (
	"fmt"
	"math/rand"
	"time"
)

func bubbleSort(array []int) {
	// 冒泡排序，交换相邻两个数，不符合就交换
	// 时间复杂度O(N^2), 空间复杂度为O(1)【原地排序】, 稳定的排序
	if len(array) <= 1 {
		return
	}
	for i := len(array) - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
}

func bubbleSort2(array []int) {
	// 冒泡排序，交换相邻两个数，不符合就交换，先交换0～N，选出最大的数放在N位置，然后选出0～N-1的位置，放在N-1位置
	// 时间复杂度O(N^2), 空间复杂度为O(1)【原地排序】, 稳定的排序
	if len(array) <= 1 {
		return
	}
	for i := len(array) - 1; i >= 0; i-- {
		swapFlag := false // 是否交换的标志，如果没有发生交换，说明数据已经有序了，可以结束后面的循环
		for j := 0; j < i; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
				swapFlag = true
			}
		}
		if !swapFlag {
			return
		}
	}
}

func selectionSort(array []int) {
	// 选择排序，每次选择当前未排序数组中最小值加入到已排序数组后面
	// 时间复杂度 O(N^2), 空间复杂度O(1)[原地排序], 非稳定排序 5，5，2，这样2和第一个5交换后就不是稳定的
	if len(array) <= 1 {
		return
	}
	for i := 0; i < len(array)-1; i++ {
		minIdx := i
		for j := i + 1; j < len(array); j++ {
			if array[j] < array[minIdx] {
				minIdx = j
			}
		}
		if i != minIdx {
			array[i], array[minIdx] = array[minIdx], array[i]
		}

	}
}

func insertionSort(array []int) {
	// 插入排序，在数组左边维护一个有序数列，刚开始是索引0的值，右边是无序数列，将无序数列中的每一个数插入到有序数列中（通过交换方式）
	if len(array) <= 1 {
		return
	}
	for i := 1; i < len(array); i++ {
		for j := i - 1; j >= 0 && array[j] > array[j+1]; j-- {
			array[j+1], array[j] = array[j], array[j+1]
		}
	}
}

func insertionSort2(array []int) {
	// 插入排序，在数组左边维护一个有序数列，刚开始是索引0的值，右边是无序数列，将无序数列中的每一个数插入到有序数列中（通过赋值方式）
	// 原地、稳定的排序
	if len(array) <= 1 {
		return
	}
	for i := 1; i < len(array); i++ {
		value := array[i]
		j := i - 1
		for ; j >= 0 && array[j] > value; j-- {
			array[j+1] = array[j]
		}
		array[j+1] = value
	}
}

func mergingSort(array []int) {
	// 归并排序，思想是将数组分成两段，每段是有序数组后再进行有序数组合并操作，时间复杂度为
	// O(NlogN),空间复杂度为O(N),非原地排序，是稳定排序
	if len(array) == 0 {
		return
	}
	mSort(array, 0, len(array)-1)
}

func mSort(array []int, left, right int) {
	if left == right {
		return
	}
	mid := left + (right-left)/2
	mSort(array, left, mid)
	mSort(array, mid+1, right)
	merge(array, left, mid, right)
}

func merge(array []int, left, mid, right int) {
	tempArray := make([]int, right-left+1)
	p1, p2, i := left, mid+1, 0
	for p1 <= mid && p2 <= right {
		if array[p1] <= array[p2] {
			tempArray[i] = array[p1]
			p1++
		} else {
			tempArray[i] = array[p2]
			p2++
		}
		i++
	}
	for p1 <= mid {
		tempArray[i] = array[p1]
		i++
		p1++
	}
	for p2 <= right {
		tempArray[i] = array[p2]
		i++
		p2++
	}
	for j := 0; j < len(tempArray); j++ {
		array[left+j] = tempArray[j]
	}
}

func quickSort(array []int) {
	// 快排，寻找一个中点，使得中点左边的数都小于中点，中点右边的数都大于中点，然后递归中点两边进行排序
	// 最差时间复杂度为O(N^2), 最好和平均时间复杂度O(NlogN),最差空间复杂度为O(N), 最好/平均为O(logN),递归栈使用。非稳定排序，6，8，6， 2， 5
	// 快排在数据大致有序的情况下时间复杂度会退化为O(N^2)
	if len(array) <= 1 {
		return
	}
	quickRecurse(array, 0, len(array)-1)
}

func quickRecurse(array []int, left, right int) {
	if left >= right {
		return
	}
	pivot := partition(array, left, right)
	quickRecurse(array, left, pivot-1)
	quickRecurse(array, pivot+1, right)
}

func partition(array []int, left, right int) int {
	p := array[right]
	i := left
	for j := left; j <= right; j++ {
		if array[j] < p {
			array[i], array[j] = array[j], array[i]
			i++
		}
	}
	array[i], array[right] = array[right], array[i]
	return i
}

func quickSort2(nums []int) []int {
	// 两个优化点
	// 1. 随机选取中间点使得中间点尽可能在中间，这样时间复杂度会为O(N*logN)
	// 2. 选取的中间点是一个中间段，中间段的数都是等于中间点的值，中间段左边小于中间点的值，中间段右边大于中间点的值
	if len(nums) <= 1 {
		return nums
	}
	quickRecurse2(nums, 0, len(nums)-1)
	return nums
}

func quickRecurse2(nums []int, left, right int) {
	if left >= right {
		return
	}
	// 选择[left, right)之间的随机数作为p
	rand.Seed(time.Now().Unix())
	r := rand.Intn(right-left) + left
	swap(nums, r, right)
	p := partition2(nums, left, right)
	quickRecurse2(nums, left, p[0]-1)
	quickRecurse2(nums, p[1]+1, right)
}

func partition2(nums []int, left, right int) []int {
	less := left - 1 // less为小于pivot区
	more := right    // more为大于pivot区
	pivot := nums[right]
	for i := left; i < more; { // i为当前处理位置
		if nums[i] < pivot { // 小于pivot右移less区，并交换，i++
			less++
			swap(nums, less, i)
			i++
		} else if nums[i] > pivot { // 大于pivot将more区左移，交换，i位置不变
			more--
			swap(nums, more, i)
		} else { // 等于pivot直接移动指针
			i++
		}
	}
	swap(nums, more, right) //将more区的第一个数和pivot交换
	return []int{less + 1, more}
}

func swap(nums []int, i, j int) {
	if i != j {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

func main() {
	array := []int{3, 6, 5, 4, 1, 2}
	quickSort2(array)
	fmt.Println(array)
}
