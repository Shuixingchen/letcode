package main

import (
	"fmt"
	"letcode/util"
)

func main(){
	aa := [][]int{
		[]int{1},
		[]int{0},
	}
	res := util.UniquePathsWithObstacles(aa)
	fmt.Println(res)
}
