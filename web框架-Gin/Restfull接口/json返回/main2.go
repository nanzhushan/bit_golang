package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// curl  127.0.0.1:8005/get?key=5

func main() {

	gin.SetMode(gin.DebugMode) //全局设置环境，此为开发环境，线上环境为gin.ReleaseMode
	router := gin.Default()    //获得路由实例

	//添加中间件
	router.Use(Middleware)
	//注册接口
	router.GET("/get", Get)
	router.GET("/getjson", Getjson)
	//监听端口
	http.ListenAndServe(":8005", router)
}

func Middleware(c *gin.Context) {
	fmt.Println("this is a middleware!")
}

func Get(c *gin.Context) {
	value, exist := c.GetQuery("key")
	if !exist {
		value = "the key is not exist!"
	}
	c.Data(http.StatusOK, "text/plain", []byte(fmt.Sprintf("get success! %s\n", value)))
	return
}

// 构造json
type JsonHolder struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func Getjson(c *gin.Context) {
	holder := JsonHolder{Id: 1, Name: "tom"}
	//若返回json数据，可以直接使用gin封装好的JSON方法
	c.JSON(http.StatusOK, holder)
	return
}

