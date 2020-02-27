## Go 基础知识

- benchmark 函数一般以 Benchmark 开头
- benchmark 的 case 一般会跑 b.N 次，而且每次执行都会如此
- 在执行过程中会根据实际情况 case 的执行时间是否稳定会增加 b.N 的次数以达到稳态
- 永远不用让被测试的函数出现不稳态的情况，不然 benchmark 可能跑不完

* 每一个 test 文件必须 import 一个 testing
* test 文件下的每一个 test case 均必须以 Test 开头并且符合 TestXxxx 形式，否则 go test 会直接跳过测试不执行
* test case 的入参为 t *testing.T 或 b *testing.B
* t.SkipNow()跳过当前的 test，必须放在 test case 第一行
* Go 的 test 不会保证多个 TestXxx 是顺序执行，但是通常会按顺序执行
* 如果想保证多个 test case 按照一定的顺序执行，就得使用 subTest：t.Run()
* t.Run 执行 subtests 可以做到控制 test 以规定好的顺序执行

* TestMain，一个特殊的 test case，作为 test 的入口，初始化 test，并使用 m.Run()来调用其他 tests

```go

func TestMain(m *testing.M) {
	fmt.Println("test main first")
	// 必须执行这函数，其他test case才能执行。不然其余tests都不会被执行
	m.Run()
}

func TestOrderPrint(t *testing.T) {
	t.Run("a1", func(t *testing.T) { fmt.Println("a1") })
	t.Run("a2", func(t *testing.T) { fmt.Println("a2") })
	t.Run("a3", func(t *testing.T) { fmt.Println("a3") })
}
```
