package main

import (
	"basic/algo/utils"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rn := rand.New(rand.NewSource(time.Now().UnixNano()))

	N := 10
	V := 100
	rounds := 10

	fmt.Println("测试开始")
	for i := 0; i < rounds; i++ {
		n := int(rn.Float32() * float32(N))
		arr := utils.RandomArray(n, V)
		fmt.Printf("第%d轮:\n", i)
		fmt.Println("排序前的数组为:", arr)
		insertSort(arr)
		fmt.Println("排序后的数组为:", arr)
		fmt.Println()
	}

	fmt.Println("测试结束")
}

func insertSort(arr []int) []int {
	length := len(arr)
	//当只有1个元素，不需要排序
	if length < 2 {
		return arr
	}

	//从第2个元素开始循环插入
	for i := 1; i < length; i++ {
		for j := i - 1; j >= 0 && arr[j] > arr[j+1]; j-- {
			arr[j+1], arr[j] = arr[j], arr[j+1]
		}
	}

	return arr
}
