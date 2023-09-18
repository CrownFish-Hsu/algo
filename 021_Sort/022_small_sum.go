// 测试链接 : https://www.nowcoder.com/practice/edfe05a1d45c4ea89101d936cac32469
package main

import (
	"fmt"
)

func main() {
	var arr []int64
	var n, tmp int64
	fmt.Scan(&n)
	for i := int64(0); i < n; i++ {
		fmt.Scan(&tmp)
		arr = append(arr, tmp)
	}
	fmt.Println(SmallSum(arr))
}

func SmallSum(nums []int64) int64 {
	if len(nums) <= 1 {
		return 0
	}

	tmp := make([]int64, len(nums))
	return mergeSort(nums, tmp, 0, len(nums)-1)
}

func mergeSort(nums, tmp []int64, left, right int) int64 {
	if left >= right {
		return 0
	}

	mid := left + (right-left)/2
	sum := mergeSort(nums, tmp, left, mid) + mergeSort(nums, tmp, mid+1, right)
	sum += mergeSmallSum(nums, tmp, left, mid, right)
	copy(nums[left:right+1], tmp[left:right+1])
	return sum
}

func mergeSmallSum(nums, tmp []int64, left, mid, right int) int64 {
	i, j, k := left, mid+1, left
	sum := int64(0)

	for i <= mid && j <= right {
		if nums[i] <= nums[j] {
			sum += nums[i] * int64(right-j+1)
			i++
		} else {
			j++
		}
	}

	i, j = left, mid+1

	for k <= right {
		if i <= mid && (j > right || nums[i] <= nums[j]) {
			tmp[k] = nums[i]
			i++
		} else {
			tmp[k] = nums[j]
			j++
		}
		k++
	}

	return sum
}
