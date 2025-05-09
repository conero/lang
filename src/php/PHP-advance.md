# PHP 扩展性学习



下载镜像：

- 搜狐php语言包下载镜像  http://mirrors.sohu.com/php/



## 基础

> 相关特点

- *语法简单，入门快。*
- *为 Web 开发设计的语言，导致其主要使用范畴为 `表现层` 的*



> 类型

- 标量类型
  - boolean           布尔类型
  - integer             整型
  - float/double    浮点类型
  - string               字符串
- 复合类型
  - array                 数组
  - object               对象
  - callable             可调用
- 特殊类型
  - resource          资源
  - NULL                无类型
- 伪类型
  - mixed               混合类型
  - number            数据类型
  - callback            回调类型
  - array|object    数组、对象类型
  - void                   无类型



### LNMP

#### CentOS 7.3

`configure` shell脚本文件是由 **Autoconf** 自动创建，用于编译前的系统环境检测，configure脚本最后环境成功会生成一个 Makefile 文件；可通过 `./configure --help` 查看支持的选项。

**.m4** 文件（marco，对应后面的4个字母），m4是一种宏处理器，它是 POSIX 标准的一部分。

Linux 编译时或者部署过程中注意系统用户权限的问题，如Nginx应用可能无法访问 root 用户下的文件/目录等。

liunx php 编译时，使用 `--with-<extension>` 为内置的编译法，即PHP内部自动包含对应扩展，也无需在 `php.ini` 内添加相应的扩展。

启动系统对应应用安装命令：`sudo apt install`。

```shell
# 安装编译环境
$ yum -y install make zlib zlib-devel gcc-c++ libtool  openssl openssl-devel

# ----------------------【nginx】------------------------
# 安装 nginx
$ wget <loaddown.net.url>
$ tar -zxvf nginx-1.6.2.tar.gz	解压
# 解压 tag 文件夹 `tar xvf file.tar`
$ cd nginx-1.6.2				目录转移
$ ./configure					编译安装
# 可能出现的依赖 
# error: the HTTP rewrite module requires the PCRE library.(pcre-devel)
yum -y install pcre-devel
# error: error: the HTTP gzip module requires the zlib library.
yum install -y zlib-devel

$ make && make install							

# 启动 nginx
$ /usr/local/nginx/sbin/nginx
# 停止服务器
$ /usr/local/nginx/sbin/nginx -s stop 或 /usr/local/nginx/sbin/nginx -s quick

# 创建 nginx 软链接
ln -s /usr/local/nginx/sbin/nginx /usr/bin/nginx

# 配置文件检测
$ /usr/local/nginx/sbin/nginx -t
# openssl 安装
$ yum -y install openssl*

# ----------------------【php】------------------------
# 依赖安装
$ yum install -y libxml2*
$ wget https://www.php.net/distributions/php-7.3.5.tar.gz
#使用：--no-check-certificate 取消证书检测
#$ wget --no-check-certificate https://www.php.net/distributions/php-7.2.34.tar.gz
$ tar -xzvf <tag.gz>
# 配置文件
$ ./buildconf --force
$ ./configure
# 编译
# 重复时需要清理编译环境 `make clean`
$ make && make install

# 具体选项，查看帮助命令
./configure --help
# [TIP] 此方法安装的 php-fpm 版本可能存在与当前PHP版本不对应（不推荐此方法）
# 可使用编译选项（如）
./configure \
--with-openssl \
--enable-mbstring \
--with-pdo-mysql \
--enable-fpm \
--with-gd

# 示例2
./configure \
--enable-fpm \
--with-openssl \
--enable-mbstring \
--with-pdo-mysql \
--with-bz2 \
--with-curl \
--enable-exif \
--enable-mbstring \
--enable-sockets \
--enable-mysqlnd \
--with-gd
# ./configure --with-freetype-dir --with-jpeg-dir --with-png-dir

# configure: error: Please reinstall the libcurl distribution
yum -y install curl curl-devel

# 启动 php-fpm
/usr/local/sbin/php-fpm
#
# 将 php-fpm 作为系统服务。 
# 创建服务脚本
# cp ~/sapi/fpm/init.d.php-fpm /etc/init.d/php-fpm
# 创建系统默认配置
# cp ~/sapi/fpm/php-fpm.conf.in /usr/local/etc/php-fpm.conf
# cp /usr/local/etc/php-fpm.d/www.conf.default /usr/local/etc/php-fpm.d/www.conf

# 杀死进程
$ killall php-fpm
# 检测是否已经正确启动
$ netstat -tln | grep 9000
# php-fpm 的进程号
$ ps -aux|grep php-fpm
```



`--with-gd` 检测错误时：

```shell
# 安装图片库支持
yum -y install libpng libpng-devel
```



Bzip2

```shell
# configure: error: Please reinstall the BZip2 distribution
yum -y install bzip2-devel
```



> linux php 扩展编译

转到ext所在目录，并使用 phpize 初始化；执行 `configure`命令，最后进行 `make && make install` 即可。

```shell
# 如编译 mysqli
cd ext/mysqli

# 使用 phpize 初始化
/usr/local/bin/phpize

# 执行 configure
./configure --with-php-config=/usr/local/php/bin/php-config --with-mysqli=/usr/local/mysql/bin/mysql_config

# 进行 make
make && make install

# 如在[/usr/local/lib/php/extensions/no-debug-non-zts-20100525/]中生成`mysqli.so`
```





windows 启动 php-fpm。win不支持php-fpm，因为php-fpm是使用Linux的fork()来做的，所以win下面基本上还是使用php-cgi。

```shell
# cd 到php二进制所在目录
# 常驻内存
php-cgi -b 127.0.0.1:9000
```



#### Ubuntu

Ubuntu 安装与 CentOs 类似，常见的错误修复如下：

```shell
# CentOS 使用 [yum install] 进行安装
# bzip2 libbz2-dev
sudo apt-get install bzip2 libbz2-dev

# 安装 - libssl-dev
sudo apt-get install libssl-dev

# curl-devel
apt-get install libcurl4-openssl-dev curl-devel -y

# php 7.3 easy.h should be in <curl-dir>/include/curl/
cd /usr/local/include sudo ; ln -s /usr/include/x86_64-linux-gnu/curl curl
sudo apt-get install libcurl4-gnutls-dev

# curl 升级
# centOS
# yum update curl -y
# yum install curl-devel -y

# libpng-dev
sudo apt-get install libpng-dev

# 其他依赖
# jpeglib.h not found
# Ubuntu: 
sudo apt-get install libxml2-dev libjpeg-dev 
# centos: 
# libjpeg-devel freetype-devel libjpeg-devel
yum install libjpeg libpng freetype libjpeg-devel libpng-devel freetype-devel -y

# error: png.h not found
yum install libpng libpng-devel -y
```





*编译项*

```shell
./configure --with-pdo-mysql --enable-bcmath --enable-mbstring --enable-sockets --with-gd --with-gettext --with-config-file-path=/usr/local/php/etc
```



也可直接去Ubuntu官网拉取对应的应用：

```shell
# 安装默认的 nginx 版本
apt install nginx
```





**配置文件**

```shell
php --ini
/php-7.3.5/php.ini-production

# 检测PHP语法是否正确
php -d display_startup_errors=1 -d error_reporting=-1 -d display_errors -c "/etc/php.ini" -m
```



**扩展编译**

```shell
# 转入扩展页面
cd ext-src

# 查看扩展信息，不存在需要安装（/usr/local/php/bin/phpize）
phpize

# 调用配置命令，注意错误信息（可以由此知道扩展所需的依赖）
./configure

# 编译并安装
make && make install
```





#### 集成环境

- lamp      https://www.lamp.sh/
- lampp    由Linux+Apache+Mysql+PHP构成，常用来运行web服务器。
- xampp   Apache + MySQL + PHP + Perl环境的一键安装程序。





##### linux 安装 SQL server 驱动

- 安装 SQLserver 驱动 https://github.com/microsoft/msphpsql
- 安装 msodbcsql17  https://pkgs.org/download/msodbcsql17
-  安装 mssql-tools-17 https://rhel.pkgs.org/7/microsoft-prod-x86_64/mssql-tools-17.10.1.1-1.x86_64.rpm.html



注意：msodbcsql18 客户端可能因 SSL 版本报错，旧版建议使用17版本。

参照：

1. [liunx ODBC 驱动安装](https://learn.microsoft.com/en-us/sql/connect/odbc/linux-mac/installing-the-microsoft-odbc-driver-for-sql-server?view=sql-server-ver16#redhat18)
2. [linux安装 odbc mysql_LINUX安装ODBC驱动](https://blog.csdn.net/weixin_42495556/article/details/113400413)
3. [Linux PHP连接sqlServer](https://blog.csdn.net/y_w_x_k/article/details/118213505)







### Nginx

nginx sever 配置

```nginx
 server {
        listen       80;
        server_name  localhost;
		
    	# 访问首页跳转其他页面
        rewrite ^/(.*) https://aurora.conero.cn/ permanent;
        #charset koi8-r;

        #access_log  logs/host.access.log  main;

        location / {
            #proxy_pass https://aurora.conero.cn/;
            #root   html;
            #index  index.html index.htm;
        }

        #error_page  404              /404.html;

        # redirect server error pages to the static page /50x.html
        #
        error_page   500 502 503 504  /50x.html;
        location = /50x.html {
            root   html;
        }

        # proxy the PHP scripts to Apache listening on 127.0.0.1:80
        #
        #location ~ \.php$ {
        #    proxy_pass   http://127.0.0.1;
        #}

        # pass the PHP scripts to FastCGI server listening on 127.0.0.1:9000
        #
        #location ~ \.php$ {
        #    root           html;
        #    fastcgi_pass   127.0.0.1:9000;
        #    fastcgi_index  index.php;
        #    fastcgi_param  SCRIPT_FILENAME  /scripts$fastcgi_script_name;
        #    include        fastcgi_params;
        #}

        # deny access to .htaccess files, if Apache's document root
        # concurs with nginx's one
        #
        #location ~ /\.ht {
        #    deny  all;
        #}
 }

```



nginx http 配置，使用 **include** 命令：

```nginx
http{
	# 引入 conf/conf.d 下所有 "*.conf" 配置文件
    include conf.d/*.conf;
}
```



> nginx linux部署主要（常见）问题

- nginx 服务用户权限问题，
- nginx `open_basedir ` 配置项处理。参考：[PHP防跨站之open_basedir目录设置](https://www.cnblogs.com/crxis/p/12720257.html)





#### 访问密码设置

使用 **htpasswd** 创建密码

```shell
# 安装 http 相关工具
yum  -y install httpd-tools

# 创建用户名和密码（conero）
htpasswd -c /usr/local/src/nginx/conero-passwd conero
```



配置文件更新

```nginx
server {
	# 登录密码认证要求
    auth_basic "It is Lab Private Net will need command for verify. ---Joshua Conero"; #这里是验证时的提示信息 
    auth_basic_user_file /usr/local/src/nginx/conero-passwd; 
}
```



重启系统即可。



### Apache

#### Windows

##### 安装

```powershell
# cd 到目标路径
# ./httpd -k install [<server_name>]
./httpd -k install

# 启动服务
net start <server_name>
```



##### 卸载

```powershell
# 删除服务名
sc delete Apache2.4
```



#### linux

```shell
# 查看编译时配置文件
httpd -V

# 查看配置文件语法情况
httpd -t
```





### 变量

#### 可变变量

*一个变量的变量名可以动态的设置和使用，`$$a`*

```php
<?php
$a = 'joshua';
$aa = 'esme';

// 三者相同 
echo $joshua;	//  输出：'esme'
echo ${$a};
echo $$a;
```



### session

*对于`单入口文件`+`地址重写` 等模式，可以仅仅在入口文件进行 `session_start()`操作即可。而多文件访问项目中，在涉及session相关的页面需要单独申明。*

 

### php-fpm

php-fpm  为 *FastCGI Process Manager* 的简称，即FastCGI 进程管理器。cgi   为 *Common Gateway Interface*，即通用网关接口，作为 Web 服务器调用外部程序时所使用的一种服务端应用的规范。cgi 可实现静态资源的web页面的动态请求，fpm是 php 版本的 FastCGI 协议实现。





### 配置(php.ini)

*memory_limit*    php 运行的内存限制。设置过大，系统对其设置的允许最大内存，反而导致 fastcgi 无响应，或者说 PHP 出错。以此该配置需要合理设置



```shell
# 查看当前的配置模块
php -i
# 查看服务的模块
php -m
```





#### 文件上传

php 网站文件上传一般受两个因素控制：

- php 配置，其设置过小时会出现报错。

  ```shell
  # php.ini
  # 开启文件上传
  file_uploads = On;
  
  # 文件上传大小
  upload_max_filesize = 1000M;
  
  # 文件上传数
  max_file_uploads = 20;
  ```

  

- nginx/Apache等服务配置

```shell
#nginx
client_max_body_size 5m;
```





## 工具

### PHPunit 单元测试框架

> 地址 https://phpunit.de/
>
> - 中文文档： https://phpunit.readthedocs.io/zh_CN/latest/index.html
>
> 中文 http://www.phpunit.cn/



*PHPUnit是一个面向PHP程序员的测试框架，这是一个xUnit的体系结构的单元测试框架。*





> 安装

通过 `composer` 包管理安装，以及`phar` 分发







### Xdebug 调试扩展

> 地址 https://xdebug.org/

windows 包安装：

- http://xdebug.org/download   下载 dll 安装即可
- 工具：dbgpClient (Command Line Debug Client)       IDE调试工具，如变量跟踪



修改配置

```ini
; 通过 zend 扩展安装
zend_extension = xdebug
```



安装成功检查

```shell
# 输入即可看到 xdebug 的版本信息
php -v
```



### composer

*PHP 包管理工具*

*国内镜像： https://developer.aliyun.com/composer*

国外官网： https://getcomposer.org/

包搜索地址：https://packagist.org



```powershell
# 默认拉取最新
composer require phpoffice/phpspreadsheet

# 指定版本
composer require 'phpoffice/phpspreadsheet:1.19'
```





### opcache

- @todo





## 资源

- 网站
  - [版本支持时间表(官网)](http://cn2.php.net/supported-versions.php)



### PHP-FIG/PSR

**PSR** *PHP Standard Recommendations 由 PHP FIG组织规定的 PHP 规范*

**PHP FIG**  *Framework Interoperability Group（框架可互用性小组），由多位开源框架开发者成立于 2009 年。*

网站： [中文网站](http://psr.phphub.org/)、[官网](https://www.php-fig.org/)



####  PSR-1 基础编码规范

- PHP代码文件 **必须** 以 `<?php` 或 `<?=` 标签开始；
- PHP代码文件 **必须** 以 `不带 BOM 的 UTF-8` 编码；
- PHP代码中 **应该** 只定义类、函数、常量等声明，或其他会产生 `副作用` 的操作（如：生成文件输出以及修改 .ini 配置文件等），二者只能选其一；
- 命名空间以及类 **必须** 符合 PSR 的自动加载规范： [[PSR-4](https://github.com/php-fig/fig-standards/blob/master/accepted/PSR-4-autoloader.md)] 。
- 类的命名 **必须** 遵循 `StudlyCaps` 大写开头的驼峰命名规范；
- 类中的常量所有字母都 **必须** 大写，单词间用下划线分隔；
- 方法名称 **必须** 符合 `camelCase` 式的小写开头驼峰命名规范。



#### PSR-2 编码风格规范

- 代码 **必须** 使用 4 个空格符而不是「Tab 键」进行缩进。
- 每行的字符数 **应该** 软性保持在 80 个之内，理论上 **一定不可** 多于 120 个，但 **一定不可** 有硬性限制。
- 每个 `namespace` 命名空间声明语句和 `use` 声明语句块后面，**必须** 插入一个空白行。
- 类的开始花括号（`{`） **必须** 写在函数声明后自成一行，结束花括号（`}`）也 **必须** 写在函数主体后自成一行。
- 方法的开始花括号（`{`） **必须** 写在函数声明后自成一行，结束花括号（`}`）也 **必须** 写在函数主体后自成一行。
- 类的属性和方法 **必须** 添加访问修饰符（`private`、`protected` 以及 `public`），`abstract` 以及 `final` **必须**声明在访问修饰符之前，而 `static` **必须** 声明在访问修饰符之后。
- 控制结构的关键字后 **必须** 要有一个空格符，而调用方法或函数时则 **一定不可** 有。
- 控制结构的开始花括号（`{`） **必须** 写在声明的同一行，而结束花括号（`}`） **必须** 写在主体后自成一行。
- 控制结构的开始左括号后和结束右括号前，都 **一定不可** 有空格符。



#### PSR-4 自动加载规范

*PSR-4 描述了从文件路径中 [自动加载](http://php.net/autoload) 类的规范。 它拥有非常好的兼容性，并且可以在任何自动加载规范中使用，包括 [PSR-0](https://github.com/php-fig/fig-standards/blob/master/accepted/PSR-0.md)。 PSR-4 规范也描述了放置 autoload 文件（就是我们经常引入的 `vendor/autoload.php`）的位置。*



- 术语「class」指的是类（classes）、接口（interfaces）、特征（traits）和其他类似的结构。
- 全限定类名具有以下形式：`\<NamespaceName>(\<SubNamespaceNames>)*\<ClassName>`
  - 全限定类名**必须**拥有一个顶级命名空间名称，也称为供应商命名空间（vendor namespace）。
  - 所有类名的引用**必须**区分大小写。
- 自动加载文件**禁止**抛出异常，**禁止**出现任何级别的错误，也**不建议**有返回值。



### 相关阅读资源

- [php 语言扩展阅读](http://www.nowamagic.net/librarys/veda/cate/PHP)
- [编译安装 PHP 7.0.3 亲测全攻略 & 以及如何单独的安装某个模块](https://blog.51cto.com/frankch/1746232)

