### security

error  `unable to write 'random state'`

```shell
export RANDFILE=.rnd
```



```shell
# 创建 CA 私钥 ca-key.pem
sudo openssl genrsa -aes256 -out ca-key.pem 4096

#  创建用来签名的公钥 ca.pem
sudo openssl req -new -x509 -days 365 -key ca-key.pem -sha256 -out ca.pem
```



```shell
Country Name (2 letter code) [AU]:CN  # 输入国家代码
State or Province Name (full name) [Some-State]:Beijing
Locality Name (eg, city) []:Beijing
Organization Name (eg, company) [Internet Widgits Pty Ltd]:Lvsoso # 输入组织名
Organizational Unit Name (eg, section) []:Lvsoso
Common Name (e.g. server FQDN or YOUR name) []:localhost # 输入服务器域名
Email Address []:wuqize5109@163.com
```

#### 服务端证书

```shell
sudo openssl genrsa -out server-key.pem 4096
sudo openssl req -subj "/CN=localhost" -sha256 -new -key server-key.pem -out server.csr
sudo echo subjectAltName = IP:127.0.0.1 >> extfile.cnf

# 设置仅用于服务器身份验证
sudo echo extendedKeyUsage = serverAuth >> extfile.cnf

sudo openssl x509 -req -days 365 -sha256 -in server.csr -CA ca.pem -CAkey ca-key.pem -CAcreateserial -out server-cert.pem -extfile extfile.cnf
```

#### 客户端证书

```shell
sudo openssl genrsa -out client-key.pem 4096
sudo openssl req -subj '/CN=client' -new -key client-key.pem -out client.csr

sudo echo extendedKeyUsage = clientAuth > client-extfile.cnf

sudo openssl x509 -req -days 365 -sha256 -in client.csr -CA ca.pem -CAkey ca-key.pem -CAcreateserial -out cert.pem -extfile client-extfile.cnf
```

#### 控制权限

```shell
sudo chmod -v 0400 ca-key.pem server-key.pem client-key.pem
sudo chmod -v 0444 ca.pem server-cert.pem cert.pem
```

#### 配置

```shell
sudo dockerd --tlsverify --tlscacert=ca.pem --tlscert=server-cert.pem --tlskey=server-key.pem -H=0.0.0.0:2376
```

#### 连接

```shell
sudo docker --tlsverify --tlscacert=ca.pem --tlscert=cert.pem --tlskey=client-key.pem -H=127.0.0.1:2376 image ls
```

#### alias

```shell
alias docker='sudo docker --tlsverify --tlscacert=ca.pem --tlscert=cert.pem --tlskey=client-key.pem -H=127.0.0.1:2376'

docker image ls
```

### `--privileged=true`

```
docker inspect lvsoso | grep Privileged
docker inspect prilvsoso | grep Privileged
```

### `--cap-add`

```shell
docker container inspect -f {{.HostConfig.Privileged}} caplvsoso
false

docker container inspect -f {{.HostConfig.CapAdd}} caplvsoso
{[NET_ADMIN]}
```



### Docker Bench Security

[https://github.com/docker/docker-bench-security/](https://github.com/docker/docker-bench-security/)