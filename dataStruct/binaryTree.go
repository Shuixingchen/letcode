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

/*
向二叉搜索树添加元素，小的放左边，大的放右边
*/
func (node *binarynode) Add(value int) {
	if node.value == -1 {
		node.value = value
		return
	}
	if node.value > value {
		if node.left != nil {
			node.left.Add(value)
		}else{
			node.left = NewBinaryNode(value)
		}
		return
	}
	if node.value < value {
		if node.right != nil {
			node.right.Add(value)
		}else{
			node.right = NewBinaryNode(value)
		}
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

