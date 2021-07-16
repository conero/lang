# PHP 扩展性学习



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

**CentOS 7.3**

```shell
# 安装编译环境
$ yum -y install make zlib zlib-devel gcc-c++ libtool  openssl openssl-devel

# ----------------------【nginx】------------------------
# 安装 nginx
$ wget <loaddown.net.url>
$ tar zxvf nginx-1.6.2.tar.gz	解压
# 解压 tag 文件夹 `tar xvf file.tar`
$ cd nginx-1.6.2				目录转移
$ ./configure					编译安装
$ make							安装/ make install

# 启动 nginx
$ /usr/local/nginx/sbin/nginx
# 停止服务器
$ /usr/local/nginx/sbin/nginx -s stop 或 /usr/local/nginx/sbin/nginx -s quick

# 配置文件检测
$ /usr/local/nginx/sbin/nginx -t
# openssl 安装
$ yum -y install openssl*

# ----------------------【php】------------------------
# 依赖安装
$ yum install -y libxml2*
$ wget https://www.php.net/distributions/php-7.3.5.tar.gz
$ tar -xzvf <tag.gz>
# 配置文件
$ ./buildconf --force
$ ./configure
# 编译
# 有事需要清理编译环境 `make clean`
$ make && make install


# 安装/启动 php-fpm
$ yum install php-fpm
# 启动 php-fpm 服务
$ service php-fpm start
# php-fpm 详情
$ systemctl status php-fpm.service
# 杀死进程
$ killall php-fpm
# 检测是否已经正确启动
$ netstat -tln | grep 9000
# php-fpm 的进程号
$ ps aux|grep php-fpm
```



*编译项*

```shell
./configure --with-pdo-mysql --enable-bcmath --enable-mbstring --enable-sockets --with-gd --with-gettext --with-config-file-path=/usr/local/php/etc
```





**配置文件**

```shell
php --ini
/php-7.3.5/php.ini-production

# 检测PHP语法是否正确
php -d display_startup_errors=1 -d error_reporting=-1 -d display_errors -c "/etc/php.ini" -m
```



### Nginx

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

 

### 配置(php.ini)

*memory_limit*    php 运行的内存限制。设置过大，系统对其设置的允许最大内存，反而导致 fastcgi 无响应，或者说 PHP 出错。以此该配置需要合理设置



```shell
# 查看当前的配置模块
php -i
# 查看服务的模块
php -m
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



### composer

*PHP 包管理工具*

*国内镜像： https://developer.aliyun.com/composer*

国外官网： https://getcomposer.org/



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

