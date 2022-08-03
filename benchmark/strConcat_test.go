package benchmark

import "testing"

/*
benchmark 和普通的单元测试用例一样，都位于 _test.go 文件中。
函数名以 Benchmark 开头，参数是 b *testing.B。和普通的单元测试用例很像，单元测试函数名以 Test 开头，参数是 t *testing.T。
测试函数要放在for循环中，循环执行b.N次。

go test -bench .  # 运行当前目录的bench测试用例
go test -bench='Fib$' -cpu=2,4 .  # 运行当前目录Fib结尾的测试用例, cpu设置为2和4执行
go test -bench=. -benchmem  # 显示内存申请情况
*/

func benchmarkPlusConcat(b *testing.B) {
	str := randomString(10)
	for i := 0; i < b.N; i++ {
		plusConcat(1000, str)
	}
}

func benchmark(b *testing.B, f func(int, string) string) {
	var str = randomString(10)
	for i := 0; i < b.N; i++ {
		f(10000, str)
	}
}

func BenchmarkPlusConcat(b *testing.B)    { benchmark(b, plusConcat) }
func BenchmarkSprintfConcat(b *testing.B) { benchmark(b, sprintfConcat) }
func BenchmarkBuilderConcat(b *testing.B) { benchmark(b, builderConcat) }
func BenchmarkBufferConcat(b *testing.B)  { benchmark(b, bufferConcat) }
func BenchmarkByteConcat(b *testing.B)    { benchmark(b, byteConcat) }
