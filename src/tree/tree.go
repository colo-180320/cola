//第一节：结构体与方法
/*
	注意：
		1.go语言仅支持封装，不支持继承和多态
		2.go语言没有class，只有struct
	结构的定义：
		type TreeNode struct {
			Value int
			Left,Right *TreeNode
		}

	值接受者 VS 指针接收者
		1.要改变内容必须使用指针接收者
		2.结构过大也考虑使用指针接收者
		3.一致性：如有指针接收者，最好都是指针接收者
		4.值接受者，是go语言特有的
		5.值/指针接受者均可接受值/指针
*/

package tree

import (
	"fmt"
)

type TreeNode struct {
	Value int
	Left,Right *TreeNode
}
//使用自定义工厂函数：返回局部变量的地址
func CreateNode(Value int) *TreeNode {
	return &TreeNode{Value:Value}
}

//打印值：方法一
func (Node TreeNode) Print_One() {
	fmt.Println(Node.Value)
}
//打印值：方法二
func Print_Two(Node TreeNode)  {
	fmt.Println(Node.Value)
}

//设置值：方法一 【注意：这个要用到指针：* 】
func (Node *TreeNode)SetValue(Value int) {
	Node.Value = Value
}
//设置值：方法二 只有使用到指针才能改变结构的内容
func SetValueTwo(Node *TreeNode,Value int)  {
	Node.Value = Value
}


