package main

import "fmt"

type Number struct {
	val int
}

func main() {
	number := &Number{100}

	//number没有变
	changeNil(number)
	fmt.Println(number)

	//number val被改变
	change(number)
	fmt.Println(number)

	//number 会被指向nil
	changeRealNil(&number)
	fmt.Println(number)
}

// 修改的是本地指针本身的副本, 这个修改不会影响到 main 函数中的 number 指针
func changeNil(number *Number) {
	number = nil
}

func change(number *Number) {
	number.val = 99
}

// 要改变 number 指针本身的值（而不仅仅是它所指向的对象），你需要传递一个指向指针的指针
func changeRealNil(number **Number) {
	*number = nil
}
