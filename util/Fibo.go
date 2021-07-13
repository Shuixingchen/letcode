package util

/*
动态规划问题
斐波那契数列
*/

func Fibo(n int) int{
	res := make([]int,n+1)
	res[0] = 0
	res[1] = 1
	if n < 2{
		return n
	}
	for i:=2;i<=n;i++{
		res[i] = res[i-1]+res[i-2]
	}
	return res[n]
}

//不同路径问题
/*
思路：表格用table[row][col]二维数组标识一个坐标，值表示该点到达目标的路径数
目标的坐标table[0][0],起点的坐标table[n,m]
dp方程为：
table[i][j] = table[i+1][j]+table[i][j+1]
https://leetcode-cn.com/problems/unique-paths-ii/
*/
func UniquePaths(m int, n int) int {
	row,column := m,n
	table := make([][]int, row)
	for i:=0;i<row;i++{
		table[i] = make([]int, column)
		for j:=0;j<column;j++{
			if i==0 || j==0 {
				table[i][j] = 1
			}else{
				table[i][j] = table[i-1][j]+table[i][j-1]
			}
		}
	}
	return table[row-1][column-1]
}

