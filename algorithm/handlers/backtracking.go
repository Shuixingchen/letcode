package handlers

import "fmt"

/*
回溯算法
全排列问题,给一个列表，找出所有的排列组合
*/
func Permute(list []int) {
	res := make([][]int, 0)
	path := make([]int, 0) // 记录已经使用过的元素
	BackTrace(list, path, &res)
	fmt.Println(res)
}

func BackTrace(list []int, path []int, res *[][]int) {
	// 可选择列表为空，说明到这里就结束了
	if len(path) == len(list) {
		p := make([]int, len(path))
		copy(p, path) // 必须使用深拷贝
		*res = append(*res, p)
		return
	}
	// 遍历所有选择
	for _, val := range list {
		if IsContain(path, val) {
			continue
		}
		path = append(path, val)
		BackTrace(list, path, res)
		path = path[:len(path)-1]
	}

}

func IsContain(path []int, val int) bool {
	for _, i := range path {
		if i == val {
			return true
		}
	}
	return false
}
