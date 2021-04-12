package dataStruct

//基础算法

/*
二分查找，给一个有序的数组，找到目标元素
时间复杂度 logn
*/
func BinarySearch(data []int, target int) (int){
	if len(data) < 0 {
		return -1
	}
	right := len(data)
	left := 0
	for i:=0; i<len(data);i++{
		mid := (left+right)/2
		if data[mid] == target {
			return mid
		}
		if data[mid] > target{
			right = mid
		}
		if data[mid] < target {
			left = mid
		}
	}
	return -1
}


