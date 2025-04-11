// 2025年4月10日
// 通过 zig 来编译 c 语言代码
// 
// 编译代码：  
// zig cc hi-c-ps-zig.c
// zig cc hi-c-ps-zig.c -o out/hi-c-ps-zig.exe
//
// Windows 中文乱码: [Console]::OutputEncoding = [System.Text.Encoding]::UTF8
#include <stdio.h>

int main () {
    printf("hello world\n");
    printf(" 💗 贵阳 中国.\n");
    printf(" 💗 使用用 zig 编译的c语言代码\n");
    return 0;
}