package main

import (
	"log"
	"regexp"
	"fmt"
	"encoding/json"
	"github.com/gocolly/colly"
)

func csvSave(fName string, data []byte) {
	var re = regexp.MustCompile(`(\{.*?\})`)
	var priceResult []map[string]interface{}


	for _, match := range re.FindAllString(string(data), -1) {
		// fmt.Println(match, "found at index", i)
		var priceItem map[string]interface{}
		if err := json.Unmarshal([]byte(match), &priceItem); err != nil {
			panic(err)
		}
		priceResult = append(priceResult, priceItem)
	}

	fmt.Println(len(priceResult))

	for i, v := range priceResult {
		fmt.Println(i, v["date"])
	}

	// out, _ := json.Marshal(priceResult)
    // println(string(out))


	
}


func main() {
	c := colly.NewCollector(
		colly.Async(true),
		colly.UserAgent("Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)"),
	)

	c.Limit(&colly.LimitRule{DomainGlob:  "*.sge.*", Parallelism: 5})

	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting:", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	c.OnResponse(func(r *colly.Response) {
		// log.Println("Response:", string(r.Body))
		csvSave("./xpd.csv" ,r.Body)
	})

	c.Visit("https://stock2.finance.sina.com.cn/futures/api/jsonp.php/var%20_XPD2020_4_24=/GlobalFuturesService.getGlobalFuturesDailyKLine?symbol=XPD&_=2020_4_24&source=web")
	c.Wait()
}