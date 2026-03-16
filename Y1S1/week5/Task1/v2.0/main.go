package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

type User struct {
	LStudentID string `json:"id"`
	PStudentID string
	Name       string `json:"name"`
	Grade      string `json:"Pid"`
}

var wg sync.WaitGroup

func main() {
	//输入用户名密码
	var username, password string
	fmt.Scanf("%s", &username)
	fmt.Scanf("%s", &password)

	//创建cookiejar
	jar, _ := cookiejar.New(nil)

	//创建客户端
	client := &http.Client{
		Jar: jar,
	}

	//创建第一次请求
	loginurl := "https://account.ccnu.edu.cn/cas/login?service=http://kjyy.ccnu.edu.cn/loginall.aspx?page="
	req, err := http.NewRequest("GET", loginurl, nil)
	if err != nil {
		fmt.Println("创建新请求失败:", err)
		return
	}

	//发送第一次请求
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("发送请求失败:", err)
		return
	}
	defer resp.Body.Close()

	//打印状态码
	fmt.Println("GET请求状态码:", resp.StatusCode)

	//解析HTML
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Println("解析HTML失败:", err)
		return
	}

	//获取lt
	ltSel := doc.Find(`input[name="lt"]`)
	lt, ok := ltSel.Attr("value")
	if !ok {
		fmt.Println("获取lt失败")
		return
	}
	// fmt.Println("lt:", lt)

	// 获取execution
	execSel := doc.Find(`input[name="execution"]`)
	execution, ok := execSel.Attr("value")
	if !ok {
		fmt.Println("获取execution失败")
		return
	}
	// fmt.Println("execution:", execution)

	data := url.Values{}
	data.Set("username", username)
	data.Set("password", password)
	data.Set("lt", lt)
	data.Set("execution", execution)
	data.Set("_eventId", "submit")
	data.Set("submit", "登录")

	payload := strings.NewReader(data.Encode())

	// var jsessionid string
	// re := regexp.MustCompile(`JSESSIONID=([^;]+)`)
	// for _, sc := range resp.Header.Values("Set-Cookie") {
	// 	if m := re.FindStringSubmatch(sc); m != nil {
	// 		fmt.Println("JSESSIONID:", m[1])
	// 		jsessionid = m[1]
	// 	}
	// }

	// loginurl = loginurl + ";jsessionid=" + jsessionid
	// fmt.Println(loginurl)

	loginurl = "https://account.ccnu.edu.cn/cas/login?service=http://kjyy.ccnu.edu.cn/loginall.aspx?page="

	req, err = http.NewRequest("POST", loginurl, payload)
	if err != nil {
		fmt.Println("创建请求失败:", err)
		return
	}

	// req.Header.Add("Cookie", "JSESSIONID="+jsessionid)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err = client.Do(req)
	if err != nil {
		fmt.Println("发送请求失败:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("POST请求状态码:", resp.StatusCode)

	// for _, cookie := range resp.Header.Values("Set-Cookie") {
	// 	fmt.Println(cookie)
	// }
	// fmt.Println(req.Header.Get("Cookie"))
	// text := req.Header.Get("Cookie")
	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Println("读取响应体失败:", err)
	// 	return
	// }
	// fmt.Println(string(body))

	//第三次请求，查询
	key := make(chan struct{}, 20)
	for i := 0; i < 20; i++ {
		key <- struct{}{}
	}

	f, err := os.Create("result.txt")
	if err != nil {
		fmt.Println("创建文件失败:", err)
		return
	}

	var mu sync.Mutex

	for i := 2025000000; i < 2026000000; i += 200 {
		wg.Add(1)
		go func(start int) {
			<-key
			defer func() { key <- struct{}{} }()

			defer wg.Done()

			end := start + 200

			for i := start; i < end; i++ {
				serchurl := "http://kjyy.ccnu.edu.cn/ClientWeb/pro/ajax/data/searchAccount.aspx"
				term := strconv.Itoa(i)
				// fmt.Println(term)
				paydata := url.Values{}
				paydata.Set("type", "logonname")
				paydata.Set("ReservaApply", "ReservaApply")
				paydata.Set("term", term)
				paydata.Set("_", "1764209114550") //这个好像写什么都行
				paydataStr := paydata.Encode()

				req, err := http.NewRequest("GET", serchurl+"?"+paydataStr, nil)
				if err != nil {
					fmt.Println("创建请求失败:", err)
					return
				}

				resp, err := client.Do(req)
				if err != nil {
					fmt.Println("发送请求失败:", err)
					return
				}

				body, err := io.ReadAll(resp.Body)
				if err != nil {
					fmt.Println("读取响应体失败:", err)
					return
				}
				resp.Body.Close()
				//fmt.Println(string(body)) 数组[{"id":"ID","Pid": "校园卡号","name": "名字","label": "名字","szLogonName": "xxxx","szHandPhone": "","szTel": "","szEmail": ""}]

				var user []User
				err = json.Unmarshal(body, &user)
				if err != nil {
					fmt.Println("解析JSON失败:", err)
					return
				}
				if len(user) != 0 {
					// fmt.Println("查询请求状态码:", resp.StatusCode)
					user[0].PStudentID = term
					fmt.Println(user)
					mu.Lock()
					f.WriteString(fmt.Sprintf("Name: %s, LStudentID: %s, PStudentID: %s, Grade: %s级\n", user[0].Name, user[0].LStudentID, user[0].PStudentID, user[0].Grade[0:4]))
					mu.Unlock()
				}
			}
		}(i)
	}
	wg.Wait()
	f.Close()
}
