## 目录
- 服务器启动与配置
- 路由与处理器
- 响应写入（ResponseWriter）
- 请求读取（*http.Request）
- JSON 读写
- 文件与静态资源
- 常见 Header 与缓存
- Cookie 与会话
- 中间件与安全
- 并发 / Context / 取消
- HTTP 客户端
- 实用小技巧
- 常用代码片段

---

## 服务器启动与配置
| 函数 / 类型                                                  | 所属 / 位置 | 作用 / 常用场景              | 注意点                                   |
| -------------------------------------------------------- | ------: | ---------------------- | ------------------------------------- |
| http.ListenAndServe(addr, handler)                       |      包级 | 在 addr 上启动 HTTP 服务（阻塞） | 例：`http.ListenAndServe(":8080", mux)` |
| http.ListenAndServeTLS(addr, certFile, keyFile, handler) |      包级 | 启动 HTTPS               | 需要证书与私钥文件                             |
| &http.Server{...} / Server.Shutdown(ctx)                 |   类型/方法 | 更细粒度配置与优雅关机            | 用于 graceful shutdown                  |
| http.TimeoutHandler(h, timeout, msg)                     |      包级 | 给单个 handler 设置超时       | 超时后自动返回 503 或自定义 msg                  |

---

## 路由与处理器
| 项目 | 说明 |
|---|---|
| http.HandleFunc(pattern, func(w,r)) | 使用 DefaultServeMux 快速注册 |
| http.Handle(pattern, handler) | 注册实现了 http.Handler 的 handler |
| http.NewServeMux() | 创建独立路由表（推荐用于复杂项目） |
| http.Handler 接口 | 自定义处理器需实现 `ServeHTTP(w, r)` |
| http.HandlerFunc | 把 `func(w,r)` 转为 Handler（常用） |

---

## 响应写入（ResponseWriter）
| 方法 | 作用 | 注意 |
|---|---|---|
| w.Header() | 返回 header map（map[string][]string） | 在写 body 之前设置 header |
| w.Header().Set(k,v) | 设置/覆盖 header | 覆盖已有值 |
| w.Header().Add(k,v) | 追加值 | 适用于允许多个值的 header（如 Set-Cookie） |
| w.Header().Get(k) | 读取 header 值 | 只读视角 |
| w.WriteHeader(code) | 显式发送状态码 | 必须在写 body 之前调用 |
| w.Write([]byte) | 写响应体并触发 header/status 发送 | 若未写 status，首次写入会隐式 200 |
| http.Error(w,msg,code) | 便捷返回错误 | 会设置 Content-Type 为 text/plain |
| http.Redirect(w,r,url,code) | 重定向（设置 Location + 状态码） | 常用 301/302/307/308 |

---

## 响应写入（ResponseWriter）

| 方法 | 作用 | 注意 |
|---|---|---|
| w.Header() | 返回 header map（map[string][]string） | 在写 body 之前设置 header |
| w.Header().Set(k, v) | 设置/覆盖 header | 覆盖已有值 |
| w.Header().Add(k, v) | 追加值 | 适用于允许多个值的 header（如 Set-Cookie） |
| w.Header().Get(k) | 读取 header 值 | 只读视角 |
| w.WriteHeader(code) | 显式发送状态码 | 必须在写 body 之前调用 |
| w.Write([]byte) | 写响应体并触发 header/status 发送 | 若未写 status，首次写入会隐式 200 |
| http.Error(w, msg, code) | 便捷返回错误 | 会设置 Content-Type 为 text/plain |
| http.Redirect(w, r, url, code) | 重定向（设置 Location + 状态码） | 常用 301/302/307/308 |

---

## JSON 读写
| 场景 | 代码示例 / 说明 |
|---|---|
| 返回 JSON | w.Header().Set("Content-Type","application/json; charset=utf-8")<br>json.NewEncoder(w).Encode(v) |
| 解析 JSON 请求体 | json.NewDecoder(r.Body).Decode(&v) |
注意：在编码前先设置 Content-Type；若 Encode 失败且已开始写入，处理会复杂。

示例：
```go
w.Header().Set("Content-Type", "application/json; charset=utf-8")
json.NewEncoder(w).Encode(resp)
```

---

## 文件与静态资源
| 函数 | 说明 |
|---|---|
| http.ServeFile(w,r,path) | 发送本地文件（自动处理 Content-Type、Range） |
| http.FileServer(http.Dir("static")) | 静态文件服务器（与 http.StripPrefix 配合） |
| http.StripPrefix("/static/", fs) | 去掉 URL 前缀后交给 FileServer |
| http.ServeContent(w, r, name, modtime, content) | 更低级的文件发送，支持条件请求 |

---

## 常见 Header 与缓存
| Header                                         | 说明                                              |
| ---------------------------------------------- | ----------------------------------------------- |
| Content-Type                                   | 必须在写 body 前设置（资源类型）                             |
| Content-Length                                 | 通常由写入器自动设置                                      |
| Cache-Control / Expires / ETag / Last-Modified | 缓存控制，与 If-None-Match/If-Modified-Since 配合       |
| Vary                                           | 告诉缓存依据哪些请求 Header 返回不同结果                        |
| Access-Control-Allow-*                         | CORS：Allow-Origin, Allow-Methods, Allow-Headers |

注意：生产环境不要随意使用 Access-Control-Allow-Origin: *（会带来安全风险）

---

## Cookie 与会话
| 方法 / 字段                              | 说明                                                                     |
| ------------------------------------ | ---------------------------------------------------------------------- |
| http.SetCookie(w, &http.Cookie{...}) | 设置 Set-Cookie（推荐）                                                      |
| r.Cookie(name)                       | 读取 Cookie                                                              |
| Cookie 字段                            | Name, Value, Path, Domain, Expires, MaxAge, Secure, HttpOnly, SameSite |
安全建议：生产环境设置 Secure、HttpOnly、SameSite=Lax/Strict

示例：
```go
http.SetCookie(w, &http.Cookie{
  Name: "session_id", Value: "abc", HttpOnly: true, Secure: true, Path: "/",
})
```

---

## 中间件与安全
中间件模式（常见）：
```go
func myMiddleware(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    // 前置
    next.ServeHTTP(w, r)
    // 后置
  })
}
```
常见中间件：日志、恢复 panic（recover）、限速、超时、认证、CORS

panic 恢复：在顶层 middleware 使用 recover 避免服务崩溃。

安全 header 建议：
- X-Content-Type-Options: nosniff
- X-Frame-Options: DENY
- Content-Security-Policy: ...（按需）

---

## 并发 / Context / 取消
| 项目 | 说明 |
|---|---|
| r.Context() | 请求的 context，支持取消/超时/值传递 |
| 使用场景 | 将 context 传给 DB/远端请求，以便在客户端断开时取消操作 |
示例：
```go
select {
case <-r.Context().Done():
  // 请求已取消
  return
default:
  // 正常进行
}
```

---

## HTTP 客户端（发起请求）
| 用法 | 说明 |
|---|---|
| http.Get / http.Post | 简单快捷，适合快速测试 |
| client := &http.Client{Timeout:10*time.Second} | 自定义客户端（应显式设置超时） |
| req, _ := http.NewRequest(method, url, body); client.Do(req) | 构造并发送请求 |
建议：复用 Transport 以重用连接池，避免频繁创建 client/transport 导致资源浪费。

---

## 实用小技巧与注意事项（速览）
- 必须在写响应体之前设置 Header；一旦写入（w.Write / Encoder.Encode）header 已发送，无法再改大多数 header 或状态码。
- Header 名称不区分大小写，但用标准写法（如 "Content-Type"）。
- Set 会覆盖已有值，Add 会追加值（适用于 Set-Cookie）。
- 使用 http.MaxBytesReader 限制请求体大小防止 DoS。
- 对外 API 谨慎设置 CORS，不要生产环境随意用 "*"
- 优雅关闭：监听 SIGINT/SIGTERM，调用 Server.Shutdown(ctx)

---

## 常用代码片段（打印友好）
文本响应：
```go
w.Header().Set("Content-Type", "text/plain; charset=utf-8")
w.Header().Set("X-App-Name", "demo")
w.WriteHeader(http.StatusOK)
w.Write([]byte("hello"))
```

返回 JSON：
```go
w.Header().Set("Content-Type", "application/json; charset=utf-8")
json.NewEncoder(w).Encode(map[string]string{"msg": "ok"})
```

设置 Cookie：
```go
http.SetCookie(w, &http.Cookie{
  Name: "sid", Value: "abc123", HttpOnly: true, Path: "/", Expires: time.Now().Add(24*time.Hour),
})
```

CORS 预检处理：
```go
w.Header().Set("Access-Control-Allow-Origin", "*")
w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
if r.Method == http.MethodOptions {
  w.WriteHeader(http.StatusNoContent)
  return
}
```

限制请求体大小：
```go
r.Body = http.MaxBytesReader(w, r.Body, 10<<20) // 10 MiB
defer r.Body.Close()
```

---

## 导出为 PDF（几种方法）
1. 使用 Pandoc（需安装 pandoc）：
   - 命令：pandoc HTTP_Cheatsheet.md -o HTTP_Cheatsheet.pdf
   - 若需更好排版并支持代码高亮，可加参数与 pdf 引擎，例如：
     pandoc HTTP_Cheatsheet.md -o HTTP_Cheatsheet.pdf --pdf-engine=xelatex

2. 使用 VS Code：
   - 打开文件 -> File -> Print -> 选择 Save as PDF（或使用 Markdown Preview -> 打印）。

3. 使用在线编辑器（如 GitHub、StackEdit）：
   - 在浏览器打开 Markdown，使用浏览器的“打印为 PDF”功能。

---

如果你想，我可以：
- 直接把这份文件另存为 PDF 并提供下载（若你允许我生成并返回 PDF 文件）；
- 或把内容按你项目场景（API/静态服务器/文件上传）进一步精简为单页一栏式速查卡以便打印。告诉我你偏好哪种（生成 PDF / 更紧凑的单页版 / 添加公司样式 header）。 