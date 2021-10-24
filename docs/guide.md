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
### Terminator
[https://blog.csdn.net/zhangkzz/article/details/90524066](https://blog.csdn.net/zhangkzz/article/details/90524066)
[https://www.cnblogs.com/xiazh/articles/2407328.html](https://www.cnblogs.com/xiazh/articles/2407328.html)

安装
```shell
sudo apt-get install terminator
```

快捷键
```plaintext
Ctrl+Shift+T	新建一个terminator窗口
Ctrl+Shift+E	对terminator窗口垂直分割
Ctrl+Shift+O	对terminator窗口水平分割
Ctrl+Shift+W	关闭当前的terminator窗口
Ctrl+Shift+I	另外创建一个新的terminator窗口
Ctrl+Shift+N	在分割的不同窗口之间切换（向后）
Ctrl+Shift+P	在分割的不同窗口之间切换（向前）
Ctrl+Shift+T	新建一个Tab窗口
Ctrl+Tab	切换下一个窗口
Ctrl+Shift+T	新建一个terminator窗口
Ctrl+Shift+F	当前terminator窗口进行搜索
Ctrl+Shift+X	当前分割的窗口最大化
Alt+A	将所有分割terminator同步
Alt+O	关闭分割terminator同步
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

### 拼音输入法
[https://blog.csdn.net/wu10188/article/details/86540464](https://blog.csdn.net/wu10188/article/details/86540464)
```shell
sudo apt-get install ibus ibus-clutter ibus-gtk ibus-gtk3 ibus-qt4
im-config -s ibus
sudo apt-get install ibus-pinyin
sudo ibus-setup
#有时候需要重启生效
```


### C/C++ 编译基础

```shell
sudo apt-get install build-essential
```

### Rust 安装
[https://www.rust-lang.org/tools/install](https://www.rust-lang.org/tools/install)

```shell
curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh
```

### neovim
[https://github.com/neovim/neovim](https://github.com/neovim/neovim)

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

### Kdeconnect

https://community.kde.org/KDEConnect#Linux_Desktop

```shell
sudo apt-get install kdeconnect
```

### AppImageLauncher

https://github.com/TheAssassin/AppImageLauncher

#### Gnome Tweak

浏览器插件：[chrome](https://chrome.google.com/webstore/detail/gnome-shell-integration/gphhapmejobijbbhgpjhcjognlahblep),[firefox](https://addons.mozilla.org/en-US/firefox/addon/gnome-shell-integration/),[opera](https://addons.opera.com/en/extensions/details/gnome-shell-integration/)



```shell
sudo apt-get install gnome-tweaks
```

#### ssh-server

```shell
sudo apt-get install openssh-server
```

#### cmake

```shell
sudo apt install cmake
```

#### go dlv
[https://github.com/go-delve/delve/blob/master/Documentation/installation/linux/install.md](https://github.com/go-delve/delve/blob/master/Documentation/installation/linux/install.md)
```shell
git clone https://github.com/go-delve/delve.git $GOPATH/src/github.com/go-delve/delve
cd $GOPATH/src/github.com/go-delve/delve
make install
```

#### go todo
```shell
go get github.com/mattn/todo
```

#### go 文档工具
```shell
git clone https://github.com/go101/golds.git
cd golds
go install
```

#### jetbrain 中文输入
大致路径：/home/lv/.local/share/JetBrains/Toolbox/apps/Goland/ch-0/203.7148.71/bin/


ibus
```shell
export GTK_IM_MODULE=ibus
export QT_IM_MODULE=ibus
export XMODIFIERS=@im=ibus
```

fcitx
```shell
export GTK_IM_MODULE=fcitx
export QT_IM_MODULE=fcitx
export XMODIFIERS=@im=fcitx
```

#### ubutu（微信、QQ、腾讯会议）wine

```shell
wget -O- https://deepin-wine.i-m.dev/setup.sh | sh

sudo apt-get install com.qq.weixin.deepin
sudo apt-get install  com.tencent.meeting.deepin
```

[https://my.oschina.net/frank1126/blog/4947220](https://my.oschina.net/frank1126/blog/4947220)
[https://github.com/zq1997/deepin-wine](https://github.com/zq1997/deepin-wine)


### reveal.js

```shell
git clone https://github.com/hakimel/reveal.js.git
cd reveal.js && npm install
```