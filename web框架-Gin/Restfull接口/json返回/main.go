package main

import "github.com/gin-gonic/gin"

// 定义map字典
var person = make(map[string]string)

func main() {
	person["name"] = "tom"
	person["sex"] = "man"
	person["city"] = "cs"

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": person,
		})
	})
	r.Run(":8081") // listen and serve on 0.0.0.0:8080
}
