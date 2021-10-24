## 系统设计-API 限速

### 场景
流量整形、流量控制。

### 阅读资料
[https://zhuanlan.zhihu.com/p/20872901](https://zhuanlan.zhihu.com/p/20872901)
[https://www.itread01.com/content/1545066784.html](https://www.itread01.com/content/1545066784.html)
[https://www.binpress.com/rate-limiting-with-redis-1/](https://www.binpress.com/rate-limiting-with-redis-1/)

#### 计数器
- 单位时间内记录次数每次进行对比；
- 相邻单位时间内可以突发两倍限制数量的请求；

#### 令牌桶
- 令牌按固定速率放入桶，桶的容量即API调用速率；
- 每次请求消耗令牌；
- 通过令牌放入速率控制请求速率；
- 也有突发的情况；

##### 实现一
1. 按照需求设置多个桶；
2. 每个桶有各自的定时器；

##### 实现二
1. 存放
##### 优劣

#### 漏桶算法
- 固定容量的漏桶，按照固定速率流出水滴；
- 空则不流出水滴；
- 请求到来视为水滴流入；
- 流出速率固定；

