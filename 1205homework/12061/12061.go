package main

import (
	"fmt"
)

func main() {
	//var name string = "Chen"
	//var p.Weight float64 = 58.5
	//var p.Tall float64 = 1.68
	//var p.Age int = 32
	//
	//var SexWeight int
	//var p.Sex string = "male"
	var p Person
	p.Age = 18
	p.Tall = 1.65
	p.Sex = "male"
	p.Weight = 58.5
	p.Name = "CHEN"

	if p.Sex == "male" { //==要区分理解=和==的区别和用法.
		p.SexWeight = 0
	} else {
		p.SexWeight = 1
	}

	var bmi = (p.Weight / (p.Tall * p.Tall))
	var (
		fatRate float64 = (1.2*bmi + 0.23*float64(p.Age) - 5.4 - 10.8*float64(p.SexWeight)) / 100
	)
	fmt.Printf("%s的bmi指数为%.2f\n", p.Name, bmi)
	fmt.Printf("%s的体脂率为%.2f\n", p.Name, fatRate)

	if p.Sex == "male" {
		//编写男性体脂率与体脂状态
		if p.Age >= 18 && p.Age <= 39 {
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
		} else if p.Age > 39 && p.Age <= 59 { //对比作业后,发现要与第2个if平齐
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
				//}else if p.Age >=60
				//fmt.Printf(fatRate)

			} //对比作业,少了左括号.少了括号,代码逻辑全部变了.所以if语句的读取顺序和条件没有和预想一致.
			//数据范围内的没有完全覆盖,所有数据无法带入if语句.
		} else {
			fmt.Println("不评估18岁以下,因为波动太大,效果没有意义")
		}
	}
}

type Person struct {
	Name      string
	Age       int
	Tall      float64
	Weight    float64
	Sex       string
	SexWeight int
}
