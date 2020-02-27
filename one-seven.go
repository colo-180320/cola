//第七节：指针
package main

import (
	"fmt"
)

//指针：
func changPa()  {
	var a int  = 2
	var pa *int  = &a  //&a : 获取到的是 a 变量存储数据的地址
	fmt.Println(pa)
	//结果：0xc00000a0c8
	fmt.Println(*pa)  //这个时候 还没有给 *pa 赋值 所以 他还是 2
	*pa = 3  //将3 指向这个变量地址
	fmt.Println(a)  //因而 a 现在的值获取到的是 3
	//结果：3
}


func main()  {
	/*
		Go语言只有一种 值传递 的方式
	*/
	changPa()
}