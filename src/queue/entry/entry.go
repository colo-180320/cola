package main

import (
	"fmt"
	"queue"
)
/*
	包：
		如何扩充系统类型或者别人的类型
			==》
					定义别名
					使用组合
*/
func main()  {
	q := queue.Queue{}
	q.Push(2)
	q.Push(3)
	fmt.Println(q)  //结果：[2,3]
	q.Pop()
	fmt.Println(q.IsEmpty()) //false
	q.Pop()
	fmt.Println(q)  //结果：[]
	fmt.Println(q.IsEmpty()) //true

}