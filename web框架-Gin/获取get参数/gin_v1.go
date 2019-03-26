package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func getParams(c *gin.Context)  {
	// http://127.0.0.1:8888/?name=knight&city=cs
	name := c.Query("name")
	city := c.Query("city")
	c.String(http.StatusOK, name+city)
}

func main() {
	router := gin.Default()
	// 获取路由参数
	router.GET("/", getParams)

	router.Run(":8888")
}