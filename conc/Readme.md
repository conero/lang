### concurrent

> 2025-09-26
>
> Joshua Conero
>
> 多语言并发测试



#### simple 基础测试

指定数并多线程执行，含加减的数据值



```shell
# ubuntu 中运行测试用例
# go 测试版本 1.25
go run golang/simple.go

# zig 测试版本 0.15.1
zig run simple.zig

# rustc 测试代码 rustc 1.90.0
rustc simple.rs;./simple;rm ./simple

# python
python simple.py
```





