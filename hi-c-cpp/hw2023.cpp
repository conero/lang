#include <iostream>
using namespace std;

// C++ hello word
// 2023年12月5日 星期二
// 使用 gcc 编译器
// 
// 1)
// 编译并运行c++文件
// g++ .\hw.cpp -o hw-gcc;./hw-gcc
//
// 中文乱码处理
// -std 指定版本号， -finput-charset 指定源码编码  -fexec-charset 可执行文件的字符集
// g++ '-std=c++17' '-finput-charset=UTF-8' '-fexec-charset=GBK' .\hw.cpp -o hw-gcc;./hw-gcc
// 
// 2) 使用 cl 编译文件，锁定文件所在目录
//   在windows搜索中找到安装的对应 vsc 版本，并执行其对应的命令行窗口。cd到源目录，然后执行
// cl /EHsc ./hw.cpp -o hw-cl-v2019
// 
// 或在其他目录下指定c++运行目录编译
// cl /EHsc /I"D:\Program Files (x86)\Microsoft Visual Studio\2019\Community\VC\Tools\MSVC\14.29.30133\include\" ./hw.cpp
// 
int main() {
  cout << "Hello, World, By Joshua Conero!\n";
  cout << " .. 你好，中国!\n";
  cout << "-then\n";
  cout << " ..这是一个中文构建的世界!\n";
  cout << "-end\n";
  return 0;
}