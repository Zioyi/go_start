
1.Go语言变量的声明

Go语言的基本类型:
	bool
	string
	int int8 int16 int32 int64
	uint uint8 uint16 uint32 uint64 uintptr
	byte // uint8的别名
	rune // int32的别名 代表一个Unicode码
	float32 float64
	complex64 complex128

变量的声明形式:
	标准格式:
		var 变量名 变量类型	
			- var a int
			- var a, b int * // a和b都是整型指针
	批量格式:
		var (
			a int
			b string
			c []float32
			d func() bool
			e struct {
				x int
			}
		)
	简短格式:
		名字 := 表达式
		限制：
			至少存在一个未被定义过的变量
			不能提供函数类型
			只能用在函数内部
		i, j := 0, 1
