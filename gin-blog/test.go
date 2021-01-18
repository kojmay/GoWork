package main

import "github.com/gin-gonic/gin"

<<<<<<< HEAD
func main1() {
	//Default返回一个默认的路由引擎
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		//输出json结果给调用方
=======
func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
>>>>>>> 12b4a013a020075c40b82fc83cf68f325a1f2ebe
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
<<<<<<< HEAD
	r.Run() // listen and serve on 0.0.0.0:8080
=======
	router.Run(":8081")
>>>>>>> 12b4a013a020075c40b82fc83cf68f325a1f2ebe
}
