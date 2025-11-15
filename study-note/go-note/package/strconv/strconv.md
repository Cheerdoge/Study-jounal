这个包提供字符串相关类型的转换
1. 数字转字符串
	1. `stronv.itoa(i int) string`返回i对应的字符串型数字
	2. `stronv.FormatFloat(f float64, fmt byte, prec int, bitSize int)`将64位浮点数换为字符串，其中fmt（b或e或f或g），prec表示精度，bitSize表示32或64位
2. 字符串转数字
	1. -`strconv.Atoi(s string) (i int, err error)` 将字符串转换为 `int` 型
	2. `strconv.ParseFloat(s string, bitSize int) (f float64, err error)` 将字符串转换为 `float64` 型
3. 字符串和字节切片互转
	无需导入`strconv`包即可使用
	1. 字符串->字节切片
		`[]byte(str string)`
	2. 字节切片->字符串
		`string(slice []byte)`