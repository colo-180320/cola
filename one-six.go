//第六节：函数
package main

import (
	"fmt"
)

//swich
func grande(a,b int,p string) int  {
	switch p {
		case "+":
			return  a + b
		case "-":
			return a - b
		case "*":
			return a * b
		case "/":
			return a / b
		case "&":
			q,_:=eval(a,b)  //因为 eval是返回两个值得，如果不要其中的那个参数，那么就用 _ 来表示
			return q
		default:
			panic("错误参数:"+p)
	}
}

//函数，q,r是返回值，那么需要在函数体外接受这个q,r
func eval(a,b int)(q,r int)  {
	return a/b, a*b
}



func main()  {
	/*
		函数：
			1.函数返回多个值时可以起名字
			2.仅用于非常简单的函数
			3.对于调用者而言没有区别
	*/
	var result = grande(1,2,"&")
	fmt.Println(result)
	q,r := eval(1,2)
	fmt.Printf("%d,%d",q,r)
	fmt.Println()

}