package dataStruct

import "fmt"

//单向链表  每个节点包含数据域和下一个节点的指针

type node struct {
	value interface{}
	next *node
}

func newNode(value interface{}) *node{
	return &node{value:value,next:nil}
}

type chain struct {
	root *node
	len int
}
func NewChain() *chain{
	return &chain{root:newNode(nil),len:0}
}

//在尾部节点增加一个节点
func (c *chain)Insert(value interface{}) {
	node := c.root
	if c.len == 0 {
		node.value = value
		c.len = 1;
		return
	}
	for node.next != nil {
		node = node.next
	}
	node.next = newNode(value)
	c.len++
	return
}

func (c *chain)Print(){
	m := make(map[int]interface{}, c.len)
	node := c.root
	i := 0
	m[i] = node.value
	for node.next != nil{
		i++
		node = node.next
		m[i] = node.value
	}
	fmt.Println(m)
}

