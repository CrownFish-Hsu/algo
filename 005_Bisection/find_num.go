package main

import (
	"basic/algo/utils"
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	rn := rand.New(rand.NewSource(time.Now().UnixNano()))

	N := 10
	V := 10
	rounds := 100

	fmt.Println("测试开始")
	for i := 0; i < rounds; i++ {
		n := int(rn.Float32() * float32(N))
		arr := utils.RandomArray(n, V)
		sort.Ints(arr)
		num := int(rn.Float32() * float32(V))
		if right(arr, num) != find(arr, num) {
			fmt.Printf("第%d轮测试失败\n", i)
			fmt.Println(arr, num)
		}
	}
	fmt.Println("测试结束")
}

// 二分查找 有序数组
func find(arr []int, num int) bool {
	l := 0
	r := len(arr) - 1

	//!!! l <= r
	for l <= r {
		mid := l + (r-l)/2
		if arr[mid] < num {
			l = mid + 1
		} else if arr[mid] > num {
			r = mid - 1
		} else {
			return true
		}
	}

	return false
}

// 暴力解法
func right(arr []int, num int) bool {
	for _, v := range arr {
		if num == v {
			return true
		}
	}

	return false
}
