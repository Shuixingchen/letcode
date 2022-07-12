package handlers

import "fmt"

/*
二分查找：二分查找依赖的是顺序表结构，简单点说就是数组。 O(logn)
*/

func BinarySearch(list []int, target int) int {

	left := 0
	right := len(list) - 1
	for left <= right {
		mid := (left + right) / 2
		if list[mid] == target {
			fmt.Println(mid)
			return mid
		}
		if list[mid] > target {
			right = mid - 1
		}
		if list[mid] < target {
			left = mid + 1
		}
	}
	fmt.Println(-1)
	return -1
}
