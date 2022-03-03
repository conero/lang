# Java Visual  Machine(JVM)

> 2018年10月16日 星期二

## Java

LTS 版本有 8(14-22/30), 11(18-23/26), 17(21-26/29). 主要发布历史

| 版本      | GA     | 功能维护 | 安全维护 | 备注     |
| --------- | ------ | -------- | -------- | -------- |
| 7**LTS**  | 2011-7 | 2019-7   | 2022-7   |          |
| 8**LTS**  | 2014-3 | 2022-3   | 2030-12  | 应用最广 |
| 9         | 2017-9 | 2018-3   |          |          |
| 10        | 2018-3 | 2018-9   |          |          |
| 11**LTS** | 2018-9 | 2023-9   | 2026-9   |          |
| 12        | 2019-3 | 2019-9   |          |          |
| 13        | 2019-9 | 2020-3   |          |          |
| 14        | 2020-3 | 2020-9   |          |          |
| 15        | 2020-9 | 2021-3   |          |          |
| 16        | 2021-3 | 2022-9   |          |          |
| 17**LTS** | 2021-9 | 2026-9   | 2029-9   |          |
| 18        | 2022-3 | 2022-9   |          |          |
| 19        | 2022-9 | 2023-3   |          |          |
| 20        | 2023-3 | 2023-9   |          |          |
| 21**LTS** | 2023-9 | 2028-9   | 2021-9   |          |





### 基础

java版本以 11 版本作为标准。



#### 起步

- 思想：一切皆是对象。



java 程序无全局变量，其全部由类构成。其应用程序入口为类的 `main` 函数：

源码示例： [Simple.java](https://gitee.com/conero/lang/blob/hi-lang/hi-java/learning/Simple.java)

```java
// SimpleTest.java
class SimpleTest{
    // 静态代码块
    static {
        System.out.println("mian before. ");
    }

    // 程序入口文件
    public static void main(String[] args){
        System.out.println(" >> Hello World!");

        // for 循环
        for(var s: args){
            System.out.println("  ..."+s);
        }
    }
}
```

执行java即可：

```shell
# 执行文件
java .\SimapleTest.java  Hacker-Conero

#输入如下：
#mian before.
# >> Hello World!
#  ...Hacker-Conero

# 编译为 <.class> 文件
javac .\SimapleTest.java

# 指定 .class 所在目录进行 jar 打包
jar -cvf SimpleTest.jar .\SimapleTest.class

# 执行 jar, 需指定类路径
java -cp SimpleTest.jar SimpleTest
```



#### 类型

类型分为：基本类型、集合（复合）类型。如下

- 基本类型
  - 字符：char(16)
  - 布尔：boolean(1)
  - 数值(数据容量：`[0, (2^bit)-1]`, 有符号: `[-2^(bit-1), 2^(bit-1)-1]` )
    - 整数(位)：byte(8), short(16), int(32), long(64),
    - 浮点数：float(32), double(64)
- 符合类型：list, map



> Java中定义的简单类型、占用二进制位数及对应的封装器类。

| 简单类型   | boolean | byte | char      | short | Int     | long | float | double | void |
| ---------- | ------- | ---- | --------- | ----- | ------- | ---- | ----- | ------ | ---- |
| 二进制位数 | 1       | 8    | 16        | 16    | 32      | 64   | 32    | 64     | --   |
| 封装器类   | Boolean | Byte | Character | Short | Integer | Long | Float | Double | Void |



Java基本类型存储在栈中，因此它们的存取速度要快于存储在堆中的对应包装类的实例对象；其数据存储相对简单，运算效率比较高。



##### 常量

源码示例： [VariableConst.java](https://gitee.com/conero/lang/blob/hi-lang/hi-java/learning/VariableConst.java)

```java
class VariableConst{
    public static final double PI = 3.1415926535898;	// 静态类常量
    final int Index = 220303; // 类成员常量
    
    public static void main(String[] args){
        // 局部常量
        final int V16 = 0xff;   // 0x/0X  16 进制
        final byte V8 = 013;     // 8进制
        final long vLong = 12L;

        // 自动推断
        var f16 = 3.51f;
    }
}
```







#### 语句





## 应用工具

### Java 构建工具

#### Ant, Maven, Gradle

> Ant - 2000 年/ 格式-XML

*Ant是第一个“现代”构建工具，在很多方面它有些像Make。Ant的主要优点在于对构建过程的控制上*



> Maven - 2004 年/  格式-XML

*目的是解决码农使用Ant所带来的一些问题。*

*Maven具备从网络上自动下载依赖的能力（Ant后来通过Ivy也具备了这个功能），这一点革命性地改变了我们开发软件的方式。*



> Gradle - 2012 年/ DSL

*Gradle结合了前两者的优点，在此基础之上做了很多改进。它具有Ant的强大和灵活，又有Maven的生命周期管理且易于使用。*





## 附录

### 参考

- [java 基本类型](https://www.cnblogs.com/doit8791/archive/2012/05/25/2517448.html)
