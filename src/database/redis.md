## Redis

> 2022年6月12日 星期日
>
> Joshua Conero



### 启动

```shell
# 使用 nohup 启动
nohup redis-server &
```







### 命令

```shell
# 查看当前所有变量
keys *
 
# 模糊搜索
keys "_x_62*"

# 删除变量（清空所有缓存）
flushall

# 删除指定的数据
del key_name
```

