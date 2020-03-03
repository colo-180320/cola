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

	c := []int{1,2,3,4}
	d := []int{7,6,5}
	copy(c,d)
	fmt.Println(c)  //结果：[7 6 5 4 ]  会将原来的c中 1,2,3的值覆盖掉
}
//切片的删除:
func DelAction()  {
	a := make([]int,10)
	//给切片赋值：
	for i:=0; i<10; i++ {  //这里 这个i 的范围  因为切片就跟array 一样，是从下表 0 开始计算的 所以 最大值 9
		a[i]=i
	}
	fmt.Println(a) //结果： [0,1,2,3,4,5,6,7,8,9]
	//删除第一位：
	fmt.Println(a[1:])  //结果：[1,2,3,4,5,6,7,8,9]
	//删除最后一位：
	fmt.Println(a[:len(a)-1]) //结果：[0,1,2,3,4,5,6,7,8]
	//删除中间的数据：
	p := []int{1,2,3,4,5}
	fmt.Println(append(p[:3],p[4:]...))  //切记：这个 切片连接使用append 一定要在最后：输入三个 ...  否则会报错
	/*
		删除某个元素的规则：
		fmt.Println(append(p[:index],p[index+1:]...))  ==> 移除 p 下标 index 的元素
	*/

}

//二维切片：
func TwoSlice()  {
	 s := [][]int{{10},{1,2,3}}
	 fmt.Println(s) //结果：[[10] [1 2 3]]
	 fmt.Println(s[1][1]) //结果：2
}


func main()  {
	/*
		注意：
			1.添加元素时如果超越原来的cao,系统会重新分配更大的底层数据(自动判断要增加多少)  一般，如果len等于cap 那就会将cap = len * 2 来扩大
	*/
	AppendAction()
	CopyAction()
	DelAction()
	TwoSlice()
}


