package programMod

/*
map,reduce,filter处理数组
map:处理的元素和返回的元素一一对应。
reduce:返回一个计算结果
filter:只返回符合条件的元素
*/
func MapCase(arr []string, fn func(item string) string) []string {
	var newArr = []string{}
	for _, item := range arr {
		newArr = append(newArr, fn(item))
	}
	return newArr
}

func ReduceCase(arr []string, fn func(item string) int) int {
	var res int
	for _, item := range arr {
		res = res + fn(item)
	}
	return res
}

func FilterCase(arr []string, fn func(item string) bool) []string {
	var newArr = []string{}
	for _, item := range arr {
		if fn(item) {
			newArr = append(newArr, item)
		}
	}
	return newArr
}

//如果数组元素是指针，可以原地处理
func MapCaseLocal(arr []*string, fn func(item *string)) {
	for _, item := range arr {
		fn(item)
	}
}
