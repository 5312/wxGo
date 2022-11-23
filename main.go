package main

import (
	"fmt"
	"github.com/gin-gonic/gin"

	"net/http"
	 
	. "wxGo/httpRequest"
)
type BackData struct{
	Event string `json:"event"`
	Wxid string `json:"wxid"`
	Data interface{} `json:"data"`
}
func main() {

	GetList()
	
	r := gin.Default()
	// var json BackData
	json := make(map[string]interface{})

	r.POST("/back", func(c *gin.Context) {

		c.BindJSON(&json)

		fmt.Println("请求开始----------")

		fmt.Println(json["event"])

		fmt.Println("请求结束--------")
		c.JSON(http.StatusOK, gin.H{
			"id":       "1",
			"page":     "",
			"username": "",
			"password": "",
		})
	})
	r.Run(":9999")
}

