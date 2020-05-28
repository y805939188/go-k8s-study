package main

import (
	"fmt"
	"unsafe"
)

func Hello(name string) string {
	return "Hello, world" + name
}

func main() {
	var a uint8 = 1 // 无符号整形
	var b int32 = 1 // 4字节的整数
	var c int = 1 // 根据系统是64或32位改变
	fmt.Print(unsafe.Sizeof(a)) // 打印1 占用1字节
	fmt.Print(unsafe.Sizeof(b)) // 打印4
	fmt.Print(unsafe.Sizeof(c)) // 打印8 如果是32位的话就打印4

	var d float32 = 1
	var e float64 = 1 // 浮点型float后面必须加数字
	fmt.Print(unsafe.Sizeof(d)) // 4
	fmt.Print(unsafe.Sizeof(e)) // 8

	var f bool = false
	fmt.Print(unsafe.Sizeof(f)) // 1

	var g byte = 1
	var h rune = 1 // 基本等同于int32
	fmt.Print(unsafe.Sizeof(g)) // 1
	fmt.Print(unsafe.Sizeof(h)) // 4

	var i string = "666"
	fmt.Print(unsafe.Sizeof(i)) // 16



	// 派生类型
	//Pointer // 指针类型
	// 数组类型
	// struct // 结构体
	// chan // channel 类型
	// func // 函数类型
	// slice // 切片类型
	// interface // 接口类型
	// map // map 类型

	// 类型0值 就是某个变量被声明后的默认类型
	// 一般值类型默认值是0 布尔是false string是空字符串

	// 可以对类型设置别名
	type ding int32 // 给 int32 类型定义了一个别名叫 ding
	var j ding = 1
	fmt.Println(j)


	// 数组类型
	//var k
	name := [5] byte {'a', 'b', 'c', 'd', 'e'} // { 97 98 99 100 101 }
	fmt.Println(name)

	num := [4] int { 1,2,3 } // [ 1 2 3 0 ]
	fmt.Println(num)


	// 单个变量的声明
	// var <变量名称> [变量类型]
	// <变量名称> = <值, 表达式, 函数等>
	// 或 var <变量名称> [变量类型] = <值, 表达式, 函数等>

	// 分组声明格式
	//var (
	//	a int
	//	b float32
	//	c string
	//)

	// 同时声明多个
	// var a, b, c int = 1, 2, 3 或 a, b := 1,2
	// var a, b, c = 1, 2, 3 不写类型go会自动推断
	// a, b, c := 1, 2, 3 这样可以不写 var
	// 冒号的这种写法 只能用在函数的局部变量中 不能用在全局

	// 全局变量必须加 var  局部变量可以省略 var

	// go 是强数据类型 不存在隐式转换 并且只能在兼容类型互相转换
	// 类型转换格式 <变量名称> [:]= <目标类型>(<需要转换的变量>)
	// 冒号可有可无
	// var a int = 3 var b float32 = 6.6 c := float32(a)

	// 变量的可见性
	// 大写字母开头的变量可以到处 其他包课件
	// 小写开头的就是不可导出的私有变量

	// 引用其他包
	// import (
	//   "xxx/yyy/packageName"
	// )
	// packageName.变量名


	// 定义常量
	// const 常量名 [type] = value type 可以省略 无类型常量
	// 常量只支持布尔 字符串 整数 和 浮点
	// 可以通过 len 取size
	// const ding string = "牛逼" len(ding) // 6 因为采用utf8所以每个中文占3字节

	// iota 是特殊常量 只能在常量中使用
	// iota在遇到const的时候被重置为 0
	// 在const中每新增一行常量声明 iota 会 +1
	// 调制使用法
	// 插队使用法
	// 表达式隐式使用法
	// 当行使用法

	 const aa = iota
	 const bb = iota
	 // 以上这俩如果直接打印都是 0

	 const (
	 	cc = iota
	 	dd = iota // 这样打印的话 dd 就会变成 1
	 )

	 const (
	 	ee = iota // 0
	 	ff = iota // 1
	 	_ // _ 下划线是个特殊变量 所有给 _ 的value都会被扔进垃圾桶 跳值
	 	zz = 6.6 // 插队
	 	gg = iota // 4 由于上面有个 _ 和 zz 所以这里直接变成 4
	 )

	 const (
	 	kk = iota * 2 // 0
	 	ll // 2
	 	mm // 4
	 	// ll 会自动继承 kk 的iota * 2 mm 又继承 ll 的iota * 2
	 )

	var ding1 interface{}
	 ding1 = 66
	switch ding1.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	}



	ding2 := []string {"ding1", "ding2", "ding3"}
	for _, value := range ding2 {
		fmt.Println(value)
	}


	goto One
	fmt.Println("8888")
	One:
		fmt.Println("使用goto可以直接跳转到这里")








	fmt.Println(Hello("ding"))
}
