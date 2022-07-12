package handlers

import (
	"letcode/algorithm/dataStruct"
)

/*
链表的算法

1.单链表反转（中间变量保存preNode, flag, nextNode）
2.链表中环的检测
3.两个有序的链表合并(归并排序会用到)
4.删除链表倒数第 n 个结点
求链表的中间结点
*/

func SingleChainReverse() {
	sc := dataStruct.NewSingleChain([]string{"1", "2", "3"})
	flag := sc.Root
	var preNode *dataStruct.Snode
	for flag.Next != nil {
		nextNode := flag.Next
		flag.Next = preNode
		preNode = flag
		flag = nextNode
	}
	// 最后一个节点
	flag.Next = preNode
	sc.Root = flag
	sc.Print()
}

func MergeChain() {

}
