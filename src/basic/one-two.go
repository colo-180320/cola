//第二节：内建变量类型
package main
// bool string
// int int8 int16 int32 int64 uintptr
/*
	int / unit 在32位操作系统  32位（4个字节）  64位 64位（8个字节）
	与操作系统无关的整数int
		int8 (-128 -> 127)
		int16 (-32768 -> 32767)
		int32 (-2,147,483,648 -> 2,147,483,647)
		int64 (-9,223,372,036,854,775,808 -> 9,223,372,036,854,775,807)
	与操作系统无关的无符号整数：
		uint8 (0 -> 255)
		uint16 (0 -> 65,535)
		uint32 (0 -> 4,294,967,295)
		uint64 (0 -> 18,446,744,073,709,551,615)
*/
// byte,rune
// float32 float64 complex64 complex128
/*
	浮点数
		float32 (+- 1e-45 -> +- 3.4 * 1e38)
		float64 (+- 5 * 1e-324 -> 107 * 1e308)  ==>尽可能地使用 float64，因为 math 包中所有有关数学运算的函数都会要求接收这个类型
	复数
		complex64 (32 位实数和虚数)
		complex128 (64 位实数和虚数)
*/

/*
	注意：
		1、Go中不允许不同类型之间混合使用  方法：variableType()
			===》 可以使用显示转化 方法：variableCanType()
*/
import (
	"fmt"
	"math"
	"math/cmplx"
)


//不是一种类型的不可以混合使用
func variableType()  {
	/*var n int
	var m int32
	n = 15
	m = n + n
	fmt.Println(m)*/
	//结果：.\one-two.go:36:4: cannot use n + n (type int) as type int32 in assignment
}

//显示转化类型
func variableCanType()  {
	var n int = 15
	var m int32
	m = int32(n)
	fmt.Println(m)
	//结果：15
}

//内置数据处理
func euler()  {
	c := 3 + 4i
	fmt.Println(cmplx.Abs(c))
	fmt.Println(cmplx.Pow(math.E,1i*math.Pi)+1)
	fmt.Println(cmplx.Exp(1i*math.Pi)+1)
}

func main()  {
	variableType()
	variableCanType()
	euler()
}