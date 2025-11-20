Marshal和UnMarshal已经过时了，现在向我们走来的是`json.NewEncoder(w).Encode()`

# 序列化
`json.NewEncoder(w io.Writer).Encode(v interface{}) error`
w只管填`http.ResponseWriter`就对了
v 是需要序列化的数据
# 反序列化
`json.NewDecoder(r io.Reader).Decode(v interface{}) error`
用来解析请求体，即r是`r.Body`，v是用来储存反序列化数据的，必须是响应的指针类型