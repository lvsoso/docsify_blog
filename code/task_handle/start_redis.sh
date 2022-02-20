#! /bin/bash


# docker run -it -d --rm --network host --name myredis redis:5.0.14

# redis:5.0
# redis:5.0.14
# redis:6.0.8
# (base) lv@lv:task_handle$ docker exec -it myredis /bin/bash
# root@lv:/data# redis-cli 
# 127.0.0.1:6379>  zadd "asynq:{bundle}:lease"  NX 1645334435  "11a5aed1-ec41-4e63-9c14-91b834bac2e1"
# (integer) 1
# 127.0.0.1:6379> zrange "asynq:{bundle}:lease" 0 100000000000000 WITHSCORES
# 1) "11a5aed1-ec41-4e63-9c14-91b834bac2e1"
# 2) "1645334435"
# 127.0.0.1:6379> zadd "asynq:{bundle}:lease"  XX GT 1645334450 "11a5aed1-ec41-4e63-9c14-91b834bac2e1"
# (error) ERR syntax error
# 127.0.0.1:6379> zadd "asynq:{bundle}:lease"  XX 1645334450 "11a5aed1-ec41-4e63-9c14-91b834bac2e1"
# (integer) 0
# 127.0.0.1:6379> zadd "asynq:{bundle}:lease"  GT 1645334455 "11a5aed1-ec41-4e63-9c14-91b834bac2e1"
# (error) ERR syntax error
# 127.0.0.1:6379> 

# redis:6.2.0