package dataStruct

/*
二叉查找树(有顺序的二叉树)
*/

type binarynode struct {
	value int
	left *binarynode
	right *binarynode
}
type BinaryTree struct {
	Root *binarynode
}

func NewBinaryNode(value int) *binarynode{
	return &binarynode{
		value : value,
		left : nil,
		right : nil,
	}
}
func NewBinaryTree() *BinaryTree{
	return &BinaryTree{
		Root: NewBinaryNode(-1),
	}
}

func (node *binarynode) Add(value int) {
	if node.value == -1 {
		node.value = value
		return
	}
	if node.value > value {
		node.left.Add(value)
		return
	}
	if node.value < value {
		node.right.Add(value)
		return
	}
}
//中序遍历,根在中
func (node *binarynode) Print(res *[]int) {
	if node.left != nil {
		node.left.Print(res)
	}
	*res = append(*res,node.value)
	if node.right != nil {
		node.right.Print(res)
	}
	return
}

