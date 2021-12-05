package main

import (
	"fmt"
)

func main() {
	var fruit = "六个苹果"
	var watermelon = true
	if watermelon {
		fruit = "一个苹果"
	}
	fmt.Println(fruit)

	var money = 200 //修改它,看看它会做什么?
	if money == 20 {
		fmt.Println("点个外卖")
	} else if money == 200 {
		fmt.Println("去下馆子")

	} else if money == 2000 {
		fmt.Println("到米其林体验体验")

	} else if money == 20000 {
		fmt.Println("出国游玩一圈")

	} else {
		fmt.Println("容我想一想")
	}

	var money1 = 218
	switch money1 {
	case 20:
		fmt.Println("点个外卖")
	case 200:
		fmt.Println("去下馆子")
	case 2000:
		fmt.Println("到米其林体验体验")
	case 20000:
		fmt.Println("出国游玩一圈")
	default:
		fmt.Println("容我想想")
	}

	var money2 = 200
	var busy bool = true
	switch {
	case money2 >= 0 && money2 <= 20:
		fmt.Println("点1个外卖")
		if busy {
			break
		}
		fmt.Println("再吃个零食")
	case money2 > 20 && money2 <= 200:
		fmt.Println("去下馆子")
	default:
		fmt.Println("容我想想")
	}
	fmt.Println("结束")

	var money3 = 0
	var busy1 bool = false
	switch {
	case money3 >= 0 && money3 <= 20:
		fmt.Println("点1个外卖")
		if busy1 {
			break
		}
		fmt.Println("再吃个零食")
		fallthrough //直接并入下一个处理分支而无需判断条件
	case money3 > 20 && money3 <= 200:
		fmt.Println("去下馆子")
	default:
		fmt.Println("容我想想")

	}
	fmt.Println("结束")

}
