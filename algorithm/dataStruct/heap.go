package dataStruct

import (
	"errors"
	"fmt"
)

/*
堆是一种完全二叉树
堆是用数组存储的，而且 0 下标不存，从 1 开始存储，建堆就是在原地通过交换位置，达到建堆的目的。
插入元素： 数据存在数组最后面，通过自下往上的堆化。
删除元素： 删除对顶的元素，然后把最后一个元素放到对顶中，然后自上而下的堆化。

堆排序：数据放入数组中，堆化为一个大顶堆，把堆顶与最后一个元素替换，则最后的元素n就是最大值，再对前面n-1数组堆化，最后数组就是有序的

多有序文件合并问题，比如10个有序文件，最终合并一个有序文件
10个文件各取一个数据，建小顶堆(堆大小为10),则堆顶元素是最小的，最先写入新文件，然后再读取新数据写入堆顶，堆化后，再把堆顶元素写入新文件。

求TopK问题：一堆无序数据，找出最大的10个数据
建立小顶堆(堆大小为10)，继续遍历数据，如果数据大于堆顶，就替换堆顶，然后堆化。如果小于堆顶，就直接丢弃，最后堆中的10数就是top10


小顶堆为例子
某个节点的值，总是小于子节点的值。
*/

type Heap struct {
	Data     []int //存储堆
	Capacity int   // 堆的容量
	Len      int   // 当前堆存放的量
}

func NewHeap() *Heap {
	return &Heap{
		Data:     make([]int, 5),
		Capacity: 5,
		Len:      0,
	}
}

// 使用堆求TopK问题，遍历所有数据，往堆插入数据，一开始放在最后，然后自下而上堆化，如果堆满了，插入数据与堆顶比较，如果小于堆顶，直接丢弃
// 如果大于堆顶，直接放到堆顶，然后自上而下堆化
func (h *Heap) Insert(value int) {
	if h.Len == h.Capacity {
		if value > h.Data[1] {
			h.Data[1] = value
			h.HeapfyDown(1)
		} else {
			return
		}
	}
	h.Len++
	h.Data[h.Len] = value
	h.HeapfyUp(h.Len)
	return
}

// 删除堆顶数据，先把最后一个数据移动到堆顶，然后自上而下的堆化
func (h *Heap) Delete() error {
	if h.Len == 0 {
		return errors.New("empty")
	}
	h.Data[1] = h.Data[h.Len]
	h.Data[h.Len] = 0
	h.Len--
	h.HeapfyDown(1)
	return nil
}

func (h *Heap) Print() {
	fmt.Println(h.Data)
}

// 自下往上的堆化
func (h *Heap) HeapfyUp(i int) {
	for {
		f := i / 2
		if f > 0 && h.Data[f] > h.Data[i] {
			h.Swap(i, f)
			i = f
		} else {
			break
		}
	}
}

// 自上而下的堆化，就是比较顶点，左子树，右子树，把最大的放到顶点。如果i是叶子节点，就不用再往下堆化
func (h *Heap) HeapfyDown(i int) {
	for {
		left := 2 * i
		right := 2*i + 1
		if left > h.Len {
			break
		}
		if h.Data[left] < h.Data[right] && h.Data[left] < h.Data[i] {
			h.Swap(left, i)
			i = left
		} else {
			break // i已经是最小，不需要交换
		}

		if h.Data[right] < h.Data[left] && h.Data[right] < h.Data[i] {
			h.Swap(right, i)
			i = right
		} else {
			break
		}
	}
}

// 交互下标为i,f的值
func (h *Heap) Swap(i, f int) {
	h.Data[i], h.Data[f] = h.Data[f], h.Data[i]
}
