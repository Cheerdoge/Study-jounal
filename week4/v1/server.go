package main

import (
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

var Userinfomation map[string]*User
var Sessions map[string]Session

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

func WriteError(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(Response{
		Code:    code,
		Message: message,
	})
}

func WriteSuccess(w http.ResponseWriter, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Response{
		Code:    200,
		Message: message,
		Data:    data,
	})
}

func CreateNewUser(newuser /*Registerreq*/ *User) bool {
	userMux.RLock()
	defer userMux.RUnlock()

	if _, ok := Userinfomation[newuser.Username]; ok {
		return false
	}
	Userinfomation[newuser.Username] = newuser
	return true
}

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

func GetSession(sessionID string) (*Session, bool) {
	sessionsMux.RLock()
	defer sessionsMux.RUnlock()

	session, ok := Sessions[sessionID]

	if !ok || time.Now().After(session.Expiry) {
		DeleteSession(sessionID)
		//这里不需要立刻进行返回，完全丢给处理器函数来做这个事情
		return nil, false
	}
	return &session, true
}

func DeleteSession(sessionID string) {
	delete(Sessions, sessionID)
}
