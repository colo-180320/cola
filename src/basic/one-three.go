//第三节：常量与枚举
package main

import (
	"fmt"
	"math"
)
//定义常量： const 数值可作为各种类型使用 方法：defConst
const fileName  = "abc.txt"
const a  = 1 //常量如果下面不去调用，是不会报错的
/*
	定义多个常量：
	const(
		fileName = "abc.txt"
		a = 1
		b = 2
	)
	const a,b = 1,2
*/

//数值可作为各种类型使用
func defConst()  {
	const c,d  = 3,4
	var f int = int(math.Sqrt(c * c + d* d)) //勾股定理
	fmt.Println(f)
}

//枚举
func enums()  {
	/*const  (
		cpp = 0
		java = 1
		php = 2
		python = 3
	)*/
	const  (
		cpp = iota
		java
		php
		python
	)
	fmt.Println(cpp,java,php,python)
	//结果：0 1 2 3

	const  (
		A = iota
		_  //注意这个
		B
		C
	)
	fmt.Println(A,C,B)
	//结果：0 3 2

	const  (
		a = 1 << (10*iota)
		b
		c
		d
	)
	fmt.Println(a,b,c,d)
	//结果：1  1024  1048576  1073741824
}

func main()  {
	/*
		变量定义要点回顾：
		1.变量类型卸载变量名之后  var a int = 1
		2.编译器可推测变量类型 （局部变量） a := 1
		3.没有char，只有 rune
	*/
	fmt.Println(fileName)
	defConst()
	enums()
}