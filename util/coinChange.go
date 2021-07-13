package util

import (
	"sort"
)

/*
https://leetcode-cn.com/problems/coin-change/
贪心法，最优解就是每一步都尽量用上最大面值的硬币，硬币数量才是最少的情况
*/
type intSlice []int

func (p intSlice) Len() int           { return len(p) }
func (p intSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p intSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func CoinChange(coins []int, amount int) int {
	res := 0
	var dd intSlice
	dd = coins
	sort.Sort(sort.Reverse(dd))

	for _,val := range dd{
		if amount ==0 {
			break
		}
		if n := amount/val;n > 0 {
			res = res + n
			amount = amount - n*val
		}

	}
	if amount > 0 {
		return -1
	}else{
		return res
	}

}