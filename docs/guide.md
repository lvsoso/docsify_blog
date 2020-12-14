## Guide

> guide.



#### git 

##### 缓存密码

```shell
vim ~/.gitconfig

[credential "http://10.0.0.20:3680"]
username = wuqz

[credential "https://github.com/lvoooo"]
username = lvoooo

[credential "https://github.com/lvsoso"]
username = lvsoso


[credential]
    helper = cache --timeout 86400
    helper = store --file ~/.git-credentials

```



#### node js

##### 配置

```shell
# 列出基本配置
npm config list

# 获取全局包的位置
npm config get prefix

# 设置全局包的位置
npm config set prefix "" 

# 设置cache的位置
npm config set cache ""
```



