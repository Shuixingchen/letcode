package test

import (
	"fmt"
	"letcode/dataStruct"
	"testing"
)

func TestBinarySearch(t *testing.T){
	arr := []int{1,2}
	key := dataStruct.BinarySearch(arr,5)
	fmt.Println(key)
}

func TestBubble() {

}
