package main

import (
	"fmt"
)

func demonstratePlaceholders() {
	fmt.Println("========== Go 占位符详解 ==========\n")

	// 1. 通用占位符
	fmt.Println("【通用占位符】")
	value := 42
	fmt.Printf("%%v - 默认格式: %v\n", value)
	fmt.Printf("%%#v - Go语法格式: %#v\n", []int{1, 2, 3})
	fmt.Printf("%%T - 类型信息: %T\n", value)
	fmt.Printf("%%%% - 百分号输出: 100%%\n")
	fmt.Println()

	// 2. 布尔占位符
	fmt.Println("【布尔占位符】")
	flag := true
	fmt.Printf("%%t - 布尔值: %t\n", flag)
	fmt.Printf("%%t - 布尔值: %t\n", false)
	fmt.Println()

	// 3. 整数占位符
	fmt.Println("【整数占位符】")
	num := 255
	fmt.Printf("%%d - 十进制: %d\n", num)
	fmt.Printf("%%b - 二进制: %b\n", num)
	fmt.Printf("%%o - 八进制: %o\n", num)
	fmt.Printf("%%x - 十六进制(小写): %x\n", num)
	fmt.Printf("%%X - 十六进制(大写): %X\n", num)
	fmt.Printf("%%c - 对应Unicode字符: %c\n", 65) // 65 对应 'A'
	fmt.Printf("%%q - 转义的字符: %q\n", 65)
	fmt.Println()

	// 4. 浮点数占位符
	fmt.Println("【浮点数占位符】")
	pi := 3.14159265
	fmt.Printf("%%f - 小数点表示: %f\n", pi)
	fmt.Printf("%%.2f - 保留2位小数: %.2f\n", pi)
	fmt.Printf("%%e - 科学记数法(小写): %e\n", pi)
	fmt.Printf("%%E - 科学记数法(大写): %E\n", pi)
	fmt.Printf("%%g - 自动选择(小写): %g\n", pi)
	fmt.Printf("%%G - 自动选择(大写): %G\n", pi)
	fmt.Println()

	// 5. 字符串占位符
	fmt.Println("【字符串占位符】")
	str := "Hello"
	fmt.Printf("%%s - 字符串: %s\n", str)
	fmt.Printf("%%q - 带引号的字符串: %q\n", str)
	fmt.Printf("%%x - 十六进制: %x\n", str)
	fmt.Printf("%%X - 十六进制(大写): %X\n", str)
	fmt.Println()

	// 6. 字节占位符
	fmt.Println("【字节占位符】")
	bytes := []byte{72, 101, 108, 108, 111} // "Hello"
	fmt.Printf("%%s - 字节切片: %s\n", bytes)
	fmt.Printf("%%q - 字节切片(转义): %q\n", bytes)
	fmt.Printf("%%x - 字节切片(十六进制): %x\n", bytes)
	fmt.Println()

	// 7. 指针占位符
	fmt.Println("【指针占位符】")
	ptr := &value
	fmt.Printf("%%p - 指针地址: %p\n", ptr)
	fmt.Printf("%%v - 指针值: %v\n", ptr)
	fmt.Println()

	// 8. 宽度和精度修饰符
	fmt.Println("【宽度和精度修饰符】")
	fmt.Printf("%%5d - 右对齐宽度5: [%5d]\n", 42)
	fmt.Printf("%%-5d - 左对齐宽度5: [%-5d]\n", 42)
	fmt.Printf("%%05d - 零填充宽度5: [%05d]\n", 42)
	fmt.Printf("%%.3s - 字符串精度3: [%.3s]\n", "Hello")
	fmt.Printf("%%5.2f - 宽度5,精度2: [%5.2f]\n", pi)
	fmt.Println()

	// 9. 标志符
	fmt.Println("【标志符】")
	fmt.Printf("%%+d - 带符号: %+d\n", 42)
	fmt.Printf("%%+d - 带符号: %+d\n", -42)
	fmt.Printf("%% d - 空格表示正数: [% d]\n", 42)
	fmt.Printf("%%#x - 显示0x前缀: %#x\n", 255)
	fmt.Printf("%%#o - 显示0前缀: %#o\n", 64)
	fmt.Println()

	// 10. 实际应用示例
	fmt.Println("【实际应用示例】")
	type User struct {
		ID   int
		Name string
		Age  float32
	}

	user := User{ID: 1, Name: "张三", Age: 28.5}
	fmt.Printf("用户信息 - ID: %d, 名字: %s, 年龄: %.1f\n", user.ID, user.Name, user.Age)
	fmt.Printf("完整结构体: %#v\n", user)
	fmt.Printf("结构体类型: %T\n", user)
}
