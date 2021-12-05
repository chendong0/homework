package main

import (
	"fmt"
)

func main() {
	var name string = "Chen"
	var weight float64 = 58.5
	var tall float64 = 1.68
	var age int = 30
	var sexWeight int
	var sex string = "male"

	if sex == "male" { //==要区分理解=和==的区别和用法.
		sexWeight = 0
	} else {
		sexWeight = 1
	}

	var bmi = (weight / (tall * tall))
	var (
		fatRate float64 = (1.2*bmi + 0.23*float64(age) - 5.4 - 10.8*float64(sexWeight)) / 100
	)
	fmt.Printf("%s的bmi指数为%.2f\n", name, bmi)
	fmt.Printf("%s的体脂率为%.2f\n", name, fatRate)

	if sex == " male" {
		//编写男性体脂率与体脂状态
		if age >= 18 && age <= 39 {
			if fatRate >= 0.05 && fatRate <= 0.1 { //5%报错,只能用0.15
				fmt.Println("评估:您的状态偏瘦")
			} else if fatRate >= 0.11 && fatRate <= 0.16 {
				fmt.Println("评估:您目前标准,请保持")
			} else if fatRate >= 0.17 && fatRate <= 0.21 {
				fmt.Println("评估:您达到标准警戒线")
			} else if fatRate >= 0.22 && fatRate <= 0.26 {
				fmt.Println("评估:您轻度肥胖")
			} else if fatRate >= 0.27 && fatRate <= 0.45 {
				fmt.Println("评估:您重度肥胖")
			} else if age >= 40 && age <= 59 {
				if fatRate >= 0.05 && fatRate <= 0.11 {
					fmt.Println("评估:偏瘦")
				} else if fatRate >= 0.12 && fatRate >= 0.17 {
					fmt.Println("评估:您标准健康型.")
				} else if fatRate >= 0.18 && fatRate >= 0.22 {
					fmt.Println("评估:您标准临界线.")
				} else if fatRate >= 0.23 && fatRate >= 0.27 {
					fmt.Println("评估:您标准健康型.")
				} else if fatRate >= 0.28 && fatRate >= 0.45 {
					fmt.Println("评估:您偏胖")

					//}else if age >=60
					//fmt.Println(fatRate)
				}
			}
		}
	} //fmt.Println("不评估18岁以下,因为波动太大,效果没有意义")
}