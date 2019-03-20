# PHP 扩展性学习



## 基础

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

