package algorithm

import "fmt"

/*
单链表
考察算法：链表反转
*/
type node struct {
	value interface{}
	next  *node
}

func newNode(value interface{}) *node {
	return &node{value: value, next: nil}
}

type chain struct {
	root *node
	len  int
}

func NewChain() *chain {
	return &chain{root: newNode(nil), len: 0}
}

//在尾部节点增加一个节点
func (c *chain) Insert(value interface{}) {
	node := c.root
	if c.len == 0 {
		node.value = value
		c.len = 1
		return
	}
	for node.next != nil {
		node = node.next
	}
	node.next = newNode(value)
	c.len++
	return
}

func (c *chain) Print() {
	m := make(map[int]interface{}, c.len)
	node := c.root
	i := 0
	m[i] = node.value
	for node.next != nil {
		i++
		node = node.next
		m[i] = node.value
	}
	fmt.Println(m)
}

//反转链表
func ReverseList(c *chain) *chain {
	var pre *node
	cur := c.root
	for cur != nil {
		cur.next, pre, cur = pre, cur, cur.next
	}
	n := NewChain()
	n.root = pre
	return n
}
