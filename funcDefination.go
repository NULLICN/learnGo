package main

import "fmt"

// 函数定义
// 可变长参数为切片
func sliceArgs(x ...int) {
	fmt.Printf("值：%v 类型：%T \n", x)
}

// 剩余参数
func restArgs(x int, rest ...int) {
	fmt.Printf("值：%v 类型：%T \n", x, x)
	fmt.Printf("值：%v 类型：%T \n", rest, rest)
}

// 规定返回值 通过var x, y = returnValues()可以分别拿到两个值
func returnValues() (x string, y int) {
	x = "hello"
	y = 100
	return // 不需要额外在return后写明，但确保函数体内有对应的变量
}

// 返回一个切片
func returnSlice() []int {
	x := 233
	y := 2778
	return []int{x, y}
}

// 函数类型与变量

type calc func(int, int) int

// 定义自定义类型
type myMap map[string]int

func add(x, y int) int {
	return x + y
}

// 匿名函数
func anonyFunc() {
	// 匿名自执行函数
	func() {
		fmt.Println("匿名自执行函数")
	}()

	var fn = func() {
		fmt.Println("匿名函数赋值给变量")
	}
	fn()
}

// 匿名defer返回
func anonyDeferRet() (a int) {
	defer func() {
		a += 1 // 能直接参与修改返回值
	}()
	return 5 // 6
}
func anonyDeferRet2() int {
	a := 0
	defer func() {
		a += 1 // 返回值已确立通过变量a返回，defer函数无法修改返回值
	}()
	return a // 0
}

// 闭包函数
func packedFunc() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main2() {
	var c calc = add
	result := c(1, 2)
	fmt.Println(result)
	var myM myMap = myMap{"c": 1, "a": 2, "b": 3}
	fmt.Printf("myM的值：%v，类型为:%T\n", myM, myM)

	var packedfn = packedFunc()
	fmt.Println("闭包函数：")
	fmt.Println(packedfn())
	fmt.Println(packedfn())
	fmt.Println(packedfn())

	fmt.Println(anonyDeferRet())
	fmt.Println(anonyDeferRet2())
}
