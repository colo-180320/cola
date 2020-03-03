//第五节：map例题 ==>  寻找最长不含重复的字符串
package main

import "fmt"

//寻找最长不含重复的字符串

//PHP 写法：
/*
	PHP:
	//查找最长的字符串：
    public function getlengstring(){
        $s = "caascsbcabc";
        if(strlen($s) ==0){
            return 0;
        }
        //把字符串转换为数组
        $arr = str_split($s);
        $num = [];
        $arr2=[];
        foreach($arr as $k=>$v){
            //判断该值是否在数组中
            if(in_array($v,$num)){
                //基类重复到这里时的数量
                $arr2[] = count($num);
                array_splice($num,0,array_search($v,$num)+1);
            }
            //数组中给值
            $num[] = $v;
            var_dump($num);
        }
        var_dump($num);

    }
*/
func FindOnlyString(a string) int {
	fmt.Println(a)
	lastOccurred := make(map[rune]int) //make(map[byte]int)
	start := 0
	maxLength := 0
	for k,v:=range []rune(a) {
		fmt.Println(k,v)
		if lastK,ok := lastOccurred[v];ok && lastK >= start{
			start = lastK + 1
		}
		if k -start + 1 > maxLength{
			maxLength = k-start+1
		}
		lastOccurred[v] = k
	}
	return maxLength
}

func main()  {
	fmt.Println(FindOnlyString("你好呀你好呀"))  //只需将字符串转为rune类型就可以使用utf-8了。
	fmt.Println(FindOnlyString("123123"))
	fmt.Println(FindOnlyString("abcabc"))
}
