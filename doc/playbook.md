# <span id="playbook">playbook</span>
什么是playbook?
> 
Playbooks 的格式是YAML（详见:YAML 语法）,语法做到最小化,意在避免 playbooks 成为一种编程语言或是脚本,但它也并不是一个配置模型或过程的模型.playbook 由一个或多个 ‘plays’ 组成.它的内容是一个以 ‘plays’ 为元素的列表.在 play 之中,一组机器被映射为定义好的角色.在 ansible 中,play 的内容,被称为 tasks,即任务.在基本层次的应用中,一个任务是一个对 ansible 模块的调用,这在前面章节学习过.‘plays’ 好似音符,playbook 好似由 ‘plays’ 构成的曲谱,通过 playbook,可以编排步骤进行多机器的部署,比如在 webservers 组的所有机器上运行一定的步骤, 然后在 database server 组运行一些步骤,最后回到 webservers 组,再运行一些步骤,诸如此类.“plays” 算是一个体育方面的类比,你可以通过多个 plays 告诉你的系统做不同的事情,不仅是定义一种特定的状态或模型.你可以在不同时间运行不同的 plays.

对初学者,这里有一个 playbook,其中仅包含一个 play:
```txt
---
- hosts: webservers
  remote_user: root
  tasks:
  - name: ensure apache is at the latest version
    yum: pkg=httpd state=latest
  - name: write the apache config file
    template: src=/srv/httpd.j2 dest=/etc/httpd.conf
    notify:
    - restart apache
  - name: ensure apache is running
    service: name=httpd state=started
  handlers:
    - name: restart apache
      service: name=httpd state=restarted
```
## <span id="helloworld">Hello World<span>
我们来编写一个Hello World版本的playbook
```txt
---
- hosts: aliyun
  remote_user: root
  tasks:
  - name: Hello World版本的playbook
    shell: /bin/echo Hello World
```

## <span id="bianliang">变量</span>
&nbsp;&nbsp;&nbsp;&nbsp;上面的例子中，有些东西我们是写死了，不方复用。比如说hosts就可以写成变量形式
```txt
---
- hosts: {{host}}
  remote_user: root
  tasks:
  - name: Hello World版本的playbook
    shell: /bin/echo Hello World
```
这样我们可以在执行过程中，通过命令行参数吧host传进来，也可以吧Host变量的值写当前文件的preferences.fact文件中。
```shell
ansible-playbook helloworld.yum -extra-vars="host=aliyun"
```
这里我们-extra-vars参数传递了host参数到yml文件中,在Ansible的1.2版本以后是支持传json的
```txt
–extra-vars ‘{“pacman”:”mrs”,”ghosts”:[“inky”,”pinky”,”clyde”,”sue”]}’
```
在Ansible的1.3版本以后，可以直接加载一个json文件
```txt
–extra-vars “@some_file.json”
```

## <span id="tiaojian">条件</span>
Ansible在执行task的时候，是可以按照条件决定是否执行的
```shell
---
- name: if条件
  hosts: localhost
  tasks:
  - name:
    shell: echo "Hello World"
    when: a == '10'

```
执行
```shell
 ~/C/M/a/res  ansible-playbook if.yml --extra-vars "a=100" -v                                    1406ms  一 10/ 9 18:38:54 2017
Using /Users/jukay/Code/Markdown/ansible/res/ansible.cfg as config file

PLAY [if条件] *********************************************************************************************************************

TASK [Gathering Facts] **********************************************************************************************************
ok: [localhost]

TASK [command] ******************************************************************************************************************
skipping: [localhost] => {"changed": false, "skip_reason": "Conditional result was False", "skipped": true}

PLAY RECAP **********************************************************************************************************************
localhost                  : ok=1    changed=0    unreachable=0    failed=0
```
这里提示条件不满足，跳过了
如果换一下参数
```shell
Using /Users/jukay/Code/Markdown/ansible/res/ansible.cfg as config file

PLAY [if条件] *********************************************************************************************************************

TASK [Gathering Facts] **********************************************************************************************************
ok: [localhost]

TASK [command] ******************************************************************************************************************
changed: [localhost] => {"changed": true, "cmd": "echo \"Hello World\"", "delta": "0:00:00.007619", "end": "2017-10-09 18:41:01.170398", "rc": 0, "start": "2017-10-09 18:41:01.162779", "stderr": "", "stderr_lines": [], "stdout": "Hello World", "stdout_lines": ["Hello World"]}

PLAY RECAP **********************************************************************************************************************
localhost                  : ok=2    changed=1    unreachable=0    failed=0
```
这里标准输出就能成功输出Hello World。类型的条件判断还有

* in 判断一个字符串是否存在另外一个字符串中
* is define 判断一个变量是否定义
* not 判断布尔类型是否为false
* a|b 判断两个布尔类型是否至少一个为true
* \> 或者 < 判断大小关系
* or 和 | 一样
* 等等。。。

## <span id="diedai">迭代</span>
Ansible的参数是可以传列表的
```shell
---
- name: loop
  hosts: localhost
  remote_user: jukay
  tasks:
  - name: 
    shell: echo Hello {{item}}
    with_items: [licong,caobo]
```
或者
```shell
---
- name: loop
  hosts: localhost
  remote_user: jukay
  tasks:
  - name: 
    shell: echo Hello {{item}}
    with_items:
    - licong
    - caobo
```
迭代map和object数组都是可以的。