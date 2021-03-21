package util

/**
回溯算法（本质就是决策树的遍历）
1.全排列
2.树的深度优先遍历
 */

/**
全排列：[1,2,3,4]
当前路径：[1,2]
选择列表：全部元素-当前路径元素
结束条件：选择列表为空
 */
func Permutations(list []int) [][]int{
	result := make([][]int,0)
	path := make([]int,0) //已走路径
	i := 0
	choice := list
	backtrace(i,choice,&result,path)
	return result
}
func backtrace(i int, choice []int, result *[][]int, path []int) {
	if len(choice) <=0 {
		path = append(path, i)
		*result = append(*result,path)
		return
	}
	for k,val:= range choice {
		newChoice := Slicedelete(choice,k)
		newPath := append(path,val)
		backtrace(val,newChoice,result,newPath)
	}
}

func Slicedelete(list []int, k int) []int{
	newS := make([]int,0)

	res := append(newS[:k],newS[k+1:]...)
	return res
}