package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	//输入用户名密码
	var username, password string
	fmt.Scanf("%s\n", &username)
	fmt.Scanf("%s\n", &password)

	//创建cookiejar
	jar, _ := cookiejar.New(nil)

	//创建客户端
	client := &http.Client{
		Jar: jar,
	}

	//创建第一次请求
	loginurl := "http://kjyy.ccnu.edu.cn/clientweb/xcus/ic2/Default.aspx"
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
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应体失败:", err)
		return
	}
	fmt.Println(string(body))
}
