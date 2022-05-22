package main

import "fmt"

func binarySearch(array []int, ele int) int {
	// 二分查找，适用于在有序数组中查找某个元素，时间复杂度为O(log2N,以2为底的N)
	if len(array) == 0 {
		return -1
	}
	left, right := 0, len(array)-1
	for left <= right { // 比较条件必须是<=
		// mid := (left + right) / 2  当left和right比较大的时候，容易溢出
		mid := left + (right-left)/2 // 或者 mid := left + ((right - left)>>1)
		if array[mid] < ele {
			left = mid + 1
		} else if array[mid] > ele {
			right = mid - 1
		} else {
			return mid
		}
	}
	return -1
}

func binarySearchFirst(array []int, ele int) int {
	// 寻找有序数组中第一个出现的值
	if len(array) == 0 {
		return -1
	}
	left, right := 0, len(array)-1
	for left <= right {
		mid := left + (right-left)/2
		if array[mid] == ele {
			if mid == 0 || array[mid-1] != ele {
				return mid
			} else {
				right = mid - 1
			}
		} else if array[mid] > ele {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

func binarySearchLast(array []int, ele int) int {
	// 寻找有序数组中最后一个出现的值
	if len(array) == 0 {
		return -1
	}
	left, right := 0, len(array)-1
	for left <= right {
		mid := left + (right-left)/2
		if array[mid] == ele {
			if mid == len(array)-1 || array[mid+1] != ele {
				return mid
			} else {
				left = mid + 1
			}
		} else if array[mid] > ele {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

func binaryFirstGe(array []int, ele int) int {
	// 查找第一个大于等于某个值的下标
	if len(array) == 0 {
		return 0
	}
	left, right := 0, len(array)-1
	for left <= right {
		mid := left + (right-left)/2
		if array[mid] >= ele {
			if mid == 0 || array[mid-1] < ele {
				return mid
			} else {
				right = mid - 1
			}
		} else {
			left = mid + 1
		}
	}
	return -1
}

func binaryLastLe(array []int, ele int) int {
	// 查找最后一个小于等于某个值的下标
	if len(array) == 0 {
		return 0
	}
	left, right := 0, len(array)-1
	for left <= right {
		mid := left + (right-left)/2
		if array[mid] <= ele {
			if mid == len(array)-1 || array[mid+1] > ele {
				return mid
			} else {
				left = mid + 1
			}
		} else {
			right = mid - 1
		}
	}
	return -1
}

func binarySearch2(array []int, ele int) int {
	// 二分查找递归方法
	return binaryRecursion(array, ele, 0, len(array)-1)
}

func binaryRecursion(array []int, ele, left, right int) int {
	if left > right {
		return -1
	}
	mid := left + (right-left)/2
	if array[mid] == ele {
		return mid
	} else if array[mid] > ele {
		binaryRecursion(array, ele, left, mid-1)
	} else {
		binaryRecursion(array, ele, mid+1, right)
	}
	return -1
}

func main() {
	a := []int{1, 2, 3, 3, 3, 3, 7, 7, 7, 14, 19}
	fmt.Println(binaryLastLe(a, 8))
}
