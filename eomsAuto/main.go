package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
)

// 全局cookie
var gCurCookies []*http.Cookie
var gCurCookieJar *cookiejar.Jar

/** 请求
 * reqType		请求类型 "GET" "POST"
 * httpClient	请求的client
 * urlStr		待请求的url
 * postData		post请求带有的数据，当请求为GET时，为空
 * hostURL		请求过程中的host
 * referURL		请求过程中的refer
 * errMsg		请求过程中出错后的标准化输出带有的特征信息
 * saveFilePath	请求过程中存储response body的文件路径
 * return (string, error) 返回(请求获得的body, 错误信息)
**/
func requestURLWithInfo(reqType string, httpClient *http.Client, urlStr string, postData *url.Values,
	hostURL string, referURL string, errMsg string, saveFilePath string) (string, error) {

	var req *http.Request
	var err error
	//新建请求
	if postData == nil {
		// GET
		req, err = http.NewRequest(reqType, urlStr, nil)
	} else { // POST
		req, err = http.NewRequest(reqType, urlStr, strings.NewReader(postData.Encode()))
	}
	throwError(err, reqType+"错误：NewRequest，"+errMsg)
	setRequestsHeadersWithHostAndRefer(req, hostURL, referURL)

	resp, err := httpClient.Do(req)
	throwError(err, reqType+"错误: Do Get，"+errMsg)
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	//读取错误，则返回错误信息
	if err != nil {
		return "", err
	}

	// 如果 saveFilePath 不为nil,则存储read的信息
	if saveFilePath != "" {
		go saveFileToPath(respBody, saveFilePath)
	}

	return string(respBody), nil

}

// 处理工地那流程
func handleTickets() {
	// Fiddler proxy
	proxy := func(_ *http.Request) (*url.URL, error) {
		return url.Parse("http://127.0.0.1:8888")
	}
	httpTransport := &http.Transport{Proxy: proxy}

	httpClient := &http.Client{
		Transport: httpTransport,
		Jar:       gCurCookieJar, //新的请求时，可以指定相同的jar，就可以发送cookies
	}

	// Step 0: 试图通过本地Cookies请求首页，成功则不需要重新请求登录
	gCurCookies = loadCookies()
	// dbgPrintCurCookies(gCurCookies)
	baseURL, _ := url.Parse("http://10.14.2.168")
	gCurCookieJar.SetCookies(baseURL, gCurCookies)
	fmt.Println("Step1 尝试通过加载本地cookie登录")
	content, err := requestURLWithInfo("GET", httpClient, "http://10.14.2.168/cmcceoms/common/jsp/index.jsp", nil,
		"10.14.2.168", "http://10.14.2.168/cmcceoms/roles/login.jsp", "Step 1 尝试通过存储的cookie加载主页", "./HTMLs/00index.txt")

	if err != nil || strings.Contains(content, "登录名或密码错误") {
		fmt.Println("通过加载本地cookies登录失败，重新进行登录请求")
		//Step 1: 请求登录页
		_, err := requestURLWithInfo("GET", httpClient, "http://10.14.2.168/cmcceoms/roles/login.jsp", nil,
			"10.14.2.168", "", "Step 1 获取登录页", "./HTMLs/01loginIndex.txt")
		throwError(err, "Step 1 获取登录页失败")
		fmt.Println("Step 1: 获取登录页成功！")

		//step 2: 登录
		reqData := url.Values{}
		reqData.Set("key1", "meihaifeng")
		reqData.Set("key2", "Password;1")

		content, err := requestURLWithInfo("POST", httpClient, "http://10.14.2.168/cmcceoms/roles/login.do",
			&reqData, "10.14.2.168", "http://10.14.2.168/cmcceoms/roles/login.jsp", "Step 2 请求登录", "./HTMLs/02login.txt")
		throwError(err, "Step 2 请求登录失败")
		if !strings.Contains(content, "<root success=\"true\">") {
			fmt.Println("Step 2: 登录失败！")
			return
		}
		// fmt.Println(content)
		fmt.Println("Step 2: 登录成功！")

	} else {
		fmt.Println("通过加载本地cookies登录成功！")
	}

	// Step 3 : 获取工单页面， todo：可能不止一页，全部获取
	//listCount=100&pageNumber=1&pageCount=163&pageSize=100&isfirst=2&BaseSummary=&
	//BaseStatus=&BaseSN=&BaseCreateDate1=&BaseCreateDate2=&txtSortfiled=StDate&sortType=1&action=DEAL
	reqData3 := url.Values{}
	reqData3.Set("listCount", "100")
	reqData3.Set("pageNumber", "1")
	reqData3.Set("pageCount", "163")
	reqData3.Set("pageSize", "10") // 每页请求数量
	reqData3.Set("isfirst", "2")
	reqData3.Set("BaseSummary", "")
	reqData3.Set("BaseStatus", "")
	reqData3.Set("BaseSN", "")
	reqData3.Set("BaseCreateDate1", "")
	reqData3.Set("BaseCreateDate2", "")
	reqData3.Set("txtSortfiled", "StDate")
	reqData3.Set("sortType", "1")
	reqData3.Set("action", "DEAL")

	content, err = requestURLWithInfo("POST", httpClient, "http://10.14.2.168/cmcceoms/UltraProcess/BaseQuery/QueryWaitDealNew.jsp?action=DEAL",
		&reqData3, "10.14.2.168", "http://10.14.2.168/cmcceoms/UltraProcess/BaseQuery/QueryWaitDealNew.jsp?action=DEAL", "Step 3 请求待处理工单", "./HTMLs/03tickets.txt")
	throwError(err, "Step 3 请求待处理工单失败")
	// fmt.Println(content)
	fmt.Println("Step 3: 登录成功！")

	// Step 4: 处理订单信息
	ticketArray := parseTickets(content)
	fmt.Println(len(ticketArray))
	for i, item := range ticketArray {
		// fmt.Printf("%d %+v\n", i, item)
		fmt.Println(i, item.ticketID, item.ticketType, item.url, item.status, item.title)
	}

	// fmt.Println(req.URL)
	gCurCookies = gCurCookieJar.Cookies(baseURL)
	// dbgPrintCurCookies(gCurCookies)
	saveCookies(gCurCookies)

}

func main() {
	//init cookie tool
	gCurCookies = loadCookies()
	gCurCookieJar, _ = cookiejar.New(nil)

	// handleTickets()

	content := loadFile("./HTMLs/03tickets.txt")
	ticketArray := parseTickets(content)
	fmt.Println(len(ticketArray))
	for i, item := range ticketArray {
		// fmt.Printf("%d %+v\n", i, item)
		fmt.Println(i, item.ticketID, item.ticketType, item.url, item.status, item.title)
	}

}
