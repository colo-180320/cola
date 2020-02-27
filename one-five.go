//第五节：循环
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//普通for循环
func NumFor()  {
	for i := 0;i < 5;i++{
		fmt.Printf("%s-%d\n","嗯",i)
	}
	//结果
	/*
		嗯-0
		嗯-1
		嗯-2
		嗯-3
		嗯-4
	*/
}

//九九乘法表（双层for循环）
func Num()  {
	for i:=1; i<=9;i++  {
		for q:=1; q<=i;q++  {
			fmt.Printf("%d*%d=%d ",i,q,i*q)
		}
		fmt.Println()
	}
}

//二进制：
func converToBin(n int ) string  {
	result := ""
	for ; n > 0 ; n /= 2  {
		lsb := n % 2
		result = strconv.Itoa(lsb)+result
	}
	return result
}

//慎用，一直for循环  可以作为一个消息接收，执行任务的方法
func forever()  {
	for {
		fmt.Println("abc")
	}
}

//打开文件，遍历里面的内容
func printFile(filename string)  {
	file,err := os.Open(filename)
	if err != nil{
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

//for循环之break break 只会退出最内层的循环
func forBreak()  {
	for i:=0;i<=3;i++{
		for j:=0; j<10; j++{
			if j>5{
				break
			}
			print(j)
		}
		fmt.Println()
	}
}

//for循环之continue 关键字 continue 忽略剩余的循环体而直接进入下一次循环的过程
func forContinue()  {
	for i:=0;i<=3;i++{
		for j:=0; j<10; j++{
			if j==5{
				continue
			}
			print(j)
		}
		fmt.Println()
	}
}

//for循环 之标签 for、switch 或 select 语句都可以配合标签  continue 语句指向 LABEL1，当执行到该语句的时候，就会跳转到 LABEL1 标签的位置。
//如果将 continue 改为 break，则不会只退出内层循环，而是直接退出外层循环了。
func Label()  {
	LABEL1:
		for i:=1; i<4;i++  {
			for j:=1; j<4;j++  {
				if j == 2{
					continue LABEL1
				}
				fmt.Println(j)
				//结果：1 1 1
			}
		}
}




func main()  {
	/*
		基本语法要点回顾
			1.for,if后面的条件是没有括号的
			2.if条件里也可以定义变量
			3.没有while
			4.switch 不需要 break,也可以直接switch多个条件
	*/

	NumFor()
	Num()
	fmt.Println(converToBin(13))
	printFile("abc.txt")
	forBreak()
	forContinue()
	Label()
}
