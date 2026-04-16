package main

import (
	"encoding/json"
	"fmt"
	"learnGo/calculation"
	"strconv"

	"github.com/shopspring/decimal"
	"github.com/tidwall/gjson"
)

type User struct {
	Id   string
	Name string
}

type Computer struct{}

func (c Computer) work(usb Usber) {
	fmt.Println(usb.(Usber))
	usb.Start()
	usb.Stop()
}

type Usber interface {
	Start()
	Stop()
}

type Phone struct {
	System string
}

func (p Phone) Start() {
	fmt.Printf("系统：%v 已启动\n", p.System)
}
func (p Phone) Stop() {
	fmt.Printf("系统：%v 已停止\n", p.System)
}

// 空接口
type A interface{}

type Address struct {
	Name        string
	PhoneNumber int
}

func main() {
	// 结构体与结构体指针
	{
		fmt.Println("===结构体")
		type User struct {
			id   int
			name string
		}

		var u User
		u.id = 1
		u.name = "jack"
		fmt.Println(u)

		var ptr *User
		ptr = &u
		ptr.id = 2
		fmt.Println(u)

		var u2 User = User{3, "Jassy"}
		fmt.Println(u2.id, u2.name)
	}

	// 数组
	{
		fmt.Println("===数组")
		var array = [2]int{1, 2}

		var arrayAutoLength = [...]int{1, 2, 3}
		fmt.Printf("值：%v 类型：%T 长度：%d \n", arrayAutoLength, arrayAutoLength, len(arrayAutoLength))
		fmt.Println(array, arrayAutoLength)
	}

	// 切片数组
	{
		fmt.Println("===切片数组")
		var arr []int /* 不指定长度即被认为是切片 反之则被认为是数组 */
		/* 下面一行也为切片声明方式 */
		var arr2 = []int{4, 5, 6}

		for k, v := range arr2 {
			fmt.Printf("| k:%d v:%d \n", k, v)
		}
		fmt.Printf("%v %T %d \n", arr, arr, len(arr))

		a := [3]int{1, 2, 3}
		b := a[:] // 获取数组所有元素作为一个切片
		fmt.Printf("%v-%T\n", b, b)
		c := a[1:2] // 左闭右开原则
		fmt.Printf("%v-%T\n", c, c)
		d := a[1:] // 从下标1开始到最后
		fmt.Printf("%v-%T\n", d, d)
		e := a[:2] // 获取下标2以前的所有元素
		fmt.Printf("%v-%T\n", e, e)

		var sliceA = make([]int, 4, 8) // make方式声明的切片才具有默认值
		fmt.Printf("sliceA切片具有长度为%d，容量为%d，初始值为%v\n", len(sliceA), cap(sliceA), sliceA)
		var arr3 []int // 此方式声明不具有默认值
		fmt.Printf("arr3不具有初始值，为空：%v，且与nil进行布尔为：%t\n", arr3, arr3 == nil)

		// 下标方式适用于修改值
		var sliceB = []string{"a", "b", "c", "d"}
		sliceB[0] = "A"

		// 扩容则使用append
		var emptySlice []string
		emptySlice = append(emptySlice, "E1", "E2")
		emptySlice = append(emptySlice, "E3")
		printLenCapValue(emptySlice)

		// 切片合并
		var sliceString = []string{"a", "b", "c", "d"}
		sliceString[0] = "中文"
		emptySlice = append(emptySlice, sliceString...)
		fmt.Print("中文：")
		printLenCapValue(emptySlice)

		// 切片是引用类型，修改切片会影响底层数组
		var sliceC = []int{1, 2, 3, 4}
		var sliceD = sliceC[1:3] // sliceD 包含 sliceC 的元素 2 和 3
		fmt.Printf("sliceC: %v, sliceD: %v\n", sliceC, sliceD)
		sliceD[0] = 20 // 修改 sliceD 中的第一个元素
		fmt.Printf("修改后 sliceC: %v, sliceD: %v\n", sliceC, sliceD)

		// 使用copy函数互不影响
		var sliceE = []int{1, 2, 3}
		var sliceF = make([]int, 3, 3)
		copy(sliceF, sliceE) // 把sliceE复制给sliceF
		sliceF[0] = 100
		printLenCapValue(sliceE)
		printLenCapValue(sliceF)

		var chineseString = "中文"
		var sliceG = []rune(chineseString)
		fmt.Printf("sliceG: %v, sliceG: %v\n", string(sliceG), sliceG)
	}

	// for range
	{
		for k, v := range "hello世界" {
			fmt.Printf("k:%d v:%c\n", k, v)
		}

		// map键值数据遍历键与值
		map1 := make(map[string]string)
		map1["name"] = "杰克"
		map1["age"] = "18"
		fmt.Println("_k ")
		for _k, v := range map1 {
			fmt.Printf("k:%d v:%c\n", _k, v)
		}
		fmt.Println()
		for k := range map1 {
			fmt.Printf("%s \n", k)
		}
		fmt.Println()

		// 通道数据的便遍历
		ch := make(chan string, 3)
		ch <- "hello"
		ch <- "world"
		ch <- "nullicn!"
		close(ch)

		for v := range ch {
			fmt.Println(v)
		}
	}

	// map
	{
		// 空创建
		var mapEmpty = make(map[string]int)

		// 带初始容量的创建
		var mapCap = make(map[string]int, 4)

		// 字面量创建
		var mapLiteral = map[string]int{
			"one":   1,
			"two":   2,
			"three": 3,
		}

		// 获取
		v1 := mapEmpty["one"]
		fmt.Println("v1 from mapLiteral: ", v1)

		// 修改
		mapLiteral["one"] = 100

		// 删除键值对
		delete(mapLiteral, "two")

		fmt.Printf("%v \n %v \n %v \n", mapEmpty, mapCap, mapLiteral)

		v, ok := mapLiteral["one"]
		fmt.Printf("v: %v ok: %v\n", v, ok)
		v, ok = mapLiteral["four"]
		fmt.Printf("v: %v ok: %v\n", v, ok)

		// 切片中存放map
		sliceMap := make([]map[string]string, 3)
		sliceMap[0] = make(map[string]string)
		sliceMap[0]["name"] = "nullicn"
		sliceMap[0]["age"] = "21"

		// map中存放切片
		mapSlice := make(map[string][]string, 3)
		mapSlice["TODO"] = []string{"学习Go", "练习Go", "掌握Go"}
		mapSlice["TODO"] = append(mapSlice["TODO"], "GO 乔尼 GO GO! LESSION 5 开始了！")
		mapSlice["DATA"] = []string{"2", "7", "7", "8"}

		fmt.Printf("\n切片map map切片\n")
		fmt.Printf("sliceMap: %v\n", sliceMap)
		fmt.Printf("mapSlice: %v\n\n", mapSlice)

	}

	// 类型转换
	{
		// 来点数值类型
		i := int(1)
		f := float64(i)
		fmt.Printf("f: %f\n", f)

		// 字符串引用类型
		str := "114514"
		strToInt, err := strconv.Atoi(str)
		fmt.Printf("转换后的值为：%v，类型为%T，错误为：%v\n", strToInt, strToInt, err)

		strFloat := "27.78"
		strToFloat, err2 := strconv.ParseFloat(strFloat, 64)
		fmt.Printf("转换后的值为：%v，类型为%T，错误为：%v\n", strToFloat, strToFloat, err2)

		toStr := strconv.FormatFloat(strToFloat, 'f', 2, 64)
		fmt.Printf("转换后的值为：%v，类型为%T\n", toStr, toStr)

		var interStr interface{} = "hello nullicn!"
		strFromInterface, ok := interStr.(string)
		if ok {
			fmt.Printf("从接口类型转换到字符串成功，值为：%s\n", strFromInterface)
		} else {
			fmt.Println("从接口类型转换到字符串失败")
		}
	}

	// 函数声明
	{
		main2()
	}

	// 调用其它包
	{
		fmt.Println(calculation.Sum(1, 2))
	}

	// 使用decimal包
	{
		number, err := decimal.NewFromString("277.8")
		fmt.Printf("decimal包NewFromString结果：%v, %v", number, err)
	}

	// 使用gjson包
	{
		var user = User{
			"1145",
			"NULLICN",
		}
		userStr, _ := json.Marshal(user)
		fmt.Println("userStr：", string(userStr))
		jsonStr := `{"Id":"1145","Name":"NULLICN"}`
		gResult := gjson.Get(jsonStr, "Id")
		value := gResult.String()
		fmt.Printf("value：%v, %T", value, value)
	}

	// 结构体实现接口方法
	{
		var myPhone = Phone{
			"XiaoMi Hyper OS 4",
		}

		//var p1 Usber = myPhone
		//p1.Start()
		//p1.Stop()
		var computer = Computer{}
		computer.work(myPhone)

		// 空接口没有任何约束,可以表示任意类型
		var num A = 1
		fmt.Printf("空接口实现，类型为：%T\n", num)
		var str A = "string"
		fmt.Printf("空接口实现，类型为：%T\n", str)
		var slice A = []string{"a", "b", "c"}
		fmt.Printf("空接口实现，类型为：%T\n", slice)

		// 空接口类型转对应类型
		var userInfo = make(map[string]interface{})
		address := Address{
			"NULLICN",
			114514,
		}
		userInfo["user1"] = address
		user1, _ := userInfo["user1"].(Address)
		fmt.Printf("回到正确的类型：%v %T\n", user1, user1)
		fmt.Printf("此时可以使用点号属性获取值：%v\n", user1.Name)

		userInfo["nums"] = []int{1, 2, 3}
		nums, _ := userInfo["nums"].([]int)
		fmt.Printf("回到正确的类型：%v %T\n", nums, nums)
		fmt.Printf("此时可以使用下标获取切片元素：%v\n", nums[0])

	}

	// Go 占位符演示
	{
		//demonstratePlaceholders()
	}

	// printLenCapValue 函数演示（支持多种数据类型）
	{
		/*fmt.Println("===printLenCapValue 函数演示（泛型版本）")

		// int 类型切片
		fmt.Println("--- int 类型 ---")
		slice1 := make([]int, 3, 5)
		printLenCapValue(slice1)
		slice2 := []int{1, 2, 3}
		printLenCapValue(slice2)

		// string 类型切片
		fmt.Println("--- string 类型 ---")
		slice3 := []string{"hello", "world"}
		printLenCapValue(slice3)

		// float64 类型切片
		fmt.Println("--- float64 类型 ---")
		slice4 := []float64{3.14, 2.71, 1.41}
		printLenCapValue(slice4)

		// bool 类型切片
		fmt.Println("--- bool 类型 ---")
		slice5 := make([]bool, 2, 4)
		slice5[0] = true
		slice5[1] = false
		printLenCapValue(slice5)

		// byte 类型切片
		fmt.Println("--- byte 类型 ---")
		slice6 := []byte{'A', 'B', 'C'}
		printLenCapValue(slice6)*/
	}
}

// printLenCapValue 打印任意类型切片的长度、容量和值（使用泛型）
func printLenCapValue[T any](slice []T) {
	fmt.Printf("长度为%d 容量为%d 值为%v\n", len(slice), cap(slice), slice)
}

// init
func init() {
	fmt.Printf("init 总是最先执行\n\n")
}
