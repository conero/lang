# compiler
>2018年11月13日 星期二



### gcc

```ini
# 文件编译 <file> -> <name>
# c 语言
gcc <file> -o <name>
./<name>


# c++ 
g++ <file> -o <name>
./<name>

# 如编译 hw2023.cpp
# 可指定源代码编码-finput-charset, 命令行输出编码 -fexec-charset ， 以及 std c++版本号等
g++ '-std=c++17' '-finput-charset=UTF-8' '-fexec-charset=GBK' ./hw2023.cpp -o hw2023-gcc;./hw2023-gcc
```





### vc编译器

因其需要依赖于已存在的构建环境，需在其执行的命令行中执行，如下：

如 vc2019



- 搜索 vc2019，并打开命令行程序如 `x64 Native Tools CommandPrompt for VS 2019`
- 在命令行中cd到源码目录，`cd /todir/source`
- 执行编译命令 `cl /EHsc ./hw2023.cpp -o hw2023-cl`

