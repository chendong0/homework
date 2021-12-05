package main

import "fmt"

func main() {

	//var money interface{} = 10
	// var money interface{} = 10.0
	var money interface{} = "10"
	switch money.(type) {
	case int:
		fmt.Println("money是 int")
	case int64:
		fmt.Println("money是 int64")
	case float64:
		fmt.Println("money是 float64")
	case float32:
		fmt.Println("money是 float32")
	default:
		fmt.Println("money是未知类型")
	}
	fmt.Println("结束", money)

	//var money1 interface{} = 10
	var money1 interface{} = 10.0
	//var money1 interface{} = "10"
	switch newmoney := money1.(type) {
	case int:
		fmt.Println("money1是 int", newmoney+11)
	case int64:
		fmt.Println("money1是 int64", newmoney+22)
	case float64:
		fmt.Println("money1是 float64", newmoney+33)
	case float32:
		fmt.Println("money1是 float32", newmoney+44)
	default:
		fmt.Println("money1是未知类型")
	}
	fmt.Println("结束", money1)

	//var money2 interface{} = 10
	// var money2 interface{} = 10.0
	var money2 interface{} = "10"
	switch money2.(type) {
	case int, int64, int16, int32:

		fmt.Println("money3是 整数")
	case float64, float32:
		fmt.Println("money2是 小数")

	default:
		fmt.Println("money2是未知类型")
	}
	fmt.Println("结束", money)
}
