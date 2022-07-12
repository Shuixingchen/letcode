package handlers

import "fmt"

/*
滑动窗口——使用双指针，left,right控制窗口
在 S中找出包含 T所有字母的最短子串
*/

// 默认字符都是英文字符，所以一个字符占一个字节
func FindSub(s, t string) {
	tmap := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		tmap[t[i]]++
	}
	var left, right int
	count := len(t)
	for right < len(s) {
		if v, ok := tmap[s[right]]; ok && v > 0 { //right的字符是在t中
			tmap[s[right]]--
			count--
		}
		for count == 0 { // 窗口位置已经包含了所有的t字符,就移动left，直到找到最短的窗口
			fmt.Println(left)
		}
	}
}
