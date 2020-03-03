//第六节：字符与字符串
package main

import (
	"fmt"
	"unicode/utf8"
)

//中文字符串的长度识别
/*
	注意：rune 相当于 go 的 char
		使用 range 可以遍历 pos,rune
		使用 uft8.RuneCountInString 获取字符串的长度
		使用len是获取字节的长度
		使用[]byte 是获取字节
*/

func ChinaAction()  {
	OneString := "你好，世界!" //UTF-8的格式
	fmt.Println(len(OneString))  //16个长度
	fmt.Printf("%X\n",[]byte(OneString)) //E4BDA0E5A5BDEFBC8CE4B896E7958C21
	fmt.Printf("%X\n",[]rune(OneString)) //[4F60 597D FF0C 4E16 754C 21]
	fmt.Printf("%d\n",len([]rune(OneString))) // 6个长度
	fmt.Println(utf8.RuneCountInString(OneString))  // 6个长度

	bytes := []byte(OneString)
	for len(bytes) > 0 {
		ch,size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c",ch) //你好，世界!
	}
	fmt.Println()

	for i,ch := range []rune(OneString){
		fmt.Printf("(%d %c)",i,ch)  //(0 你)(1 好)(2 ，)(3 世)(4 界)(5 !)
	}
}

func main()  {
	ChinaAction()
}