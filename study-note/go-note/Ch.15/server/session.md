```go
type session struct {
ID string
Usename string
Expiry time.Time //过期时间
}
```
1. 使用map存储，key->value为session.ID->*session
2. 分为三个，创建（create）、获取（get）、删除过期（delete）
	1. **创建**
		1. 首先创建ID，可以用`sessionID := fmt.Sprintf("session_%d", time.Now().UnixNano())`
			`Sprintf`可以格式化写入字符串并返回
		2. 随后可以储存入map中，注意要取地址，过期时间可用`time.Now().Add(24 * time.Hour)`表示24小时有效期
		3. 最后返回sessionID，因为获取会话和map的key都是sessionID
		```go
		func CreateSession(username string) string {

    sessionsMux.Lock()

    defer sessionsMux.Unlock()

    // 生成会话ID

    sessionID := fmt.Sprintf("session_%d", time.Now().UnixNano())

    // 创建并存储会话

    sessions[sessionID] = &Session{

        ID:       sessionID,

        Username: username,

        Expiry:   time.Now().Add(24 * time.Hour), // 24小时有效期

    }

    return sessionID

}
		```
	2. **获取**
		1. 首先检查是否存在这个session，以及这个session是否过期（`time.Now().After(session.Expiry)`，返回值是bool值)
		2. 如果检查通过，则返回相应的会话（session）和bool值
		```go
		func GetSession(sessionID string) (*Session, bool) {

    sessionsMux.RLock()

    defer sessionsMux.RUnlock()

    session, exists := sessions[sessionID]

    if !exists || time.Now().After(session.Expiry) {

        return nil, false

    }

    return session, true

}
		```
	3. **删除**
		也可以只写删除代码然后放在get检查session下
		```go
		func CleanupExpiredSessions() {
    sessionMux.Lock()
    defer sessionMux.Unlock()
    
    now := time.Now()
    expiredCount := 0
    
    for sessionID, session := range sessions {
        if now.After(session.Expiry) {
            delete(sessions, sessionID)
            expiredCount++
        }
    }
		```
3. 需要一个并发锁保护，最好也是读写锁
4. **认证中间件**
   就是检测用户是否认证的，是个函数，一般签名为`AuthMIddleware(next http.Handler) http.Handler`，其中会涉及到上下文，点[这里](./Context.md)了解
```go
func RequireAuth(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // 从Cookie获取会话ID
        cookie, err := r.Cookie("session_id")
        if err != nil {
            writeError(w, "请先登录", http.StatusUnauthorized)
            return
        }
        // 验证会话
        session, valid := GetSession(cookie.Value)
        if !valid {
            writeError(w, "会话已过期，请重新登录", http.StatusUnauthorized)
            return
        }
        // 将用户名存入上下文，供后续处理器使用
        ctx := context.WithValue(r.Context(), "username", session.Username)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}
```
`http.Handler`是一个函数类型，实现了ServeHTTP方法，即实际处理函数
`ServeHTTP(w, r)`是将请求传递给下一个处理器的,其中w为`http.ResponseWrite`,r为`*http.Request`
`r.WithContext(ctx)`表示返回基于原有请求，但上下文被替换为ctx的请求体
`r.Context`是个上下文类型
获取cookie时返回的第二个值是错误类型
思路：
1. 闭包获取外部的参数，next，因为不能把函数当作Handler返回
2. 获取cookie，没有则返回
3. 验证session，成功则将session获取
4. 传入上下文，r中自带`r.Context()`
使用认证中间件的路由设置如下
`http.Handle("/password", RequireAuth(http.HandlerFunc(ChangePassword)))`
# 如何设置cookie返回
以上只是如何创建、存储和管理session，现在我们要在服务器运行时使得用户**登录**时得到一个cookie
## Cookie的属性
```go
type Cookie struct {
    Name     string //cookie名称
    Value    string //cookie值(session.ID)
    Path     string //cookie有效的路径
    Domain   string //cookie有效的域名
    Expires  time.Time //过期时间，-1表示立即过期
    MaxAge   int //最大存活时间，相比过期时间优先级更高，-1表示关闭浏览器消失
    Secure   bool //指定是否只能https传输，true为是
    HttpOnly bool //防止xss攻击
    SameSite SameSite //防止CSRF攻击
}
```

## 设置Cookie
使用`http.SetCookie(w &http.Cookie{})`函数设置
直接在括号内cookie结构体的花括号内填写即可

## 删除Cookie
当用户**登出**时，需要清除cookie
```go
 http.SetCookie(w, &http.Cookie{
        Name:     "session_id",
        Value:    "",
        Expires:  time.Now().Add(-1 * time.Hour), // 设置过去时间或-1会立即删除
        HttpOnly: true,
        Secure:   false,
        Path:     "/",
    })
```