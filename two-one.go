//第一节：数组
package main

import (
	"fmt"
)

func SetArray()  {
	/*
		注意：
			1.如果没有定义值的数组，是不可以用简写方式来定义的  必须 var array [3]int
			2.定义已知值得二维数组，要写明多少位数，否则编译错误
			3.数量要写在类型前面
	*/
	var arr_one [5]int  //不定义具体内容，五个位置的数组
	arr_two := [3]int{1,3,5}   //明确知道 3个数组的值
	arr_three := [...]int{2,4,6,8,10} //知道一些值，但是不知道具体多少位
	var arr_four  [3][4]int  //两位数组，如果用 arr_four ：= [3][4]int 来定义，会报错
	arr_five := [3][4]int{
		{0,1,2,3},
		{1,2,3,4},
		{2,3,4,5},
	}

	fmt.Println(arr_one) //结果：[0,0,0,0,0]
	fmt.Println(arr_two) //结果：[1,3,5]
	fmt.Println(arr_three) //结果：[2,4,6,8,10]
	fmt.Println(arr_four) //结果：[2,4,6,8,10]
	fmt.Println(arr_five) //结果：[[0,1,2,3] [1,2,3,4] [2,3,4,5]]
}

//数组遍历：range
func ToRange()  {
	arr_two := [3]int{1,3,5}
	for i,v:=range arr_two {
		fmt.Println(arr_two[i])
		fmt.Println(v)  //这个 v 就是 值 等于当前的 arr_two[i]
	}
	/*
	第一种：
		for i:=range arr_two {
			fmt.Println(arr_two[i])
		}
	第二种：
		for i,_:=range arr_two {
			fmt.Println(arr_two[i])
		}
	第三种：
		for i,v:=range arr_two {
			fmt.Println(arr_two[i])
			fmt.Println(v)  //这个 v 就是 值 等于当前的 arr_two[i]
		}
	第四种：
		for _,v:=range arr_two {
			fmt.Println(v)  //这个 v 就是 值
		}
	*/
}

//数组遍历：for
func ToFor()  {
	arr_two := [3]int{1,3,5}
	for i:=0; i<len(arr_two);i++  {
		fmt.Println(arr_two[i])
	}
}

//拷贝数组
func parintArray_One(arr [5]int)  {
	for _,v :=range arr{
		fmt.Println(v)
	}
	arr[0] = 100
}

//指针数组
func parintArray_Two(arr *[5]int)  {
	for _,v :=range arr{
		fmt.Println(v)
	}
	arr[0] = 50


}

func main()  {
	/*
		数组遍历：建议使用 range方法  意义明确，美观
		数组的值，是值类型
		在go语言中一般不直接使用数组
	*/
	SetArray()
	ToRange()
	//ToFor()
	arr_one := [...]int{1,2,3,4,5}
	parintArray_One(arr_one) //这里方法中已经明确说了 5位数的参数数组，所以如果不符合要求，会报错！
	fmt.Println(arr_one) //即使 函数中：printArray已经修改了值，但是因为是值传递，所以不会对arr_one造成影响的 [拷贝数组]

	//可以更改值：
	/*
		注意：
			参数：&arr_one 是  数组arr_one的地址
			函数中parintArray_Two 用 arr *[5]int 就是 arr_one 值
			通过在函数中更改数组的值，是会对整个数组产生印象的
	*/
	parintArray_Two(&arr_one) //结果： 1,2,3,4,5
	fmt.Println(arr_one)	  //结果： 50,2,3,4,5
}
