package main

import "fmt"

func main() {
	for i := 1; i >= 1 && i <= 10; i++ {
		fmt.Print("你好,Golang!")

	}
	//	for m := 1;m >=1 && m <= 100; m++ {
	//
	//		fmt.Println("你好,Golang!")
	//}
	//
	//	for m := 1;m >=1 && m <= 1000; m++ {
	//
	//	fmt.Print("你好,Golang!")
	//}
	//for m := 1;m >=1 && m <= 1383474637; m++ {
	//
	//	fmt.Print("你好,Golang!")
	//}
	f := 0
	for f < 3 {
		fmt.Println("2222")
		f++
	}

	q := 0
	for {
		fmt.Println("qqqq")
		q++
		if q > 3 {
			break
		}

	}
	for t := 0; t < 5; t++ {
		if t%3 == 0 {
			continue
		}
		fmt.Println("8888")
	}
}
