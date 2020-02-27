//第二节：切片的概念
/*
	切片（slice）是对数组一个连续片段的引用（该数组我们称之为相关数组，通常是匿名的），
	    所以切片是一个引用类型（因此更类似于 C/C++ 中的数组类型，或者 Python 中的 list 类型）。
	    这个片段可以是整个数组，或者是由起始和终止索引标识的一些项的子集。需要注意的是，终止索引标识的项不包括在切片内。
	    切片提供了一个相关数组的动态窗口。
*/
package main

import (
	"fmt"
)

//获取数组内容
func getSlice()  {
	arr := [...]int{0,1,2,3,4,5}
	/*
		2 ==> 数组下的第二位
		4 ==> 数组下的第四位的前一个
	*/
	s := arr[2:4]  // s 现在就属于切片了
	//结果： [2,3]
	fmt.Println(s)

	/*
		s ==> 表面看到的数据 就： [2,3]
			实际，他会默认得到 arr 从 位数2 后面的所有值都会获取到 所以： [2,3,4,5]
																下标：	0,1,2,3
	*/
	s_one := s[0:4]
	fmt.Println(s_one)
	//结果：[2,3,4,5]
}

//数组数据获取
func getSliceDetail()  {
	arr := [...]int{0,1,2,3,4,5,6,7}
	fmt.Println("arr[2:6]",arr[2:6])  //结果：[2,3,4,5]
	fmt.Println("arr[:6]",arr[:6])	  //结果：[0,1,2,3,4,5]
	fmt.Println("arr[2:]",arr[2:])	  //结果 : [2,3,4,5,6,7]
	fmt.Println("arr[:]",arr[:])	  //结果： [0,1,2,3,4,5,6,7]
}

//切片传递给函数
func sum(a []int) int {
	s := 0
	for _,v :=range a {
		s += v
	}
	a[1] = 50 //切片更改值,可以作用到整个数据 a
	return s
}

//在没有数组的情况下，如果创建切片：make/new
/*
	new([100]int)[0:50]:为每个新的类型T分配一片内存，初始化为0并且返回类型为 *T 的内存地址：这种方法返回一个执行类型为 T，值为0的地址的指针
	make([]int,10,50)：返回一个类型为T的初始值，它只适用于3种内建的引用类型：切片，map和channel
	总结：new函数分配内存，make函数初始化
*/
func makeSlice()  {
	//slice1 := make([]type, start_length, capacity)
	var arr_one []int= make([]int,10,50)
	fmt.Printf("原数据：%d-长度len：%d-cap长度cap：%d",arr_one,len(arr_one),cap(arr_one))

	fmt.Println()

	var arr_two []int = new([100]int)[0:5]
	fmt.Printf("原数据：%d-长度len：%d-cap长度cap：%d",arr_two,len(arr_two),cap(arr_two))
}

//关于 byte 这个要注意一下
func getBty()  {
	a := make([]byte,5)
	fmt.Println()
	fmt.Printf("原数据：%d-长度len：%d-cap长度cap：%d",a,len(a),cap(a))

	b :=[]byte{0,1,2,3,4}
	fmt.Println(b[2:])  //结果：[2,3,4]

}

func main()  {
	getSlice()
	getSliceDetail()

	array_one := [...]int{1,2,3}
	s := sum(array_one[:])
	fmt.Println(s)  //结果：6
	fmt.Println(array_one) //结果：[1,50,3]
	makeSlice()
	getBty()
}
