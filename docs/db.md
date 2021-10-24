## 数据库

### 数据结构和查询

范围查询        记录查询
好               差
                ^
|       B+树    |
|       B树     |
|       hash    |
^                

### SQL
MySQL、Oracle、PostgreSQL、DB2、sqlite

### NoSQl
MongoDB、Redis、DynamoDB2、Elasticsearch



### MysqlDB

### MongoDB
#### 存储引擎
1. WiredTiger
2. InMemory
3. Encrypted

#### WiredTiger使用B树
1. 注重单条数据的修改和读取,但也需要较好的遍历数据；
2. 叶子节点也存储数据；
3. MongoDB并不适合遍历索引去收集指定条件的数据，更适合直接把数据集合到一起；
4. LSM树牺牲读性能，将随机写转换成顺序些以优化数据的写入；

[https://blog.csdn.net/weixin_41987908/article/details/105255119](https://blog.csdn.net/weixin_41987908/article/details/105255119)