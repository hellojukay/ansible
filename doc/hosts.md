# <span id="hosts">Hosts文件介绍</span>
&nbsp;&nbsp;&nbsp;&nbsp;Anasible管理者海量的机器，这些机器的ip和用户密码都配置在Ansible的hosts文件中，hosts文件默认是存在/etc/ansible/hosts这个地址，也是可以在运行过程中指定的,比如
```shell
bash-3.2$ ansible huanggai -m command -a "echo hello world" -v -i /Users/jukay/Desktop/hosts
No config file found; using defaults
huanggai | SUCCESS | rc=0 >>
hello world
```

## <span id="xiehosts">编写host文件</span>
&nbsp;&nbsp;&nbsp;&nbsp;[Ansible](http://docs.ansible.com/)的hosts文件格式如下:


```txt

[test]
aliyun       ansible_host=47.92.102.184 ansible_user=root ansible_ssh_private_key_file=~/.ssh/aliyun
huanggai        ansible_host=192.168.1.197 ansible_user=deploy ansible_ssh_port=9122

[defualt]
ansible_ssh_private_key_file=~/.ssh/aliyun
baochai      ansible_host=192.168.1.191 ansible_user=deploy ansible_ssh_port=9122
daiyu        ansible_host=192.168.1.192 ansible_user=deploy ansible_ssh_port=9122
xiangyun     ansible_host=192.168.1.193 ansible_user=deploy ansible_ssh_port=9122
xifeng       ansible_host=192.168.1.195 ansible_user=deploy ansible_ssh_port=9122
xiren        ansible_host=192.168.1.196 ansible_user=deploy ansible_ssh_port=9122
huanggai        ansible_host=192.168.1.197 ansible_user=deploy ansible_ssh_port=9122
yuanchun     ansible_host=192.168.1.224 ansible_user=deploy ansible_ssh_port=9122
puppetmaster ansible_host=192.168.1.224 ansible_user=deploy ansible_ssh_port=9122
hp           ansible_host=192.168.1.225 ansible_user=deploy ansible_ssh_port=9122
jira         ansible_host=192.168.1.227 ansible_user=deploy ansible_ssh_port=9122
caiyun       ansible_host=192.168.1.234 ansible_user=deploy ansible_ssh_port=9122
caohong      ansible_host=192.168.2.49 ansible_user=deploy ansible_ssh_port=9122
caozhi       ansible_host=192.168.2.50 ansible_user=deploy ansible_ssh_port=9122
caopi        ansible_host=192.168.2.51 ansible_user=deploy ansible_ssh_port=9122
caochong     ansible_host=192.168.2.52 ansible_user=deploy ansible_ssh_port=9122
ceshi        ansible_host=47.92.102.184 ansible_user=deploy ansible_ssh_port=9122

```

[test]是什么意思?
> [test]表示对机器进行了分组，这里aliyun和huanggai作为一组机器，组的名字叫test，其他机器座位一组，名字叫defualt，我们可以直接对着一组机器进操作,比如，我们直接调用ping模块来查看一组机器的网络情况
```shell
bash-3.2$  ansible test -m  ping -v -i /Users/jukay/Desktop/hosts
No config file found; using defaults
huanggai | SUCCESS => {
    "changed": false,
    "ping": "pong"
}
aliyun | SUCCESS => {
    "changed": false,
    "ping": "pong"
}
```
我们也可以一组机器加若干台机器一起执行命令
```shell
bash-3.2$  ansible test,daiyu -m  ping -v -i /Users/jukay/Desktop/hosts
No config file found; using defaults
daiyu | UNREACHABLE! => {
    "changed": false,
    "msg": "Failed to connect to the host via ssh: Permission denied (publickey,gssapi-keyex,gssapi-with-mic,password).\r\n",
    "unreachable": true
}
huanggai | SUCCESS => {
    "changed": false,
    "ping": "pong"
}
aliyun | SUCCESS => {
    "changed": false,
    "ping": "pong"
}
```


ansible_host是什么意思？
> ansible_host表示目标机器的ip地址.如果这个地址如当前网络不通，那么ansible命令会返回UNREACHABLE！

ansible_user是什么意思？
> ansible_user表示登录到目标机器时候使用的用户账号。


ansible_ssh_port是什么意思?
> ansible_ssh_port是ansible通过ssh连接到目标机器时候使用的端口，默认是22号端口。