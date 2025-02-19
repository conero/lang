# WEB

> 万维网(World Wide Web/全球广域网)学习
>
> 2019年1月8日 星期二





## 概述

- 发明者: *Tim Berners-Lee*
- 表现形式: 
  - *超文本(Hyper text)*
  - *超媒体(Hyper Media)*
  - *超文本传输协议(HyperText Transfer Protocol)*
  - *超文本标记语言(Hyper Text Markup Language，HTML)*



> _**W3C**_

*World Wide Web Consortium*

- 相关协议
  - *HTTP*
    - *HTTP/1.0*
    - *HTTP/1.1*
    - *HTTP/2*
  - *SOAP*



### Cgi 和 FastCgi

- _**CGI** 是HTTP Server和一个独立的进程之间的协议，把HTTP Request的Header设置成进程的环境变量，HTTP Request的正文设置成进程的标准输入，而进程的标准输出就是HTTP Response包括Header和正文_
- *__php-fpm__相当于是Apache+mod_php。无非php-fpm自带了FASTCGI Server，而Apache是HTTP Server*



> CGI(*Common Gateway Interface*)  通用网关接口

*CGI 是Web 服务器运行时外部程序的规范,按CGI 编写的程序可以扩展服务器功能。CGI 应用程序能与浏览器进行交互,还可通过数据库API 与数据库服务器等外部数据源进行通信,从数据库服务器中获取数据。格式化为HTML文档后，发送给浏览器，也可以将从浏览器获得的数据放到数据库中。*



*CGI是外部应用程序（CGI程序）与WEB服务器之间的接口标准，是在CGI程序和Web服务器之间传递信息的过程。CGI规范允许Web服务器执行外部程序，并将它们的输出发送给Web浏览器，CGI将Web的一组简单的静态超媒体文档变成一个完整的新的交互式媒体。*



> FastCGI (*Fast Common Gateway Interface*)

_FastCGI像是一个常驻(long-live)型的CGI，它可以一直执行着，只要激活后，不会每次都要花费时间去fork一次(这是CGI最为人诟病的fork-and-execute 模式)。它还支持分布式的运算, 即 FastCGI 程序可以在网站服务器以外的主机上执行并且接受来自其它网站服务器来的请求。_



## WEB

### *Web1.0*

*本质是__互联(联合)__，信息共享。信息通常以单向传播的方式广播*



### *Web 2.0*

*本质是__互动__，信息共建。用户与服务端(服务商)之间的交互性*



### *Web3.0*

*暴风雨前的平静，区块链的诞生引发了一场旨在颠覆整个科技行业的运动。区块链和加密货币的拥趸们称之为Web 3.0，这场运动有望使所有传统商业模式不复存在。那是由于，简而言之，这项技术将便于万维网（WWW）的去中心化，因而将控制权和所有权从唯利是图的大公司手中夺回来，还给大众。*

*实现去中心化，让互联网回归自由交流的“数据广场”*

*可以理解为“__个性化定制__”*

*2009年，中本聪提出了区块链的概念，用虚拟货币当做帮助存储的奖励，激励更多的人参与进来，可以说，是很天才地解决了分散存储数据的成本问题。区块链就是数据库，虚拟货币就是一种Dapps。*

*但本质上来说，区块链只是实现Web 3.0的其中一种解决方案，并不是唯一的，其他的解决方案还有data pods技术和星际文件系统IPFS。*



> *__Web 3.0__ 区块链技术和加密货币等的技术总称*

世界正在慢慢向 Web 3.0 迈进。未来 Web 将由云解决方案和机器学习主导。





### 基础知识

#### 屏幕宽度

可使用浏览器的屏幕模拟工具调整

| PC   | 尺寸        | 备注 |
| ---- | ----------- | ---- |
|      | 1920 × 1080 |      |
|      | 1440 × 900  |      |
|      | 1366 × 768  |      |
|      | 1280 × 800  |      |
|      | 1280 × 720  |      |
|      | 1366 × 768  |      |
|      | 1024 × 768  |      |





### 会话及令牌



#### JWT

Json Web Token, 跨域身份验证解决方案。

由 *JWT头*、*有效荷载*、*签名* 等组成。



处理过程：

- 浏览器向服务器发送用户及密码，获取 token参数
- 请求服务器时将 token通过 `Authorization: Bearer {token}` 传递给服务进行认证即可。





## HTTP

*超文本传输协议 (HyperText Transfer Protocol / 超文本传输协议) 伴随着计算机网络和浏览器的诞生，HTTP1.0也随之而来，处于计算机网络中的应用层，HTTP是建立在TCP协议之上*

![](./image/web/https-history.png)



*影响一个HTTP网络请求的因素主要有两个：__带宽__ 和 __延迟__。*



*HTTP1.0最早在网页中使用是在1996年，那个时候只是使用一些较为简单的网页上和网络请求上，而HTTP1.1则在1999年才开始广泛应用于现在的各大浏览器网络请求中，同时HTTP1.1也是当前使用最为广泛的HTTP协议。*



*HTTP/1 建立在 TCP 协议之上。 HTTP 1.0 与 1.1 的主要区别在于长连接支持、多路复用、带宽节约与数据压缩等。其主要缺陷：连接无法复用、Head-Of-Line Blocking（HOLB，队头阻塞）、协议开销大（header 内容过大）、安全因素（明文传输）。*



![](./image/web/http-diff-1.0-and-1.1.png)



*HTTP 对照*

- HTTP/1 有连接无法复用、队头阻塞、协议开销大和安全因素等多个缺陷。（`read`）
- HTTP/2 通过多路复用、二进制流与 Header 压缩等技术，极大地提高了性能，但是还是存在一些问题（`read/write`）
- HTTP/3 抛弃 TCP 协议，以全新的视角重新设计 HTTP。其底层支撑是 QUIC 协议，该协议基于 UDP，有 UDP 特有的优势，同时它又取了 TCP 中的精华，实现了即快又可靠的协议（`read/write/own`）



**里程碑**

- HTTP 0.9（1991年）只支持get方法不支持请求头；
- HTTP 1.0（1996年）基本成型，支持请求头、富文本、状态码、缓存、连接无法复用；
- HTTP 1.1（1999年）支持连接复用、分块发送、断点续传；
- HTTP 2.0（2015年）二进制分帧传输、多路复用、头部压缩、服务器推送等；
- HTTP 3.0（2018年）QUIC 于2013年实现、2018年正式更名为HTTP3；



#### 请求

分别使用ip4/ip6进行本地地址请求，支持谷歌浏览器

- [http://[::1]:3000/](http://[::1]:3000/)                                ipv6
- http://127.0.0.1:3000/                       ipv4



##### post

```yaml
# post请求/常见的 form 数据提交类型
ContentType: application/x-www-form-urlencoded; charset=UTF-8
# Body 为，query 数据（url编码数据）
```



较快的请求库：

- rapidoid              https://github.com/rapidoid/rapidoid java 库
- tokio
  - mio               https://github.com/tokio-rs/mio ，  rust 轻量集合 io 库
  - tokio             https://github.com/tokio-rs/tokio ， rust http 异步事件库
  - client
    - reqwest       https://github.com/seanmonstar/reqwest ， rust http 客户端
    - hyper           https://github.com/hyperium/hyper ， rust http 客户端
    - rust-curl      https://github.com/alexcrichton/curl-rust , curl 绑定
- fasthttp               https://github.com/valyala/fasthttp ， go 版本请求库





### HTTPS

*为了解决以上问题，网景在1994年创建了HTTPS，并应用在网景导航者浏览器中。 最初，HTTPS是与[SSL](https://zh.wikipedia.org/wiki/%E5%82%B3%E8%BC%B8%E5%B1%A4%E5%AE%89%E5%85%A8%E5%8D%94%E8%AD%B0)一起使用的；在SSL逐渐演变到[TLS](https://zh.wikipedia.org/wiki/%E5%82%B3%E8%BC%B8%E5%B1%A4%E5%AE%89%E5%85%A8%E5%8D%94%E8%AD%B0#TLS_1.0)时（其实两个是一个东西，只是名字不同而已），最新的HTTPS也由在2000年五月公布的RFC 2818正式确定下来。简单来说，HTTPS就是安全版的HTTP，并且由于当今时代对安全性要求更高，chrome和firefox都大力支持网站使用HTTPS，苹果也在ios 10系统中强制app使用HTTPS来传输数据，由此可见HTTPS势在必行。*



**HTTPS与HTTP的一些区别**

1. HTTPS协议需要到CA申请证书，一般免费证书很少，需要交费。
2. HTTP协议运行在TCP之上，所有传输的内容都是明文，HTTPS运行在SSL/TLS之上，SSL/TLS运行在TCP之上，所有传输的内容都经过加密的。
3. HTTP和HTTPS使用的是完全不同的连接方式，用的端口也不一样，前者是80，后者是443。
4. HTTPS可以有效的防止运营商劫持，解决了防劫持的一个大问题。

![](./image/web/http-diff-http-and-https.png)





### HTTP/2

> [版本协议状态](https://datatracker.ietf.org/doc/rfc7540/)

- 协议代码
  - http:  
  - https:  **TLS(_Transport Layer Security_)**

*参考* __SPDY/2__



*2015 年，继承于 SPDY 的 HTTP/2 协议发布了。HTTP/2 是 HTTP/1 的替代品，但它不是重写，协议中还保留着第一代的一些内容，比如 HTTP 方法、状态码与语义等都与 HTTP/1 一样。*

*HTTP/2 基于SPDY3，专注于**性能**，最大的一个目标是在用户和网站间只用一个连接。*



**主要新特性**

- *新的二进制格式*
- *多路复用*
- *header 压缩*
- *服务端推送*



_HTTP/1 的请求和响应报文，都是由起始行、首部和实体正文（可选）组成，各部分之间以文本换行符分隔。**HTTP/2 将请求和响应数据分割为更小的帧，并且它们采用二进制编码**。_





*相关的概念*

- **流（stream）：**流是连接中的一个虚拟信道，可以承载双向的消息；每个流都有一个唯一的整数标识符（1、2…N）
- **消息（message）：**指逻辑上的 HTTP 消息，比如请求、响应等，由一或多个帧组成
- **帧（frame）：**HTTP/2 通信的最小单位，每个帧包含帧首部，至少也会标识出当前帧所属的流，承载着特定类型的数据，如 HTTP 首部、负荷等





*HTTP2.0可以说是SPDY的升级版（其实原本也是基于SPDY设计的），但是，HTTP2.0 跟 SPDY 仍有不同的地方，主要是以下两点：*

- *HTTP2.0 支持明文 HTTP 传输，而 SPDY 强制使用 HTTPS*
- *HTTP2.0 消息头的压缩算法采用 [HPACK](http://http2.github.io/http2-spec/compression.html)，而非 SPDY 采用的 [DEFLATE*](http://zh.wikipedia.org/wiki/DEFLATE)



![](./image/web/http-diff-1.x-and-2.0.png)



### HTTP/3

> HTTP-over-QUIC

*Google 就自己架起炉灶搞了一个基于 UDP 协议的 QUIC 协议，并且使用在了 HTTP/3 上，HTTP/3 之前名为 HTTP-over-QUIC，从这个名字中我们也可以发现，HTTP/3 最大的改造就是使用了 QUIC。*





### QUIC

*Quick UDP Internet Connections*。是谷歌在2013年发明的一种基于UDP的全新的低延时互联网传输协议。

QUIC是基于UDP协议的，我们知道UDP是面向无连接、不可靠协议。QUIC既然作为HTTP3的底层协议，保证可靠是基础。因此QUIC实现了TCP的ACK确认、seq序号、排序、重传、差错控制、流量控制。





## HTML/CSS

> 关于css中的计量单位: *px*, *em*, *rem*

- *`px`          px像素 (Pixel)，像素px是相对于显示器屏幕分辨率而言的。*
- *`em`        相对长度单位。相对于当前对象内文本的字体尺寸*
- *`rem`    相对单位（root em，根em)；与`em`的区别在于使用rem为元素设定字体大小时，仍然是相对大小，但相对的只是HTML根元素*



其他（绝对单位）

- in       英寸(inches), `1in = 2.54cm = 96px` 
- cm     厘米(centimeters), `1cm = 10mm = 96px/2.54 = 37.8px`
- mm   毫米(millimeters), `1mm = 0.1cm = 3.78px`
- pt      点(points), `1pt = 1/72in = 0.0139in = 1/72*2.54cm = 1/72*96px = 1.33px`
- pc     派卡(picas)，`1pc = 12pt = 1/6in = 1/6*96px = 16px`



相对单位

- em
- rem
- ex                所用字体中小写x的高度。ex在实际中常用于微调
- ch                ch与ex类似，被定义为数字0的宽度。当无法确定数字0宽度时，取em值的一半作为ch值



> 字号与pt值对比，如word中

- wps中字体对应关系
  - 初体/42，小初/36，一号/26，小一/24，二号/22，小二/18，三号/16，小三/15，四号/14，小四/12，五号/10.5，小五/9





### HTML5

> [html 5.1](https://www.w3.org/TR/html51/index.html)



@TODO   [html5 中文文档](https://whatwg-cn.github.io/html/)

#### history API

*更改浏览器 URL 页面不刷新，通过内部链接的点击！*



> 使用库

- reactjs
  - [react-router](https://github.com/ReactTraining/react-router)  支持 browser/hash 引擎
- vuejs
  - [vue-router](https://github.com/vuejs/vue-router) 支持 browser/hash 引擎





## 浏览器

浏览器页面加载过程：

 ```
浏览器			->					DNS 服务器		->  IP服务器(网站)	-> HTTP 请求
IP 服务器		->  HTTP 相应
浏览器			->  获得内容并 渲染
 ```





### 内核

*浏览器的内核是指支持浏览器运行的最核心的程序，分为两个部分的，一是渲染引擎，另一个是JS引擎。渲染引擎在不同的浏览器中也不是都相同的。*



常用浏览器内核： **Trident（IE）、Gecko（火狐）、Blink（Chrome、Opera）、Webkit（Safari）**。





### 版本

*chrome 浏览器版本介绍。软件内访问: [chrome://version/](chrome://version/)*

| chrome 版本   | 发布日期   | 依赖信息      | js等支持版本 |
| ------------- | ---------- | ------------- | ------------ |
| 81.0.4044.138 | 2020-04-08 | V8 8.1.307.32 |              |
| 80.0.3987.87  | 2020-02-06 |               |              |
|               |            |               |              |





## 服务器

### nginx

*[engine X] 由俄罗斯人 Igor Sysoev编写的 HTTP、反向代理、邮件代理服务器。* https://github.com/nginx/nginx



> **Igor Sysoev** *出生于1970年，在苏联的阿拉木图（现在的哈萨克斯坦阿拉木图）长大。我于1994年毕业于莫斯科国立鲍曼技术大学(BMSTU, Bauman Moscow State Technical University)。毕业后我继续住在莫斯科，目前在[NGINX，Inc](http://nginx.com/)担任首席技术官*



主要模块：

- core
- event
- http
- mail
- stream



**tengine** 由淘宝开发的基于nginx子集服务器，在兼容nginx的基础上提供新功能。源码地址 https://github.com/alibaba/tengine，网站 https://tengine.taobao.org/。



#### nginx conf

nginx 配置文件说明，默认名：`nginx.conf`。

```shell
# 当启动 nginx 后，可使用 -s 命令进行正在运行服务进程的启动、停止等
# nginx -s signal
# 重载配置文件
nginx -s reload
# signal: stop 终端退出，quit 优雅退出，reload 配置文件重载，reopen 重新打开日志文件

```



nginx 由模块组成，模块由配置文件指定的指令控制。指令分为简单指令(行)和块指令(块)。

```nginx
# 配置文件注释行

# 行指令（主上下文）
worker_processes  1;
error_log  logs/error.log;
access_log logs/access.log;

# 块指令，http
http{
    # 内部嵌套指令
    server{
        # 隐藏 nginx 版本号
        server_tokens off;
        # 内部上下文
        listen       9310;
    	server_name  localhost;
    }
}
```

##### ip 白名单

ip 白名单方法如下：

- 使用 `$remote_addr` 进行规则匹配
- 使用 `$http_x_forwarded_for` 进行规则匹配
- 使用 `allow` 配合 `deny all` 指令限制
- 系统防火墙级别处理 `iptables`



```nginx
http{
    server{
        # 1)
        # 使用 $remote_addr 规则匹配
        # 指定ip可访问
        if ($remote_addr !~ ^(218.201.194.162|218.21.184.165|127.0.0.1)){
            rewrite ^.*$ /admin.php last;
        }

        # 2)
        # 使用 $http_x_forwarded_for  规则匹配
        # 指定ip可访问
        if ($http_x_forwarded_for  !~ ^(218.201.194.162|218.21.184.165|127.0.0.1)){
            rewrite ^.*$ /admin.php last;
        }


        # 3)
        # 使用 allow/deny 指令
        allow 218.201.194.162;
        allow 218.21.184.165;
        allow 127.0.0.1;
        # 引入允许列表
        include /etc/nginx/conf.d/whitelist;
        deny all;
    }
}
```



##### http

http 服务配置

```nginx
# 负载均衡本地测试
# <loading_test_expo2022> upstream 名称
upstream loading_test_expo2022{
	server 127.0.0.1:9341 weight=3;
	server 127.0.0.1:9342 weight=1;
	server 127.0.0.1:9343 weight=1;
	server 127.0.0.1:9344 weight=1;
	server 127.0.0.1:9345 weight=1;
	server 127.0.0.1:9346 weight=1;
	server 127.0.0.1:9347 weight=2;
}
```



##### location

语法: `location [= | ~ | ~* | ~~] url { ... }`

- `location = pattern {}`    精准匹配（完全匹配）
- `location pattern {}`    一般匹配
- `location ~ pattern{}`   正则匹配
  - `location ~* pattern{}`      不区分大小写



```nginx
http{
    server{
        listen     8089;
        
        # 请求 /upload-x/v1/users.jpg -> http://127.0.0.1:8011/upload-x/v1/users.jpg
        location /upload-x/ {
            proxy_pass http://127.0.0.1:8011;
        }
        
        # 请求 /upload-y/v1/users.jpg -> http://127.0.0.1:8011/v1/users.jpg
        location /upload-y/ {
            proxy_pass http://127.0.0.1:8011/;
        }
        
        # 静态资源
        # 请求  /static-y/base/xxxxx/1.jpg -> /home/webapp/webimage/base/xxxxx/1.jpg
        location /static-y/ {
            alias /home/webapp/webimage/;
        }
        
        # 静态资源
        # 请求  /static-x/base/xxxxx/1.jpg -> /home/webapp/webimage/static-x/base/xxxxx/1.jpg
        location /static-x/ {
            root /home/webapp/webimage/;
        }
        
        # 首页
        location / {
        }
    }
}
```



http location 返回字符串内容方法

```nginx
http{
    server{        
        set $apiWeb http://$host:$server_port/;
        location / {
            return 200 "返回内容：$apiWeb .";
            
            # 如直接返回内容
            # return 200 $document_root$fastcgi_script_name;
        }
    }
}
```





##### 异常

> Request Entity Too Large

```shell
http {
    # 修改上传文件大小
    client_max_body_size 200m;
}
```



##### 指令

> worker_processes 

```nginx
# 指定nginx的进程分配数
worker_processes auto;
```









#### Linux 安装环境

**通过源码编译 Nginx**

```shell
# 准备
# 常用依赖更新/安装
# nginx 编译时依赖 gcc 环境 
sudo yum -y install gcc gcc-c++ 
# 让 nginx 支持重写功能
sudo yum -y install pcre pcre-devel 
# zlib 库提供了很多压缩和解压缩的方式，nginx 使用 zlib 对 http 包内容进行 gzip 压缩
sudo yum -y install zlib zlib-devel
# 安全套接字层密码库，用于通信加密
sudo yum -y install openssl openssl-devel

# 1、下载源码
wget -c http://nginx.org/download/nginx-1.20.0.tar.gz

# 2、解压源码
tar -zxvf nginx-1.20.0.tar.gz
cd nginx-1.20.0

# 3、 配置文件,编译
# 默认编辑
./configure	
# 支持 ssl 
./configure --prefix=/usr/local/nginx --with-http_stub_status_module --with-http_ssl_module

# 4、 编译
make && make install

# 5、创建环境变量（软件连接）
ln -s /usr/local/nginx/sbin/nginx /usr/local/bin/
```



nginx 设置站点访问密码，主要加密工具：`htpasswd` 和 `openssl`

```shell
#安装包
yum install -y httpd-tools

# 设置密码文件
htpasswd -c /etc/nginx/psswd/yang yang

# OpenSSL 版本
# 指定用户及密码 test/abc123456
echo -n "test:" > passwd
#设置密码密码  
openssl passwd abc123456 >> passwd
# 使用算法 -apr1
openssl passwd -apr1 abc123456
```



生成的密码文件：

```shell
# Comments
user1:password1
user2:password2:comment
user3:password3
```







#### issue

##### windows nginx 启动，命令行启动错误

```shell
# 不在项目所在的根目录，指定命令时报错；由于项目项目配置以根目录为标准
nginx: [alert] could not open error log file: CreateFile() "logs/error.log" failed (3: The system cannot find the path specified)
2019/07/26 22:26:49 [emerg] 20168#19264: CreateFile() "D:\tmp\CroIns-Php\resource/conf/nginx.conf" failed (3: The system cannot find the path specified)

# 需要跳转至 nginx 所在目录，可使用 nssm 来安装一个信息服务。
```

参考 nssm 服务  http://www.nssm.cc/download，安装服务如下：

```yaml
# nginx
- Application Path: D:\conero\~\nginx-1.20.2\nginx.exe
- Startup Directory: D:\conero\~\nginx-1.20.2

# php-cgi
- Application Path: D:\conero\~\php\php-cgi.exe
- Startup Directory: D:\conero\~\php
- Arguments: -b 127.0.0.1:9000

```



### apache

*与 php等相关的服务语言，主要存在两种模式：`mod_php`  和 `CGI/FastCGI` 。mod_php 是将 PHP作为apache的一个模块实现，这个是PHP最传统的裕兴方式；而 `CGI/FastCGI` 使用 CGI 协议实现。*



*两者选择的 PHP 版本不一样：FastCGI 使用的 NTS 版本（Non-Thread Safe），Mod_php 使用的是 TS(版本)。*



### Tomcat

由 Apache 开发的一个 Servlet 容器，其实是Web服务器和Servlet容器的结合体。Servlet 是Server Applet的简称，是javaEE平台下用于java编译的服务端程序。



SSM框架即 `Spring+SpringMVC+MyBatis` 组成的标准MVC框架，是较为主流的java EE企业级框架。其常由 dao(mapper)、service、controller、view 四层组成。





**部署项目**

打包项目为 war 包，然后将其复制到 webapps目录下，再执行`bin/startup.bat` 即可。



#### IntelliJ IDEA 部署

部署是有exploded与非exploded模式，前者允许在运行时实时更新资源（便于调试），而是在部署时一次性部署。





### 测试服务器

#### PHP

> php 内建服务器 测试系统

```powershell
php -S 0.0.0.0:2019
# 指定 root 目录
php -S 0.0.0.0:2019 -t ./public
```



#### python

> 使用库 `http.server` 实现测试服务的搭建，默认开启 `文件服务器`

```powershell
python -m http.server <port>

# 如 2020 服务器
python -m http.server 2020
```



#### nodeJs

> 安装 [`http-server`](<https://github.com/indexzero/http-server>) 插件实现命令式测试服务器



```powershell
# 安装全局插件
npm install http-server -g

# 开启服务器
http-server [path] [options]
```





> 使用 `npx` 启动服务

npx 即 npm exec，执行包命令，执行包时不存在则下载对应的包并执行它。

```shell
# 通过执行serve启动服务
npx serve
```



### 服务器端应用安全

#### PHP

- *PHP 框架的漏洞，导致脚本注入*
- *数据库脚本注入*



*通过漏洞：可以获取、删除、篡改数据、载入木马等*



### web应用性能指标

- 并发数（concurrency）
- 响应时间(RT/response time)
- 吞吐量(Throughput)
  - QPS            每秒查询数，每秒能够响应的查询次数
  - TPS             每秒事务数，一个事务是指一个客户机向服务器发送请求然后服务器做出反应的过程。
  - HPS            每秒HTTP请求数
- 服务器资源使用率（Resource Utilization）      服务的CPU、内存、磁盘IO使用率
- 网络宽度使用率
- 页面浏览量(PV *Page View*)
- 网站独立访客(UV  Unique Visitor) 
- 峰值QPS



**吞吐量** 是指单位时间内系统能处理的请求数量，体现系统处理请求的能力，这是目前最常用的性能测试指标。

`QPS（TPS）= 并发数 / 平均响应时间`



### 技术栈

- lamp                  linux+Apache+mysql+php
- lnmp                  linux+nginx+mysql+php
- mean                 MongoDB，Express，Angular，Node
- mern                  MongoDB、Express、React、Node.js
- menv                  MongoDB、Express、Node.js、vue.js



## 网络

### 基本知识概述

#### 网络带宽

带宽（频带宽度）是指在单位时间（一般指的是1秒钟）内能传输的数据量。网络和高速公路类似，带宽越大，就类似高速公路的车道越多，其通行能力越强。

数字信息流的基本单位是bit（比特），时间的基本单位是s（秒），因此bit/s（比特/秒）是描述带宽的单位，1bit/s是带宽的基本单位。

> **单位**

- bps    (bit per second) 值比特每秒或bit/s

```shell
1Mb/s=1024Kb/s=1024/8KB/s=128KB/s


# 理论上
# 由于受用户计算机性能、网络设备质量、资源使用情况、网络高峰期、网站服务能力、线路衰耗，信号衰减等多因素的影响，
#   实际速率大学为理论值的 80%左右
1Mb/s    速率 128 KB/s
2M       速率 256 KB/s(实际速率大约为150~240KB/s)
```



吞吐量是指在规定时间、空间及数据在网络中所走的路径（网络路径）的前提下，下载文件时实际获得的带宽值。由于多方面的原因，实际上吞吐量往往比传输介质所标称的最大带宽小得多。




## 附录

### 名称解释

- 数据交换格式
  - [*__XML__*](https://www.w3.org/XML/)     *eXtensible Markup Language, 可扩展标记语言, https://www.w3.org/TR/xml11/*
  - [*__JSON__*](https://www.json.org/)   *JavaScript Object Notation*
  - [__*Toml*__](https://github.com/toml-lang/toml)  *Tom's Obvious, Minimal Language.*
  - [__*YAML*__](https://yaml.org/)  *YAML Ain't Markup Language*
  - [__*Protocol Buffers*__](https://github.com/protocolbuffers/protobuf)


### 参考

- [HTTPWG](https://httpwg.org/)
- [IETF](https://www.ietf.org/)
  - *文档索引：https://www.ietf.org/download/rfc-index.txt*
  - 文档类型： (`$index = INDEX`)
    - pdf https://tools.ietf.org/pdf/rfc-INDEX.pdf
    - html  https://tools.ietf.org/html/rfc-INDEX
    - text https://tools.ietf.org/rfc/rfc-INDEX.txt
    - https://datatracker.ietf.org/doc/rfc-INDEX/ 
    - https://www.rfc-editor.org/info/rfc-INDEX
- [RFC-Editor](https://www.rfc-editor.org/)
- [ISO](https://www.iso.org/home.html)  *the International Organization for Standardization*
- [Nginx域名访问的白名单配置 - 运维总结](https://www.cnblogs.com/kevingrace/p/6086652.html)
- [JWT全面解读、详细使用步骤](https://www.jianshu.com/p/d1644e281250)



#### 网络链接

- [HTTP,HTTP2.0,SPDY,HTTPS你应该知道的一些事（伯乐在线）](http://web.jobbole.com/87695/)
- [淘宝HTTPS探索](http://velocity.oreilly.com.cn/2015/ppts/lizhenyu.pdf)
- [NGINX白皮书](https://www.nginx.com/wp-content/uploads/2015/09/NGINX_HTTP2_White_Paper_v4.pdf)
- [NGINX配置HTTP2.0官方指南](https://www.nginx.com/blog/nginx-1-9-5/)

