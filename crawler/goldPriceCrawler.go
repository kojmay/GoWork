package main

import (
	"log"
	"regexp"
	"encoding/json"
	"encoding/csv"
	"os"
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

	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Cannot create file %q: %s\n", fName, err)
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	
	// {"date":"2020-04-24","open":"1948.000","high":"2020.000","low":"1939.400","close":"1993.000","volume":"61"}
	writer.Write([]string{"Date", "Open", "High", "Low", "Close", "volume"})

	for _, v := range priceResult {
		writer.Write([]string{v["date"].(string), v["open"].(string), v["high"].(string), v["low"].(string), v["close"].(string), v["volume"].(string)})
	}
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
		log.Println("Response:", r.Request.URL.String()[64:67])
		csvSave(r.Request.URL.String()[64:67]+".csv" ,r.Body)
	})

	c.Visit("https://stock2.finance.sina.com.cn/futures/api/jsonp.php/var%20_XPD2020_4_24=/GlobalFuturesService.getGlobalFuturesDailyKLine?symbol=XPD&_=2020_4_24&source=web")
	c.Visit("https://stock2.finance.sina.com.cn/futures/api/jsonp.php/var%20_XAU2020_4_25=/GlobalFuturesService.getGlobalFuturesDailyKLine?symbol=XAU&_=2020_4_25&source=web")
	c.Visit("https://stock2.finance.sina.com.cn/futures/api/jsonp.php/var%20_XAG2020_4_25=/GlobalFuturesService.getGlobalFuturesDailyKLine?symbol=XAG&_=2020_4_25&source=web")
	c.Visit("https://stock2.finance.sina.com.cn/futures/api/jsonp.php/var%20_XPT2020_4_25=/GlobalFuturesService.getGlobalFuturesDailyKLine?symbol=XPT&_=2020_4_25&source=web")
	c.Wait()

}