package tree
func (Node *TreeNode) Traverser()  {
	if Node == nil{
		return
	}
	Node.Print_One()
	Node.Left.Traverser()
	Node.Right.Traverser()
}
