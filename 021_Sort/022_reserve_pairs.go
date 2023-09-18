package main

// 测试链接 : https://leetcode.cn/problems/reverse-pairs/

// 翻转对数量
// 类似于小和问题
func reversePairs(nums []int) int {
	pairNum := 0
	if len(nums) > 0 {
		pairNum = reserve(nums, 0, len(nums)-1)
	}

	return pairNum
}

func reserve(nums []int, l int, r int) int {
	if l == r {
		return 0
	}

	mid := l + (r-l)/2
	return reserve(nums, l, mid) + reserve(nums, mid+1, r) + mergeReserve(nums, l, r, mid)
}

func mergeReserve(nums []int, l int, r int, mid int) int {
	pairNum := 0
	//归并前 从 ls->m, m+1->rs 依次比较元素
	ls := l
	rs := mid + 1
	for ; ls <= mid; ls++ {
		for rs <= r && nums[ls] > nums[rs]*2 {
			rs++
		}

		//当ls滑动，统计rs的个数，就是符合的答案
		pairNum += rs - (mid + 1)
	}

	//归并排序代码
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

	return pairNum
}
