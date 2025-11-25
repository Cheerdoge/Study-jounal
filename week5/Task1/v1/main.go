package main

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	var username, password string
	fmt.Scanf("%s\n", &username)
	fmt.Scanf("%s\n", &password)

	jar, _ := cookiejar.New(nil)

	client := &http.Client{
		Jar: jar,
	}

	loginurl := "https://account.ccnu.edu.cn/cas/login"
	req, err := http.NewRequest("GET", loginurl, nil)
	if err != nil {
		fmt.Println("创建新请求失败:", err)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("发送请求失败:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("请求状态码:", resp.StatusCode)

	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Println("读取响应头失败:", err)
	// 	return
	// }

	// content := string(body)

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Println("解析HTML失败:", err)
		return
	}
	ltSel := doc.Find(`input[name="lt"]`)
	lt, ok := ltSel.Attr("value")
	if !ok {
		fmt.Println("获取lt失败")
		return
	}
	fmt.Println("lt:", lt)

	// 获取 execution 值（重要！）
	execSel := doc.Find(`input[name="execution"]`)
	execution, ok := execSel.Attr("value")
	if !ok {
		fmt.Println("获取execution失败")
		return
	}
	fmt.Println("execution:", execution)

	data := url.Values{}
	data.Set("username", username)
	data.Set("password", password)
	data.Set("lt", lt)
	data.Set("execution", execution)
	data.Set("_eventId", "submit")
	data.Set("submit", "登录")

	payload := strings.NewReader(data.Encode())
	// re := regexp.MustCompile(`^JSESSIONID=([^; \t\r\n]+)$`)
	// // fmt.Println(re.FindAllString(resp.Header["Set-Cookie"]["JSESSIONID"], -1))
	// fmt.Printf("%s", resp.Header["Set-Cookie"][0])
	// if m := re.FindStringSubmatch(resp.Header["Set-Cookie"][0]); m != nil {
	// 	fmt.Println(m[1])
	// }

	var jsessionid string
	re := regexp.MustCompile(`JSESSIONID=([^;]+)`)
	for _, sc := range resp.Header.Values("Set-Cookie") {
		if m := re.FindStringSubmatch(sc); m != nil {
			// fmt.Println("匹配到(JSESSIONID 头解析):", m[0])
			fmt.Println("JSESSIONID:", m[1])
			jsessionid = m[1]
		}
	}

	loginurl = loginurl + ";jsessionid=" + jsessionid
	fmt.Println(loginurl)
	req, err = http.NewRequest("POST", loginurl, payload)
	if err != nil {
		fmt.Println("创建请求失败:", err)
		return
	}

	req.Header.Add("Cookie", "JSESSIONID="+jsessionid)

	resp, err = client.Do(req)
	if err != nil {
		fmt.Println("发送请求失败:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("请求状态码:", resp.StatusCode)

	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Println("读取响应体失败:", err)
	// 	return
	// }

	// content := string(body)

	// fmt.Println(content)
	for _, cookie := range resp.Header.Values("Set-Cookie") {
		fmt.Println(cookie)
	}
}
