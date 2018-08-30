# Ansible
这个项目讲述了Ansible的基本使用方法，和开发自定义模块的方式和例子
# 目录
* [Ansible简介](./doc/introduction.md#desc)
  * [Ansible的特点](./doc/introduction.md#tedian)
  * [Hello World](./doc/introduction.md#helloworld)
  * [安装方式](./doc/introduction.md#anzhuang)
* [Hosts文件](./doc/hosts.md#hosts)
  * [编写host文件](./doc/hosts.md#xiehosts)
* [playbook](./doc/playbook.md#playbook)
  * [helloworld](./doc/playbook.md#helloworld)
  * [变量](./doc/playbook.md#bianliang)
  * [条件](./doc/playbook.md#tiaojian)
* [自定义模块](./doc/modeul.md#modeul)
  * [参数](./doc/modeul.md#canshu)
  * [返回值](./doc/modeul.md#fanhuizhi)
  * [使用自定义模块](./doc/modeul.md#shiyon)
* [Ansible配置](./doc/config.md)

# demo
这里编写一个名字叫做`fuck`的模块，他的功能是往标准输出里面输出`fuck`.
```golang
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
```
编译之后放在模块的目录中
```shell
mac-pro:res jukay$ ansible dev -m fuck -u root
39.106.10.228 | SUCCESS => {
    "changed": false,
    "fail": false,
    "msg": "",
    "rc": 0
}
mac-pro:res jukay$
```