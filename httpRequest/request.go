package httpRequest

import (
	"fmt"
	"bytes"
	"io/ioutil"
	"net/http"
	"encoding/json"
)

type ParamsBody struct{
	Type string `json:"type"`
	Data interface{} `json:"data"`
}
const MYWXID = "wxid_th751pkvbmi422"

func GetList() {
	BaseUrl := "http://127.0.0.1:7777/DaenWxHook/httpapi/?wxid="+MYWXID
	data:= &ParamsBody{
		Type:"X0000",
		Data:"{}",
	}
	bytesData, _ := json.Marshal(data)
	resp, _ := http.Post(BaseUrl, "application/json", bytes.NewReader(bytesData))
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
  fmt.Println(string(body))

}
