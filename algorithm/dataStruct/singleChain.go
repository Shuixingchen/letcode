package dataStruct

import "fmt"

/*
单链表 Root->next
考察算法：链表反转
*/
type Snode struct {
	Next *Snode
	Data string
}

// 单链表
type SingleChain struct {
	Length int
	Root   *Snode
}

func newNode(data string) *Snode {
	return &Snode{Next: nil, Data: data}
}

func NewSingleChain(list []string) *SingleChain {
	sc := &SingleChain{Length: 0, Root: nil}
	if len(list) > 0 {
		for _, value := range list {
			sc.Append(value)
		}
	}
	return sc
}

func (c *SingleChain) Append(data string) {
	c.Length++
	node := newNode(data)
	tmp := c.Root
	if tmp == nil {
		c.Root = node
		return
	}
	// 找到最后的一个节点
	for tmp.Next != nil {
		tmp = tmp.Next
	}
	tmp.Next = node
}

func (c *SingleChain) Print() {
	tmp := c.Root
	fmt.Println(tmp.Data)
	for tmp.Next != nil {
		fmt.Println(tmp.Next.Data)
		tmp = tmp.Next
	}
}
