package main

import (
	"fmt"
	"retriever/mock"
	"retriever/read"
)

/*
	Go一眼中接口类型的独特之处在于他是满足隐式实现的，也就是说，我们没有必要对于给定大具体类型定义所有
满足的接口类型，简单拥有一些必须的方法就足够了

接口声明的格式:
type 接口类型名 interface{
	方法名1（参数列表1）返回值列表1
	方法名2（参数列表2）返回值列表2
	...
}
说明：
		1.接口类型名：使用type将接口定义为自定义的类型名 一班都是在单词后面添加er
		2.方法名：当方法名首字母是大写是，接口类型名也是大写的，这个方法可以被接口所在的包之外的代码访问
		3.参数列表，返回值列表：参数列表和返回值列表中的参数变量是可以忽略的
*/

type Retriever interface {
//	方法名    参数    返回值
	Get(url string) (content string)
}

func download(r Retriever) string {
	return r.Get("http://www.imooc.com")
}
func main()  {
	var r Retriever
	r = mock.Retriever{"内部Contents结构体"}
	r = read.Retriever{}  //这个必须 r = 不能用其他的代替 因为上面 var r Retriever
	fmt.Println(download(r))
	fmt.Println(download(r))
}
