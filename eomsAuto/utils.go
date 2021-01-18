package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"time"
)

// 一些简化方法

// 打印cookie
func dbgPrintCurCookies(curCookies []*http.Cookie) {
	var cookieNum int = len(curCookies)
	fmt.Printf("cookieNum=%d\n", cookieNum)
	for i := 0; i < cookieNum; i++ {
		var curCk *http.Cookie = curCookies[i]
		fmt.Printf("\n------ Cookie [%d]------", i)
		fmt.Printf("\n\tName=%s", curCk.Name)
		fmt.Printf("\n\tValue=%s", curCk.Value)
		fmt.Printf("\n\tPath=%s", curCk.Path)
		fmt.Printf("\n\tDomain=%s", curCk.Domain)
		fmt.Printf("\n\tExpires=%s", curCk.Expires)
		fmt.Printf("\n\tRawExpires=%s", curCk.RawExpires)
		fmt.Printf("\n\tMaxAge=%d", curCk.MaxAge)
		fmt.Printf("\n\tSecure=%t", curCk.Secure)
		fmt.Printf("\n\tHttpOnly=%t", curCk.HttpOnly)
		fmt.Printf("\n\tRaw=%s", curCk.Raw)
		fmt.Printf("\n\tUnparsed=%s\n", curCk.Unparsed)
		fmt.Println("\n\t-------------------------")
	}
}

// 为请求设置Header的host 与 refer
func setRequestsHeadersWithHostAndRefer(req *http.Request, host string, refer string) {
	req.Header.Set("Accept-Language", "zh-Hans;q=0.5")
	req.Header.Set("User-Agent", "Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 10.0; WOW64; Trident/7.0; .NET4.0C; .NET4.0E; .NET CLR 2.0.50727; .NET CLR 3.0.30729; .NET CLR 3.5.30729)")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("Accept", "application/xml, text/xml, */*")
	req.Header.Set("Accept-Encoding", "gzip, deflate")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("Connection", "Keep-Alive")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Host", host)
	req.Header.Set("Referer", refer)
}

// error handler
func throwError(err error, msg string) {
	if err != nil {
		fmt.Println("Error, "+msg+" :", err)
		panic(err)
	}
}

// 将content存储到 filePath
func saveFileToPath(content []byte, filePath string) {
	err := ioutil.WriteFile(filePath, []byte(content), 0644)
	throwError(err, "File cannot be saved:"+filePath)
	fmt.Println("File Saved:", filePath)
}

// 保存cookies
// todo 按cookie名来存取和更新
func saveCookies(curCookies []*http.Cookie) {
	data, err := json.Marshal(curCookies)
	throwError(err, "cookies转换错误")
	saveFileToPath(data, "./loginCookieJar")
}

// 加载Cookies
func loadCookies() []*http.Cookie {
	data, err := ioutil.ReadFile("./loginCookieJar")
	throwError(err, "读取loginCookieJar错误")

	var cookiesJSON = regSearch(string(data), `\{(.*?)\}`)
	var storedCookies []*http.Cookie
	for _, cookieStr := range cookiesJSON {
		// fmt.Println(cookieStr)
		var cookie http.Cookie
		err := json.Unmarshal([]byte(cookieStr), &cookie)
		throwError(err, "error 加载出错")

		storedCookies = append(storedCookies, &cookie)
	}
	// dbgPrintCurCookies(storedCookies)
	return storedCookies
}

// 正则匹配
func regSearch(dataStr, regStr string) []string {
	var re = regexp.MustCompile(regStr)
	return re.FindAllString(dataStr, -1)
}

// load html，作测试用
func loadFile(filePath string) string {
	content, err := ioutil.ReadFile(filePath)
	throwError(err, " ,"+filePath+"加载出错")
	return string(content)
}

func parseTickets(content string) []*Ticket {
	var re = regexp.MustCompile(`openSheet\('(.*?)','(.*?)','(.*?)'.*\s*(.*)\s*</a>`)
	// for i, data := range re.FindAllStringSubmatch(content, -1) {
	// 	fmt.Println(i, data[1], data[2], data[3], data[4])
	// 	// fmt.Printf("%dssa:%q11", i, data[4])
	// }
	matchResult := re.FindAllStringSubmatch(content, -1)
	// var ticketResult []*Ticket
	ticketResult := []*Ticket{}
	for i := 0; i < len(matchResult); i = i + 8 {
		ticketItem := new(Ticket)
		ticketItem.ticketID = matchResult[i][3]
		ticketItem.ticketType = matchResult[i][2]
		ticketItem.url = string(matchResult[i][1])

		ticketItem.group = matchResult[i+1][4]
		deadline, _ := time.Parse("2006-01-02 15:04:05", strings.Split(matchResult[i+3][4], "\r")[0])
		ticketItem.deadline = deadline
		arrTime, _ := time.Parse("2006-01-02 15:04:05", strings.Split(matchResult[i+4][4], "\r")[0])
		ticketItem.arrivedTime = arrTime
		ticketItem.status = matchResult[i+5][4]
		ticketItem.title = matchResult[i+7][4]
		ticketResult = append(ticketResult, ticketItem)
		// fmt.Println(i, ticketResult[len(ticketResult)-1].ticketID, ticketResult[len(ticketResult)-1].ticketType,
		// 	ticketResult[len(ticketResult)-1].url, matchResult[i][1], ticketResult[len(ticketResult)-1].status, ticketResult[len(ticketResult)-1].title)

		fmt.Println(i, len(ticketResult[len(ticketResult)-1].url), ticketResult[len(ticketResult)-1].url, matchResult[i][1])
	}

	// for i, item := range ticketResult {
	// 	fmt.Printf("%d %+v\n", i, item)
	// 	fmt.Println(i, len(item.url), item.url)
	// }
	return ticketResult
}
