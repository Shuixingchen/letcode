package test

import (
	"fmt"
	"letcode/dataStruct"
	"testing"
)

func TestBinarySearch(t *testing.T){
	arr := []int{3,5,2,6,8,9,1}
	key := dataStruct.BinarySearch(arr,5)
	fmt.Println(key)
}
