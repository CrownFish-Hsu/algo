package main

import (
	"basic/algo/utils"
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// 有序数组中找>=num的最左位置
func main() {
	N := 10
	V := 100
	rounds := 100
	rn := rand.New(rand.NewSource(time.Now().UnixNano()))

	fmt.Println("测试开始")
	for i := 0; i < rounds; i++ {
		n := int(rn.Float32() * float32(N))
		arr := utils.RandomArray(n, V)
		sort.Ints(arr)

		num := int(rn.Float32() * float32(V))
		rp := rightLeft(arr, num)
		fl := findLeft(arr, num)
		if rp != fl {
			fmt.Printf("第%d轮测试失败, 寻找元素%d, 暴力解为%d, 最优解为%d\n", i, num, rp, fl)
			fmt.Println(arr)
		}
	}
	fmt.Println("测试结束")
}

func findLeft(arr []int, num int) int {
	l := 0
	r := len(arr) - 1
	pos := -1
	for l <= r {
		mid := l + (r-l)/2
		if arr[mid] < num {
			l = mid + 1
		} else if arr[mid] >= num {
			r = mid - 1
			pos = mid
		}
	}

	return pos
}

func rightLeft(arr []int, num int) int {
	for k, v := range arr {
		if v >= num {
			return k
		}
	}

	return -1
}
