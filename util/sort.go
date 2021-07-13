package util

/*
排序算法

*/
func BobbleSort(arr []int) []int{
	len := len(arr)
	for i:=0; i<len; i++{
		for j:=0; j<len-1; j++{
			if arr[j]>arr[j+1] {
				arr[j],arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}
