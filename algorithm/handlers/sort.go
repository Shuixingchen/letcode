package handlers

import (
	"fmt"
)

/*
冒泡排序，插入排序，选择排序  O(n2), 都是原地排序

插入排序：就是把数据分成两部分，左边部分为排序好的，右边部分为未排序的。每次从右边取出一个元素 s [j], 和左边排序好的元素依次比较。
然后合适的地方插入。

归并排序和快速排序。O(nlogn),这两种排序算法适合大规模的数据排序

归并排序：把数据分左右两部分，分别对左右部分排序，然后再把两个列表合并起来。用到了递归的思维

桶排序：适合外部排序，就是数据存储在外部磁盘中，数据量比较大，内存有限，无法将数据全部加载到内存中。把数据写入m个有序的桶，每个桶都有一个范围。
然后对桶内的数据分配排序。

*/

func Bubble() {
	list := []int{2, 76, 3, 4}
	l := len(list)
	for i := 0; i < l; i++ { //一共做l次冒泡，每次冒泡都会把最大的数移动到最右边
		flag := false
		for j := 0; j < l-i-1; j++ {
			if list[j] > list[j+1] {
				list[j+1], list[j] = list[j], list[j+1]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
	fmt.Println(list)
}

func InsertSort() {
	list := []int{2, 76, 3, 4}
	l := len(list)
	for i := 1; i < l; i++ {
		value := list[i]
		j := i - 1
		// 在左侧寻找value插入的位置
		for ; j >= 0; j-- {
			if list[j] > value {
				list[j+1] = list[j]
			} else {
				break
			}
		}
		list[j+1] = value
	}
	fmt.Println(list)
}

func MergeSort(list []int) []int {
	l := len(list)
	if l < 2 {
		return list
	}
	mid := l / 2
	left := MergeSort(list[:mid])
	right := MergeSort(list[mid:])
	return merge(left, right)
}

// 合并两个有序的列表
func merge(left, right []int) []int {
	lleft := len(left)
	lright := len(right)
	res := make([]int, 0) //申请了新空间来保存，所以不是原地排序
	l, r := 0, 0
	for l < lleft && r < lright {
		if left[l] > right[r] {
			res = append(res, right[r])
			r++
		} else {
			res = append(res, left[l])
			l++
		}
	}
	if l < lleft {
		res = append(res, left[l:]...)
	}
	if r < lright {
		res = append(res, right[r:]...)
	}
	return res
}

func QuickSort() {

}
