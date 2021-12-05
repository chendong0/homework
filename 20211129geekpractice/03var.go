package main

import (
	"fmt"
	"math"
	"reflect"
)

const pi float64 = 3.1415926

func main() {
	var hello string = "你好,golang!"
	fmt.Println(hello)

	var int2 int = 33
	fmt.Println(int2)
	var float3 float64 = 1.34
	fmt.Println(float3)

	var int4, int5 = 44, 66
	fmt.Println(int4, int5)
	float3, float5 := 4.0, 5.5
	fmt.Println(float5, float3)

	var (
		int6, int7 = 6, 7
	)
	fmt.Println(int6, int7)

	var horizontal int = 3
	var vertical int = 4
	fmt.Println(horizontal * vertical)
	var area = horizontal * vertical
	fmt.Println(area)

	var horizontal1, vertical1 int
	fmt.Println(horizontal1)
	fmt.Println(vertical1)

	var name string
	fmt.Println(name)
	var mapA map[string]string
	fmt.Println(mapA)

	a1 := "hello"
	fmt.Println(reflect.TypeOf(a1))
	a2 := 3
	fmt.Println(reflect.TypeOf(a2))
	a3 := 6.0
	fmt.Println(reflect.TypeOf(a3))
	a4 := &a3
	fmt.Println(reflect.TypeOf(a4))

	var f1 float64 = 1.234
	var i1 int = int(f1)
	fmt.Println("f1:", f1, "i1:", i1)
	var a6 uint = math.MaxUint64
	var a7 int = int(a6)
	fmt.Println(a6, a7)

	fmt.Println(pi)

	fmt.Println(pi)
}
