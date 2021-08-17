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