## 系统设计-短链接服务

### 场景
微博、推特有字数限制、营销短信等情况下可以被使用。

### 阅读资料
[https://mp.weixin.qq.com/s/80uVWV1Lh_iOOdykzMhyeQ](https://mp.weixin.qq.com/s/80uVWV1Lh_iOOdykzMhyeQ)
[https://soulmachine.gitbooks.io/system-design/content/cn/tinyurl.html](https://soulmachine.gitbooks.io/system-design/content/cn/tinyurl.html)
[https://www.zhihu.com/question/29270034](https://www.zhihu.com/question/29270034)
[https://hufangyun.com/2017/short-url/](https://hufangyun.com/2017/short-url/)
[https://xie.infoq.cn/article/483fcfbe3f942cb1fa9d9ce20](https://xie.infoq.cn/article/483fcfbe3f942cb1fa9d9ce20)

### 流程
1. 输入短链接，回车；
2. 短链接地址服务器返回`301/302`,`location`上带有真实地址；
3. 浏览器读取到`301/302`，自动跳转到`location`指向的地址；

### 问题
1. 会不会重复？
2. 数据量大的情况下url也会太长；
3. 有规律的情况下容易被遍历；
4. 是否需要反推全局ID？
5. 一定时间内是否会使用完？
6. 一个长链接对应多个短链接的情况下，如何不重复？

### 方法
1. `2^32`<当前全网网页数<`2^64`,像微博短链接长度为7，如果不能覆盖，那就会出现碰撞，如果不是要覆盖全网的，可以选择更短的长度；
2. 三列：id、key、url，url为最终访问的url；
3. id一般使用`分布式发号器`产生，key由id生成，也可以使用数据库的自增索引；如果直接将长链接进行哈希，再转62字符，这样存在哈希冲突问题，也可以通过拼接字符解决；
4. `MurmurHash`,非加密型哈希函数，性能相对加密型高，适用于一般的哈希检索操作，对于规律性较强的key随机分布特征表现良好；
5. a-z,A-Z,0-9共62个字符，对id进行base62编码得到key，覆盖`2^64`需要11个字符；
6. 一般使用一对多的形式，一个长链接对应多个短链接，这样可以进行数据分析；
7. 一般使用`302`状态码,这样可以后续进行统计；
8. `分布式发号器`： 类uuid、Redis、Snowflake、Mysql自增主键；
9. 使用Mysql多台机器的时候，可以使用发号表。

### 优化
1. 分库分表；
2. 读写分离；
3. 缓存；
4. 防止恶意攻击？限IP、设置阈值，LRU缓存`长网址->ID`;


### github 项目学习
[https://github.com/golang-design/redir.git](https://github.com/golang-design/redir.git)
