//第一节：变量定义
package main
import (
	"fmt"  //定义了包，一定要在下面代码中使用到
)
//全局变量，方法见：variableGlobal
var global_a = 12  // global_a := 12 全局变量简写 这是错误的写法
var globale_b = "Hello"
//全局变量简写：
var (
	global_a_short = 12
	globale_b_short = "Hello"
)

//方法都要先从main() 方法中开始调用

//变量初始化，不设置变量值得情况下的结果
func variableZeroValue()  {
	var a int
	var b string
	fmt.Printf("%d %q\n",a,b)
	//结果：  0 ""
}

//赋值给变量
func variableInitiaValue()  {
	var a int = 3
	var q,w int =  3,4
	var b string = "Hello World"
	fmt.Printf("%d %q\n",a,b)
	//结果 3 "Hello World"
	fmt.Println(a,b)
	//结果 3 Hello World
	fmt.Println(q,w)
	//结果 3 4
}

//自动识别变量的类型
func variableTypeDeduction()  {
	var a,b,c,d = 3,4,true,"def"
	fmt.Println(a,b,c,d)
}

//变量定义简写方式
func variableShorter()  {
	a,b,c,d :=3,4,true,"def"
	b = 5 //这里如果要更改 b 的值 要用 b=5 不能用b:=5  因为上面b 已经定义了
	q:=3 //这里可以正常将 q 变量设置为 3 简写的方式 但是不可以用 q=3 否则会报错
	fmt.Println(a,b,c,d,q,global_a)
}

//全局的变量是可以在所有的方法中使用
func variableGlobal()  {
	fmt.Println(global_a,globale_b)
}


func main()  {
	variableZeroValue()
	variableInitiaValue()
	variableTypeDeduction()
	variableShorter()
	variableGlobal()
}

