context可以实现不同处理函数之间请求信息的传递，避免因为使用全局变量导致的混乱
# 创建
`ctx := context.Background()`

# 存入
`ctx = context.WithValue(ctx, "key", value)`
key必须是字符串（？）
这个方法也可以直接创建一个上下文变量

# 取出
`number := ctx.Value("key").(type)`
**由于取出来的值是个接口类型，我们需要进行类型判断**
`username := r.Context().Value("username").(string)`