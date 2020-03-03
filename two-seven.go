//第七节：列表->list
/*
	描述：
		列表是一种非连续的存储容器，由多个节点组成，节点通过一些变量记录彼此之间的关系，列表有多种实现方法，如单链表，双链表等

		由：container/list 包来实现的，内部的实现原理是双链表，列表能够高效的进行任意位置的元素插入和删除操作

	注意：
		列表与切片和map不同的是，列表并没有具体元素类型的限制，因此：【列表的元素可以是任意类型】 问题：给列表放入一个 interface{}类型的值
		后，如果要将interface{}转化为其他类型将会发生宕机

*/
package main

import (
	"container/list"
	"fmt"
)

//初始化列表
func ListOne()  {
	/*
		变量名 := list.New()
		队前：变量名.PushFront("值")
		队后：变量名.PushBack("值")
		某个值前面增加：变量名.InsertBefore("添加值",变量名.PushBack("某个值"))
		某个值后面增加：变量名.InsertAfter("添加值",变量名.PushBack("某个值"))
		移除摸个值：变量名.Remove("值")
	*/


	OneList := list.New()
	//队列前插入元素：
	OneList.PushFront("First")
	//队列后插入元素：
	OneList.PushBack("Two")
	//在Five后面添加元素：Three
	GetTwoElement := OneList.PushBack("Five")
	OneList.InsertAfter("Three",GetTwoElement)
	//在Two之前添加元素：Four
	OneList.InsertBefore("Four",GetTwoElement)
	//删除掉Two:  不删除之前的列表数据：[First Two Four Five Three]
	OneList.Remove(GetTwoElement)  //删除后： [First Two Four Three]

	//列表变量：
	for i:= OneList.Front();i !=nil ;i = i.Next()  {
			fmt.Println(i.Value)
	}

	one := OneList.Front()  //指向的是  列表:OneList的头的第一个值：First
	two := one.Next()
	three := two.Next()
	four := three.Next()
	fmt.Println(one)
	fmt.Println(two)
	fmt.Println(three)
	fmt.Println(four)
}


func main()  {
	ListOne()
}
