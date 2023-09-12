package main

// 空间复杂度为O(N) 需要借助一个额外数组
func sortArray(nums []int) []int {
	if len(nums) > 0 {
		mergeSort1(nums, 0, len(nums)-1)
	}

	return nums
}

/**
 * 递归版本：
 * 1.停止递归条件 l==r
 * 2.存在mid 继续递归 (l, mid) (mid+1, r)
 * 3.merge(l, m, r)
 */
func mergeSort1(nums []int, l int, r int) {
	if l == r {
		return
	}

	mid := l + (r-l)/2
	mergeSort1(nums, l, mid)
	mergeSort1(nums, mid+1, r)
	merge(nums, l, r, mid)
}

func merge(nums []int, l int, r int, mid int) {
	help := make([]int, r-l+1) // 修改数组长度
	p, i := l, 0
	q := mid + 1
	for p <= mid && q <= r {
		if nums[p] <= nums[q] {
			help[i] = nums[p]
			p++
		} else {
			help[i] = nums[q]
			q++
		}
		i++
	}

	for p <= mid {
		help[i] = nums[p]
		p++
		i++
	}

	for q <= r {
		help[i] = nums[q]
		q++
		i++
	}

	for s := l; s <= r; s++ {
		nums[s] = help[s-l] // 修正索引
	}
}
