//第四节：条件语句
package main

import (
	"fmt"
	"io/ioutil"
)
//if...else  打开一个文件,并读取里面的内容
func redaFile()  {
	const FileName  = "abc.txt"
	contents,err :=ioutil.ReadFile(FileName)
	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Printf("%s\n",contents)
		//结果：
			//Hello,World!
			//你好,世界
	}
	/*
		缩略版本
		const FileName  = "abc.txt"
		if contents,err := ioutil.ReadFile(FileName);err != nil{
			fmt.Println(err)
		}else {
			fmt.Printf("%s\n",contents)
		}
	*/
}

//switch
/*
	a,b 两个int的参数
	ap  一个 stirng 类型的参数
	int 一个return 返回值是一个 int类型的
	说明：switch 会自动break 除非使用 fallthrough

*/
func grade(a,b int,op string) int {
	var result int
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		panic(fmt.Sprintf("error:%d",op))
	}
	return result
}

func main()  {
	redaFile()
	var q  = grade(2,3,"+")  //注意，调用 grand方法，只是return了值，输出的话，需要重新舒勇 下面的方法
	fmt.Println(q)
}
