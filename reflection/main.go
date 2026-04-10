package main

import (
	"fmt"
	"reflect"
)

func reflectType(x interface{}) {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)
	fmt.Printf("t:%v v:%v\n\n", t, v)
	fmt.Printf("type of x: %v\n", t)
	k := t.Kind()
	//kv := v.Kind()
	switch k {
	case reflect.Int:
		fmt.Printf("int type: %v computed value：%v\n", t, v.Int()+10)
	}
}

func reflectionSet(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Elem().Kind()
	if k == reflect.Int64 {
		v.Elem().SetInt(2778)
	} else if k == reflect.String {
		v.Elem().SetString("Hello NULLICN!")
	}
}

type User struct {
	Name string `json:"用户名" from:"用户自定义"`
}

func (u User) GetUserName() string {
	return u.Name
}
func (u *User) SetUserName(name string) {
	u.Name = name
}

func reflectionStruct(x interface{}) {
	t := reflect.TypeOf(x)
	// 处理指针类型，获取其指向的类型
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		fmt.Println("参数不是一个结构体")
		return
	}

	//v := reflect.ValueOf(x)
	//f0 := t.Field(0)
	fieldNum := t.NumField()
	fmt.Printf("结构体具有 %d 个属性\n", fieldNum)
	f0, ok := t.FieldByName("Name")
	if !ok {
		fmt.Println("特定属性不存在")
		return
	}
	fmt.Println("字段名称", f0.Name)
	fmt.Println("字段类型", f0.Type)
	fmt.Println("字段Tag", f0.Tag.Get("json"))
	fmt.Println("字段Tag", f0.Tag.Get("from"))

	methodNum := t.NumMethod()
	fmt.Printf("结构体具有 %d 个方法\n", methodNum)
	//m0 := t.Method(0)
	//fmt.Println(m0)
	v := reflect.ValueOf(x) // 获取函数
	mGet := v.MethodByName("GetUserName")
	mSet := v.MethodByName("SetUserName")
	mGetValue := mGet.Call(nil)

	params := make([]reflect.Value, 1)
	params[0] = reflect.ValueOf("nullicn")
	mSet.Call(params)
	// 在SetUserName之后再次调用GetUserName获取修改后的值
	v.Elem().FieldByName("Name").SetString("Hello NULLICN!")
	mGetValueAfter := mGet.Call(nil)
	fmt.Printf("修改前: %v, 修改后: %v\n", mGetValue, mGetValueAfter)

}

func main() {
	//var a int64 = 100
	//astr := "hello world"
	//fmt.Println(a)
	//reflectionSet(&a)
	//fmt.Println(a)

	//reflectType(a)
	//fmt.Println(a, astr)
	//b := a
	//b = 200
	//bstr := astr
	//bstr = "hello NULLICN!"
	//fmt.Println(a, b, astr, bstr)

	user := User{
		"NULLICN",
	}
	reflectionStruct(&user)
}
