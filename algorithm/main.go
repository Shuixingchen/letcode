package main

import (
	"fmt"
	"letcode/algorithm/dataStruct"
	"letcode/algorithm/handlers"
)

func main() {
	// SingleChain()
	// DoubleChain()
	// handlers.SingleChainReverse()
	// handlers.Bubble()
	// handlers.InsertSort()
	// MergeSort()
	// BinaryTree()
	// Heap()
	// handlers.Permute([]int{5, 4, 6, 2})
	// handlers.BinarySearch([]int{5, 9}, 5)
	handlers.FindSub("adfabcg", "abc")
}
func SingleChain() {
	sc := dataStruct.NewSingleChain([]string{"1", "2", "3"})
	sc.Print()
}
func DoubleChain() {
	dc := dataStruct.NewDoubleChain([]string{"11", "22", "33"})
	dc.Print()
}

func BinaryTree() {
	sc := dataStruct.NewBinaryTree()
	sc.Add(11)
	sc.Add(4)
	sc.Add(8)
	sc.Add(7)
	sc.Print()
}

func Heap() {
	h := dataStruct.NewHeap()
	h.Insert(66)
	h.Insert(62)
	h.Insert(90)
	h.Insert(1)
	h.Insert(100)
	h.Print()
}

func MergeSort() {
	res := handlers.MergeSort([]int{5, 7, 4, 2, 0, 9, 8, 7, 33, 55})
	fmt.Println(res)
}
