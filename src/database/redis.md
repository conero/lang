## Redis

> 2022年6月12日 星期日
>
> Joshua Conero

Redis 为 Remote Dictionary Service 缩写，是一个数据结构服务器。

一个基于 C 语言开发的开源数据库（BSD 许可），与传统数据库不同的是 Redis 的数据是存在内存中的（内存数据库），读写速度非常快，被广泛应用于缓存方向。并且，Redis 存储的是 KV 键值对数据。



github 地址： https://github.com/redis/redis

兼容 Redis 服务且更快的KV缓存数据库，[dragonfly](https://github.com/dragonflydb/dragonfly)

其他类似的KV缓存数据库

- valkey         https://github.com/valkey-io/valkey , redis 分支版本
- etcd             https://github.com/etcd-io/etcd
- garnet         https://github.com/microsoft/garnet



### 基础

数据类型

- string
- hash                             散列
- sorted set
- bitmap                         位存储
- list                                列表
- set                                集合
- zset                              有序集合



支持事务、持久化、Lua脚本、集群化部署。



特性

- 基于内存，访问速度比磁盘快上千倍
- 基于Reactor模式设计开发一套高效的事件处理模型
- 内置多种优化数据结构实现，性能较高
- 发布 / 订阅模式
- 支持消息持久化
- Redis 的瓶颈主要受限于内存和网络



主要用途

- (分布式)缓存
- 限流
- 消息队列



### 安装

通过编译源码安装 redis

```shell
wget https://download.redis.io/redis-stable.tar.gz
tar -zxvf redis-stable.tar.gz
cd redis-stable

# 编译以及安装
# redis 内存分配策略
# % make MALLOC=libc              .   
# % make MALLOC=jemalloc          .
make && make install
```





### 启动

```shell
# 使用 nohup 启动
nohup redis-server &
```



**dragonfly** 启动

```shell
# 指定端口号
./dragonfly --port 6380

# 使用 nohup 保持应用
nohup ./dragonfly --port 6380 &
```



#### 配置

如配置文件 /usr/local/etc/redis.conf ，修改配置文件

```shell
# 关闭保护模式
protected-mode no
```





### 命令

启动服务

```shell
# redis 链接指定 host 域名
redis-cli -h 172.27.57.98 -p 6380

# 指向端口
redis-cli -p 6380
```





常量命令操作

```shell
# 查看当前所有变量
keys *
 
# 模糊搜索
keys "_x_62*"

# 删除变量（清空所有缓存）
flushall

# 删除指定的数据
del key_name

# 查看缓存数据的过期时间(秒)
ttl key
# -1 表示永久
# -2 表已过期


# 设置参数，单个
set $key $values
# 读取参数
get $key
# 设置多个参数
mset $k1 v1 $k2 v2 $k3 v3 [$k value]
# 批量获取参数
mget $k1 $k2 [...]


# 设置值和时间期限
setex $key $sec $value
#设置期限
expire $key $sec
# 设置秒期限: ex $sex/ 
set jctest jkkkkkkk ex 120

# 查看存储值的参数
strlen $key

# 键值对应值+1
incr $key
# 整形键值减一
decr $key

# 查看键值是否存在
exits $key
```



版本信息等查看

```shell
# 版本信息查看
info

# 获取服务器版本号
info server
```



设置密码及登录

```shell
# 配置文件如
#   windows - redis.windows.conf
#   linux - redis.conf
requirepass 123456

# 使用 cli 的工具登录
redis-cli -a 123456

# 使用 redis-cli 服务的`123456`密码登录用户
auth 1234567

# 登录redis-cli后设置/查看密码
# 设置密码
config set requirepass 123456
# 清空密码
config set requirepass ""

# 查看密码
config get requirepass
```







事务控制

- `multi`                           开始事务
- 命令入队，批量操作，FIFO 顺序执行
- `exec`                            执行事务
- `discard`                      抛弃事务（取消事务）





### 附录

与之相同（类似）的产品

- memcached



#### 参考

- [图解redis](https://mp.weixin.qq.com/s/fEAWotIg2LN324wDfznpDA)
- [redis-windows](https://github.com/redis-windows/redis-windows)  windows环境下安装包
