## docker

### OCI
- 镜像标准定义应用如何打包
- 运行时标准定义如何解压应用包并运行
- 分发标准定义如何分发容器镜像

docker run xxx
|   http
daemon 
|   grpc
containerd
|   
shim
|   
runc

### 网络
- null 独立网络空间，不做任何网络配置
- host 使用主机网络名空间，复用主机网络
- container 重用其他容器的网络
- bridge 使用linux网桥和iptables提供容器互联，docker0，veth pair
- overlay  网络封包实现
- remote underlay， overlay 
### 文件

**fs**

- bootfs
  - Bootloader
  - kernel
- rootfs

docker 复用文件系统 rootfs，并设置“readonly” ，并在这个基础上进行叠加，设置权限，叠加...

对于文件，写时复制（写的时候复制一份进行修改，不影响别的），用时分配（用得时候再分配空间）。

**存储驱动**

AUFS、OverlayFS、Device Mapper、Btrfs、ZFS

**OverlayFS**

upper 层和 lower 层,Lower 层代表镜像层,upper 层代表容器可写层。

最上层merged层，对文件进行合并，形成合并的视图。


### 常用服务

```shell
docker run --name some-mysql -p3306:3306 -e MYSQL_ROOT_PASSWORD=123456 -d  mysql:latest --character-set-server=utf8mb4 --collation-server=utf8mb4_unicode_ci

docker run -d -p 6379:6379  --name redis -e ALLOW_EMPTY_PASSWORD=yes bitnami/redis:latest
```

### `ENTRYPOINT`和`CMD`
1. 容器启动后会自行运行；
2. 后边的`ENTRYPOINT`和`CMD`会覆盖前者；
3. 命令行参数会覆盖`CMD`，而`ENTRYPOINT`则不会被覆盖；
4. exec表示法`CMD ["/bin/ping","localhost"]`, `/bin/sh -c 'ping localhost'`；
5. shell表示法`CMD  "/bin/ping" "localhost"`, `"/bin/ping localhost`;


###  安装docker-engine

```shell
# https tool
sudo apt-get -y install apt-transport-https ca-certificates curl software-properties-common

# net tool
sudo apt install net-tools

# ping
sudo apt install -y iputils-ping`

# aliyun key
curl -fsSL http://mirrors.aliyun.com/docker-ce/linux/ubuntu/gpg | sudo apt-key add -

# aliyun repo
sudo add-apt-repository "deb [arch=amd64] http://mirrors.aliyun.com/docker-ce/linux/ubuntu $(lsb_release -cs) stable"

# 更新 apt 索引库
sudo apt-get update

# 安装 docker-ce
sudo apt-get install docker-ce
```

### /etc/docker/daemon.json

```json
{
  "registry-mirrors": ["https://n6syp70m.mirror.aliyuncs.com"]
}
```

### command

```shell
docker system info
docker info
docker --help
```

##### docker container

```shell
docker container ls -a -q
docker container run [OPTIONS] IMAGE [COMMAND [ARGS...]]

docker container run busybox echo "hello lvsoso"
docker container run -i -t ubuntu /bin/bash
docker container run -i -t -d  ubuntu /bin/bash
```

 [OPTIONS] 

- `-i` 或 `--interactive`， 交互模式。
- `-t` 或 `--tty`， 分配一个 `pseudo-TTY`，即伪终端。
- `-d` 或 `--detach`，在后台运行容器并输出容器的 `ID` 到终端。
- `--rm` 在容器退出后自动移除。
- `-p` 将容器的端口映射到主机。
- `-v` 或 `--volume`， 指定数据卷。

###### create

```shell
# 只创建容器，并不会运行容器。
docker container create [OPTIONS] IMAGE [COMMAND] [ARG...]

docker container create \
    --name lvsoso1 \
    --hostname lvsoso1 \
    --mac-address 00:01:02:03:04:05 \
    --ulimit nproc=1024:2048 \
    -it ubuntu /bin/bash
```

[OPTIONS]

- `--name` 指定一个容器名称，未指定时，会随机产生一个名字

- `--hostname` 设置容器的主机名

- `--mac-address` 设置 `MAC` 地址

- `--ulimit` 设置 `ulimit` 选项

  

  ulimit [https://www.ibm.com/developerworks/cn/linux/l-cn-ulimit/](https://www.ibm.com/developerworks/cn/linux/l-cn-ulimit/)

  set ulimit [https://access.redhat.com/solutions/61334](https://access.redhat.com/solutions/61334)

###### start

```shell
docker container start [OPTIONS] CONTAINER [CONTAINER...]

docker container start lvsoso1
```

###### stop

```shell
docker container stop CONTAINER [CONTAINER...]
```

###### restart

```shell
docker container restart CONTAINER [CONTAINER...]
```

###### pause

```shell
docker container pause CONTAINER [CONTAINER...]
```

###### unpause

```shell
docker container unpause CONTAINER [CONTAINER...]
```

###### ls

```shell
docker container ls [OPTIONS]
```

 [OPTIONS]

- `-a` 显示所有的容器。
- `-q` 仅显示 `ID`。
- `-s` 显示总的文件大小。

###### attach

```shell
docker attach [OPTIONS] CONTAINER
```

###### inspect

```shell
docker container inspect [OPTIONS] CONTAINER [CONTAINER...]

docker container inspect shiyanlou01 | grep "MacAddress"
```

###### logs

```shell
docker container logs [OPTIONS] CONTAINER
```

[OPTIONS]

- `-t` 或 `--timestamps` 显示时间戳
- `-f` 实时输出，类似于 `tail -f`

###### top

```shell
docker container top CONTAINER
```

###### diff

```shell
docker container diff CONTAINER
```

###### exec

```shell
docker container exec [OPTIONS] CONTAINER COMMAND [ARG...]
```

###### rm

```shell
docker container rm [OPTIONS] CONTAINER [CONTAINER...]
```

```
docker container prune [OPTIONS]
```

[OPTIONS]

- `-f` 或 `--force` 跳过警告提示
- `--filter` 执行过滤删除

##### docker image

```shell
docker image ls
docker image pull ubuntu
docker image inspect ubuntu
docker search ubuntu
```

```shell
docker image pull [OPTIONS] NAME[:TAG|@DIGEST]

docker image build [OPTIONS] PATH | URL

docker image rm ubuntu:latest
```

### storage

Docker 提供三种不同的方式将数据从 Docker 主机挂载到容器中，分别为卷（`volumes`），绑定挂载（`bind mounts`），临时文件系统（`tmpfs`）。

- `volumes` 卷存储在 Docker 管理的主机文件系统的一部分中（`/var/lib/docker/volumes/`）中，完全由 Docker 管理。
- `bind mounts` 绑定挂载，可以将主机上的文件或目录挂载到容器中。
- `tmpfs` 仅存储在主机系统的内存中，而不会写入主机的文件系统。

#### volums

```shell
docker volume ls

# 创建一个数据卷，并且会随机生成一个名称
docker volume create
# 指定名称 volume1
docker volume create volume1

```

```shell
-v 或 --volume
```

- 由三个由冒号（:）分隔的字段组成，`[HOST-DIR:]CONTAINER-DIR[:OPTIONS]`。
- `HOST-DIR` 代表主机上的目录或数据卷的名字。省略该部分时，会自动创建一个匿名卷。如果是指定主机上的目录，需要**使用绝对路径**。
- `CONTAINER-DIR` 代表将要挂载到容器中的目录或文件，即表现为容器中的某个目录或文件
- `OPTIONS` 代表配置，例如设置为只读权限（`ro`），此卷仅能被该容器使用（`Z`），或者可以被多个容器使用（`z`）。多个配置项由逗号分隔。

```shell
# 将卷`volume1`挂载到容器`/volume1`目录
# `ro,z` 代表该卷被设置为只读（`ro`），并且可以多个容器使用该卷（`z`）。
docker container run -v volume1:/volume1:ro,z 
```
#### mount

```
--mount
```

- 由多个键值对组成，键值对之间由逗号分隔。例如：`type=volume,source=volume1,destination=/volume1,ro=true`。
- `type`，指定类型，可以指定为 `bind`，`volume`，`tmpfs`。
- `source`，当类型为 `volume` 时，指定卷名称，匿名卷时省略该字段。当类型为 `bind`，指定路径。可以使用缩写 `src`。
- `destination`，挂载到容器中的路径。可以使用缩写 `dst` 或 `target`。
- `ro` 为配置项，多个配置项直接由逗号分隔一般使用 `true` 或 `false`。

```shell
docker container run -it --name lvsoso \
    -v volume1:/volume1 \
    --rm ubuntu /bin/bash
    
docker container run -it --name lvsoso \
    --mount type=volume,src=volume1,target=/volume1 \
    --rm ubuntu /bin/bash
```

#### tmpfs

```shell
docker run \
    -it \
    --mount type=tmpfs,target=/test \
    --name lvsoso \
    --rm ubuntu bash
```



#### 文件夹处理

对于数据卷来说，由 Docker 完全管理，而绑定挂载，则需要我们自己去维护。

> 文件使用挂载类型的情况下，外部修改文件，内部是无法感知的

host文件夹存在且非空

- 如果container不存在则创建，并把host对应文件夹中的内容拷贝进去；
- 如果container存在则先清空，并把host对应文件夹中的内容拷贝进去；

host文件夹存在为空

- 如果container存在文件夹，在host上创建该文件夹，并清空container文件夹；
- 如果container不存在文件夹，都创建；

#### 文件处理

- **docker 禁止用主机上不存在的文件挂载到container中已经存在的文件**
- 文件挂载不会对同一文件夹下的其他文件产生任何影响
- 用host上的文件的内容覆盖container中的文件的内容

host上文件不存在

- host不存在，container存在，异常；

host文件存在

- host文件存在，container文件存在，内容被host文件覆盖；
- host文件存在，container文件不存在，container会创建相应文件，内容和host中的文件内容一致；

#### 数据卷容器

数据卷容器就是一种普通容器，它专门提供数据卷供其它容器挂载使用。

对容器`ShareVolume`的`/vdata`目录下的文件修改是能够被test1和test2两个容器感知的。

```shell
#  创建数据卷
docker volume create vdata

# 运行数据卷容器
docker container run \
    -it \
    -v vdata:/vdata \
    --name ShareVolume ubuntu /bin/bash
   
 # 运行测试容器test1
docker container run \
    -it \
    --volumes-from ShareVolume \
    --name test1 ubuntu /bin/bash    

# 运行测试容器 test2
docker container run \
    -it \
    --volumes-from ShareVolume \
    --name test2 ubuntu /bin/bash
```

#### 数据备份

```shell
docker container run \
# 备份容器继承容器 ShareVolume 数据卷
   --volumes-from ShareVolume \
# 主机上绑定挂载方式进行挂载
   -v /home/lvsoso/backup:/backup \
# 压缩归档
   ubuntu tar cvf /backup/backup.tar /vdata/
```



#### 数据恢复

```shell
docker container run \
   --volumes-from ShareVolume \
   -v /home/lvsoso/backup:/backup \
   ubuntu tar xvf /backup/backup.tar -C /
```

###  network

```shell
docker network ls

NETWORK ID          NAME                DRIVER              SCOPE
28f9126dee48        bridge              bridge              local
3ea24cc139cd        host                host                local
bea2bd0a321e        none                null                local
a901f27cf74e        oauth20_default     bridge              local
```

#### 桥接网络

容器run的时候默认使用桥接网络`--network bridge`

```shell
# 查看 bridge 网络的详细信息，并通过 grep 获取名称项
docker network inspect bridge | grep name

# 使用 ifconfig 查看 docker0 网络
ifconfig
```

Docker 对于 bridge 网络使用端口的方式为设置端口映射，通过 `iptables` 实现。

```shell
# 查看nat转发规则
sudo iptables -t nat -nvL

# 查看所有转发规则
sudo iptables -nvL
```

#### 自定义网络

容器`link`互联

```shell
 # `name` 为容器名，`alias` 为别名。
 # /etc/hosts 下会进行配置解析
 # 但是之后并不能自动更新
 --link <name or id>:alias
```

创建自己的 bridge 或 `overlay` 网络，并且docker的DNS服务支持连接到网络的容器名解析。

```shell
docker network create network1
docker network ls
docker network rm network1
docker network create -d bridge --subnet=192.168.5.0/24 --gateway=192.168.5.1 network1
docker run -it --name lvsoso1 --network network1 --rm busybox /bin/sh
```

```shell
docker run -it --name lvsoso2 --rm busybox /bin/sh

# 将 lvsoso2 也接进来 network1
# 这时候 lvsoso2 会有两个网络接口
docker network connect network1 lvsoso2
```

```shell
# 自定义网络可以指定使用的ip
docker run -it --network network1 --ip 192.168.5.200 --rm busybox /bin/sh
```

#### host 和 none

当容器的网络为 `host` 时，容器可以直接访问主机上的网络。

`none` 网络，容器中不提供其它网络接口。

### Dockerfile

```dockerfile
# 指定基础镜像
# 后续指令将在此镜像的基础上运行
FROM ubuntu:16.04

# 维护者信息
MAINTAINER wuqilv/wuqize5109@163.com

# 指定用户，后续的 `RUN`，`CMD` 以及 `ENTRYPOINT` 指令都会使用该用户去执行，但是该用户必须提前存在。
USER lvsoso

# 命令执行时的当前目录
WORKDIR /

# 镜像操作命令
RUN \
    apt-get -yqq update && \
    apt-get install -yqq apache2

# 容器启动命令
# CMD 指令的值会被当作 ENTRYPOINT 指令的参数附加到 ENTRYPOINT 指令的后面
# 相当 `ls -a -l`
ENTRYPOINT ["ls", "-a"]
CMD ["-l"]
```

RUN

- `RUN <command>`，使用 shell 去执行指定的命令 `command`，一般默认的 shell 为 `/bin/sh -c`。
- `RUN ["executable", "param1", "param2", ...]`，使用可执行的文件或程序 `executable`，给予相应的参数 `param`。

CMD

- 一个 Dockerfile 文件中只能有一个 CMD 指令，如果有多个 CMD 指令，则只有最后一个会生效。
- 该指令为我们运行容器时提供默认的命令
- 可以在 `docker run` 时被覆盖

ENTRYPOINT

- ENTRYPOINT 指令会覆盖 CMD 指令作为容器运行时的默认指令

- 不会在 `docker run` 时被覆盖

COPY 和 ADD

```shell
ADD <src>... <dest>
ADD ["<SRC>",... "<dest>"]

COPY <src>... <dest>
COPY ["<src>",... "<dest>"]
```

- `<src>` 可以指定多个，但是其路径不能超出上下文的路径，即必须在跟 Dockerfile **同级**或**子目录**中。

- `<dest>` 不需要预先存在，不存在路径时会自动创建，如果没有使用绝对路径，则 `<dest>` 为**相对于工作目录的相对路径**。

- ADD 可以添加远程路径的文件，并且 `<src>` 为可识别的压缩格式，如 gzip 或 tar 归档文件等，ADD 会自动将其解压缩为目录。

ENV 环境变量

```dockerfile
ENV <key> <value>
ENV <key>=<value> <key>=<value>...
```

VOLUME

在容器运行时，创建两个匿名卷，并挂载到容器中的 `/data1` 和 `/data2` 目录上。

```
VOLUME /data1 /data2
```

EXPOSE

与 `docker run` 命令的 `-p` 参数不一样，并不实际映射端口，只是将该端口暴露出来，允许外部或其它的容器进行访问。

```dockerfile
EXPOSE port
```

