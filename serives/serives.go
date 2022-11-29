package serives

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BackData struct {
	Event string      `json:"event"`
	Wxid  string      `json:"wxid"`
	Data  interface{} `json:"data"`
}

func HttpServes() {

	r := gin.Default()
	// var json BackData
	json := make(map[string]interface{})

	r.POST("/back", func(c *gin.Context) {

		err := c.BindJSON(&json)
		if err != nil {
			return
		}

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
	_ = r.Run(":9999")
}
