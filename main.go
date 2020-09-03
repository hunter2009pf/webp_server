package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.LoadHTMLFiles("index.html")
	router.StaticFS("/image", http.Dir("./image"))

	router.GET("/android/appmarket", func(c *gin.Context) {
		// c.Redirect(http.StatusPermanentRedirect, "")
		c.Redirect(http.StatusPermanentRedirect, "market://details?id=com.tianxun.zhuwo")
		// c.Redirect(http.StatusPermanentRedirect, "https://www.qq.com/")
	})

	router.GET("/hello", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "hello world",
		})
	})

	router.Run(":8888")
}
