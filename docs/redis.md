
[深度剖析：Redis分布式锁到底安全吗？](https://mp.weixin.qq.com/s/VFZ2tR86kxj4AYBUnsKc2Q)
[基于Redis的分布式锁到底安全吗](http://zhangtielei.com/posts/blog-redlock-reasoning.html)



进成可能没释放锁

```shell
SETNX lock 1
DEL lock
```

设置10s后过期，不是原子性

```shell
SETNX lock 1
EXPIRE lock 10
```

存在超时后被别人加锁

```shell
SET lock 1 EX 10 NX
```

区别别人上的锁，但是在解锁时判断是否为自己的锁再删除不是原子性的。

```shell
SET lock $uuid EX 20 NX
```

lua 脚本原子性

```shell
// 判断锁是自己的，才释放
if redis.call("GET",KEYS[1]) == ARGV[1]
then
    return redis.call("DEL",KEYS[1])
else
    return 0
end
```


自动延期,但是副本情况下存在复制延迟

```shell

```

Redlock?
- 只部署主库
- 部署至少5个实例


1. 客户端在多个Redis实例上申请加锁
2. 必须保证大多数节点加锁成功
3. 总的耗时小于锁的过期时间
4. 释放锁，需要向全部节点发起释放锁请求

需要解决NPC问题
- N：Network Delay，网络延迟
- P：Process Pause，进程暂停（GC）
- C：Clock Drift，时钟漂移

 fecing token？

 分布式锁的本质，是为了「互斥」，只要能保证两个客户端在并发时，一个成功，一个失败就好了，不需要关心「顺序性」?

一个分布式锁，在极端情况下，不一定是安全的。

