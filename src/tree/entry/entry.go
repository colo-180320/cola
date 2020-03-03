//第三节：包与封装
package main

import (
	"fmt"
	"tree"
)

//第四节内容：扩展已有类型
type myTreeNode struct {
	node *tree.TreeNode
}

func (myNode *myTreeNode) postOrder()  {
	if myNode == nil || myNode.node == nil{
		return
	}
	left := myTreeNode{myNode.node.Left}
	right := myTreeNode{myNode.node.Right}
	left.postOrder()
	right.postOrder()
	myNode.node.Print_One()
}
/*
	注意：
		不论地址还是结构本身，一律使用 .(点) 来访问成员
	包：
		为结构定义的方法必须放在同一个包内
*/
func main()  {
	var Root tree.TreeNode
	Root = tree.TreeNode{
		Value: 3,
		Left:  &tree.TreeNode{
			Value: 0,
			Left:  nil,
			Right: nil,
		},
		Right: &tree.TreeNode{
			Value: 5,
			Left:  nil,
			Right: nil,
		},
	}
	Root.Right.Left = new(tree.TreeNode)
	Root.Left.Right = tree.CreateNode(2)
	fmt.Println(Root)

	/*Nodes := []TreeNode{
		{Value:3},
		{},
		{6,nil,&Root},
	}*/
	Root.Print_One()  // 结果：3
	Root.Right.Left.Print_One()  // 结果：0
	Root.Right.Left.SetValue(6)
	Root.Right.Left.Print_One() // 结果：6
	tree.Print_Two(Root)        //结果：3  [但是不可以这样子使用：Print_Two(oot.Right.Left)]
	tree.SetValueTwo(&Root,9)   //注意，这里要穿指针
	tree.Print_Two(Root)        // 结果 9
	fmt.Println("---Traverser---")
	/*
		结构图：
					9
			0				5
				2		6
	*/
	Root.Traverser()  //结果：9 0 2 5 6
	fmt.Println("---postOrder---")
	myNode := myTreeNode{&Root}
	myNode.postOrder() //结果： 2 0 6 5 9
}