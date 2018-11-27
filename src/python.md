# Python(py)

> 2018年8月9日 星期四
>
> Joshua Conero



# Python 学习笔记

- 官网:  https://www.python.org/





## 概述

​	日期： 2017年2月23日 星期四

​	特点： “优雅”、“明确”、“简单”



### 安装

安装 python 运行包以后，可以在安装的目录下查看语言的核心库等源码



## 基础

### 流程控制

#### for

```python
words = ['w', 'w1', 'w2', '...', 'wn']

# 循环
# for w in words[:]:
for w in words:
    # w,w1,w2
    print(w)
    
# 通过长度
for i in range(len(words)):
    # i 系统处理
    print(words[i])
```



*`break`、`continue` 中断/跳过当前的步骤*



