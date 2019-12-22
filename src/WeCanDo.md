# 解决方案

> 2019年4月10日 星期三



## Windows OS

### 文件操作提示“文件或文件夹正在使用”的错误

参考： https://zhidao.baidu.com/question/1500683749523466899.html



**查看或解除** *依次进入【任务管理器(K) 】》【打开资源监视器】》【CPU、关联的句柄】 。在“关联的句柄”中搜索找出并继续操作即可*





### 设备相连

*两台设备相连可通过`Internet` 、`内部局域网` 、`网线直连`等方式实现。*

**网线直连**   需要手动配置IP地址后接口连接，执行 `mstsc` (*Microsoft terminal services client*) 命令接口打开远程连接端口。



### FTP

*Windows 下可以使用 `文件资源管理器`，来读取资源以及基本的操作。*



## 内网穿透，外部网络访问本地网站

### natapp

> *使用简单，外网访问本地网页(web) 的解决方案。*
>
> [详见官网](https://natapp.cn/)







### 正则表达式

```javascript
/^\s*\n/				// 空行
```



## 微信

> *微信内部调试地址*

- <http://debugtbs.qq.com/>
- http://debugx5.qq.com/





## markdown

### markdown 中的流程图

*[mermaid](https://github.com/knsv/mermaid)*  可以实现 markdown 中的流程图





## 跨平台方案

### Windows&Linux

#### 脚本语言

*通常情况下，脚本语言都具有跨平台性。处理语言上的区别，如python2 与 python3.*



#### 编译型

*由于部分网络安全机制的限制，可能系统无法联网。因此在离线情况下，还有**脚本程序**如 python 版本不兼容的情况；或者考虑到软件的性能，而采用**编译型**语言。可选择方案如下：*

- *jvm 系列*            通过JVM上的语言，将程序编译为二进制的**jar**，**war**包实现
- *golang*               通过go编译器，在Windows下交叉编译程序实现的方法





## 网站资源

### google.cn 系列网站

- [谷歌开发中心](![1570515135884](C:\Users\11066\AppData\Roaming\Typora\typora-user-images\1570515135884.png))
- [golang](https://golang.google.cn/)
- [TensorFlow](https://tensorflow.google.cn)







#### 谷歌开发中心

![](https://www.gstatic.cn/devrel-devsite/v45ff8f523b59dba1a68309c845322ac9e5fdaa00274f34f2986d46556e823992/developers/images/lockup.svg)

https://developers.google.cn/



#### golang

![](https://golang.google.cn/lib/godoc/images/go-logo-blue.svg)



### 国内镜像

#### 大学

- [清华](https://mirrors.tuna.tsinghua.edu.cn/)
- [中国科技大学](http://mirrors.ustc.edu.cn/)
- [北理工](http://mirror.bit.edu.cn/)



#### 企业

- [搜狐](http://mirrors.sohu.com/)
- [163](http://mirrors.163.com/)
- ![](https://img.alicdn.com/tfs/TB12yDStQ7mBKNjSZFyXXbydFXa-248-72.png)[阿里云](https://opsx.alibaba.com/mirror)
- [腾讯](https://mirrors.tencent.com/)



### 其他

#### 多媒体

#### 在线文件转换 

地址 https://convertio.co/ 



#### 随机生成器

1. **头像生成**

<img src="https://thispersondoesnotexist.com/image" style="zoom:40%;" />



2. **随机图像**

   <img src="http://lorempixel.com/800/800/" style="zoom:50%;" />





## 工具

### 编程

#### Jenkins 

*Jenkins是开源CI&CD软件领导者， 提供超过1000个插件来支持构建、部署、自动化， 满足任何项目的需要。*