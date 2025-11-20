package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

// ============================ 全局变量和锁 ============================

// 存储所有用户数据，key为用户名，value为用户对象
var users = make(map[string]*User)

// 存储所有会话数据，key为会话ID，value为会话对象
var sessions = make(map[string]*Session)

// 互斥锁，用于保证并发安全
var (
	usersMux    sync.RWMutex // 用户数据的读写锁
	sessionsMux sync.RWMutex // 会话数据的读写锁
)

// ============================ 数据结构定义 ============================

// User 用户结构体
type User struct {
	Username  string    `json:"username"`   // 用户名，唯一标识
	Password  string    `json:"-"`          // 密码，不序列化到JSON
	Nickname  string    `json:"nickname"`   // 昵称
	Email     string    `json:"email"`      // 邮箱
	CreatedAt time.Time `json:"created_at"` // 创建时间
}

// Session 会话结构体
type Session struct {
	ID       string    `json:"id"`       // 会话ID
	Username string    `json:"username"` // 关联的用户名
	Expiry   time.Time `json:"expiry"`   // 过期时间
}

// 请求和响应结构体
type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ChangePasswordRequest struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

type UpdateProfileRequest struct {
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// ============================ 数据存储操作函数 ============================

// CreateUser 创建新用户
func CreateUser(user *User) bool {
	usersMux.Lock()
	defer usersMux.Unlock()

	// 检查用户名是否已存在
	if _, exists := users[user.Username]; exists {
		return false
	}

	// 存储用户
	users[user.Username] = user
	return true
}

// GetUser 根据用户名获取用户
func GetUser(username string) (*User, bool) {
	usersMux.RLock()
	defer usersMux.RUnlock()

	user, exists := users[username]
	return user, exists
}

// UpdateUser 更新用户信息
func UpdateUser(username string, updates map[string]interface{}) bool {
	usersMux.Lock()
	defer usersMux.Unlock()

	user, exists := users[username]
	if !exists {
		return false
	}

	// 更新字段
	if nickname, ok := updates["nickname"]; ok {
		user.Nickname = nickname.(string)
	}
	if email, ok := updates["email"]; ok {
		user.Email = email.(string)
	}

	return true
}

// CreateSession 创建会话
func CreateSession(username string) string {
	sessionsMux.Lock()
	defer sessionsMux.Unlock()

	// 生成会话ID
	sessionID := fmt.Sprintf("session_%d", time.Now().UnixNano())

	// 创建并存储会话
	sessions[sessionID] = &Session{
		ID:       sessionID,
		Username: username,
		Expiry:   time.Now().Add(24 * time.Hour), // 24小时有效期
	}

	return sessionID
}

// GetSession 获取会话
func GetSession(sessionID string) (*Session, bool) {
	sessionsMux.RLock()
	defer sessionsMux.RUnlock()

	session, exists := sessions[sessionID]
	if !exists || time.Now().After(session.Expiry) {
		return nil, false
	}

	return session, true
}

// DeleteSession 删除会话
func DeleteSession(sessionID string) {
	sessionsMux.Lock()
	defer sessionsMux.Unlock()

	delete(sessions, sessionID)
}

// ============================ 响应辅助函数 ============================

// writeError 返回错误响应
func writeError(w http.ResponseWriter, message string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(Response{
		Code:    code,
		Message: message,
	})
}

// writeSuccess 返回成功响应
func writeSuccess(w http.ResponseWriter, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Response{
		Code:    200,
		Message: message,
		Data:    data,
	})
}

// ============================ 认证中间件 ============================

// RequireAuth 认证中间件，保护需要登录的路由
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

// ============================ HTTP处理器函数 ============================

// Register 用户注册
func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		writeError(w, "只支持POST请求", http.StatusMethodNotAllowed)
		return
	}

	// 解析请求体
	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, "请求格式错误", http.StatusBadRequest)
		return
	}

	// 验证必要字段
	if req.Username == "" || req.Password == "" {
		writeError(w, "用户名和密码不能为空", http.StatusBadRequest)
		return
	}

	// 创建用户对象
	user := &User{
		Username:  req.Username,
		Password:  req.Password, // 注意：这里直接存储明文密码，仅用于练习
		Nickname:  req.Nickname,
		Email:     req.Email,
		CreatedAt: time.Now(),
	}

	// 尝试创建用户
	if !CreateUser(user) {
		writeError(w, "用户名已存在", http.StatusConflict)
		return
	}

	writeSuccess(w, "注册成功", nil)
}

// Login 用户登录
func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		writeError(w, "只支持POST请求", http.StatusMethodNotAllowed)
		return
	}

	// 解析登录请求
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, "请求格式错误", http.StatusBadRequest)
		return
	}

	// 查找用户
	user, exists := GetUser(req.Username)
	if !exists {
		writeError(w, "用户名或密码错误", http.StatusUnauthorized)
		return
	}

	// 验证密码（直接比较明文）
	if req.Password != user.Password {
		writeError(w, "用户名或密码错误", http.StatusUnauthorized)
		return
	}

	// 创建会话
	sessionID := CreateSession(req.Username)

	// 设置Cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		Value:    sessionID,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		MaxAge:   24 * 60 * 60, // 24小时
	})

	// 返回登录成功信息
	writeSuccess(w, "登录成功", map[string]string{
		"username": user.Username,
		"nickname": user.Nickname,
	})
}

// Logout 用户登出
func Logout(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		writeError(w, "只支持POST请求", http.StatusMethodNotAllowed)
		return
	}

	// 获取并删除会话
	cookie, err := r.Cookie("session_id")
	if err == nil {
		DeleteSession(cookie.Value)
	}

	// 清除Cookie
	http.SetCookie(w, &http.Cookie{
		Name:   "session_id",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	})

	writeSuccess(w, "登出成功", nil)
}

// ChangePassword 修改密码
func ChangePassword(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		writeError(w, "只支持PUT请求", http.StatusMethodNotAllowed)
		return
	}

	// 解析请求
	var req ChangePasswordRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, "请求格式错误", http.StatusBadRequest)
		return
	}

	// 从上下文获取当前登录的用户名
	username := r.Context().Value("username").(string)

	// 获取用户信息
	user, exists := GetUser(username)
	if !exists {
		writeError(w, "用户不存在", http.StatusNotFound)
		return
	}

	// 验证旧密码（直接比较明文）
	if req.OldPassword != user.Password {
		writeError(w, "旧密码错误", http.StatusBadRequest)
		return
	}

	// 更新密码
	user.Password = req.NewPassword

	writeSuccess(w, "密码修改成功", nil)
}

// GetProfile 获取用户信息
func GetProfile(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		writeError(w, "只支持GET请求", http.StatusMethodNotAllowed)
		return
	}

	// 从上下文获取当前登录的用户名
	username := r.Context().Value("username").(string)

	// 获取用户信息
	user, exists := GetUser(username)
	if !exists {
		writeError(w, "用户不存在", http.StatusNotFound)
		return
	}

	// 返回用户信息（不包含密码）
	userInfo := map[string]interface{}{
		"username":   user.Username,
		"nickname":   user.Nickname,
		"email":      user.Email,
		"created_at": user.CreatedAt,
	}

	writeSuccess(w, "获取用户信息成功", userInfo)
}

// UpdateProfile 更新用户资料
func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		writeError(w, "只支持PUT请求", http.StatusMethodNotAllowed)
		return
	}

	// 解析请求
	var req UpdateProfileRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, "请求格式错误", http.StatusBadRequest)
		return
	}

	// 从上下文获取当前登录的用户名
	username := r.Context().Value("username").(string)

	// 构建更新字段
	updates := make(map[string]interface{})
	if req.Nickname != "" {
		updates["nickname"] = req.Nickname
	}
	if req.Email != "" {
		updates["email"] = req.Email
	}

	// 更新用户信息
	if !UpdateUser(username, updates) {
		writeError(w, "更新用户信息失败", http.StatusInternalServerError)
		return
	}

	writeSuccess(w, "资料更新成功", nil)
}

// HealthCheck 健康检查接口
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	writeSuccess(w, "服务运行正常", map[string]interface{}{
		"timestamp": time.Now(),
		"status":    "running",
	})
}

// ============================ 主函数 ============================

func main() {
	// 注册路由
	http.HandleFunc("/register", Register)
	http.HandleFunc("/login", Login)
	http.HandleFunc("/logout", Logout)

	// 需要认证的路由（使用中间件保护）
	http.Handle("/password", RequireAuth(http.HandlerFunc(ChangePassword)))
	http.Handle("/user/profile", RequireAuth(http.HandlerFunc(GetProfile)))
	http.Handle("/user/profile/update", RequireAuth(http.HandlerFunc(UpdateProfile)))

	// 健康检查
	http.HandleFunc("/health", HealthCheck)

	// 启动服务器
	fmt.Println("用户管理系统服务启动成功！")
	fmt.Println("服务地址: http://localhost:8080")
	fmt.Println("\n可用接口:")
	fmt.Println("  POST   /register          - 用户注册")
	fmt.Println("  POST   /login             - 用户登录")
	fmt.Println("  POST   /logout            - 用户登出")
	fmt.Println("  PUT    /password          - 修改密码（需登录）")
	fmt.Println("  GET    /user/profile      - 获取用户信息（需登录）")
	fmt.Println("  PUT    /user/profile/update - 更新用户资料（需登录）")
	fmt.Println("  GET    /health            - 健康检查")

	fmt.Println("\n使用示例:")
	fmt.Println("  1. 先注册用户: curl -X POST http://localhost:8080/register -H 'Content-Type: application/json' -d '{\"username\":\"test\",\"password\":\"123\",\"nickname\":\"测试用户\",\"email\":\"test@example.com\"}'")
	fmt.Println("  2. 登录获取Cookie: curl -X POST http://localhost:8080/login -H 'Content-Type: application/json' -d '{\"username\":\"test\",\"password\":\"123\"}' -c cookies.txt")
	fmt.Println("  3. 访问受保护接口: curl -X GET http://localhost:8080/user/profile -b cookies.txt")

	// 启动HTTP服务器
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("服务器启动失败: %v\n", err)
	}
}
