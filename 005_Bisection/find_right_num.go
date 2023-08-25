package main

import (
	"basic/algo/utils"
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// 有序数组中找<=num的最右位置
func main() {
	N := 6
	V := 10
	rounds := 10
	rn := rand.New(rand.NewSource(time.Now().UnixNano()))

	fmt.Println("测试开始")
	for i := 0; i < rounds; i++ {
		n := int(rn.Float32() * float32(N))
		arr := utils.RandomArray(n, V)
		sort.Ints(arr)

		num := int(rn.Float32() * float32(V))
		rr := rightRight(arr, num)
		fr := findRight(arr, num)
		if rr != fr {
			fmt.Printf("第%d轮测试失败, 寻找元素%d, 暴力解为%d, 最优解为%d\n", i, num, rr, fr)
			fmt.Println(arr)
		}
	}
	fmt.Println("测试结束")
}

func findRight(arr []int, num int) int {
	l := 0
	r := len(arr) - 1
	pos := -1
	for l <= r {
		mid := l + (r-l)/2
		if arr[mid] <= num {
			pos = mid
			l = mid + 1
		} else if arr[mid] > num {
			r = mid - 1
		}
	}

	return pos
}

// 从右往左找到第一个 不比目标值大的，就是位置
func rightRight(arr []int, num int) int {
	pos := -1
	for k := len(arr) - 1; k >= 0; k-- {
		if arr[k] <= num {
			return k
		}
	}

	return pos
}
