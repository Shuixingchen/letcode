package util

/*
n皇后
思路：每一行都最多放一个皇后，每一行有n个位置，看成树的话，就是树高为n,每个节点下一层的选项有n种
回溯：保留已经存在的攻击范围（每次加一个皇后，都增加攻击范围），递归下去后，再恢复添加前的攻击范围
攻击范围：clo[],pie[],na[]，确定列后，pie = row+col固定，na = row-col固定
*/
type attack struct {
	clo map[int]bool
	pie map[int]bool
	na map[int]bool
}
func SolveNQueens(n int)  [][]string{
	res := make([][]int,0) //保存结果，一个元素代表对应行的clo
	path := make([]int,0)
	a := attack{
		make(map[int]bool,n),
		make(map[int]bool,2*n-1),
		make(map[int]bool,2*n-1),
	}
	Dfs6(a,&res,n,0,path)
	return show(&res,n)
}
func Dfs6(a attack,res *[][]int,n int,row int, path []int){
	//剪枝，n皇后要求每一行都要放一个Q,如果有一行没有放Q,就不用往下走
	if row>0 && len(path) < row-1 {
		return
	}
	if row == n { //行都扫描完了，树到底了,把这次遍历下来的path写入结果
		if len(path) == n {
			p := make([]int,n)
			copy(p,path)
			*res = append(*res,p)
		}
		return
	}
	//遍历所有选项,从第一列到第n列
	for clo:=1;clo<=n;clo++{
		//判断col列是否可以放入，如果可以就放一个q,然后修改攻击变量，递归进入下一层
		if a.clo[clo] || a.pie[clo+row] || a.na[row-clo] {
			continue
		}
		path = append(path,clo)
		a.clo[clo] = true
		a.pie[clo+row] = true
		a.na[row-clo] = true
		Dfs6(a,res,n,row+1,path)
		//恢复path和攻击范围
		path = path[:len(path)-1]
		delete(a.clo,clo) //删除map的键值对,也可以赋值为true
		delete(a.pie,clo+row)
		delete(a.na,row-clo)
	}
}
func show(res *[][]int, n int)[][]string{
	item := make([]string,n)
	for i:=0; i<n; i++ {
		item[i] = "."
	}
	show := make([][]string,0)
	for _,val:=range *res{
		strSlice := make([]string,0)
		for _,v := range val{
			p := make([]string,n)
			copy(p,item)
			p[v-1] = "Q"
			s := ""
			for _,e:=range p{
				s = s + e
			}
			strSlice = append(strSlice,s)
		}
		show = append(show,strSlice)
	}
	return show
}
