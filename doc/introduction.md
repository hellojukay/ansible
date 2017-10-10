# <span id="desc">Ansible简介</span>
&nbsp;&nbsp;&nbsp;&nbsp;[Ansible](http://docs.ansible.com/)是新出现的自动化运维工具，基于Python开发，集合了众多运维工具（puppet、cfengine、chef、func、fabric）的优点，实现了批量系统配置、批量程序部署、批量运行命令等功能。
## <span id="tedian">Ansible的特点</span>
&nbsp;&nbsp;&nbsp;&nbsp;[Ansible](http://docs.ansible.com/)是基于模块工作的，ansible本身是没有批量部署的能力，他只是提供了服务端通信功能，模块执行功能，以及一些核心的模块。还有一个比较特别的地方是:[Ansible](http://docs.ansible.com/)是不要安装客户端的，他只要求目标机器上安装了Python2.4以及以上版本即可。同时Ansible也是支持windows的.

## <span id="helloworld">HelloWorld</span>

```shell
 ~/C/M/ansible  ansible huanggai -m command -a "echo hello world"                                                             1009ms  六  9/30 13:39:21 2017
huanggai | SUCCESS | rc=0 >>
hello world
```
解释说明:huanggai 是指配置在/etc/ansible/hosts文件中的机器， -m 表示选择模块 command是[Ansible](http://docs.ansible.com/)的内置模块，-a表示的是command模块的参数
## <span id="anzhuang">安装Ansible</span>
* mac:<br> 
brew install ansible
* centos: <br>
yum install ansible
*  ubuntu: <br>
apt-get install ansible