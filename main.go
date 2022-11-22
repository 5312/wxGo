package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"strings"
)

func main() {
	r := gin.Default()
	r.POST("/back", func(c *gin.Context) {
		json := make(map[string]interface{})
		c.BindJSON(&json)

		fmt.Println("请求开始")

		fmt.Println(json)
		fmt.Println("请求结束")
		c.JSON(http.StatusOK, gin.H{
			"id":       "1",
			"page":     "",
			"username": "",
			"password": "",
		})
	})
	r.Run(":9999")
}

func main1() {

	url := "http://127.0.0.1:7777/DaenWxHook/httpapi/"
	method := "POST"

	payload := strings.NewReader(`{
    "type": "X0000",
    "data": {}
}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("User-Agent", "Apifox/1.0.0 (https://www.apifox.cn)")
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
