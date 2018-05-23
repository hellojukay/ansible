# <span id="config">Ansible的配置</span>
ansible一般都有一个ansible.cfg配置，可以指定ssh，hosts，library,roles等信息。这里要说明一下ansible是怎么查找这个配置文件的，查找书心虚如下:

* 从环境变量ANSIBLE_CONFIG
* 当前目录的ansible.cfg
* home目录的隐藏文件.ansible.cfg
* /etc/ansible.cf文件