package main

import (
	"fmt"
)

func main() {
	//var name string = "Chen"
	//var  Weight float64 = 58.5
	//var  Tall float64 = 1.68
	//var  Age int = 32
	//
	//var SexWeight int
	//var  Sex string = "male"
	//var p Person(
	//	 Age int
	//	 Tall float64
	//	 Sex string
	//	 Weight float64
	//	 Name string
	//)
	fmt.Println("hello")

	for {
		var Name string
		fmt.Print("姓名:")
		fmt.Scanln(&Name)

		var Weight float64
		fmt.Print("体重(KG):")
		fmt.Scanln(&Weight)

		var Tall float64
		fmt.Print("身高(米):")
		fmt.Scanln(&Tall)

		var Age int
		fmt.Print("年龄:")
		fmt.Scanln(&Age)

		var SexWeight int
		var Sex string = "male"
		fmt.Println("性别(female/male):")
		fmt.Scanln(&Sex)

		if Sex == "male" { //==要区分理解=和==的区别和用法.
			SexWeight = 0
		} else {
			SexWeight = 1
		}

		var bmi = (Weight / (Tall * Tall))
		var (
			fatRate float64 = (1.2*bmi + 0.23*float64(Age) - 5.4 - 10.8*float64(SexWeight)) / 100
		)
		fmt.Printf("%s的bmi指数为%.2f\n", Name, bmi)
		fmt.Printf("%s的体脂率为%.2f\n", Name, fatRate)

		if Sex == "male" {
			//编写男性体脂率与体脂状态
			if Age >= 18 && Age <= 39 {
				if fatRate < 0.05 {
					fmt.Printf("体脂率为%.2f 体脂率为%.2f评估:偏瘦,加强营养多锻炼", fatRate)
				} else if fatRate >= 0.05 && fatRate < 0.1 { //5%报错,只能用0.15
					fmt.Printf("体脂率为%.2f评估:您的状态偏瘦", fatRate)
				} else if fatRate >= 0.1 && fatRate < 0.17 {
					fmt.Printf("体脂率为%.2f评估:您目前标准,请保持", fatRate)
				} else if fatRate >= 0.17 && fatRate < 0.22 {
					fmt.Printf("体脂率为%.2f评估:您达到标准警戒线", fatRate)

				} else if fatRate >= 0.22 && fatRate < 0.27 {
					fmt.Printf("体脂率为%.2f评估:您轻度肥胖", fatRate)
				} else if fatRate >= 0.27 && fatRate < 0.45 {
					fmt.Printf("体脂率为%.2f评估:您重度肥胖", fatRate)
				} //对比作业,多加了左括号
			} else if Age > 39 && Age <= 59 { //对比作业后,发现要与第2个if平齐
				if fatRate >= 0.05 && fatRate < 0.12 {
					fmt.Printf("体脂率为%.2f评估:偏瘦", fatRate)
				} else if fatRate >= 0.12 && fatRate > 0.18 {
					fmt.Printf("体脂率为%.2f评估:您标准健康型.", fatRate)
				} else if fatRate >= 0.18 && fatRate > 0.23 {
					fmt.Printf("体脂率为%.2f评估:您标准临界线.", fatRate)
				} else if fatRate >= 0.23 && fatRate > 0.28 {
					fmt.Printf("体脂率为%.2f评估:您标准健康型.", fatRate)
				} else if fatRate >= 0.28 && fatRate > 0.45 {
					fmt.Printf("体脂率为%.2f评估:您偏胖", fatRate)
				} else if Age >= 60 {
					if fatRate > 0.05 && fatRate < 0.13 {
						fmt.Printf("体脂率为%.2f评估:偏瘦", fatRate)
					} else if fatRate >= 0.13 && fatRate < 0.19 {
						fmt.Printf("体脂率为%.2f 评估:您标准健康型.", fatRate)
					} else if fatRate >= 0.19 && fatRate < 0.24 {
						fmt.Printf("体脂率为%.2f评估:标准型警戒线.", fatRate)
					} else if fatRate >= 0.24 && fatRate < 0.29 {
						fmt.Printf("体脂率为%.2f 评估:您轻度肥胖.", fatRate)
					} else if fatRate >= 0.29 && fatRate < 0.45 {
						fmt.Printf("体脂率为%.2f 评估:重度肥胖", fatRate)
					}
					//fmt.Printf(fatRate)

				} //对比作业,少了左括号.少了括号,代码逻辑全部变了.所以if语句的读取顺序和条件没有和预想一致.
				//数据范围内的没有完全覆盖,所有数据无法带入if语句.

			} else if Age < 18 {
				fmt.Println("不评估18岁以下,因为波动太大,效果没有意义")
			}

			var whetherContinue string
			fmt.Print("\n是否录入下一个(y/n)?")
			fmt.Scanln(&whetherContinue)
			if whetherContinue != "y" {
				break
			}
		}
	}
}

//var {
//	Name      string
//	Age       int
//	Tall      float64
//	Weight    float64
//	Sex       string
//	SexWeight int
//}
