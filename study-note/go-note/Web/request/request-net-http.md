## 快速要点
- 使用 http.Client 执行请求（复用 Client/Transport 以重用连接池）。  
- 推荐显式超时：client := &http.Client{Timeout: 10*time.Second} 或使用 context.WithTimeout。  
- 构造请求用 http.NewRequest(method, url, body)。对于带上下文用 req = req.WithContext(ctx)。  
- 始终在读取完响应后调用 resp.Body.Close()。  
- 对文件/流式上传使用 multipart.Writer 或 io.Pipe；对大响应使用流式解码（json.Decoder）。

---

## 常用函数/字段速查表

| 功能         |                                                                            函数 / 字段 | 作用 / 常用场景                                       | 代码提示                                                        |
| ---------- | --------------------------------------------------------------------------------- | ----------------------------------------------- | ----------------------------------------------------------- |
| 构造请求       |                                                 http.NewRequest(method, url, body) | 创建 *http.Request；body 为 io.Reader（nil 表示无 body） | body 常用 bytes.NewReader / strings.NewReader / io.Pipe       |
| 带上下文请求     |                                                               req.WithContext(ctx) | 取消/超时控制（推荐）                                     | ctx := context.WithTimeout(...); req = req.WithContext(ctx) |
| 直接快捷请求     |                                               http.Get / http.Post / http.PostForm | 简单请求快速使用                                        | 不可配置复杂选项                                                    |
| 发送请求       |                                                                     client.Do(req) | 使用自定义 http.Client 发送请求                          | client := &http.Client{Timeout:...}                         |
| 简单响应检查     |                                                                          http.Head | 请求头（不取响应体）                                      | 适合检查资源存在或大小                                                 |
| 设置请求头      |                                                               req.Header.Set / Add | 设置 Content-Type、Authorization 等                 | req.Header.Set("Content-Type", "application/json")          |
| Basic Auth |                                                       req.SetBasicAuth(user, pass) | 设置 Authorization: Basic ...                     | 简便方法                                                        |
| 添加 Cookie  |                                                   req.AddCookie(&http.Cookie{...}) | 单个 cookie                                       | 或使用 client.Jar 管理 cookie                                    |
| 请求调试       |                                             httputil.DumpRequestOut / DumpResponse | 打印/记录请求与响应（含 body）                              | 注意敏感信息不记录                                                   |
| 读取响应       | resp.StatusCode / resp.Header / io.ReadAll(resp.Body) / json.NewDecoder(resp.Body) | 检查状态码并解析 body                                   | defer resp.Body.Close()                                     |
| 客户端级超时     |                                                                     client.Timeout | 总超时（包括连接、重定向、读体）                                | 适合大部分场景                                                     |
| 取消/超时      |                                                  context.Context + req.WithContext | 更细粒度控制                                          | 推荐用于 RPC/DB/外部请求                                            |
| 跟随重定向      |                                                               client.CheckRedirect | 自定义重定向策略                                        | 返回 error 可阻止重定向                                             |
| Cookie Jar |                                                                 cookiejar.New(nil) | 自动管理 Cookie（client.Jar）                         | client := &http.Client{Jar: jar}                            |
| 代理         |                     Transport.Proxy = http.ProxyURL(u) / http.ProxyFromEnvironment | 设置 HTTP 代理                                      | 使用 Transport 配置                                             |
| 传输层配置      |                                                               &http.Transport{...} | 连接复用、TLS、Idle timeout 等                         | 复用 Transport 避免频繁新建                                         |
| TLS 设置     |                                             Transport.TLSClientConfig (tls.Config) | 控制 InsecureSkipVerify、证书等                       | 仅开发/测试时慎用 InsecureSkipVerify                                |
| HTTP/2     |                                                   default Transport 支持 HTTP/2（TLS） | 若需更多控制使用 golang.org/x/net/http2                 | -                                                           |
| 流式/分块上传    |                                                      io.Pipe + multipart.NewWriter | 较大或生成中数据时使用                                     | 支持 chunked transfer                                         |
| 限制读取大小     |                                                  io.LimitReader / io.LimitedReader | 保护内存，避免读取无限大响应                                  | 结合最大允许大小使用                                                  |
| 重用连接       |                                              复用 Transport（不要频繁创建 client/transport） | 提高性能，减少资源                                       | 设置 MaxIdleConns、IdleConnTimeout                             |

---

## 常见示例片段

构造并发送请求（JSON body）：
```go
reqBody := bytes.NewBuffer(jsonBytes) // 或 bytes.NewReader
req, err := http.NewRequest("POST", "https://api.example.com/v1", reqBody)
req.Header.Set("Content-Type", "application/json")
req.Header.Set("Authorization", "Bearer "+token)
resp, err := client.Do(req)
defer resp.Body.Close()
```

使用 context 超时：
```go
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()
req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)
// or req = req.WithContext(ctx)
resp, err := client.Do(req)
```

发送表单 application/x-www-form-urlencoded：
```go
vals := url.Values{}
vals.Set("field", "value")
req, _ := http.NewRequest("POST", url, strings.NewReader(vals.Encode()))
req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
resp, _ := client.Do(req)
```

multipart 文件上传（示例）：
```go
var b bytes.Buffer
w := multipart.NewWriter(&b)
fw, _ := w.CreateFormFile("file", filename)
io.Copy(fw, fileReader)
w.WriteField("other", "value")
w.Close()
req, _ := http.NewRequest("POST", uploadURL, &b)
req.Header.Set("Content-Type", w.FormDataContentType())
resp, _ := client.Do(req)
```

流式上传（边生成边上传）：
```go
pr, pw := io.Pipe()
mw := multipart.NewWriter(pw)
go func() {
  defer pw.Close()
  part, _ := mw.CreateFormFile("file", "name")
  io.Copy(part, sourceReader) // 大数据源
  mw.Close()
}()
req, _ := http.NewRequest("POST", url, pr)
req.Header.Set("Content-Type", mw.FormDataContentType())
resp, _ := client.Do(req)
```

处理响应 JSON：
```go
defer resp.Body.Close()
if resp.StatusCode != http.StatusOK { /* handle */ }
var out MyStruct
if err := json.NewDecoder(resp.Body).Decode(&out); err != nil { /* handle */ }
```

调试请求与响应：
```go
dumpReq, _ := httputil.DumpRequestOut(req, true) // true 包含 body
log.Println(string(dumpReq))
dumpResp, _ := httputil.DumpResponse(resp, true)
log.Println(string(dumpResp))
```

自定义 Transport（示例）：
```go
tr := &http.Transport{
  Proxy: http.ProxyFromEnvironment,
  MaxIdleConns: 100,
  IdleConnTimeout: 90 * time.Second,
  TLSHandshakeTimeout: 10 * time.Second,
  ExpectContinueTimeout: 1 * time.Second,
  // TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // 测试慎用
}
client := &http.Client{Transport: tr, Timeout: 30 * time.Second}
```

阻止或控制重定向：
```go
client := &http.Client{
  CheckRedirect: func(req *http.Request, via []*http.Request) error {
    if len(via) >= 10 { return http.ErrUseLastResponse } // 停止重定向，返回最后响应
    return nil
  },
}
```

基本认证与 Bearer Token：
```go
req.SetBasicAuth("user", "pass")
req.Header.Set("Authorization", "Bearer "+token)
```

使用 Cookie Jar 自动存储/发送 Cookie：
```go
jar, _ := cookiejar.New(nil)
client := &http.Client{Jar: jar}
```

限制读取大小（防止 OOM）：
```go
lr := io.LimitReader(resp.Body, 10<<20) // 10 MiB
data, _ := io.ReadAll(lr)
```

---

## 调优与最佳实践（速览）
- 不要每次请求都创建新的 Transport/Client；复用 Transport 以复用连接。  
- 使用 client.Timeout 或 context 控制超时，防止挂起。  
- 对外部请求使用重试和指数退避策略（外部库或自实现）。  
- 对响应体较大或未知大小的情况，使用流式解码并限制可读字节数。  
- 小心记录日志时泄露敏感头（Authorization、Cookie）。  
- 对 TLS 验证不要在生产禁用（InsecureSkipVerify=false）。  
- 如果需上传/下载大文件，优先使用流式（io.Copy、io.Pipe、multipart.Writer）。  
- 对并发请求，注意最大并发数（goroutine 数 + Transport 配置）。