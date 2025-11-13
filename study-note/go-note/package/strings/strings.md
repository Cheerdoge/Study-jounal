1. 前缀后缀 Has Pre/Suf fix
	`strings.HasPrefix(s, prefix string)bool`检查s是否以prefix开头
	Pre改为Suf检查结尾
2. 包含 Contains
	`strings.Contains（s, substr string) bool`检查是否包含
3. 判断字符串或字符的位置 Index/LastIndex
	对，字符串也可以当切片或者数组一样有索引
	str在s中第一次出现的索引，-1表示不包含
	`string.Index(s, str string) int`
	如果要最后出现的索引，函数名前Last
	如果是非ASCII编码的字符，函数名最后加Rune
4. 替换 Replace
	将前n个字符串old替换为new，n=-1则全部替换
	`strings.Replace(str, old, new, n) string`
5. 统计出现次数 Count
	str在s中出现的非重叠次数（如gg在gggg出现四次）
	`strings.Count(s, str string)` int
6. 重复输出 Repeat
	重复输出n次字符串str
	`strings.Repeat(str, n int) string`
7. 修改大小写 To Lower/Uper
8. 修剪 Trim
	剔除开头结尾空白符号`strings.TrimSpace(s)`
	剔除指定的字符`strings.Trim(s, "cut")`
	如果只想剔除开头或结尾，函数名加Left或Right
9. 分割
	`strings.Fields(s)` 将会利用 1 个或多个空白符号（空格）来作为动态长度的分隔符将字符串分割成若干小块，并返回一个 slice，如果字符串只包含空白符号，则返回一个长度为 0 的 slice。
	`strings.Split(s, sep)` 用于自定义分割符号来对指定字符串进行分割（要求字符串内包含sep），同样返回 slice。
	因为这 2 个函数都会返回 slice，所以习惯使用 for-range 循环来对其进行处理
10. 拼接切片 Join
	将一个切片用分割符号拼接组成
	`Strings.Join(s1 []string, sep string)`
11. **读取 NewReader**
	`strings.NewReader(str)`会生成一个`Reader`并读取内容，返回指向该`Reader`的指针