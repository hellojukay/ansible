# <span id="modeul">自定义模块</span>
有时候官方自带的模块不能完成我们需要的功能，或者实现起来非常复杂的时候，这个时候我们也找不到开源的已有方案，我们可以考虑自己编写Ansible模块，他的使用方法和核心模块的使用方法是一样的。Ansible在设计的是并没有限制我们开发Ansible模块时候使用的语言。Ansible在2.2以后就支持二进制模块了,所以我们可以使用golang来编写Ansible模块。

首先，将模块文件读入内存，然后添加传递给模块的参数，最后将模块中所需要的类添加到内存，由zipfile压缩后，再由base64进行编码，写入到模版文件内。

通过默认的连接方式，一般是ssh。ansible通过ssh连接到远程主机，创建临时目录，并关闭连接。然后将打开另外一个ssh连接，将模版文件以sftp方式传送到刚刚创建的临时目录中，写完后关闭连接。然后打开一个ssh连接将任务对象赋予可执行权限，执行成功后关闭连接。

最后，ansible将打开第三个连接来执行模块，并删除临时目录及其所有内容。模块的结果是从标准输出stdout中获取json格式的字符串。ansible将解析和处理此字符串。如果有任务是异步控制执行的，ansible将在模块完成之前关闭第三个连接，并且返回主机后，在规定的时间内检查任务状态，直到模块完成或规定的时间超时。

## <span id="canshu">读取参数</span>
我们在使用golang开发模块的时候，面临一个问题：如何读取Ansible的参数？Ansible是将参数以json的形式写入到一个临时文件中，然后将这个临时文件的路径传给golang的命令行，我们可以用os.Args[1]来获取这个json文件。读取参数成功以后，我们就可以使用golang在目标机器上执行命令或者操作文件和网络等等。
```golang
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

```

## <span id="fanhuizhi">返回值</span>
Ansible的返回值是一个json，他一般可能包括以下几个部分。


|字段|类型|说明|
|---|---|---|
|backup_file|string|对于一些modules使用了backup变量，返回备份的文件路径|
|changed|bool|表示任务是否对服务器进行了修改|
|failed|bool|表示任务是否失败|
|invocation|map|调用模块的方法|
|msg|string|返回的消息|
|rc|int|命令行程序的返回码|
|results|map|如果该键存在，则表示该任务存在循环，并且它包含每个项目的模块“results”的列表|
|skipped|bool|是否跳过执行|
|stderr|string|标准错误输出|
|stderr_lines|list|它将stderr字符串按行分割存储在列表中|
|stdout|string|标准输出|
|stdout_lines|list|它将stdout字符串按行分割存储在列表中|


## <span id="shiyon">如何使用自己写的模块</span>
自动以的模块的使用方法和核心模块的使用方法基本是一致的，这里举个例子
```shell
---
- name:
  hosts: aliyun
  remote_user: root
  tasks:
  - name: Hello World版本的playbook
    hello:
      name: "{{ item }}"
    with_items: ["licong",caobo]
```
这里的hello就是上面的自定义模块，当然，我们也可以item参数改成从ansible参数传递一进去的形式。
使用golang开发Ansible模块的弊端就是无法跨平台，而python是可以跨平台执行的，而且golang在编译以后，二进制包的体积比较大，python开发的模块一般只有几kb大小。