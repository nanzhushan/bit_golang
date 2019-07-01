package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 定义map字典
var person = make(map[string]string)

func main() {
	person["name"] = "tom"
	person["sex"] = "man"
	person["city"] = "cs"

	r := gin.Default()
	r.GET("/ping", Getting)
	r.Run(":8081") // listen and serve on 0.0.0.0:8080
}

func Getting(c *gin.Context) {
	fmt.Print(person)
	// c.JSON(http.StatusOK, person["name"])
	c.JSON(http.StatusOK, person)
}

