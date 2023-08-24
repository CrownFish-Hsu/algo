package utils

import (
	"math/rand"
	"time"
)

// RandomArray
// 随机生成元素为N，元素大小在1-V范围内的数组
func RandomArray(n int, v int) []int {
	arr := make([]int, n)
	rn := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < n; i++ {
		arr[i] = int(rn.Float32()*float32(v)) + 1
	}

	return arr
}
