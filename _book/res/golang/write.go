package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// Response 返回值
type Response struct {
	Changed bool   `json:"changed"`
	Fail    bool   `json:"fail"`
	Msg     string `json:"msg"`
	RC      int    `json:"rc"`
}

func readJSON(f string) (string, error) {
	fh, err := os.Open(f)
	if err != nil {
		return "", err
	}
	contents, err := ioutil.ReadAll(fh)
	return string(contents), err
}

func main() {
	var f = os.Args[1]
	res := Response{}
	ff, err := os.Open(f)
	if err != nil {
		res.Changed = false
		res.Fail = true
		res.Msg = "打开参数文件" + err.Error()
		returnJSON(res, 2)
	}
	contents, err := ioutil.ReadAll(ff)
	if err != nil {
		res.Changed = false
		res.Fail = true
		res.Msg = "读取参数文件失败" + err.Error()
		returnJSON(res, 2)
	}
	var m = make(map[string]interface{})
	err = json.Unmarshal(contents, &m)
	if err != nil {
		res.Changed = false
		res.Fail = true
		res.Msg = "解析map出错" + err.Error()
		returnJSON(res, 2)
	}
	var path = m["path"]
	pathstring, _ := path.(string)
	fh, err := os.Create(pathstring)
	if err != nil {
		res.Changed = false
		res.Fail = true
		res.Msg = string(contents)
		returnJSON(res, 2)
	}
	defer fh.Close()
	var content = m["content"]
	contentstring, _ := content.(string)
	fh.Write([]byte(contentstring))
	res.Changed = true
	res.Fail = false
	res.Msg = "success"
	returnJSON(res, 0)
}

func returnJSON(res Response, status int) {
	res.RC = status
	var out io.Writer
	if status == 0 {
		out = os.Stdout
	} else {
		out = os.Stderr
	}
	contents, _ := json.Marshal(res)
	fmt.Fprint(out, string(contents))
	os.Exit(status)
}
