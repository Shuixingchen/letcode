package dataStruct

import "fmt"

/*
双向循环链表
Root->node1->node2->Root
*/
type DCnode struct {
	Data string
	Pre  *DCnode
	Next *DCnode
}
type DoubleChain struct {
	Length int
	Root   *DCnode
}

func newDCnode(data string) *DCnode {
	return &DCnode{Data: data, Pre: nil, Next: nil}
}
func NewDoubleChain(list []string) *DoubleChain {
	dc := &DoubleChain{Length: 0, Root: nil}
	for _, value := range list {
		dc.Append(value)
	}
	return dc
}

func (d *DoubleChain) Append(data string) {
	d.Length++
	node := newDCnode(data)
	flag := d.Root
	if flag == nil {
		d.Root = node
		return
	}
	// 找到最后一个节点
	for flag.Next != nil {
		flag = flag.Next
	}
	flag.Next = node
	node.Pre = flag
}

func (d *DoubleChain) Print() {
	flag := d.Root
	fmt.Println(flag.Data)
	for flag.Next != nil {
		flag = flag.Next
		fmt.Println(flag.Data)
	}
}
