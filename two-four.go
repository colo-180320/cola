//第四节：Map
/*
	描述： [类似 redis中的hash]
		Map是一种特殊的数据结构，一种元素对(pair)的无序集合，元素对应一个key(索引)和一个value(值)
		所以这个结构也成为了关联数组或字典，这是一种能够快速寻找值得理想结构，给定key，就可以迅速找到对应的value
	var mapName map[keyType]valueType
		mapName:map的变量名
		keyType:键类型
		valueType:键对应的值类型
	提示：keyType跟valueType之间允许有空格

*/
package main

import (
	"fmt"
	"sort"
	"sync"
)


//设置map  两种赋值的方式 使用map在后面赋值  调用make设置好map 后赋值
/*
	总结：mapList := map[keyType]valueType{} <==>  mapList := make(map[keyType]valueType)  //注意是有 {}  的区别
*/
func SetMap(){
	//var mapList map[string]int  //设置一个变量 mapList  它的键类型为：string  值类型为：int
	var mapList =  map[string]int{"one":1,"two":2}
	fmt.Println(mapList["one"])  //结果：1

	//先声明在赋值 主要：要结合 make 才可以这样子去赋值
	var mapListTwo = make(map[string]int)
	mapListTwo["key1"] = 1
	mapListTwo["key2"] = 2
	fmt.Println(mapListTwo) //结果：map[key1:1 key2:2]
	fmt.Println(mapListTwo["key2"]) //结果：2

	/*
		map的容量：当map增长到容量上限的时间，如果再增加新的key对应value值时，map的大小会自动增加 1
				所以出于性能的考虑，对于大的map 或者快速增加的map,即使只知道大概的cap,也是可以提前表明的

	格式：make(map[keyType]valueType,cap)
	*/
	var mapListThree = make(map[string]int,10)
	fmt.Println(mapListThree)  //结果：map[]
}

//map的数据遍历
func ForMap()  {
	var mapList = make(map[string]int,10)
	mapList["one"] = 1
	mapList["two"] = 2
	mapList["three"] = 3
	for k,v :=range mapList{
		fmt.Printf("%s==>%d",k,v)
		fmt.Println()
	}
	/*
	 	注意：无序的,每一次的结果都是会不一样的，符合hash
			解决：每次结果都是无序的方法：先将值赋值到一个切片中，再将这个切片进行排序  方法：ForMapSet()
		结果：
			key==>value
	*/
}

//map数据输出重新排序
func ForMapSet()  {
	var mapList = make(map[string]int,10)
	mapList["one"] = 1
	mapList["two"] = 2
	mapList["three"] = 3
	var newList []string
	for k :=range mapList{
		newList = append(newList,k)
	}
	fmt.Println(newList) //无序结果
	//重新排序：
	sort.Strings(newList)
	fmt.Println(newList)  //固定结果： 【one three two]
}


//二维数组赋值：
func ForMapSetTwo()  {
	var mapList = make(map[string]int,10)
	mapList["one"] = 1
	mapList["two"] = 2
	mapList["three"] = 3
	var NewList  [][]int

	for i:=0;i<len(mapList);i++ {
		var tmp []int
		for _,v :=range mapList{
			tmp = append(tmp, v)
		}
		NewList = append(NewList, tmp)
	}
	fmt.Println(NewList)
	//无序结果结果：[[1 3 2] [1 3 2] [2 1 3]]
}

/*
	map删除指定的键值对
	注意：map无法清空map所以的元素，重新用make生成一个map即可，无序考虑垃圾回收机制效率
*/
func DelMapAction()  {
	var a = make(map[string]int,5)
	a["one"] = 1
	a["two"] = 2
	a["three"] = 3
	//删除map数据用：delete
	delete(a,"one")
	//输出最新的 map 数据
	for k,v :=range a{
		fmt.Println(k,v)
	}
}


/*
map数据的 存  取   删
解决：并发问题 ==> sync包
	sync.Map 有一下特性：
1.无序初始化，直接声明即可
2.sync.Map 不能使用 map 的方式进行取值和设置等操作，而是使用 sync,Map 的方法进行调用，Store : 存储  Load : 获取 Delete : 删除
3.使用 Range 配合一个回调函数进行遍历操作，通过回调函数返回内部遍历出来的值，Range 参数中回调函数的返回值在需要继续迭代遍历是，返回 true 否则 false
*/
func MapMore()  {
	fmt.Println("方法MapMore：")
	//先声明:
	var newList sync.Map

	//存储：Store
	newList.Store("one",1)
	newList.Store("two",2)
	newList.Store("three",3)

	//取值：Load
	fmt.Println(newList.Load("one"))  //结果：1 true

	//删值：Delete
	newList.Delete("two")
	//结果：{{0 0} {{map[] true}} map[one:0xc000088020 three:0xc000088030] 1}
	fmt.Println(newList)
}

//map取值时，可以校验是不是有这个值：
func FindAction()  {
	a := map[string]string{"one":"第一个","two":"第二个"}
	s,ok := a["one"]  //map值是可以接受到两个参数值得
	if(ok){
		fmt.Println(s)
	}else {
		fmt.Println("无key:one")
	}
}


func main()  {
	/*
		Map的操作：
		 1.创建：make(map[keyType]valueType)
		 2.获取元素：m[key]
		 3.key不存在时，获得Vluee类型的初始值
		 4.用value,ok := m[key] 来判断是否存在key
	  	 5.用delete删除一个key
		 6.使用range 遍历 key,或者遍历key,value对
		 7.不保证遍历顺序，如需顺序，需手动对key排序
	     8.使用 len 获得元素个数
	     9.高并发的map 使用包：sync 进行操作
		 10.map使用哈希表，必须可以比较相等
		 11.除了 slice map function 的内建类型都可以作为key
		 12.Stuct类型不包含上述字段，也可以作为key

	*/
	SetMap()
	ForMap()
	ForMapSet()
	ForMapSetTwo()
	DelMapAction()
	MapMore()
	FindAction()
}
