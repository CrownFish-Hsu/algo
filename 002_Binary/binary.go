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
	e := ^a + 1
	printBinary(e)

	//!!!负数最小值的相反数和绝对值都是本身!!!
	var f int32 = math.MinInt32
	printBinary(f)
	printBinary(-f)
	printBinary(^f + 1)
}

// 输出二进制
func printBinary(num int32) {
	fmt.Print("num:", num, "=")
	for i := 31; i >= 0; i-- {
		//当前值与 最高位以1开头后面一串0的数 相与，最高位如果为0，则表示当前为为0；否则当前位为1，后面相与的肯定都是0
		//不可改成: (a & (1 << i)) == 1 ? "1" : "0";
		//因为a如果第i位有1，那么(a & (1 << i))是2的i次方，而不一定是1。
		//比如，a = 0010011, a的第0位是1，第1位是1，第4位是1, (a & (1<<4)) == 16（不是1），说明a的第4位是1状态
		if (num & (1 << i)) == 0 {
			fmt.Print("0")
		} else {
			fmt.Print("1")
		}
	}

	fmt.Println()
}
