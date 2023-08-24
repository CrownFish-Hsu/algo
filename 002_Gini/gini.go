package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"
)

/**
假设一开始有100个人，每个人都有100元
在每一轮都做如下的事情 :
每个人都必须拿出1元钱给除自己以外的其他人，给谁完全随机
如果某个人在这一轮的钱数为0，那么他可以不给，但是可以接收
发生很多很多轮之后，这100人的社会财富分布很均匀吗
**/

func main() {
	var round int = 500000
	var nums int = 100
	experiment(nums, round)
}

func experiment(nums int, round int) {
	fmt.Println("实验开始, 人数:", nums, ", 总计轮数:", round)
	//初始化100个人 每人100元
	//初始化hash,一开始所有人都没有钱
	wealth := make([]float64, nums)
	whoHasMoney := make([]bool, nums)
	for i := range wealth {
		wealth[i] = 100
		whoHasMoney[i] = false
	}

	//模拟Round
	for r := 0; r < round; r++ {
		//判断当前轮次下 有余额的人
		for i := 0; i < nums; i++ {
			if wealth[i] > 0 {
				whoHasMoney[i] = true
			}
		}

		//有钱的拿出1块钱 随机给另外一人
		//每一轮选一个随机种子
		rn := rand.New(rand.NewSource(time.Now().UnixNano()))
		for i := 0; i < nums; i++ {
			if wealth[i] > 0 {
				other := i
				for {
					other = rn.Intn(nums)
					if other != i {
						break
					}
				}

				wealth[i]--
				wealth[other]++
			}
		}
	}

	fmt.Println("第", round, "轮后, 每个人的财富为(贫穷到富有):")
	sort.Float64s(wealth)
	fmt.Println(wealth)

	gini := calGini(wealth)
	fmt.Println("最终的基尼系数为:", gini)
}

func calGini(wealth []float64) float64 {
	var wealthTotal float64 = 0
	var wealthDelta float64 = 0
	var wealthLen int = len(wealth)
	for i := 0; i < wealthLen; i++ {
		wealthTotal += wealth[i]
		for j := 0; j < wealthLen; j++ {
			wealthDelta += math.Abs(wealth[i] - wealth[j])
		}
	}

	return wealthDelta / (2 * float64(wealthLen) * wealthTotal)
}
