package main

import (
	"fmt"
	"math"
)

func main() {
	var a int32 = 78
	printBinary(a)

	var b int32 = -78
	printBinary(b)

	//0b表示2进制: 0110 1110
	var c int32 = 0b1001110
	printBinary(c)

	//0x表示16进制：4e = 0110(4) 1110(e)
	var d int32 = 0x4e
	printBinary(d)

	//78的相反数为-78: (取反78) + 1
	var e int32 = ^a + 1
	printBinary(e)

	//!!!负数最小值的相反数和绝对值都是本身!!!
	var f int32 = math.MinInt32
	printBinary(f)
	printBinary(-f)
	printBinary(^f + 1)
}

func printBinary(num int32) {
	fmt.Print("num:", num, "=")
	for i := 31; i >= 0; i-- {
		if (num & (1 << i)) == 0 {
			fmt.Print("0")
		} else {
			fmt.Print("1")
		}
	}

	fmt.Println()
}
