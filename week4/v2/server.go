package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

// type User struct {
// 	Username string    `json:"username"`
// 	Password string `json:"-"`
// 	Nickname string `json:"nickname"`
// 	Name     string `json:"name"`
// }

// var Userinfomation map[string]User

// func Register(w http.ResponseWriter, req *http.Request) {
// 	if  req.Method != "POST" {
// 		fmt.Fprint(w, "The request method isn't POST")
// 		return
// 	}

// 	Newusername := req.FormValue("username")
// 	Newpassword := req.FormValue("password")
// 	if _, ok := Userinfomation[Newusername]; ok {
// 		fmt.Fprint(w, "用户已存在")
// 		return
// 	}
// 	user := User{Username: Newusername, Password: Newpassword}
// 	Userinfomation[Newusername] = user
// }

// func Login(w http.ResponseWriter, req *http.Request)

// var UserInfomation map[string]*User
// var Sessions map[string]Session
var UserInfomation = make(map[string]*User)
var Sessions = make(map[string]Session)

// var userMux sync.Mutex

var userMux sync.RWMutex
var sessionsMux sync.RWMutex

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Session struct {
	ID       string    `json:"id"`
	Username string    `json:"username"`
	Expiry   time.Time `json:"expiry"`
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
	Age      int    `json:"age"`
}

type RegisterReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
	Age      int    `json:"age"`
}

func (r RegisterReq) Context() {
	panic("unimplemented")
}

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ChangeReq struct {
	Nickname string `json:"nickname"`
	Age      int    `json:"age"`
}

type ChangePasswordReq struct {
	Oldpassword string `json:"oldpassword"`
	Newpassword string `json:"newpassword"`
}

// 错误响应函数
func WriteError(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(Response{
		Code:    code,
		Message: message,
	})
}

// 成功响应函数
func WriteSuccess(w http.ResponseWriter, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Response{
		Code:    200,
		Message: message,
		Data:    data,
	})
}

// 创建新用户
func CreateNewUser(newuser /*Registerreq*/ *User) bool {
	// userMux.RLock()
	// defer userMux.RUnlock()
	userMux.Lock()
	defer userMux.Unlock()

	if _, ok := UserInfomation[newuser.Username]; ok {
		return false
	}
	UserInfomation[newuser.Username] = newuser
	return true
}

// 创建新Session
func CreateSession(username string) string {
	sessionsMux.Lock()
	defer sessionsMux.Unlock()

	sessionID := fmt.Sprintf("session.%d", time.Now().UnixNano())

	session := Session{
		ID:       sessionID,
		Username: username,
		Expiry:   time.Now().Add(24 * time.Hour),
	}

	Sessions[sessionID] = session

	return sessionID
}

// 获取Session
func GetSession(sessionID string) (*Session, bool) {
	// sessionsMux.RLock()
	// defer sessionsMux.RUnlock()

	// session, ok := Sessions[sessionID]
	sessionsMux.RLock()

	session, ok := Sessions[sessionID]

	sessionsMux.RUnlock()

	// if !ok || time.Now().After(session.Expiry) {
	// 	DeleteSession(sessionID)
	// 	//这里不需要立刻进行返回，完全丢给处理器函数来做这个事情
	// 	return nil, false
	// }
	if !ok {
		return nil, false
	}

	if time.Now().After(session.Expiry) {
		DeleteSession(sessionID)
		return nil, false
	}
	return &session, true
}

// 删除Session
func DeleteSession(sessionID string) {
	sessionsMux.Lock()
	defer sessionsMux.Unlock()

	delete(Sessions, sessionID)
}

// 获取用户
func GetUser(username string) (*User, bool) {
	userMux.RLock()
	defer userMux.RUnlock()

	user, ok := UserInfomation[username]
	if !ok {
		return nil, false
	}
	return user, true
}

// 认证中间件
func MiddleAthusation(dofunc http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		cookie, err := req.Cookie("session_id")
		if err != nil {
			WriteError(w, 404, "请先登录")
			return
		}

		// session, ok := Sessions[cookie.Value]
		// if !ok {
		// 	WriteError(w, 404, "会话已经过期,请重新登录")
		// 	return
		// }
		session, ok := GetSession(cookie.Value)
		if !ok {
			WriteError(w, 404, "会话已过期,请重新登录")
			return
		}

		ctx := context.WithValue(req.Context(), "username", session.Username)
		dofunc.ServeHTTP(w, req.WithContext(ctx))

	})
}

// 更新个人资料
func ChangeProfile(username string, changes map[string]interface{}) bool {
	//空
	userMux.Lock()
	defer userMux.Unlock()

	// user, ok := UserInfomation[username]
	// if !ok {
	// 	return false
	// }
	user, ok := GetUser(username)
	if !ok {
		return false
	}

	nickname, ok := changes["nickname"]
	if ok {
		user.Nickname = nickname.(string)
	}

	age, ok := changes["age"]
	if ok {
		user.Age = age.(int)
	}

	return true
}

// 登录函数
func Login(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		WriteError(w, 404, "方法必须是POST")
		return
	}

	var r LoginReq
	err := json.NewDecoder(req.Body).Decode(&r)
	if err != nil {
		WriteError(w, 404, "反序列化请求体时失败")
		return
	}

	User, ok := GetUser(r.Username)
	if !ok {
		WriteError(w, 400, "用户不存在")
		return
	}

	if r.Password != User.Password {
		WriteError(w, 404, "密码错误")
		return
	}

	sessionID := CreateSession(User.Username)

	http.SetCookie(w, &http.Cookie{
		Name:   "session_id",
		Value:  sessionID,
		Path:   "/",
		MaxAge: 24 * 60 * 60,
	})

	WriteSuccess(w, "登录成功", nil)
}

// 登出函数
func Logout(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		WriteError(w, 404, "必须使用POST")
		return
	}

	// var r LoginReq
	// err := json.NewDecoder(req.Body).Decode(&r)
	// if err != nil {
	// 	WriteError(w, 404, "反序列化失败")
	// 	return
	// }

	http.SetCookie(w, &http.Cookie{
		Name:   "session_id",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	})

	WriteSuccess(w, "登出成功", nil)
}

// 注册函数
func Register(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		WriteError(w, 404, "必须使用POST方法")
		return
	}

	var r RegisterReq
	err := json.NewDecoder(req.Body).Decode(&r)
	if err != nil {
		WriteError(w, 400, "反序列化失败")
		return
	}

	// newuser := User{
	// 	Username: r.Username,
	// 	Password: r.Password,
	// 	Nickname: r.Nickname,
	// 	Age:      r.Age,
	// }
	newuser := User(r)

	ok := CreateNewUser(&newuser)
	if !ok {
		WriteError(w, 401, "创建失败,用户已存在")
		return
	}

	WriteSuccess(w, "创建成功,请牢记用户名和密码", nil)
}

// 以下都是需要经过中间件认证的处理函数
// 修改昵称和年龄
func ChangeInformation(w http.ResponseWriter, req *http.Request) {
	if req.Method != "PUT" {
		WriteError(w, 404, "必须使用PUT")
		return
	}

	var r ChangeReq
	err := json.NewDecoder(req.Body).Decode(&r)
	if err != nil {
		WriteError(w, 400, "反序列化失败")
		return
	}

	// ok := ChangeProfile(r.Context().Value("username"), )
	//username := r.Context().Value("username").(string)
	username := req.Context().Value("username").(string)

	changes := map[string]interface{}{
		"nickname": r.Nickname,
		"age":      r.Age,
		//Age
	}

	ok := ChangeProfile(username, changes)
	if !ok {
		WriteError(w, 402, "获取用户失败")
		//空
		return
	}

	WriteSuccess(w, "修改成功", nil)
}

// 修改密码
// ***没写修改密码的处理函数我服了
func ChangePassword(username string, changes map[string]string) bool {
	//空
	userMux.Lock()
	defer userMux.Unlock()

	user, ok := GetUser(username)
	if !ok {
		return false
	}

	if user.Password != changes["oldpassword"] {
		return false
	}

	user.Password = changes["newpassword"]
	return true
}

func ChangeMima(w http.ResponseWriter, req *http.Request) {
	if req.Method != "PUT" {
		WriteError(w, 403, "必须使用PUT方法")
		return
	}

	var r ChangePasswordReq
	err := json.NewDecoder(req.Body).Decode(&r)
	if err != nil {
		WriteError(w, 406, "反序列化失败")
		return
	}

	username := req.Context().Value("username").(string)

	changes := map[string]string{
		"oldpassword": r.Oldpassword,
		"newpassword": r.Newpassword,
	}

	ok := ChangePassword(username, changes)
	if !ok {
		WriteError(w, 408, "旧密码错误")
		return
	}

	WriteSuccess(w, "修改成功", nil)
}

// 获取用户资料
func GetProfile(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		WriteError(w, 419, "必须使用GET方法")
		return
	}

	username := req.Context().Value("username").(string)

	user, ok := GetUser(username)
	if !ok {
		WriteError(w, 444, "获取用户失败")
		return
	}

	data := map[string]interface{}{
		"username": user.Username,
		"nickname": user.Nickname,
		"Age":      user.Age,
	}

	WriteSuccess(w, "获取成功", data)
}

func main() {
	//原本是先监听再注册，应先注册再监听
	http.HandleFunc("/login", Login)
	http.HandleFunc("/logout", Logout)
	http.HandleFunc("/register", Register)
	http.Handle("/changepassword", MiddleAthusation(http.HandlerFunc(ChangeMima)))
	http.Handle("/user/profile", MiddleAthusation(http.HandlerFunc(GetProfile)))
	http.Handle("/user/profile/change", MiddleAthusation(http.HandlerFunc(ChangeInformation)))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("启动服务器失败")
	}
}
