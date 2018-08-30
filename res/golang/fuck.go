package main

import (
	"encoding/json"
)

// Response 返回值的消息结构
type Response struct {
	Changed bool   `json:"changed"`
	Fail    bool   `json:"fail"`
	Msg     string `json:"msg"`
	RC      int    `json:"rc"`
}

func main() {
	println("fuck")
	var res = Response{
		Changed: false,
		Fail:    false,
		Msg:     "",
		RC:      0,
	}
	buf, _ := json.Marshal(res)
	println(string(buf))
}
