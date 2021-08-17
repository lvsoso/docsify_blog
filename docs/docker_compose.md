### Compose



```yaml
version: '3.0'

services:
  redis:
    image: redis:6.0
  web:
    build:
      context: /home/lvsoso/app/web
    depends_on:
      - redis
    ports:
      - 8001:80/tcp
    volumes:
      - /home/lvsoso/app/web:/web:rw

```

```shell
docker-compose up
docker-compose up -d
docker-compose down
```

### Swarm

一个服务是任务在管理节点或工作节点执行的定义，服务中运行的单个容器被称为任务。

使用集群模式运行服务时，一般有两种选项：

- `replicated services`（复制服务），根据设定的值，swarm 调度在节点之间运行指定的副本任务。
- `global services`（全局服务），集群在每个可用节点上运行一项任务。

```shell

docker swarm init --advertise-addr <IP>
docker swarm init --advertise-addr=eth0

docker node ls

# 获取添加管理节点的命令
docker swarm join-token manager

# 获取添加工作节点的命令
docker swarm join-token worker

docker node ls

# 移除节点
# NODE 为该节点的 ID
docker node rm NODE

# 提权
# 使之成为管理节点
docker node promote NODE

# 撤权
# 使之成为工作节点
docker node demote NODE
```

### 集群模式部署

```shell
docker stack deploy -c docker-compose.yml app

# 查看所有service
docker service ls

# 查看所有stack
docker stack ls

# 查看指定stack的service
docker stack services app

# 查看 app stack的任务
docker stack ps app

# 查看服务 app_redis 的任务
docker service ps app_redis

# 移除一个satck，将会移除satck中的所有services
docker stack rm app

# 移除一个或多个service
docker service rm app_redis app_web
```

