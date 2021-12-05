package main

import "fmt"

func main() {
	a := 1
	fmt.Println(a)
	//a := 2 //重复声明会报错,
	a = 2 //变量不能重名,可以重新赋值
	//：=是声明并赋值。变量重名可以解释为同一个变量名称不能声明并赋值两次
	fmt.Println(a)
}
