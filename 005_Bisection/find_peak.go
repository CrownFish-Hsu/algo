package main

import (
	"fmt"
)

// 峰值元素是指其值严格大于左右相邻值的元素
// 给你一个整数数组 nums，已知任何两个相邻的值都不相等
// 找到峰值元素并返回其索引
// 数组可能包含多个峰值，在这种情况下，返回 任何一个峰值 所在位置即可。
// 假设 nums[-1] = nums[n] = 无穷小
// 时间复杂度为 O(log n) 的算法

func main() {
	arr := []int{2, 3, 5, 3, 10, 6, 9, 8, 1, 0}
	peak := findPeak(arr)
	fmt.Println(arr)
	fmt.Printf("峰的位置为%d, 峰值为%d", peak, arr[peak])
}

func findPeak(arr []int) int {
	n := len(arr)
	if n == 0 {
		return -1
	}

	//单个元素 就是峰值
	if n == 1 {
		return 0
	}

	//至少2个元素
	//单独验证0位置 是否峰值
	if arr[0] > arr[1] {
		return 0
	}

	//单独验证n-1位置 是否峰值
	if arr[n-1] > arr[n-2] {
		return n - 1
	}

	//若0和n-1都不是峰值点，则1~n-2位置上至少存在1个峰值
	l := 1
	r := n - 2
	pos := -1
	for l <= r {
		mid := l + (r-l)/2
		if arr[mid-1] > arr[mid] {
			r = mid - 1
		} else if arr[mid] < arr[mid+1] {
			l = mid + 1
		} else {
			pos = mid
			break
		}
	}

	return pos
}
