package dataStruct

import "fmt"

/*
堆栈
考察算法：括号字符串是否非法
*/
//使用切片来实现栈
type Stack struct{
	Len int
	s []interface{}
}
func NewStack() *Stack{
	return &Stack{0,make([]interface{},0)}
}
func (stack *Stack)Pop() interface{}{
	if stack.Len == 0{
		return nil
	}else{
		res := stack.s[stack.Len-1] //获取栈顶元素
		stack.s = stack.s[:stack.Len-1] //删除栈顶元素
		stack.Len--
		return res
	}
}
func (stack *Stack)Push(data interface{}){
	defer func() {
		stack.Len++
	}()
	stack.s = append(stack.s, data)
}
func (stack *Stack)Print(){
	fmt.Println(stack.s)
}

/*
判断括号字符串是否合法？
思路：
1. 遍历字符串，如果是左括号，就压入栈中，如果是右括号，就判断
2. 判断的方法，首先看栈是否为空，如果为空，则说明字符串非法。再看看 Pop () 出来的元素是否与之配对，不配对就是非法，如果配对就进行下一循环。
3. 字符串遍历完后，如果栈还是不为空，则说明字符串非法
*/
func CheckStr(s string) bool{
	stack := NewStack();
	paren_map := map[string]string{"]":"[", "}":"{", ")":"("}
	str := []rune(s)
	for i:=0;i<len(str);i++{
		char := string(str[i])
		if value, ok := paren_map[char]; !ok{ //左括号
			stack.Push(value)
		}else if stack.Len == 0 || stack.Pop() != value {
			return false
		}
	}
	return true
}