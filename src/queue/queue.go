package queue

type Queue []int

//添加元素
func (q *Queue) Push(v int)  {
	*q = append(*q,v)
}

//移除第一个元素
func (q *Queue) Pop() int  {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

//判断是否为空
func (q *Queue) IsEmpty() bool  {
	return len(*q) == 0
}