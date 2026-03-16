# 一、创建客户端
## 创建客户端变量
Client被声明为一个结构体，声明一个Client类型变量来使用Client的方法: `clent := &http.Client`
	可能用到的成员
	* Jar CookieJar
		设置为nil，全程忽略cookie
			CookieJar
				接口，包含`SetCookies`和`Cookies`两种方法
				其中SetCookies管理在回复收到的cookie，选择是否储存cookie
	* Timeout time.Duration
		指定请求时间限制

# 创建请求
`req := http.NewRequest("GET", "www.xxxx", body/nil)`创建请求
1. 设置Cookie
	`req.AddCookie(cookie)`添加cookie
	Cookie是一个结构体
2. 设置请求头
	`req.Header.Add("key", "value")`

# 发送请求
`client.Do(req)`发送请求，返回响应和err

# 检查状态码
resp内成员StatusCode可以获取，对其进行处理

# 读取响应体
读取时必须defer关闭响应体的方法`defer resp.Body.Close()`
1. 解析为能够阅读的源码
	`ioutil.ReadAll(resp.Body)`
	返回的值为字节切片，需要进行类型转换
2. 快速寻找需要的字段
	Ctrl+F
	- `action=` - 找到表单提交目标
	- `name="lt"` - 找到一次性令牌
	- `name="execution"` - 找到执行参数
	- `hidden` - 找到所有隐藏字段
3. 获取需要的内容
	使用正则表达式: [笔记](study-note/go-note/package/regexp/regexp.md)
	

# 读取响应头
响应头是元素为字符串切片的字典
1. Values方法返回key下的所有元素，不是切片副本

# 设置表单数据
url包内的结构体Values`map[string][]string`，用于查询的参数和表单属性
方法有Get、Set、Add、Del、Encode