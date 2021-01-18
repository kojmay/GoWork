package main

import (
<<<<<<< HEAD
	"net/http"

	"github.com/gin-gonic/gin"
=======
	"fmt"
	"net/http"

	"github.com/kojmay/GOWORK/gin-blog/routers"
	"github.com/kojmay/GoWork/gin-blog/pkg/setting"
>>>>>>> 12b4a013a020075c40b82fc83cf68f325a1f2ebe
)

var router *gin.Engine

func main() {
<<<<<<< HEAD
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	initRoutes()

	router.Run()
}

func showIndexPage(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"index.html",
		gin.H{
			"title": "Home Page",
		},
	)
=======
	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
>>>>>>> 12b4a013a020075c40b82fc83cf68f325a1f2ebe
}
