package main

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/gin-gonic/gin"
)

// MD5
func MD5(in string) (string, error) {
	hash := md5.Sum([]byte(in))
	return hex.EncodeToString(hash[:]), nil
}

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"hello": "hello world",
		})
	})

	router.LoadHTMLFiles("html/index.html")
	router.GET("/html", func(c *gin.Context) {
		c.HTML(200, "index.html", "meihaifneg")
	})

	router.Run(":8081")
}
