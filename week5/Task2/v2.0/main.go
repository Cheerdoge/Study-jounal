package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type User struct {
	StudentID string `json:"id"`
	Name      string `json:"name"`
	Grade     string `json:"pid"`
}

type ReserveMsg struct {
	Ret int    `json:"ret"`
	Msg string `json:"msg"`
}

type SearchData struct {
	Data []Seat `json:"data"`
}

type Seat struct {
	DevID   string `json:"devid"`
	DevName string `json:"devname"`
}

func main() {
	//输入用户名密码
	var username, password string
	fmt.Printf("请输入用户名:")
	fmt.Scanf("%s\n", &username)
	fmt.Printf("请输入密码:")
	fmt.Scanf("%s\n", &password)

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
	resp1, err := client.Do(req)
	if err != nil {
		fmt.Println("发送请求失败:", err)
		return
	}
	defer resp1.Body.Close()

	//打印状态码
	fmt.Println("模拟登录GET请求状态码:", resp1.StatusCode)

	//解析HTML
	doc, err := goquery.NewDocumentFromReader(resp1.Body)
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

	resp2, err := client.Do(req)
	if err != nil {
		fmt.Println("发送请求失败:", err)
		return
	}
	defer resp2.Body.Close()

	fmt.Println("模拟登录POST请求状态码:", resp2.StatusCode)

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

	//第三次请求，进入座位表页面
	zuoweibiaourl := "http://kjyy.ccnu.edu.cn/ClientWeb/pro/ajax/device.aspx"

	fmt.Printf("请输入去图书馆内卷的日期(格式2025-11-29):")
	var date string
	fmt.Scanf("%s\n", &date)
	fmt.Printf("请输入开始内卷的时间(格式18:00):")
	var fr_start string
	fmt.Scanf("%s\n", &fr_start)
	fmt.Printf("请输入结束内卷的时间(格式18:00):")
	var fr_end string
	fmt.Scanf("%s\n", &fr_end)
	fmt.Printf("请输入预约启动的日期(格式2025-11-29):")
	var reservedate string
	fmt.Scanf("%s\n", &reservedate)
	fmt.Printf("请输入预约启动的时间(格式18:00):")
	var reservetime string
	fmt.Scanf("%s\n", &reservetime)
	reservedate = reservedate + " " + reservetime + ":00"
	starttime, err := time.ParseInLocation("2006-01-02 15:04:05", reservedate, time.Local)
	if err != nil {
		fmt.Println("时间格式错误:", err)
		return
	}
	fmt.Println("预约启动时间为:", starttime)

	paydata1 := url.Values{}
	paydata1.Set("byType", "devcls")
	paydata1.Set("classkind", "8")
	paydata1.Set("display", "fp")
	paydata1.Set("md", "d")
	paydata1.Set("room_id", "101699179")
	paydata1.Set("purpose", "")
	paydata1.Set("selectOpenAty", "")
	paydata1.Set("cld_name", "default")
	paydata1.Set("selectOpenAty", "")
	paydata1.Set("date", date)
	paydata1.Set("fr_start", fr_start)
	paydata1.Set("fr_end", fr_end)
	paydata1.Set("act", "get_rsv_sta")
	paydata1.Set("_", "1764209114550") //这个好像写什么都行
	req, err = http.NewRequest("GET", zuoweibiaourl+"?"+paydata1.Encode(), nil)
	if err != nil {
		fmt.Println("创建请求失败:", err)
		return
	}

	resp3, err := client.Do(req)
	if err != nil {
		fmt.Println("发送请求失败:", err)
		return
	}

	fmt.Println("获取座位表请求状态码:", resp3.StatusCode)

	body3, err := io.ReadAll(resp3.Body)
	if err != nil {
		fmt.Println("读取响应体失败:", err)
		return
	}
	resp3.Body.Close()

	var searchData SearchData
	err = json.Unmarshal(body3, &searchData)
	if err != nil {
		fmt.Println("解析响应体失败:", err)
		return
	}
	// fmt.Println("可用座位列表:")

	seatlist := make(map[string]string)
	for _, seat := range searchData.Data {
		// fmt.Printf("座位ID: %s, 座位名称: %s\n", seat.DevID, seat.DevName)
		seatlist[seat.DevName] = seat.DevID
	}
	//第四次请求，预定座位
	reserveurl := "http://kjyy.ccnu.edu.cn/ClientWeb/pro/ajax/reserve.aspx"

	fmt.Printf("请输入要预约的座位名称(N1199~N1328):")
	var seatName string
	fmt.Scanf("%s\n", &seatName)
	seatID, ok := seatlist[seatName]
	if !ok {
		fmt.Println("座位不存在")
		return
	}

	start := date + " " + fr_start
	end := date + " " + fr_end

	start_time := fr_start[0:2] + fr_start[3:5]
	end_time := fr_end[0:2] + fr_end[3:5]

	paydata2 := url.Values{}
	paydata2.Set("dialogid", "")
	paydata2.Set("dev_id", seatID)
	paydata2.Set("lab_id", "")
	paydata2.Set("kind_id", "")
	paydata2.Set("room_id", "101699179")
	paydata2.Set("type", "dev")
	paydata2.Set("prop", "")
	paydata2.Set("test_id", "")
	paydata2.Set("term", "")
	paydata2.Set("Vnumber", "")
	paydata2.Set("test_name", "")
	paydata2.Set("start", start)
	paydata2.Set("end", end)
	paydata2.Set("start_time", start_time)
	paydata2.Set("end_time", end_time)
	paydata2.Set("up_file", "")
	paydata2.Set("memo", "")
	paydata2.Set("act", "set_resv")
	paydata2.Set("_", "1764209114550") //这个好像写什么都行
	req, err = http.NewRequest("GET", reserveurl+"?"+paydata2.Encode(), nil)
	if err != nil {
		fmt.Println("创建请求失败:", err)
		return
	}

	waitDuration := time.Until(starttime)
	if waitDuration > 0 {
		fmt.Println("开始等待", waitDuration)
		time.Sleep(waitDuration)
	}

	fmt.Println("等待结束，开始预约")

	for i := 0; i < 3; i++ {

		fmt.Printf("第%d次预约尝试\n", i+1)

		resp4, err := client.Do(req)
		if err != nil {
			fmt.Println("发送请求失败:", err)
			return
		}

		body4, err := io.ReadAll(resp4.Body)
		if err != nil {
			fmt.Println("读取响应体失败:", err)
			return
		}
		resp4.Body.Close()
		// fmt.Println(string(body))
		/*
			{"ret":0,"act":"set_resv","msg":"当前时间预约冲突","data":null,"ext":null}
			{"ret":0,"act":"set_resv","msg":"2025-11-30您选择的时间内成员[2025211393/陈博文]已有预约[176190643],不得再预约","data":null,"ext":null}
			{"ret":1,"act":"set_resv","msg":"操作成功！","data":null,"ext":null}
		*/
		var reservemsg ReserveMsg
		err = json.Unmarshal(body4, &reservemsg)
		if err != nil {
			fmt.Println("解析响应体失败:", err)
			return
		}
		if reservemsg.Ret == 1 {
			fmt.Println("预约成功:", reservemsg.Msg)
			break
		} else {
			fmt.Println("预约失败:", reservemsg.Msg)
		}
	}
}
