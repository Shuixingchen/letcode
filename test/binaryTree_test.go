package test

import (
	"fmt"
	"letcode/dataStruct"
	"testing"
)

func TestTree(t *testing.T) {
	data := []int{21,5,76,2,6}
	tree := dataStruct.NewBinaryTree()
	for _,value := range data {
		tree.Root.Add(value)
	}
	res := make([]int,0)
	tree.Root.Print(&res)
	fmt.Println(res)
}
