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
- *__php-fpm__相当于是Apache+mod_php。无非php-fpm自带了FASTCGI Server，而Apache是HTTP Serve*



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



## HTTP

*超文本传输协议 (HyperText Transfer Protocol / 超文本传输协议) 伴随着计算机网络和浏览器的诞生，HTTP1.0也随之而来，处于计算机网络中的应用层，HTTP是建立在TCP协议之上*

![](./image/web/https-history.png)



*影响一个HTTP网络请求的因素主要有两个：__带宽__ 和 __延迟__。*



*HTTP1.0最早在网页中使用是在1996年，那个时候只是使用一些较为简单的网页上和网络请求上，而HTTP1.1则在1999年才开始广泛应用于现在的各大浏览器网络请求中，同时HTTP1.1也是当前使用最为广泛的HTTP协议。*



*HTTP/1 建立在 TCP 协议之上。 HTTP 1.0 与 1.1 的主要区别在于长连接支持、多路复用、带宽节约与数据压缩等。其主要缺陷：连接无法复用、Head-Of-Line Blocking（HOLB，队头阻塞）、协议开销大（header 内容过大）、安全因素（明文传输）。*



![](./image/web/http-diff-1.0-and-1.1.png)



*HTTP 对照*

- HTTP/1 有连接无法复用、队头阻塞、协议开销大和安全因素等多个缺陷
- HTTP/2 通过多路复用、二进制流与 Header 压缩等技术，极大地提高了性能，但是还是存在一些问题
- HTTP/3 抛弃 TCP 协议，以全新的视角重新设计 HTTP。其底层支撑是 QUIC 协议，该协议基于 UDP，有 UDP 特有的优势，同时它又取了 TCP 中的精华，实现了即快又可靠的协议



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







## HTML/CSS

> 关于css中的计量单位: *px*, *em*, *rem*

- *`px`          px像素 (Pixel)，像素px是相对于显示器屏幕分辨率而言的。*
- *`em`        相对长度单位。相对于当前对象内文本的字体尺寸*
- *`rem`    相对单位（root em，根em)；与`em`的区别在于使用rem为元素设定字体大小时，仍然是相对大小，但相对的只是HTML根元素*





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







## 服务器

### nginx

*[engine X] 由俄罗斯人 Igor Sysoev编写的 HTTP、反向代理、邮件代理服务器。*



> **Igor Sysoev** *出生于1970年，在苏联的阿拉木图（现在的哈萨克斯坦阿拉木图）长大。我于1994年毕业于莫斯科国立鲍曼技术大学(BMSTU, Bauman Moscow State Technical University)。毕业后我继续住在莫斯科，目前在[NGINX，Inc](http://nginx.com/)担任首席技术官*





### 测试服务器

#### PHP

> php 内建服务器 测试系统

```powershell
php -S 0.0.0.0:2019
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



### 服务器端应用安全

#### PHP

- *PHP 框架的漏洞，导致脚本注入*
- *数据库脚本注入*



*通过漏洞：可以获取、删除、篡改数据、载入木马等*




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



#### 网络链接

- [HTTP,HTTP2.0,SPDY,HTTPS你应该知道的一些事（伯乐在线）](http://web.jobbole.com/87695/)
- [淘宝HTTPS探索](http://velocity.oreilly.com.cn/2015/ppts/lizhenyu.pdf)
- [NGINX白皮书](https://www.nginx.com/wp-content/uploads/2015/09/NGINX_HTTP2_White_Paper_v4.pdf)
- [NGINX配置HTTP2.0官方指南](https://www.nginx.com/blog/nginx-1-9-5/)

