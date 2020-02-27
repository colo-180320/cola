//第三节：切片的操作
package main

import (
	"fmt"
)

//切片添加数据：append 只能添加到末端
func AppendAction()  {
	a := make([]int,5)
	fmt.Println(cap(a))   //在没有给切片添加元素之前，cap数为 5
	a_one := append(a,1)
	fmt.Println(a_one)
	fmt.Println(cap(a_one))  //添加一位  之后：cap 数为 10
	//结果：[0 0 0 0 0 1]
}

//切片的复制：copy
func CopyAction()  {
	a := make([]int,9)  //切片
	fmt.Println(a)
	b := []int{1,2,3,4,5,6}	//数组
	copy(a,b)
	fmt.Println(a)
	//结果  [1,2,3,4,5,6,0,0,0]
	/*
		解读：
			a 首先是 [0,0,0,0,0,0,0,0,0]
			通过：copy(a,b)  //将b的值，复制到a,的前面，但是不可以超过a 设置 的len 长度
	*/

}

func main()  {
	/*
		注意：
			1.添加元素时如果超越原来的cao,系统会重新分配更大的底层数据(自动判断要增加多少)  一般，如果len等于cap 那就会将cap = len * 2 来扩大
	*/
	AppendAction()
	CopyAction()
}


