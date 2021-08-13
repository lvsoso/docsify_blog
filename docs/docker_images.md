### Image

### 基本框架

```dockerfile
# Version 0.1

# 基础镜像
FROM ubuntu:14.04

# 维护者信息
MAINTAINER wuqize5109@163.com

# 镜像操作命令
RUN echo "deb http://mirrors.cloud.aliyuncs.com/ubuntu/ trusty main universe" > /etc/apt/sources.list
RUN apt-get update && apt-get install -yqq supervisor && apt-get clean

# 容器启动命令
CMD ["supervisord"]
```

#### SSH

```dockerfile
RUN apt-get install -yqq openssh-server openssh-client

RUN mkdir -p /var/run/sshd

RUN echo 'root:lvsoso' | chpasswd
RUN sed -i 's/PermitRootLogin without-password/PermitRootLogin yes/' /etc/ssh/sshd_config
```



```dockerfile
# 指定基础镜像
FROM ubuntu:14.04

# 安装软件
RUN apt-get update && apt-get install -y openssh-server && mkdir /var/run/sshd

# 添加用户 lvsoso 及设定密码
RUN useradd -g root -G sudo lvsoso && echo "lvsoso:123456" | chpasswd lvsoso

# 暴露 SSH 端口
EXPOSE 22

CMD ["/usr/sbin/sshd", "-D"]
```

```shell
docker build -t sshd:test .

docker run -itd -p 10001:22 sshd:test
```



#### MongoDB

supervisord.conf

```conf
[supervisord]
nodaemon=true

[program:mongodb]
command=/opt/mongodb/bin/mongod

[program:ssh]
command=/usr/sbin/sshd -D
```



```dockerfile
# Version 0.1

# 基础镜像
FROM ubuntu:14.04

# 维护者信息
MAINTAINER wuqize5109@163.com

# 镜像操作命令
#RUN echo "deb http://mirrors.cloud.aliyuncs.com/ubuntu/ trusty main universe" > /etc/apt/sources.list
RUN apt-get -yqq update && apt-get install -yqq supervisor
RUN apt-get install -yqq openssh-server openssh-client

# ssh
RUN mkdir -p /var/run/sshd
RUN echo 'root:lvsoso' | chpasswd
RUN sed -i 's/PermitRootLogin without-password/PermitRootLogin yes/' /etc/ssh/sshd_config

RUN mkdir -p /opt
ADD https://labfile.oss.aliyuncs.com/courses/498/mongodb-linux-x86_64-ubuntu1404-3.2.3.tgz /opt/mongodb.tar.gz
RUN cd /opt && tar zxvf mongodb.tar.gz && rm -rf mongodb.tar.gz
RUN mv /opt/mongodb-linux-x86_64-ubuntu1404-3.2.3 /opt/mongodb

RUN mkdir -p /data/db
ENV PATH=/opt/mongodb/bin:$PATH

COPY supervisord.conf /etc/supervisor/conf.d/supervisord.conf

EXPOSE 27017 22

# 容器启动命令
CMD ["supervisord"]



```



```shell
docker build -t mongodb:0.1 .

#  `-P` 参数将容器 EXPOSE 的端口随机映射到主机的端口。
docker run -P -d --name mongodb mongodb:0.1

docker exec -it mongodb /bin/bash
mongo --host 127.0.0.1 --port 27017
```




#### Redis

supervisord.conf

```conf
[supervisord]
nodaemon=true

[program:redis]
command=/usr/bin/redis-server

[program:ssh]
command=/usr/sbin/sshd -D
```

```dockerfile
# Version 0.1

# 基础镜像
FROM ubuntu:14.04

# 维护者信息
MAINTAINER wuqize5109@163.com

# 镜像操作命令
#RUN echo "deb http://mirrors.cloud.aliyuncs.com/ubuntu/ trusty main universe" > /etc/apt/sources.list
RUN apt-get -yqq update && apt-get install -yqq supervisor
RUN apt-get install -yqq openssh-server openssh-client

# ssh
RUN mkdir -p /var/run/sshd
RUN echo 'root:lvsoso' | chpasswd
RUN sed -i 's/PermitRootLogin without-password/PermitRootLogin yes/' /etc/ssh/sshd_config

# redis
RUN apt-get install -yqq redis-server

COPY supervisord.conf /etc/supervisor/conf.d/supervisord.conf


EXPOSE 6379 22

# 容器启动命令
CMD ["supervisord"]
```

```shell
docker image build -t redis:0.1 .
docker run -P -d --name redis redis:0.1

ssh root@127.0.0.1 -p 32771

redis-cli -h 127.0.0.1 -p 6379
```

#### WordPress

supervisord.conf

```conf
[supervisord]
nodaemon=true

[program:php5-fpm]
command=/usr/sbin/php5-fpm -c /etc/php5/fpm
autorstart=true

[program:mysqld]
command=/usr/bin/mysqld_safe

[program:nginx]
command=/usr/sbin/nginx
autorstart=true

[program:ssh]
command=/usr/sbin/sshd -D
```

nginx-config

```conf
server {
 listen *:80;
 server_name localhost;

    root /var/www/wordpress;
    index index.php;

    location / {
        try_files $uri $uri/ /index.php?$args;
    }

    location ~ \.php{
        try_files $uri =404;
        include fastcgi_params;
        fastcgi_pass unix:/var/run/php5-fpm.sock;
    }
}
```

```dockerfile
# Version 0.1
FROM ubuntu:14.04

# 维护者信息
MAINTAINER wuqize5109@163.com

# 镜像操作命令
#RUN echo "deb http://mirrors.cloud.aliyuncs.com/ubuntu/ trusty main universe" > /etc/apt/sources.list
RUN apt-get -yqq update && apt-get install -yqq supervisor
RUN apt-get install -yqq openssh-server openssh-client

RUN apt-get -yqq install nginx supervisor wget php5-fpm php5-mysql
# 将 nginx 设置成非守护模式
RUN echo "daemon off;" >> /etc/nginx/nginx.conf

RUN mkdir -p /var/www
ADD https://labfile.oss-cn-hangzhou-internal.aliyuncs.com/courses/498/wordpress-4.4.2.tar.gz /var/www/wordpress-4.4.2.tar.gz
RUN cd /var/www && tar zxvf wordpress-4.4.2.tar.gz && rm -rf wordpress-4.4.2.tar.gz
RUN chown -R www-data:www-data /var/www/wordpress

RUN mkdir /var/run/sshd
RUN apt-get install -yqq openssh-server openssh-client
RUN echo 'root:shiyanlou' | chpasswd
RUN sed -i 's/PermitRootLogin without-password/PermitRootLogin yes/' /etc/ssh/sshd_config

RUN echo "mysql-server mysql-server/root_password password shiyanlou" | debconf-set-selections
RUN echo "mysql-server mysql-server/root_password_again password shiyanlou" | debconf-set-selections
RUN apt-get  install -yqq mysql-server mysql-client

EXPOSE 80 22

COPY nginx-config /etc/nginx/sites-available/default
COPY supervisord.conf /etc/supervisor/conf.d/supervisord.conf
RUN service mysql start && mysql -uroot -pshiyanlou -e "create database wordpress;"
RUN sed -i 's/database_name_here/wordpress/g' /var/www/wordpress/wp-config-sample.php
RUN sed -i 's/username_here/root/g' /var/www/wordpress/wp-config-sample.php
RUN sed -i 's/password_here/shiyanlou/g' /var/www/wordpress/wp-config-sample.php
RUN mv /var/www/wordpress/wp-config-sample.php /var/www/wordpress/wp-config.php

CMD ["/usr/bin/supervisord"]
```



```shell
docker build -t wordpress:0.1 .

docker run -d -p 80:80 --name wordpress wordpress:0.1
```

### Docker Registry

```shell
# 运行 registry
docker run -d -p 5000:5000 --restart=always --name registry registry:2.7.0

# tag redis
docker pull redis
docker tag redis:latest localhost:5000/my-redis

# 推送 redis
docker push localhost:5000/my-redis

# 删除 本地原来的 redis
docker image remove redis:latest
docker image remove localhost:5000/my-redis

# 测试拉取
docker pull 127.0.0.1:5000/my-redis
```

#### 镜像存储

```shell
docker run -d -p 5000:5000 \
	--restart=always  \
	--name registry  \
	-v /home/lvsoso/data:/var/lib/registry   registry:2.7.0
```

#### 认证机制

```shell
docker run --entrypoint htpasswd registry:2.7.0 -Bbn user pass > auth/htpasswd
```

```shell
docker run -d \
  -p 5000:5000 \
  --restart=always \
  --name registry \
  -v $(pwd)/auth:/auth \
  -e "REGISTRY_AUTH=htpasswd" \
  -e "REGISTRY_AUTH_HTPASSWD_REALM=Registry Realm" \
  -e REGISTRY_AUTH_HTPASSWD_PATH=/auth/htpasswd \
  registry:2.7.0
```

```shell
docker login localhost:5000
```

#### 详细配置

`/etc/docker/registry/config.yml`

##### 日志级别

设置 `REGISTRY_LOG_LEVEL`，`REGISTRY_LOG_FORMATTER` 来设置日志输出级别和格式，级别可以设置为 `error`，`warn`，`info`，`debug`，默认为 `info` 级别。输出格式可以为 `json`，`text`，`logstash`。

##### Hook

```yaml
hooks:
  - type: mail
    levels:
      - panic
    options:
      smtp:
        addr: smtp.sendhost.com:25
        username: sendername
        password: password
        insecure: true
      from: name@sendhost.com
      to:
        - name@receivehost.com
```

##### 存储设置

```yaml
storage:
  filesystem:
    rootdirectory: /var/lib/registry
  azure:
    accountname: accountname
    accountkey: base64encodedaccountkey
    container: containername
  gcs:
    bucket: bucketname
    keyfile: /path/to/keyfile
    rootdirectory: /gcs/object/name/prefix
  s3:
    accesskey: awsaccesskey
    secretkey: awssecretkey
    region: us-west-1
    bucket: bucketname
    encrypt: true
    secure: true
    v4auth: true
    chunksize: 5242880
    rootdirectory: /s3/object/name/prefix
  rados:
    poolname: radospool
    username: radosuser
    chunksize: 4194304
  swift:
    username: username
    password: password
    authurl: https://storage.myprovider.com/auth/v1.0 or https://storage.myprovider.com/v2.0 or https://storage.myprovider.com/v3/auth
    tenant: tenantname
    tenantid: tenantid
    domain: domain name for Openstack Identity v3 API
    domainid: domain id for Openstack Identity v3 API
    insecureskipverify: true
    region: fr
    container: containername
    rootdirectory: /swift/object/name/prefix
  oss:
    accesskeyid: accesskeyid
    accesskeysecret: accesskeysecret
    region: OSS region name
    endpoint: optional endpoints
    internal: optional internal endpoint
    bucket: OSS bucket
    encrypt: optional data encryption setting
    secure: optional ssl setting
    chunksize: optional size valye
    rootdirectory: optional root directory
  inmemory:
  delete:
    enabled: false
  cache:
    blobdescriptor: inmemory
  maintenance:
    uploadpurging:
      enabled: true
      age: 168h
      interval: 24h
      dryrun: false
  redirect:
    disable: false
```

##### 认证设置

```yaml
auth:
  token:
    realm: token-realm
    service: token-service
    issuer: registry-token-issuer
    rootcertbundle: /root/certs/bundle
```

##### HTTP 设置

```yaml
http:
  addr: localhost:5000
  net: tcp
  prefix: /my/nested/registry/
  host: https://myregistryaddress.org:5000
  secret: asecretforlocaldevelopment
  tls:
    certificate: /path/to/x509/public
    key: /path/to/x509/private
    clientcas:
      - /path/to/ca.pem
      - /path/to/another/ca.pem
  debug:
    addr: localhost:5001
  headers:
    X-Content-Type-Options: [nosniff]
```

