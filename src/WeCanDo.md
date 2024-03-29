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



### 软链接与Linux `ln -s` 类似

*使用 windows 系统下的 `linkd` 进行设置，其形式与快捷方式一样。只是前者可实现程序正常读取，与 linux下的`ln -s` 一致。*



如果命令不存在，需要下载windows 资源工具集(Resource Kit Tools)，安装即可，如目录 [Windows Server 2003 Resource Kit Tools](https://www.microsoft.com/en-us/download/details.aspx?id=17657) 。

```powershell
# 此方法编写的软链接可供程序使用
linkd <Source> <Destination>
# 如下图示意图，再当前目录下生成“../uploads”的软链接地址
linkd uploads-rename ../uploads

# 手动创建的快捷文件无法被访问（程序）
```



<img src="./image/wecando/linkd-vs-lnk-icon.png" style="zoom:67%;" />

<img src="./image/wecando/linkd-files.png" style="zoom: 50%;" />





liunx 命令行

```shell
# linux 建议使用，觉得地址 - source_file
# 创建-软链接
ln -s source_file softlink_file

# 查看的软链接信息
ls -l
```





### 微信、火狐浏览器、系统搜索十分缓慢的解决方案

因系统使用专门的清理工具对系统文件清理，操作部分软件网络连接十分缓慢；此外 chm 页面打开也存在很大的问题。

解决方案是：

- 删除或重命名文件夹 `C:\Users\用户\AppData\Local\Microsoft\Windows\WebCache`。 WebCache清除后恢复正常。





## 内网穿透，外部网络访问本地网站

### natapp

> *使用简单，外网访问本地网页(web) 的解决方案。*
>
> [详见官网](https://natapp.cn/)





### ngrok

> 反向代理，[网站](https://ngrok.com/)



## 正则表达式

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



#### 具体应用

- [tortoisesvn](https://mirrors.tuna.tsinghua.edu.cn/osdn/storage/g/t/to/tortoisesvn/)  svn 下载镜像



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





#### IDE

- vscode





**vscode**

多行编辑：

- `Alt+光标`                       当行选择
- `Alt+Shift+光标`          光标批量选择



### 文档或编码

ecma-376 即 *[Office Open XML](https://ecma-international.org/publications-and-standards/standards/ecma-376/)* 规范，规范会后的文档如：

```shell
# 最初为微软开发实现，后转为公开的 ecma-376 协议
doc,xls,ppt
docx,xlsx,pptx

# 类似的 pdf 最初有 abode开发，后期交给 iso标准
```

类似 pdf 文档标准为 ISO 32000-1





#### pandoc

文档转换工具，安装地址 [pandoc](https://github.com/jgm/pandoc)。生成 pdf 中文乱码的处理方式，https://zhuanlan.zhihu.com/p/477441931 。



```shell
# 将 test.docx word 文档; 并转化为 md
pandoc test.docx --extract-media=. -o test.md

# 将md文件转化为 docx word 文件
pandoc test.md -o test.docx

#
# 安装 latex 后生成 PDF
# 生成模板
pandoc -D latex > template.tex
# 使用模板生成PDF
pandoc test.md -o test.pdf --template=template.tex --pdf-engine=xelatex
pandoc test.md -o test.pdf --template=template.tex
```



参考 pdf 文件：[《PDF解析》](https://github.com/zxyle/PDF-Explained)，阅读链接 https://github.com/zxyle/PDF-Explained 。



#### pdf

文档合并

- merge2pdf   基于 unidoc框的 pdf 文件合并工具（带水印/需购买），  https://github.com/ajaxray/merge2pdf
- pypdf   纯 python 库， https://github.com/py-pdf/pypdf





**Latex**

latex 项目，https://www.latex-project.org/ 。是一种基于tex的编排系统，由美国计算机学家莱斯利·兰伯特（Leslie Lamport）在20世纪80年代初期开发。

Markdown to PDF，主要是通过模板生成tex文件，然后渲染成pdf。因此需要安装Latex。在windows下推荐安装[miktex](http://miktex.org/)，Miktex是Latex的轻量化打包，并且提供了缺包就自动安装的功能。



### 网络

#### 代理服务器

*如由于网络的限制，部分计算机无法上网；此此时可以通过代理服务器进行代理上网。*



##### CCProxy

[网站](http://www.ccproxy.com/)



**相关帮助链接**

1. [启用后无法上网](http://www.ccproxy.com/support_c01.htm)
2. 下载，[以及使用方法介绍](https://www.onlinedown.net/soft/4156.htm)





### 代码

- 代码统计工具
  - https://github.com/XAMPPRocky/tokei                         rust 编写，计算快。
  - https://github.com/AlDanial/cloc                                    perl编写



### 实操

#### 打印机身份证双面复印

方法一：手动重复打印，即打印后将纸放回后再次打印。原件放置一般都是靠左上角，

<img src="./image/wecando/print-id.png" style="zoom: 50%;" />
