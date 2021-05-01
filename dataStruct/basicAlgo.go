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
	right := len(data)-1
	left := 0
	for left <= right {
		mid := (left+right)/2
		if data[mid] == target {
			return mid
		}
		if data[mid] > target{
			right = mid-1
		}
		if data[mid] < target {
			left = mid+1
		}
	}
	return -1
}

/*
冒泡排序
*/
func Bubble(data []int) {
	l := len(data)
	for i:=0; i<l; i++{
		for j:=0; j<l-1; j++ {
			if data[j] > data[j+1] {
				data[j+1],data[j] = data[j],data[j+1]
			}
		}
	}
	return
}


