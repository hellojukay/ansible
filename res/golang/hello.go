package main

import (
	"encoding/json"
	"fmt"
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

func main() {
	var args = os.Args[:]
	var response Response
	if len(args) != 2 {
		response.Fail = true
		response.Msg = "失败"
		response.Changed = false
	}
	contents, err := readJSON(args[1])
	if err != nil {
		response.Msg = err.Error()
		response.Fail = true
		response.Changed = false

	} else {
		var m = make(map[string]string)
		json.Unmarshal([]byte(contents), &m)
		response.Msg = " Hello " + m["name"]
		response.Fail = false
		response.Changed = false
	}

	buffer, _ := json.Marshal(response)
	fmt.Println(string(buffer))
}

func readJSON(f string) (string, error) {
	fh, err := os.Open(f)
	if err != nil {
		return "", err
	}
	contents, err := ioutil.ReadAll(fh)
	return string(contents), err
}
