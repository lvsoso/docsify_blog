## Guide

> guide.


### 软件源（Ubuntu）

[清华软件源](https://mirrors.tuna.tsinghua.edu.cn/help/ubuntu/)

```text
# 默认注释了源码镜像以提高 apt update 速度，如有需要可自行取消注释
deb https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ focal main restricted universe multiverse
# deb-src https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ focal main restricted universe multiverse
deb https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ focal-updates main restricted universe multiverse
# deb-src https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ focal-updates main restricted universe multiverse
deb https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ focal-backports main restricted universe multiverse
# deb-src https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ focal-backports main restricted universe multiverse
deb https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ focal-security main restricted universe multiverse
# deb-src https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ focal-security main restricted universe multiverse

# 预发布软件源，不建议启用
# deb https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ focal-proposed main restricted universe multiverse
# deb-src https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ focal-proposed main restricted universe multiverse
```


### Git 

##### 提交代码

```shell
echo "# rustbook" >> README.md
git init
git add README.md
git commit -m "first commit"
git branch -M main
git remote add origin https://github.com/lvsoso/rustbook.git
git push -u origin main
```

```shell
git remote add origin https://github.com/lvsoso/rustbook.git
git branch -M main
git push -u origin main
```



##### 缓存密码

```shell
vim ~/.gitconfig

[credential "http://1.2.3.4:5678"]
username = xxx

[credential "https://github.com/yyy"]
username = yyy

[credential "https://github.com/zzz"]
username = zzz


[credential]
    helper = cache --timeout 86400
    helper = store --file ~/.git-credentials

[alias]
    co = checkout                                       
    br = branch                                         
    ci = commit                                         
    st = status

```



#### Node js

##### 配置

```shell
# 列出基本配置
npm config list

# 获取全局包的位置
npm config get prefix

# 设置全局包的位置
npm config set prefix "/home/lv/soft/node_module" 

# 设置cache的位置
npm config set cache "/home/lv/soft/node_cache" 
```


### Docker安装（Ubuntu）

[官方文档](https://docs.docker.com/engine/install/ubuntu/)

```shell
sudo apt-get remove docker docker-engine docker.io containerd runc
sudo apt-get update
sudo apt-get install \
    apt-transport-https \
    ca-certificates \
    curl \
    gnupg-agent \
    software-properties-common
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
sudo apt-key fingerprint 0EBFCD88
sudo add-apt-repository \
   "deb [arch=amd64] https://download.docker.com/linux/ubuntu \
   $(lsb_release -cs) \
   stable"
sudo apt-get update
sudo apt-get install docker-ce docker-ce-cli containerd.io
```

Docker安装（CentOS8.2）

[官方文档](https://docs.docker.com/engine/install/centos/)

```shell
sudo yum remove docker \
                  docker-client \
                  docker-client-latest \
                  docker-common \
                  docker-latest \
                  docker-latest-logrotate \
                  docker-logrotate \
                  docker-engine

sudo yum install -y yum-utils

sudo yum-config-manager \
    --add-repo \
    https://download.docker.com/linux/centos/docker-ce.repo
    
sudo yum install -y  docker-ce docker-ce-cli containerd.io


```



### Proxychain-ng

```shell
git clone https://github.com/rofl0r/proxychains-ng.git
cd proxychains-ng
./config
make
sudo make install
sudo make install-config
```

### Docsify

```shell
docker build -f Dockerfile -t docsify/demo .
docker run -itdp 3000:3000 --restart=always --name=docsify -v $(pwd):/docs docsify/demo
```

### Kedconnect

https://community.kde.org/KDEConnect#Linux_Desktop

```shell

```

### AppImageLauncher

https://github.com/TheAssassin/AppImageLauncher