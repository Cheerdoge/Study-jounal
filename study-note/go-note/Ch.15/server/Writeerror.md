需要辅助响应函数返回给客户端成功连接信息等，即设置响应包中的一些东西

# 设置响应头
`w.Header().Set()`进行设置，注意它是响应写的值的方法，不是指针
* `Header()`返回一个`map[string][]string`
* `Set(key, value)`设置响应头中某个字段，没有则添加，有则覆盖，如`Set("Content-type", "application/json")`

# 设置HTTP状态码
`w.WriteHeader()`其中`w`是`*http.ResponseWriter`
在响应出错时需要，括号内可以直接写相应的代码，也可以设置一个变量，根据辅助响应函数在不同情况被传入的参数而变化，
	http中已经设置常量如下，使用`http.StatusCOntinue`
		```go
		StatusContinue           = 100 
		StatusSwitchingProtocols = 101 
		StatusProcessing         = 102 
		StatusEarlyHints         = 103 
		StatusOK                   = 200 
		StatusCreated              = 201 
		StatusAccepted             = 202 
		StatusNonAuthoritativeInfo = 203 
		StatusNoContent            = 204 
		StatusResetContent         = 205
		StatusPartialContent       = 206 
		StatusMultiStatus          = 207 
		StatusAlreadyReported      = 208 
		StatusIMUsed               = 226 
		StatusMultipleChoices  = 300 
		StatusMovedPermanently = 301 
		StatusFound            = 302 
		StatusSeeOther         = 303 
		StatusNotModified      = 304 
		StatusUseProxy         = 305 
		StatusTemporaryRedirect = 307 
		StatusPermanentRedirect = 308 
		StatusBadRequest                   = 400 
		StatusUnauthorized                 = 401 
		StatusPaymentRequired              = 402 
		StatusForbidden                    = 403 
		StatusNotFound                     = 404 
		StatusMethodNotAllowed             = 405 
		StatusNotAcceptable                = 406 
		StatusProxyAuthRequired            = 407 
		StatusRequestTimeout               = 408 
		StatusConflict                     = 409 
		StatusGone                         = 410 
		StatusLengthRequired               = 411 
		StatusPreconditionFailed           = 412 
		StatusRequestEntityTooLarge        = 413 
		StatusRequestURITooLong            = 414 
		StatusUnsupportedMediaType         = 415 
		StatusRequestedRangeNotSatisfiable = 416 
		StatusExpectationFailed            = 417 
		StatusTeapot                       = 418 
		StatusMisdirectedRequest           = 421 
		StatusUnprocessableEntity          = 422 
		StatusLocked                       = 423 
		StatusFailedDependency             = 424 
		StatusTooEarly                     = 425 
		StatusUpgradeRequired              = 426 
		StatusPreconditionRequired         = 428 
		StatusTooManyRequests              = 429 
		StatusRequestHeaderFieldsTooLarge  = 431 
		StatusUnavailableForLegalReasons   = 451 
		StatusInternalServerError           = 500
		StatusNotImplemented                = 501 
		StatusBadGateway                    = 502 
		StatusServiceUnavailable            = 503 
		StatusGatewayTimeout                = 504 
		StatusHTTPVersionNotSupported       = 505 
		StatusVariantAlsoNegotiates         = 506 
		StatusInsufficientStorage           = 507 
		StatusLoopDetected                  = 508 
		StatusNotExtended                   = 510 
		StatusNetworkAuthenticationRequired = 511 
		```

# json序列化
如果在设置响应头时设置了结构为json，就要对内容进行序列化，使用一个新的函数，点[这里](./json-Decoder.md)

这是一个示例
```go
func WriteError(w http.ResponseWriter, message string, code string) {
//设置响应头中使用的类型
w.Header().Set("Content-type", "application/json")
//设置http状态码，因为是返回错误
w.WriteHead(code)
//对返回的响应包进行序列化
json.NewEncoder(w).Encode(Respone{
//具体看响应结构体怎么设置的

//如果式成功响应还需要返回可能的Data
})
}
```

如果是成功返回函数，传入参数应是一个空接口，方便各种类型的数据传入